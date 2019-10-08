create user if not exists test;
grant select on netschool.user to 'test'@'%' identified by 'test';
create database if not exists netschool;
use netschool;
create table if not exists user(
  id int not null primary key auto_increment,
  login varchar(20) not null unique,
  hashed_password varchar(200) not null,
  realname varchar(30) not null,
  roles varchar(255) not null
);
insert into user(login,hashed_password,realname,roles) values("hello",password("hello"),"zxy","teacher,ds");