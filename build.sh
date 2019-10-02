#!/bin/bash
echo "Make Grpcd Service File"
protoc -I=grpcd --go_out=plugins=grpc:grpcd grpcd/grpcd.proto 

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o "dist/server" 
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o "dist/client" client/client.go

upx "dist/server"
upx "dist/client"
cd dist

chmod 0755 server
mkdir -p myauthd/usr/local/bin/
cp server myauthd/usr/local/bin/myauthd
dpkg-deb --build myauthd
