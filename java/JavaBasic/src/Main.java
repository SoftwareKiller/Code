import com.jayden.service.Card;
import com.jayden.service.GoldCard;

public class Main {
    public static void main(String[] args) {
        // 目标：加油站支付小程序
        // 1. 创建卡片类，以便创建金卡或者银卡对象，封装车主数据
        // 2. 定义一个卡片父类：Card，定义金卡和银卡的共同属性和方法
        // 3. 定义一个金卡类：GoldCard，继承Card，定义自己的属性和方法
        // 4. 定义一个银卡类：SilverCard，继承Card，定义自己的属性和方法
        GoldCard gc = new GoldCard("123456", "张三", 1000, "13800000000");
        pay(gc, 1000);
    }

    public static void pay(Card c, double money) {
        System.out.println("使用" + c.getCardName() + "支付了" + money + "元");
        c.consume(money);
    }
}