package com.jayden.jerrymouse.connector;

import com.sun.net.httpserver.Headers;
import com.sun.net.httpserver.HttpExchange;

import java.io.IOException;
import java.io.OutputStream;
import java.net.URI;

public class HttpExchangeAdapter implements HttpExchangeRequest, HttpExchangeResponse{
    final HttpExchange exchange;
    public HttpExchangeAdapter(HttpExchange exchange) {
        this.exchange = exchange;
    }

    @Override
    public String getRequestMethod() {
        return exchange.getRequestMethod();
    }

    public URI getRequestURI() {
        return exchange.getRequestURI();
    }

    @Override
    public Headers getResponseHeaders() {
        return exchange.getResponseHeaders();
    }

    public OutputStream getResponseBody() {
        return exchange.getResponseBody();
    }

    public void sendResponseHeaders(int rCode, long responseLength) throws IOException {
        this.exchange.sendResponseHeaders(rCode, responseLength);
    }
}
