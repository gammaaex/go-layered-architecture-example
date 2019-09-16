CREATE SCHEMA IF NOT EXISTS `example` DEFAULT CHARACTER SET utf8mb4;
USE `example`

CREATE TABLE IF NOT EXISTS `users` (
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
	name varchar(191) NOT NULL COMMENT 'name',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
	PRIMARY KEY (id)
) COMMENT = 'ユーザー' ENGINE = InnoDB;
