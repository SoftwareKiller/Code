package org.jayden.sprintbootmybits;

import org.jayden.sprintbootmybits.domain.App;
import org.jayden.sprintbootmybits.mapper.AppMapper;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import java.util.List;

@SpringBootTest
class SprintBootMybitsApplicationTests {

    @Autowired
    private AppMapper appMapper;

    @Test
    void testFindAll() {
        List<App> apps = appMapper.findAll();
        System.out.println(apps);
    }

}
