const cbor = require('cbor')
const http2 = require('http2');
const fs = require('fs');
const client = http2.connect('https://localhost:8443', {
  ca: fs.readFileSync('server.crt')
});
client.on('error', (err) => console.error(err));

const req = client.request({ ':path': '/' });

req.on('response', (headers, flags) => {
  for (const name in headers) {
    console.log(`${name}: ${headers[name]}`);
  }
});

let data = [];
req.on('data', (chunk) => { 
  const decoded = cbor.decode(Buffer.from(chunk));
  console.log(`rcv: ${JSON.stringify(decoded)}`)
  data.push(decoded) 
});
req.on('end', () => {
  console.log(`\n${JSON.stringify(data)}`);
  client.close();
});
req.end();