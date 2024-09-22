🌐 TCP Socket Project with Golang and JavaScript

This project implements TCP communication between a server written in Go and a client in JavaScript—perfect for learning about networking and protocols!

📁 Project Structure

server.go: Implements the TCP server in Go.
client.js: JavaScript client script that connects to the server.
🚀 How to Run the Project

Clone the repository:
bash
Copy code
git clone https://github.com/BrenoD/socketTCP-golang.git
Run the server (Go):
Make sure you have Go installed, then execute:

bash
Copy code
go run server.go
Run the client (JavaScript):
You can use Node.js to run the client:

bash
Copy code
node client.js
💬 How It Works

The server listens on a specific TCP port and waits for connections.
The client sends messages to the server, which responds back.
📚 Requirements

Go 1.16+
Node.js (for running the JavaScript client)
🛠️ Customization

Modify the code to test different ports or message formats.
Experiment with multiple clients connecting simultaneously.
🤝 Contributing

Feel free to submit pull requests or open issues to improve the project!
