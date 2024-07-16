package com.jayden.HelloSprintBoot;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.redis.core.RedisTemplate;

@SpringBootTest
class HelloSprintBootApplicationTests {
	@Autowired
	private RedisTemplate redisTemplate;

	@Test
	void testSet() {
		redisTemplate.boundValueOps("name").set("luca");
	}

	@Test
	void testGet() {
		Object name = redisTemplate.boundValueOps("name").get();
		System.out.println(name);
	}

}
