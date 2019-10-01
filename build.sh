#!/bin/bash
echo "Make Grpcd Service File"
protoc -I=grpcd --go_out=plugins=grpc:grpcd grpcd/grpcd.proto 

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o "dist/myauthd_linux" 

upx "dist/myauthd_linux"
cd dist

chmod 0755 myauthd_linux
mkdir -p myauthd/usr/local/bin/
cp myauthd_linux myauthd/usr/local/bin/myauthd
dpkg-deb --build myauthd
