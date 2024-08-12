const http = require('http');

// dados para a req
const SERVER_HOST = 'localhost';
const SERVER_PORT = 8080;

// dados req
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

// parse de respostas HTTP
function parseHTTPResponse(response) {
    const lines = response.split('\n');
    const [version, statusCode, statusMessage] = lines[0].split(' ');
    const headers = {};

    for (let i = 1; i < lines.length; i++) {
        const line = lines[i].trim();
        if (line === '') break;
        const [key, value] = line.split(': ');
        headers[key] = value;
    }

    const body = lines.slice(lines.indexOf('') + 1).join('\n');
    return { version, statusCode, statusMessage, headers, body };
}

// req
const req = http.request(options, (res) => {
    let rawData = '';
    res.setEncoding('utf8');
    res.on('data', (chunk) => { rawData += chunk; });
    res.on('end', () => {
        const parsedResponse = parseHTTPResponse(`HTTP/${res.httpVersion} ${res.statusCode} ${res.statusMessage}\n${rawData}`);
        console.log(parsedResponse);
    });
});

// error
req.on('error', (e) => {
    console.error(`Problema com a requisição: ${e.message}`);
});

// dados da req
req.write(postData);
req.end();
