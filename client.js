const net = require('net');

// Configuração do servidor
const SERVER_HOST = 'localhost';
const SERVER_PORT = 8090;

// Estabelecendo conexão
const client = new net.Socket();
client.connect(SERVER_PORT, SERVER_HOST, () => {
    console.log('Conectado ao servidor');

    // Enviando dados
    client.write('interessante, nao?');
});

// Lidando com dados recebidos do servidor
client.on('data', (data) => {
    console.log('Recebido: ' + data.toString());

    // Fechando a conexão após receber a resposta
    client.destroy();
});

// Lidando com fechamento da conexão
client.on('close', () => {
    console.log('Conexão fechada');
});

// Lidando com erros
client.on('error', (err) => {
    console.error('Erro: ' + err.message);
});