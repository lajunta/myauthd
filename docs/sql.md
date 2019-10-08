# Sql Usage

## create table

```
create table user(id int not null primary key auto_increment,
  login varchar(20), password(30),realname varchar(10),roles varchar(255));
```

## add more columns

```
ALTER TABLE table_name ADD column_name datatype;
```

## create a user

```
create user "abc" identified by "abc";
```

## grant a user 

```
grant select on db.table to 'test';
```

## update value

```
UPDATE 表名称 SET 列名称 = 新值 WHERE 列名称 = 某值
```

## show grants
```
show grants
```

## revoke grants
```
revoke all on db.table from 'test'@'%'
```