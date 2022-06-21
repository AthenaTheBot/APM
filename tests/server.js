const http = require("http");

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.write("Hello World!");
  res.end();
});

server.listen(4040, undefined, undefined, () => {
  console.log("Started server!");
});
