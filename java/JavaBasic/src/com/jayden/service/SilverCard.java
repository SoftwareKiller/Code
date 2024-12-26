package com.jayden.service;

public class SilverCard extends Card{
    public SilverCard(String cardId, String cardName, double money, String phone) {
        super(cardId, cardName, phone, money);
    }
    public void consume(double money) {
        double discount = money * 0.9;
        if (discount > 0) {
            if (discount > this.getMoney()) {
                System.out.println("余额不足");
            } else {
                this.setMoney(this.getMoney() - discount);
                System.out.println("消费成功，余额为：" + this.getMoney());
            }
        } else {
            System.out.println("消费金额必须大于0");
        }
    }

    public void preStore(double money) {}
}
