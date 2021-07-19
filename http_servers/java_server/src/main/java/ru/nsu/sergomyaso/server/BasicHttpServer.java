package ru.nsu.sergomyaso.server;

import com.sun.net.httpserver.HttpExchange;

import java.io.IOException;
import java.io.OutputStream;

public class BasicHttpServer {

    public static final int SERVER_PORT = 8080;

    public static void handleRequest(HttpExchange exchange) throws IOException {
        String response = "Hola, mundo from java";
        exchange.sendResponseHeaders(200, response.getBytes().length);//response code and length
        OutputStream os = exchange.getResponseBody();
        os.write(response.getBytes());
        os.close();
    }
}
