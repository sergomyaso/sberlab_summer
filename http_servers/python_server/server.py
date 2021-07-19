from http.server import HTTPServer, BaseHTTPRequestHandler


class SimpleHTTPRequestHandler(BaseHTTPRequestHandler):

    def do_GET(self):
        self.send_response(200)
        self.end_headers()
        self.wfile.write(b'Hola, mondo from python')

server_port = 8080
httpd = HTTPServer(('0.0.0.0', server_port), SimpleHTTPRequestHandler)
httpd.serve_forever()