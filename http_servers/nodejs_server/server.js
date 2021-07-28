
function loadGenartor() {
  var countIteration = 150;
  var sum = 0;
  for(var i = 0; i < countIteration; i++) {
    for(var j = 0; j < countIteration; j++) {
      sum += 1;
    }
  }
  return sum;
}

//Load HTTP module
const http = require("http");
const hostname = '0.0.0.0';
const port = 8080;

//Create HTTP server and listen on port 3000 for requests
const server = http.createServer((req, res) => {

  //Set the response HTTP header with HTTP status and Content type
  var strRes = "Hola, mondo from nodeJs\n Your sum, brother, =" + loadGenartor();
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/plain');
  res.end(strRes);
});

//listen for request on port 3000, and as a callback function have the port listened on logged
server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});