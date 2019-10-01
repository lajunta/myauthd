# Mysql Auth

This program is used to authenticate user from a remote grpc request in local mysql server .

## make a config file

~/.mysql_auth/config.toml

```
DbAddress =  "127.0.0.1:4000"
DbName  = "netschool" 
TableName  = "user" 
DbUser  = "test"
DbPass  = "test"
LoginFieldName = "login"
PassFieldName  = "password"
RealNameFieldName = "realname"
RolesFieldName  = "roles"
```

## Sql Usage

### create table

```
create table user(id int not null primary key auto_increment,
  login varchar(20), password(30),realname varchar(10),roles varchar(255));
```

### add more columns

```
ALTER TABLE table_name ADD column_name datatype;
```

### create a user

```
create user "abc" identified by "abc";
```

### grant a user 

```
grant select on db.table to 'test';
```

### update value

```
UPDATE 表名称 SET 列名称 = 新值 WHERE 列名称 = 某值
```