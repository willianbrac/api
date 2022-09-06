
CREATE DATABASE IF NOT EXISTS devbook;

USE devbookdb;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(10) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;