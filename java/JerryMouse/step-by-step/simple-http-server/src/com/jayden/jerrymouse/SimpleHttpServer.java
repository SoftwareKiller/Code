package com.jayden.jerrymouse;

import com.sun.net.httpserver.Headers;
import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.net.URI;

public class SimpleHttpServer implements HttpHandler,AutoCloseable {

    final Logger logger = LoggerFactory.getLogger(getClass());

    public static void main(String[] args) {
        String host = "0.0.0.0";
        int port = 8080;
        try (SimpleHttpServer server = new SimpleHttpServer(host, port)) {
            for (;;) {
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    break;
                }
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }

    final HttpServer server;
    final String host;
    final int port;

    public SimpleHttpServer(String host, int port) throws IOException {
        this.host = host;
        this.port = port;
        this.server = HttpServer.create(new InetSocketAddress(host, port), 0);
        this.server.createContext("/", this);
        this.server.start();
    }

    @Override
    public void close() throws Exception {
        this.server.stop(3);
    }

    @Override
    public void handle(HttpExchange exchange) throws IOException {
        String method = exchange.getRequestMethod();
        URI uri = exchange.getRequestURI();
        String path = uri.getPath();
        String query = uri.getRawQuery();
        logger.info("{}: {}?{}", method, path, query);
        Headers rspHeaders = exchange.getResponseHeaders();
        rspHeaders.set("Content-Type", "text/html; charset=utf-8");
        rspHeaders.set("Cache-Control", "no-cache");
        exchange.sendResponseHeaders(200, 0);

        String rspBody = "<html><body><h1>Hello, World!</h1></body></html>";
        try (OutputStream out = exchange.getResponseBody()){
            out.write(rspBody.getBytes());
        }
    }
}