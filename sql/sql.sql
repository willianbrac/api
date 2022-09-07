USE devbookdb;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    name varchar(50) not null,
    email varchar(50) not null unique,
    password varchar(180) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE books(
    id int auto_increment,
    title varchar(50) not null unique,
    category varchar(60) not null unique,
    synopsis varchar(250) not null,
    author_id int not null,
    createdAt timestamp default current_timestamp(),
  
  	PRIMARY KEY (id),
    CONSTRAINT FK_AuthorBook FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=INNODB;