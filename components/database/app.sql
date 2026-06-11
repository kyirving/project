CREATE TABLE `user_index` (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
    user_id bigint(20) unsigned NOT NULL COMMENT '用户ID',
    index_value varchar(128) NOT NULL DEFAULT '' COMMENT '索引值',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id` (`user_id`),
)COMMENT '用户索引表';


CREATE TABLE `user_0` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

CREATE TABLE `user_1` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';


CREATE TABLE `user_2` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';


CREATE TABLE `user_3` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

CREATE TABLE `user_4` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

CREATE TABLE `user_5` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';


CREATE TABLE `user_6` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

CREATE TABLE `user_7` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

CREATE TABLE `user_8` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';


CREATE TABLE `user_9` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';


CREATE TABLE `user_10` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';


CREATE TABLE `user_11` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

CREATE TABLE `user_12` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

CREATE TABLE `user_13` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';

CREATE TABLE `user_14` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';


CREATE TABLE `user_15` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键索引',
  `user_id` bigint(20) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `mobile` varchar(128) DEFAULT '' COMMENT '手机号',
  `nikename` varchar(128) DEFAULT '' COMMENT '用户别名',
  `register_type` tinyint(1) unsigned DEFAULT '1' COMMENT '注册类型 1:账号密码 2:手机号',
  `balance` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '账户余额',
  `is_deleted` tinyint(1) unsigned DEFAULT '0' COMMENT '是否删除 1:删除 0:未删除',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),   
  UNIQUE KEY `uk_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户信息表';
