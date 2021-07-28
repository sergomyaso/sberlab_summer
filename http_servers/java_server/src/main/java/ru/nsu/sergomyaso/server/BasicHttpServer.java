package ru.nsu.sergomyaso.server;

import com.sun.net.httpserver.HttpExchange;


import java.io.IOException;
import java.io.OutputStream;
import java.lang.Math;


public class BasicHttpServer {

    public static final int SERVER_PORT = 8080;
    public static final int COUNT_OPERATIONS = 10000;

    public static void handleRequest(HttpExchange exchange) throws IOException {
        double sinResult = 0;
        double arg = 0.2;

        for (int i = 0; i < COUNT_OPERATIONS; i++) {
            for (int j = 0; j < COUNT_OPERATIONS; j++) {
                sinResult += Math.sin(arg);
                sinResult += Math.sin(arg);
                sinResult += Math.sin(arg);
                sinResult += Math.sin(arg);
                sinResult += Math.sin(arg);
                sinResult += Math.sin(arg);
                sinResult = Math.sin(sinResult);
            }
        }
        String response = "Hola, mundo from java " + Double.toString(sinResult);
        exchange.sendResponseHeaders(200, response.getBytes().length);//response code and length
        OutputStream os = exchange.getResponseBody();
        os.write(response.getBytes());
        os.close();
    }
}
