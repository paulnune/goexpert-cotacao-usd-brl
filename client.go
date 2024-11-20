package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const serverURL = "http://localhost:8080/cotacao"

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	// Timeout de 1 segundo para a requisição HTTP
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	log.Println("Iniciando requisição ao servidor...")
	req, err := http.NewRequestWithContext(ctx, "GET", serverURL, nil)
	if err != nil {
		log.Fatalf("Erro ao criar requisição: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Erro ao realizar requisição ao servidor: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Resposta inválida do servidor: %s", resp.Status)
	}

	log.Println("Resposta recebida com sucesso do servidor. Processando...")
	var cotacao Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		log.Fatalf("Erro ao decodificar resposta: %v", err)
	}

	// Salvar a cotação no arquivo
	valor := fmt.Sprintf("Dólar: %s", cotacao.Bid)
	if err := ioutil.WriteFile("cotacao.txt", []byte(valor), 0644); err != nil {
		log.Fatalf("Erro ao salvar cotação no arquivo: %v", err)
	}

	log.Println("Cotação salva no arquivo 'cotacao.txt' com sucesso!")
}
