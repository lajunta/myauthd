package main

import (
	"context"
	"flag"
	"log"
	"os"
	"path"
	"time"

	"github.com/lajunta/myauthd/grpcd"
	"google.golang.org/grpc"
)

var (
	host         string
	port         string
	address      string
	testUser     string
	testPasswd   string
	certFilePath string
)

func init() {
	flagParse()
	address = host + ":" + port
	//setCertPath()
}

func flagParse() {
	flag.StringVar(&port, "port", "5050", "server running port")
	flag.StringVar(&host, "host", "127.0.0.1", "server host address")
	flag.StringVar(&testUser, "u", "hello", "a test login name")
	flag.StringVar(&testPasswd, "p", "hello", "a test password")
	flag.Parse()
}

func setCertPath() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Can't Open User Home Dir")
		os.Exit(1)
	}
	certFilePath = path.Join(home, ".myauthd", "cert", "server.crt")
}

func main() {

	// creds, err := credentials.NewClientTLSFromFile(certFilePath, "")
	// if err != nil {
	// 	log.Fatalf("could not load tls cert: %s", err)
	// }

	// conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		log.Println("Dial Server Successfully!")
	}

	defer conn.Close()
	c := grpcd.NewGrpcdClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	authed, err := c.Auth(ctx, &grpcd.AuthRequest{Login: testUser, Password: testPasswd})
	if err != nil {
		log.Println("Auth Failed: ", err.Error())
	} else {
		log.Println(authed.Logined, authed.Login, authed.RealName, authed.Tags)
	}
}
