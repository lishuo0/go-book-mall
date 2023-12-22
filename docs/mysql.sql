create table mall_user (
    `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户自增ID',
    `nick_name` varchar(64) not null default '' comment '昵称',
    `account` varchar(32) not null default '' comment '账号',
    `password` varchar(64) not null default '' comment '密码',
    `icon` varchar(256) not null default '' comment '头像',
    `gender` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1男；2女',
    `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态：1正常2、冻结、3注销',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_account` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

create table mall_goods (
    `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `name` varchar(100) not null default '' comment '商品名称',
    `description` varchar(255) not null default '' comment '商品描述',
    `tags` varchar(255) not null default '' comment '商品标签',
    `detail` TEXT not null  comment '商品详情',
    `category_id` int NOT NULL default 0 COMMENT '商品所在类别',
    `small_image`  varchar(255) not null default '' comment '缩略图',
    `detail_image` varchar(255) not null default '' comment '商品详情页展示图',
    `price` int NOT NULL default 0 COMMENT '商品默认价格',
    `status` tinyint unsigned NOT NULL DEFAULT 1 COMMENT '1: 有效; 2: 失效',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `mall_goods_sku` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `goods_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '所属的商品id',
  `attribute_ids` varchar(255) NOT NULL DEFAULT '' COMMENT 'sku属性id列表, 逗号分隔',
  `spend_price` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '成本价',
  `price` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '原价',
  `discount_price` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '优惠价',
  `left_store` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '剩余库存',
  `all_store` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '总库存',
  `status` tinyint unsigned NOT NULL DEFAULT 1 COMMENT '1: 有效; 2: 失效',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_goods_id` (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `mall_order_0` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint NOT NULL DEFAULT '0' COMMENT 'id',
  `order_id` varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL DEFAULT '' COMMENT 'id',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '1 2 3 4 11 12',
  `pay_id` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL DEFAULT '' COMMENT 'id',
  `pay_status` tinyint NOT NULL DEFAULT '0' COMMENT '0 1 2 3 4:',
  `pay_type` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '1,2,3,4,5',
  `source` tinyint NOT NULL DEFAULT '0' COMMENT '1 2PC',
  `serial` varchar(64) NOT NULL DEFAULT '',
  `total_amount` int NOT NULL DEFAULT '0' COMMENT ',',
  `goods_num` int unsigned NOT NULL DEFAULT '1' COMMENT '1',
  `pay_time` datetime NOT NULL DEFAULT '1000-10-10 10:00:00',
  `cancel_time` datetime NOT NULL DEFAULT '1000-10-10 10:00:00',
  `order_expand` varchar(512) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_serial` (`serial`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_orderid` (`order_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `mall_order_detail` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '子订单，自增id',
  `order_id` varchar(64) NOT NULL DEFAULT '' COMMENT '主订单id',
  `sku_id` int(11) NOT NULL DEFAULT '0' COMMENT 'sku_id',
  `price` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '价格，单位（分）',
  `num` int(11) unsigned NOT NULL DEFAULT '1' COMMENT '商品购买数量，默认1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `order_expand` varchar(512)  CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '订单信息扩展字段',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态：1正常 2、取消',
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_order_id` (`order_id`),
  KEY `idx_sku_id` (`sku_id`),
  UNIQUE KEY `idx_order_sku_id` (`order_id`,`sku_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单详细表';