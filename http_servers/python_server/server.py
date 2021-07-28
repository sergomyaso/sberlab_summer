from http.server import HTTPServer, BaseHTTPRequestHandler


class SimpleHTTPRequestHandler(BaseHTTPRequestHandler):
    count_iteration = 150

    def create_load(self):
        sum_res = 0
        for i in range(self.count_iteration):
            for j in range(self.count_iteration):
                sum_res += 1
        return sum_res

    def do_GET(self):
        self.send_response(200)
        self.end_headers()
        sum_res = self.create_load()
        resp_string = 'Hola, mondo from python!\n' + "Your sum, brother, = " + str(sum_res)
        self.wfile.write(resp_string.encode())


server_port = 8080
httpd = HTTPServer(('0.0.0.0', server_port), SimpleHTTPRequestHandler)
httpd.serve_forever()
