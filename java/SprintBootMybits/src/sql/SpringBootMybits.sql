CREATE TABLE `t_app` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `username` varchar(64) DEFAULT NULL,
    `app_id` varchar(64) DEFAULT NULL,
    primary key (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

INSERT INTO `t_app` (id, username, app_id) VALUES (1, '流行蝴蝶剑', 1);
