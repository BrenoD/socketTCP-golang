const http = require('http');

// configuracao
const SERVER_HOST = 'localhost';
const SERVER_PORT = 8080;

// dados para a req
const postData = JSON.stringify({ message: 'interessante, nao?' });

// config http
const options = {
    hostname: SERVER_HOST,
    port: SERVER_PORT,
    path: '/send',
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
        'Content-Length': Buffer.byteLength(postData)
    }
};

// req
const req = http.request(options, (res) => {
    console.log(`STATUS: ${res.statusCode}`);
    res.setEncoding('utf8');
    res.on('data', (chunk) => {
        console.log(`Corpo: ${chunk}`);
    });
});

// error
req.on('error', (e) => {
    console.error(`Problema com a requisição: ${e.message}`);
});

// dados da req
req.write(postData);
req.end();
