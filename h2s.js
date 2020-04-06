const cbor = require('cbor')
const http2 = require('http2');
const fs = require('fs');

const server = http2.createSecureServer({
  key: fs.readFileSync('server.key'),
  cert: fs.readFileSync('server.crt')
});
server.on('error', (err) => console.error(err));

server.on('stream', (stream, headers) => {
  stream.setEncoding('binary')
  stream.setDefaultEncoding('binary')
  // stream is a Duplex
  stream.respond({
    'content-type': 'application/cbor',
    ':status': 200
  });

  stream.write(cbor.encode([42, "col1", "col2", "col3", 15123.15123]))
  stream.end(cbor.encode([43, "col1", "col2", "col3", 756783.0987]));
});

server.listen(8443);