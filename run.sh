#!/bin/bash

# Verificar se a porta 8080 está em uso e matar o processo, se necessário
PORT=8080
if lsof -i :$PORT &>/dev/null; then
    echo "Porta $PORT já está em uso. Finalizando processo..."
    PID=$(lsof -ti :$PORT)
    kill $PID
    sleep 1
fi

# Iniciar o servidor em background
echo "Iniciando o servidor..."
go run server.go &
SERVER_PID=$!

# Esperar o servidor inicializar completamente
sleep 2

# Executar o cliente
echo "Executando o cliente..."
go run client.go

# Parar o servidor
echo "Parando o servidor..."
kill $SERVER_PID
