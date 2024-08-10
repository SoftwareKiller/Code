package com.jayden.jerrymouse.connector;

import com.jayden.jerrymouse.engine.HttpServletRequestImpl;
import com.jayden.jerrymouse.engine.HttpServletResponseImpl;
import jakarta.servlet.ServletException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;

import java.io.IOException;
import java.io.PrintWriter;
import java.net.InetSocketAddress;
import java.net.http.HttpRequest;

public class HttpConnector implements HttpHandler, AutoCloseable {
    final Logger logger = LoggerFactory.getLogger(getClass());

    final HttpServer server;

    public HttpConnector() throws IOException {
        String host = "0.0.0.0";
        int port = 8080;
        server = HttpServer.create(new InetSocketAddress(host, port), 0, "/",this);
        server.start();
        logger.info("HttpConnector started on " + host + ":" + port);
    }

    @Override
    public void close() throws Exception {
        server.stop(3);
    }

    @Override
    public void handle(HttpExchange exchange) throws IOException {
        logger.info("{}: {}?{}", exchange.getRequestMethod(), exchange.getRequestURI().getPath(),
                exchange.getRequestURI().getRawQuery());
        var adapter = new HttpExchangeAdapter(exchange);
        var request = new HttpServletRequestImpl(adapter);
        var response = new HttpServletResponseImpl(adapter);
        // process
        try {
            process(request, response);
        } catch (Exception e) {
            logger.error(e.getMessage(), e);
        }
    }

    void process(HttpServletRequestImpl request, HttpServletResponseImpl response) throws ServletException, IOException {
        String name = request.getParameter("name");
        String html = "<html><body><h1>Hello, " + name + "</h1></body></html>";
        response.setContentType("text/html");
        PrintWriter pw = response.getWriter();
        pw.print(html);
        pw.close();
    }
}
