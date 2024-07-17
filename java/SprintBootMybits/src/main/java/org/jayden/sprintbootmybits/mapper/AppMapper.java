package org.jayden.sprintbootmybits.mapper;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import org.jayden.sprintbootmybits.domain.App;
import org.springframework.stereotype.Repository;

import java.util.List;

@Mapper
@Repository
public interface AppMapper {

    @Select("select * from `t_app`")
    public List<App> findAll();
}
