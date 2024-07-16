package com.jayden.HelloSprintBoot.service;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.List;


@Controller
public class Hello {
//    @PostConstruct
//    public void hello(String name) {
//        System.out.println(name);
//    }

    @GetMapping("hello")
    @ResponseBody
    public String hello(@RequestParam(value = "name", defaultValue = "World") String name) {
        return String.format("Hello %s!", name);
    }

    @Value("${name}")
    private String name;

    @Value("${car.name}")
    private String carName;

    @Value("${car.price}")
    private String carPrice;

    @GetMapping("name")
    @ResponseBody
    public String cName() {
        return name;
    }

    @GetMapping("car")
    @ResponseBody
    public String CarInfo() {
        return "name:" + carName + " price:" + carPrice;
    }


    @Override
    public String toString() {
        return "Hello";
    }

//    public void ListPrint() {
//        List<Integer> list = List.of(3, 5 , 7);
//        for (int i = 0; i < list.size(); i++) {
//            int item = list.get(i);
//            System.out.println(item);
//        }
//    }
}
