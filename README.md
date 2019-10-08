# MyAuthd 

[中文](docs/readme.cn.md)

This program is used to authenticate user in local mysql server via remote grpc request . It uses grpc and protobuf to format and transport data

## Installation

This utility build binary and make a deb file for convience

### 1. Install protoc 

Go to https://github.com/protocolbuffers/protobuf to download protoc

### 2. Install go protoc package

```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

### 3. Using build.sh to build 

```
./build.sh
```

## Usage

### 1. Make a config file within the running user home dir

~/.myauthd/config.toml

```toml
  DbAddress =  "127.0.0.1:3306"
  DbName  = "netschool" 
  TableName  = "user" 
  DbUser  = "test"
  DbPass  = "test"
  LoginFieldName = "login"
  PassFieldName  = "hashed_password"
  RealNameFieldName = "realname"
  RolesFieldName  = "roles"
  CryptMethod = "" 
  ToUTF8 = false
```

### 2. Modify systemd file and Start it

Open File `/etc/systemd/system/myauthd.conf`
Modify User and Group 

```
sudo vi /etc/systemd/system/myauthd
sudo systemctl enable myauthd
sudo systemctl start myauthd
```

### 3. Test it

- Insert database some record
- Grant select right to some user (like test)
- Test it 
- If using tsl , add -h servername 

```
dist/client -h servername -u test -p test
```

It will println user realname and roles

### 4. Http restful test

```
curl -X POST -k http://localhost:8081/v1/auth -d '{"Login":"xxx","Password":"xxx"}'
```