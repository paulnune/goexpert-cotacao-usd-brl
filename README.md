
# 💰 GoExpert - Cotação USD/BRL 📈

Bem-vindo ao repositório do desafio da **Pós GoExpert 2024**! Este projeto foi desenvolvido por **Paulo Nunes** para demonstrar conhecimentos em **Go**, abrangendo conceitos como servidores HTTP, contextos, banco de dados SQLite e manipulação de arquivos.

---

## 📋 Descrição do Projeto

Este projeto consiste em dois sistemas principais, desenvolvidos em **Go**:

1. **`server.go`**: Um servidor HTTP que consome a API pública de câmbio [AwesomeAPI](https://economia.awesomeapi.com.br/json/last/USD-BRL), retorna a cotação atual do dólar para o cliente e salva os dados em um banco de dados SQLite.

2. **`client.go`**: Um cliente HTTP que realiza uma requisição ao servidor, recebe a cotação do dólar e salva o valor em um arquivo chamado `cotacao.txt`.

---

## 🚀 Funcionalidades

- **Server**:
  - Requisição à API pública de câmbio com timeout.
  - Persistência da cotação no banco de dados SQLite com controle de timeout.
  - Resposta em formato JSON no endpoint `/cotacao`.

- **Client**:
  - Requisição ao servidor com controle de timeout.
  - Salvamento da cotação recebida no arquivo `cotacao.txt`.

---

## 🛠️ Tecnologias Utilizadas

- **Linguagem:** Go
- **Banco de Dados:** SQLite
- **Bibliotecas:** 
  - `github.com/mattn/go-sqlite3`
  - Pacotes nativos como `net/http`, `context`, `encoding/json`, `io/ioutil`, entre outros.

---

## ⚙️ Como Executar o Projeto

### Pré-requisitos
- [Go instalado](https://golang.org/dl/) (versão 1.20 ou superior)
- Conexão com a internet para acessar a [API AwesomeAPI](https://economia.awesomeapi.com.br/json/last/USD-BRL)

### Passo a Passo

1. Clone o repositório:
   ```bash
   git clone https://github.com/paulnune/goexpert-cotacao-usd-brl.git
   cd goexpert-cotacao-usd-brl
   ```

2. Inicialize o módulo do Go:
   ```bash
   go mod init goexpert-cotacao-usd-brl
   ```

3. Instale a dependência do SQLite:
   ```bash
   go get github.com/mattn/go-sqlite3
   ```

4. Organize as dependências:
   ```bash
   go mod tidy
   ```

5. Inicie o servidor:
   ```bash
   go run server.go
   ```
   O servidor estará disponível em http://localhost:8080/cotacao.

6. Em outro terminal, execute o cliente:
   ```bash
   go run client.go
   ```

7. Verifique os resultados:
   - O arquivo `cotacao.txt` conterá a cotação no formato:
     ```
     Dólar: {valor}
     ```
   - O banco de dados SQLite (`cotacoes.db`) terá o histórico de cotações salvas.

---

## 📂 Estrutura do Projeto

```
├── client.go       # Sistema cliente que consome o servidor
├── server.go       # Servidor HTTP que consome a API e persiste no banco
├── cotacoes.db     # Banco de dados SQLite (gerado automaticamente)
├── cotacao.txt     # Arquivo contendo a cotação atual (gerado pelo client)
└── go.mod          # Arquivo de dependências do Go
```

---

## 📖 Referências

- [AwesomeAPI](https://docs.awesomeapi.com.br/api-de-moedas)  
- [Documentação oficial do Go](https://golang.org/doc/)

---

## 👨‍💻 Autor

### **Paulo Henrique Nunes Vanderley**  
- 🌐 [Site Pessoal](https://www.paulonunes.dev/)  
- 🌐 [GitHub](https://github.com/paulnune)  
- ✉️ Email: [paulo.nunes@live.de](mailto:paulo.nunes@live.de)  
- 🚀 Aluno da Pós **GoExpert 2024** pela [FullCycle](https://fullcycle.com.br)

---

## 🎉 Agradecimentos

Este repositório foi desenvolvido com muita dedicação para a **Pós GoExpert 2024**. Agradeço à equipe da **FullCycle** por proporcionar uma experiência incrível de aprendizado! 🚀