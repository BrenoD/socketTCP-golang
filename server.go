package main

import (
	"fmt"
	"net"
)

type Route struct {
	Path    string
	Method  string
	Handler func(map[string]interface{}) Response
}

// struct response
type Response struct {
	Status  int
	Headers map[string]string
	Body    string
}

var routes []Route

func main() {
	routes = append(routes, Route{
		Path:   "/",
		Method: "GET",
		Handler: func(data map[string]interface{}) Response {
			return Response{
				Status: 200,
				Headers: map[string]string{
					"Content-Type": "text/html; charset=utf-8",
				},
				Body: "<h1>Olá, cliente!</h1>",
			}
		},
	})

	//servidor
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server listening on port 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}

		go handleConnection(conn)
	}
}

//processa a conexão recebida.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	//obj
	obj := map[string]interface{}{
		"headers":     map[string]string{},
		"status":      200,
		"httpVersion": "1.1",
		"path":        "/",
		"body":        "Olá, cliente!",
	}

	for _, route := range routes {
		if route.Path == obj["path"] && route.Method == "GET" {
			response := route.Handler(obj)
			conn.Write([]byte(toHTTP(response)))
			return
		}
	}

	//404 se a rota não for encontrada
	notFoundResponse := Response{
		Status:  404,
		Headers: map[string]string{},
		Body:    "Not found",
	}
	conn.Write([]byte(toHTTP(notFoundResponse)))
}

// transformando para string
func toHTTP(response Response) string {
	statusLine := fmt.Sprintf("HTTP/1.1 %d %s\r\n", response.Status, getStatusText(response.Status))
	headers := ""
	for key, value := range response.Headers {
		headers += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	return statusLine + headers + "\r\n" + response.Body
}

func getStatusText(status int) string {
	switch status {
	case 200:
		return "OK"
	case 404:
		return "Not Found"
	default:
		return "Unknown Status"
	}
}
