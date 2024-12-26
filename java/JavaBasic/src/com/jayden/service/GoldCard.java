package com.jayden.service;

// 金卡重写消费方法
public class GoldCard extends Card{
    public GoldCard(String cardId, String cardName, double money, String phone) {
        super(cardId, cardName, phone, money);
    }

    @Override
    public void consume(double money) {
        System.out.println("您当前消费：" + money);
        double discount = money * 0.8;
        System.out.println("您当前消费折扣后：" + discount);
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
}
