
# 💰 GoExpert - Cotação USD/BRL 📈

Bem-vindo ao repositório do desafio da **Pós GoExpert 2024**! Este projeto foi desenvolvido por **Paulo Nunes** para demonstrar conhecimentos em **Go**, abrangendo conceitos como servidores HTTP, contextos, banco de dados SQLite e manipulação de arquivos.

---

## 📋 Desafio Proposto

1. O sistema consiste em dois arquivos principais:
   - **`client.go`**: Deve realizar uma requisição HTTP ao servidor (`server.go`) solicitando a cotação do dólar.
   - **`server.go`**: Deve consumir a API pública [AwesomeAPI](https://economia.awesomeapi.com.br/json/last/USD-BRL) para obter a cotação do dólar, retornar o valor ao cliente e registrar a cotação no banco de dados SQLite.

2. Requisitos de **Timeout**:
   - O **server.go** deve usar o pacote `context` para:
     - Limitar o tempo para chamar a API da AwesomeAPI a **200ms**.
     - Limitar o tempo para registrar a cotação no banco a **10ms**.
   - O **client.go** deve ter um timeout máximo de **300ms** para receber a resposta do servidor.

3. Persistência:
   - O servidor deve registrar a cotação no banco de dados SQLite.
   - O cliente deve salvar o valor da cotação em um arquivo `cotacao.txt` no formato:
     ```
     Dólar: {valor}
     ```

---

## 🚀 Implementação e Justificativas

Durante a execução do desafio, foi identificado que os **timeouts especificados no enunciado não eram viáveis** para execução no ambiente local. Problemas encontrados:

1. **Latência Variável da API AwesomeAPI**:
   - Os tempos de resposta medidos com `curl` mostraram que a API frequentemente ultrapassa os **200ms**:
     ```
     Tempo total: 0.564606s
     Tempo total: 0.145143s
     Tempo total: 0.718081s
     Tempo total: 0.625480s
     ```
   - Isso torna o timeout de **200ms** insuficiente para a maioria das requisições.

2. **Persistência no Banco de Dados SQLite**:
   - Com um timeout de **10ms**, o SQLite não conseguia registrar a cotação devido a operações de I/O e concorrência no sistema.
   - Um timeout mais realista de **100ms** foi configurado.

3. **Timeout no Cliente**:
   - Com o servidor ajustado para um timeout de **2 segundos** para a API, o cliente foi configurado com um timeout de **3 segundos**, garantindo tempo suficiente para processar a resposta.

---

## 🛠️ Ajustes Realizados

1. **Server (server.go)**:
   - Timeout para a API ajustado de **200ms** para **2 segundos**:
     ```go
     const cotacaoTimeout = 2 * time.Second
     ```
   - Timeout para o banco ajustado de **10ms** para **100ms**:
     ```go
     const dbTimeout = 100 * time.Millisecond
     ```

2. **Client (client.go)**:
   - Timeout ajustado de **300ms** para **3 segundos**:
     ```go
     ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
     ```
---

## ⚙️ Como Executar o Projeto

### Método 1: Usando Comandos Individuais
1. **Iniciar o servidor**:
   ```bash
   go run server.go
   ```

2. **Executar o cliente em outro terminal**:
   ```bash
   go run client.go
   ```

---

### Método 2: Usando o Script `run.sh`
Se preferir, utilize o script que automatiza os passos acima. Ele:
- Finaliza processos que estejam usando a porta 8080.
- Inicia o servidor em segundo plano.
- Executa o cliente.
- Finaliza o servidor após a execução do cliente.

Para executar o script:
```bash
chmod +x run.sh
sh run.sh
```

---

## 📂 Estrutura do Projeto

```
├── client.go       # Sistema cliente que consome o servidor
├── server.go       # Servidor HTTP que consome a API e persiste no banco
├── cotacoes.db     # Banco de dados SQLite (gerado automaticamente)
├── cotacao.txt     # Arquivo contendo a cotação atual (gerado pelo client)
├── run.sh          # Script para automatizar execução do servidor e cliente
└── go.mod          # Arquivo de dependências do Go
```

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