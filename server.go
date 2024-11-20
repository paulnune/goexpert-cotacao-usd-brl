package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	apiURL         = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
	databaseFile   = "cotacoes.db"
	serverPort     = ":8080"
	cotacaoTimeout = 2 * time.Second        // Timeout para a chamada à API
	dbTimeout      = 100 * time.Millisecond // Timeout para persistência no banco
)

// Reutilizar conexões HTTP com Keep-Alive
var httpClient = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: false,
	},
	Timeout: cotacaoTimeout,
}

type CotacaoAPIResponse struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	// Configuração do banco de dados SQLite
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		log.Fatalf("Erro ao abrir banco de dados: %v", err)
	}
	defer db.Close()

	// Criação da tabela, se não existir
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY, bid TEXT, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatalf("Erro ao criar tabela no banco de dados: %v", err)
	}

	// Configuração do endpoint
	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		// Timeout para a operação de obtenção de cotação
		ctx, cancel := context.WithTimeout(r.Context(), cotacaoTimeout)
		defer cancel()

		cotacao, err := obterCotacao(ctx)
		if err != nil {
			http.Error(w, "Erro ao obter cotação", http.StatusInternalServerError)
			log.Printf("Erro ao obter cotação: %v", err)
			return
		}

		// Timeout para persistência no banco de dados
		ctxDB, cancelDB := context.WithTimeout(context.Background(), dbTimeout)
		defer cancelDB()

		if err := salvarCotacao(ctxDB, db, cotacao.Bid); err != nil {
			http.Error(w, "Erro ao salvar cotação no banco", http.StatusInternalServerError)
			log.Printf("Erro ao salvar cotação no banco: %v", err)
			return
		}

		// Retornar a cotação em JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cotacao)
	})

	// Inicializar o servidor
	log.Printf("Servidor rodando na porta %s", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, nil))
}

// Função para obter a cotação da API externa
func obterCotacao(ctx context.Context) (Cotacao, error) {
	log.Println("Iniciando chamada para a API externa...")
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		log.Printf("Erro ao criar requisição: %v", err)
		return Cotacao{}, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Erro ao realizar requisição: %v", err)
		return Cotacao{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Resposta inválida da API: %d", resp.StatusCode)
		return Cotacao{}, err
	}

	log.Println("Resposta da API recebida com sucesso...")
	var apiResponse CotacaoAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Printf("Erro ao decodificar resposta: %v", err)
		return Cotacao{}, err
	}

	log.Println("Cotação obtida com sucesso:", apiResponse.USDBRL.Bid)
	return Cotacao{Bid: apiResponse.USDBRL.Bid}, nil
}

// Função para salvar a cotação no banco de dados
func salvarCotacao(ctx context.Context, db *sql.DB, bid string) error {
	query := "INSERT INTO cotacoes (bid) VALUES (?)"
	done := make(chan error)

	go func() {
		_, err := db.Exec(query, bid)
		done <- err
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}
