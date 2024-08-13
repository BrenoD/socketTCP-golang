const net = require('net');

//config do cliente
const options = {
  host: 'localhost',
  port: 8080
};

//conexão TCP com o servidor
const client = net.createConnection(options, () => {
  console.log('Conectado ao servidor!');

  //envio de uma requisição HTTP para o servidor
  const request = `GET / HTTP/1.1\r\nHost: localhost\r\n\r\n`;
  client.write(request);
});

//exec a resposta do servidor
client.on('data', (data) => {
  console.log('Resposta do servidor:');
  console.log(data.toString());

  //fim da conexao
  client.end();
});

// Lidando com erros na conexão
client.on('error', (err) => {
  console.error('Erro na conexão:', err);
});
