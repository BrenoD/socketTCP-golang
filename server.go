package main

import (
	"fmt"
	"net"
)

// Configuração do servidor
const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8090"
	SERVER_TYPE = "tcp"
)

func main() {
	listener, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Servidor ouvindo em", SERVER_HOST+":"+SERVER_PORT)

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}

		go handleConnection(connection)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	mLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erro ao ler dados:", err)
		return
	}

	fmt.Println("Recebido:", string(buffer[:mLen]))

	// Enviando uma resposta de volta ao cliente
	_, err = conn.Write([]byte("Mensagem recebida"))
	if err != nil {
		fmt.Println("Erro ao enviar dados:", err)
	}
}
