package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

const (
	SERVER_HOST      = "localhost"
	SERVER_PORT      = "8090"
	SERVER_TYPE      = "tcp"
	HTTP_SERVER_PORT = "8080"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	go startHTTPServer()

	listener, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(connection)
	}
}

func startHTTPServer() {
	http.HandleFunc("/send", handleHTTP)
	http.ListenAndServe(":"+HTTP_SERVER_PORT, nil)
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var msg Message
	json.Unmarshal(body, &msg)

	response, _ := sendToTCPServer(msg.Message)

	w.Write([]byte(response))
}

func sendToTCPServer(message string) (string, error) {
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	conn.Write([]byte(message))

	buffer := make([]byte, 1024)
	mLen, _ := conn.Read(buffer)

	return string(buffer[:mLen]), nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	mLen, _ := conn.Read(buffer)

	response := "Mensagem recebida via TCP"

	// Parse da requisição HTTP
	parseHTTPRequest(string(buffer[:mLen]))

	conn.Write([]byte(response))
}

func parseHTTPRequest(request string) {
	lines := strings.Split(request, "\n")
	method, path, version := parseRequestLine(lines[0])
	headers := parseHeaders(lines[1:])
	fmt.Println("Método:", method)
	fmt.Println("Caminho:", path)
	fmt.Println("Versão HTTP:", version)
	fmt.Println("Cabeçalhos:", headers)
}

func parseRequestLine(line string) (string, string, string) {
	parts := strings.Split(line, " ")
	return parts[0], parts[1], parts[2]
}

func parseHeaders(lines []string) map[string]string {
	headers := make(map[string]string)
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			break
		}
		parts := strings.SplitN(line, ": ", 2)
		headers[parts[0]] = parts[1]
	}
	return headers
}
