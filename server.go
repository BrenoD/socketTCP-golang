package main

import (
	"fmt"
	"net"
	"strings"
)

// info da rota
type Route struct {
	Path    string
	Method  string
	Handler func() string
}

// var global em lista
var routes []Route

func main() {

	//preenche a lista
	routes = append(routes, Route{
		Path:   "/",
		Method: "GET",
		Handler: func() string {
			return `HTTP/1.1 200 OK
		Content-Type: text/html; charset=utf-8
		<h1>Acho q aprendi rsrs</h1>`
		},
	})
	//a partir do net.Listen ele ouve a conexao tcp
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Servidor ouvindo na porta 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	//cria, le e armazena no buffer
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Erro ao ler dados:", err)
		return
	}

	request := string(buffer[:n])
	requestLines := strings.Split(request, "\r\n")

	//caso a conexao tenha menos de uma linha, f
	if len(requestLines) < 1 {
		conn.Write([]byte("HTTP/1.1 400 Bad Request\r\n\r\n"))
		return
	}

	//divide a linha, se tiver menos de duas, f
	firstLine := strings.Split(requestLines[0], " ")
	if len(firstLine) < 2 {
		conn.Write([]byte("HTTP/1.1 400 Bad Request\r\n\r\n"))
		return
	}

	//extração
	method := firstLine[0]
	path := firstLine[1]

	//se a rota for igual ao caminho e o method for igual o method, resposta sera enviada
	for _, route := range routes {
		if route.Path == path && route.Method == method {
			response := route.Handler()
			conn.Write([]byte(response))
			return
		}
	}

	// 404 se nenhuma rota for encontrada
	conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\nNot found"))
}
