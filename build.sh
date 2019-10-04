#!/bin/bash
echo "Make Grpcd Service File"
protoc -I /usr/local/include -I. -I $GOPATH/src -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 -I grpcd --go_out=plugins=grpc:. grpcd/grpcd.proto 

echo "Generate reverse-proxy gateway file"
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  grpcd/grpcd.proto

go get .

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o "dist/server" 
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o "dist/client" client/client.go

if [[ $1 == "prod" ]];then
  upx "dist/server"
  upx "dist/client"

  cd dist
  chmod 0755 server
  mkdir -p myauthd/usr/local/bin/
  cp server myauthd/usr/local/bin/myauthd
  dpkg-deb --build myauthd
fi
