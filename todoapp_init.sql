DROP DATABASE IF EXISTS todoapp;
CREATE DATABASE IF NOT EXISTS todoapp;
USE todoapp;

DROP TABLE IF EXISTS todos;
DROP TABLE IF EXISTS users;

CREATE TABLE `users`
(
 `id`                 int NOT NULL AUTO_INCREMENT ,
 `name`               varchar(32) NOT NULL ,
 `password_hash`      varchar(256) NOT NULL ,
 `email`              varchar(255) NOT NULL ,
 `session_key`        varchar(36) ,
 `session_started_at` datetime ,

PRIMARY KEY (`id`)
);

INSERT INTO users (name, password_hash, email) VALUES ("User", "1234", "user@mail.com");

CREATE TABLE `todos`
(
 `id`          int NOT NULL AUTO_INCREMENT ,
 `user_id`     int NOT NULL ,
 `title`       varchar(64) NOT NULL ,
 `description` varchar(196) NULL ,
 `complete`    tinyint NOT NULL ,

PRIMARY KEY (`id`),
KEY `FK_user_id` (`user_id`),
CONSTRAINT `FK_user` FOREIGN KEY `FK_user_id` (`user_id`) REFERENCES `users` (`id`)
);