package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lajunta/myauthd/grpcd"

	"google.golang.org/grpc"
)

func client() {
	// Set up a connection to the server.

	// creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	// if err != nil {
	// 	log.Fatalf("could not load tls cert: %s", err)
	// }

	//conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	fmt.Println(address)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := grpcd.NewGrpcdClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	authed, err := c.Auth(ctx, &grpcd.AuthRequest{Login: "zxy", Password: "opendoor"})
	if err != nil {
		log.Println("Auth Failed")
	} else {
		log.Println(authed.Login, authed.RealName, authed.Tags)
	}
}
