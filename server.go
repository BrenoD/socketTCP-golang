package main

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "8090"
	TCP_PORT    = "8080"
)

type Message struct {
	Message string `json:"message"`
}

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

	fmt.Println("Servidor TCP simulando HTTP ouvindo em localhost:" + TCP_PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}

		go handleHTTPRequest(conn)
	}
}

func handleHTTPRequest(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erro ao ler dados:", err)
		return
	}

	data := string(buffer[:n])
	lines := strings.Split(data, "\r\n")
	if len(lines) < 2 {
		fmt.Println("Requisição inválida")
		return
	}

	body := lines[len(lines)-1]

	// parsing do JSON
	var msg Message
	err = json.Unmarshal([]byte(body), &msg)
	if err != nil {
		fmt.Println("Erro ao fazer parse do JSON:", err)
		return
	}

	response, err := sendToTCPServer(msg.Message)
	if err != nil {
		fmt.Println("Erro ao enviar dados para o servidor TCP:", err)
		return
	}

	httpResponse := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: " + fmt.Sprintf("%d", len(response)) + "\r\n" +
		"\r\n" +
		response
	conn.Write([]byte(httpResponse))
}

func sendToTCPServer(message string) (string, error) {
	conn, err := net.Dial("tcp", SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message))
	if err != nil {
		return "", err
	}

	buffer := make([]byte, 1024)
	mLen, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:mLen]), nil
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

	// resposta de volta ao cliente
	_, err = conn.Write([]byte("Mensagem recebida via TCP"))
	if err != nil {
		fmt.Println("Erro ao enviar dados:", err)
	}
}
