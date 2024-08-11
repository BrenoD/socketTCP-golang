package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

// Configuração do servidor
const (
	SERVER_HOST     = "localhost"
	SERVER_PORT     = "8090"
	SERVER_TYPE     = "tcp"
	HTTP_SERVER_PORT = "8080"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	// Iniciando o servidor HTTP em uma nova goroutine
	go startHTTPServer()

	// Configurando o servidor TCP
	listener, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
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

func startHTTPServer() {
	http.HandleFunc("/send", handleHTTP)
	fmt.Println("Servidor HTTP ouvindo em localhost:" + HTTP_SERVER_PORT)
	if err := http.ListenAndServe(":"+HTTP_SERVER_PORT, nil); err != nil {
		fmt.Println("Erro ao iniciar o servidor HTTP:", err)
	}
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// parse do JSON
	var msg Message
	err = json.Unmarshal(body, &msg)
	if err != nil {
		http.Error(w, "Erro ao fazer parse do JSON", http.StatusBadRequest)
		return
	}

	// mensagem para o servidor TCP
	response, err := sendToTCPServer(msg.Message)
	if err != nil {
		http.Error(w, "Erro ao enviar dados para o servidor TCP", http.StatusInternalServerError)
		return
	}

	// respondendo de volta ao cliente HTTP
	w.Write([]byte(response))
}

func sendToTCPServer(message string) (string, error) {
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
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
