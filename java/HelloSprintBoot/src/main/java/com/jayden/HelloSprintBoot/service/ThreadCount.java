package com.jayden.HelloSprintBoot.service;

public class ThreadCount {
    public void ThreadStaticCount() {
        var add = new AddThread();
        var dec = new DecThread();
        add.start();
        dec.start();
        try {
            add.join();
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }
        try {
            dec.join();
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }

        System.out.println(StaticCounter.count);
    }

    public void ThreadCounter() throws InterruptedException {
        var c = new Counter();
        var tAdd = new Thread(() -> {
            for(int i = 0; i < 10000; i++){
                c.add(i);
            }
		});

        var tDec = new Thread(() -> {
            for(int i = 0; i < 10000; i++){
                c.dec(i);
            }
        });

        tAdd.start();
        tDec.start();
        tAdd.join();
        tDec.join();
        System.out.println(c.get());
    }
}

class StaticCounter {
    public static final Object lock = new Object();
    public static int count = 0;
}

class AddThread extends Thread {
    public void run() {
        for (int i = 0;i < 10000; i++) {
            synchronized (StaticCounter.lock) {
                StaticCounter.count++;
            }
        }
    }
}

class DecThread extends Thread {
    public void run() {
        for (int i = 0;i < 10000; i++) {
            synchronized (StaticCounter.lock) {
                StaticCounter.count--;
            }
        }
    }
}
