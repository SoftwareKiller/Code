package com.jayden.service;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data // @Data注解可以自动生成getter/setter/toString/equals/hashCode
@NoArgsConstructor
@AllArgsConstructor
public class Card {
    private String cardId;
    private String cardName;
    private String phone;
    private double money;

    // 预存金额
    public void preStore(double money) {
        this.money += money;
    }

    // 消费金额
    public void consume(double money) {
        this.money -= money;
    }
}
