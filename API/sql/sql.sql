CREATE DATABASE IF NOT EXISTS healthfit; 
USE healthfit;

DROP TABLE IF EXISTS usuarios; 

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    peso int not null,
    altura int not null, 
    idade int not null, 
    senha varchar(100) not null
) ENGINE=INNODB;