package com.jayden.HelloSprintBoot.service;

public class MyRunnable implements Runnable {
    @Override
    public void run() {
        System.out.println("start new runnable thread");
    }
}
