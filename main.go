package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/lajunta/myauthd/grpcd"
	"github.com/lajunta/myauthd/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	host    string
	port    string
	gc      bool
	daemon  bool
	address string
)

func init() {
	flagParse()
	address = host + ":" + port
}

func flagParse() {
	flag.StringVar(&port, "port", "5005", "server running port")
	flag.StringVar(&host, "host", "127.0.0.1", "server host address")
	flag.BoolVar(&daemon, "d", false, "server daemon mode")
	flag.BoolVar(&gc, "c", false, "grpc client mode(for test only)")

	flag.Parse()
}

func main() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		s := <-sigc
		fmt.Println(s.String())
		err := utils.DB.Close()
		if err != nil {
			log.Fatalln(err.Error())
		} else {
			log.Println("Program Exit Ok and DB Close Wonderful")
			os.Exit(0)
		}
	}()

	// Above code handle program exit and close db connection.

	if gc {
		client()
	} else {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		//creds, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.key")

		if err != nil {
			log.Fatalf("could not load TLS keys: %s", err)
		}
		//opts := []grpc.ServerOption{grpc.Creds(creds)}

		s := grpc.NewServer()
		grpcd.RegisterGrpcdServer(s, &grpcd.Server{})
		// Register reflection service on gRPC server.
		reflection.Register(s)

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}
}
