package com.jayden.HelloSprintBoot.service;

public class MyThread extends Thread{
    @Override
    public void run() {
        System.out.println("Start new Thread");
    }
}
