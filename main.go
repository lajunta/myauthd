package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/lajunta/myauthd/grpcd"
	"github.com/lajunta/myauthd/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var (
	host         string
	port         string
	certFilePath string
	certKeyPath  string
)

func init() {
	flagParse()
	setCertPath()
}

func flagParse() {
	flag.StringVar(&port, "port", "5050", "server running port")
	flag.StringVar(&host, "host", "0.0.0.0", "server host ip ")
	flag.Parse()
}

func setCertPath() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Can't Open User Home Dir")
		os.Exit(1)
	}
	certFilePath = path.Join(home, ".myauthd", "cert", "server.crt")
	certKeyPath = path.Join(home, ".myauthd", "cert", "server.key")
}

// handle program exit event
func handleExit() {
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
}

func main() {

	handleExit()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(certFilePath, certKeyPath)

	if err != nil {
		log.Fatalf("could not load TLS keys: %s", err)
	}
	opts := []grpc.ServerOption{grpc.Creds(creds)}
	s := grpc.NewServer(opts...)
	grpcd.RegisterGrpcdServer(s, &grpcd.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	fmt.Printf("Server is running on %s:%s \n", host, port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
