package main

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8090"
	TCP_PORT    = "8080"
)

func main() {
	go startTCPServer()

	listener, err := net.Listen("tcp", SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Servidor TCP ouvindo em", SERVER_HOST+":"+SERVER_PORT)

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}

		go handleConnection(connection)
	}
}

func startTCPServer() {
	listener, err := net.Listen("tcp", SERVER_HOST+":"+TCP_PORT)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor TCP:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Servidor TCP ouvindo em", SERVER_HOST+":"+TCP_PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}

		go handleTCPConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erro ao ler dados:", err)
		return
	}

	fmt.Println("Recebido:", string(buffer[:n]))

	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 23\r\n" +
		"\r\n" +
		"Mensagem recebida via TCP"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Erro ao enviar dados:", err)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erro ao ler dados:", err)
		return
	}

	data := string(buffer[:n])
	fmt.Println("Recebido no TCP:", data)

	httpResponse := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: " + fmt.Sprintf("%d", len(data)) + "\r\n" +
		"\r\n" +
		data

	_, err = conn.Write([]byte(httpResponse))
	if err != nil {
		fmt.Println("Erro ao enviar resposta HTTP:", err)
	}
}
