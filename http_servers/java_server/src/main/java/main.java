import com.sun.net.httpserver.HttpContext;
import com.sun.net.httpserver.HttpServer;
import ru.nsu.sergomyaso.server.BasicHttpServer;

import java.io.IOException;
import java.net.InetSocketAddress;

public class main {
    public static void main(String[] args) throws IOException {
        HttpServer server = HttpServer.create(new InetSocketAddress(BasicHttpServer.SERVER_PORT), 0);
        HttpContext context = server.createContext("/");
        context.setHandler(BasicHttpServer::handleRequest);
        server.start();
        System.out.println("Server is start on "+ BasicHttpServer.SERVER_PORT + " port");
    }
}
