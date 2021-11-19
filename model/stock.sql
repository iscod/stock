SET @PREVIOUS_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_
CHECKS = 0;

DROP TABLE IF EXISTS `symbol`;
DROP TABLE IF EXISTS `quote`;
DROP TABLE IF EXISTS `comment`;

CREATE TABLE `comment` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `symbol` varchar(128) DEFAULT NULL,
  `user_id` bigint(20) unsigned DEFAULT '0',
  `title` varchar(256) NOT NULL DEFAULT '',
  `description` text NOT NULL,
  `source` varchar(128) NOT NULL DEFAULT '',
  `text` text NOT NULL,
  `user` text NOT NULL,
  `type` varchar(10) NOT NULL DEFAULT '0',
  `view_count` int(10) unsigned NOT NULL DEFAULT '0',
  `created_at` bigint(20) unsigned NOT NULL,
  `updated_at` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4

CREATE TABLE `quote` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(128) DEFAULT NULL,
  `symbol` varchar(128) NOT NULL DEFAULT '',
  `comment_count` int(10) unsigned NOT NULL DEFAULT '0',
  `comment_count1` int(10) unsigned DEFAULT '0',
  `comment_count2` int(10) unsigned DEFAULT '0',
  `comment_count3` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '官方认证大v评论数量',
  `current` decimal(10,5) unsigned NOT NULL DEFAULT '0.00000' COMMENT '当前价格',
  `open` decimal(10,5) unsigned NOT NULL DEFAULT '0.00000',
  `avg_price` decimal(10,5) unsigned NOT NULL DEFAULT '0.00000' COMMENT '今日均价',
  `low` decimal(10,5) unsigned NOT NULL DEFAULT '0.00000' COMMENT '最低价',
  `high` decimal(10,5) unsigned NOT NULL DEFAULT '0.00000' COMMENT '最高价',
  `detail` text,
  `volume` int(10) unsigned NOT NULL DEFAULT '0',
  `s_volume` int(10) unsigned NOT NULL DEFAULT '0',
  `b_volume` int(10) unsigned NOT NULL DEFAULT '0',
  `m_volume` int(10) unsigned NOT NULL DEFAULT '0',
  `exec_at` date NOT NULL DEFAULT '0000-00-00',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_symbol_created_at` (`symbol`,`exec_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `symbol` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `symbol` varchar(128) NOT NULL DEFAULT '',
  `name` varchar(128) NOT NULL DEFAULT '',
  `status` int(10) unsigned NOT NULL DEFAULT '0',
  `background_color` varchar(128) NOT NULL DEFAULT '',
  `border_color` varchar(128) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4

SET FOREIGN_KEY_CHECKS = @PREVIOUS_FOREIGN_KEY_CHECKS;

SET @PREVIOUS_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_CHECKS = 0;


LOCK TABLES `comment` WRITE;
ALTER TABLE `comment` DISABLE KEYS;
ALTER TABLE `comment` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `quote` WRITE;
ALTER TABLE `quote` DISABLE KEYS;
ALTER TABLE `quote` ENABLE KEYS;
UNLOCK TABLES;


LOCK TABLES `symbol` WRITE;
ALTER TABLE `symbol` DISABLE KEYS;
ALTER TABLE `symbol` ENABLE KEYS;
UNLOCK TABLES;




SET FOREIGN_KEY_CHECKS = @PREVIOUS_FOREIGN_KEY_CHECKS;


