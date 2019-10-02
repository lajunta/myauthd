# Secure the communication

From this [url](https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2)

Make crt and key under your config dir

```
openssl genrsa -out cert/server.key 2048
openssl req -new -sha256 -key cert/server.key -out cert/server.csr
openssl x509 -req -sha256 -in cert/server.csr -signkey cert/server.key -out cert/server.crt -days 3650
```

When making csr file , you should write the FQDN or computer name