package com.jayden.HelloSprintBoot;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@RestController
public class HelloSprintBootApplication {

	public static void main(String[] args) {
//		Hello h = new Hello();
//		h.hello("你好啊");
//		h.ListPrint();
		SpringApplication.run(HelloSprintBootApplication.class, args);
//		MyThread t = new MyThread();
//		t.run();
//		Thread t = new Thread(new MyRunnable());
//		t.start();
//
//		Thread lam = new Thread(() -> {
//			System.out.println("Start new lambda thread");
//		});
//		lam.start();
//		ThreadCount tc = new ThreadCount();
//		tc.ThreadStaticCount();
//		ThreadCount tc = new ThreadCount();
//        try {
//            tc.ThreadCounter();
//        } catch (InterruptedException e) {
//            throw new RuntimeException(e);
//        }
    }
}
