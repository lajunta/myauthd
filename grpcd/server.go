package grpcd

import (
	"context"
	"errors"
	fmt "fmt"
	"log"
	"net"
	"os"
	"path"

	"github.com/lajunta/myauthd/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	c            = utils.CFG
	certFilePath string
	certKeyPath  string
)

// Server used in main package
type Server struct{}

func init() {
	//setCertPath()
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

// Auth for grpc client
func (s *Server) Auth(ctx context.Context, in *AuthRequest) (*AuthReply, error) {

	u := AuthReply{Logined: false}
	login := utils.Filter(in.Login)
	if len(login) == 0 {
		return &u, nil
	}
	passd := utils.Filter(in.Password)
	passd = utils.Crypt(passd)
	qstr := fmt.Sprintf("select %s,%s from %s where %s='%s' and %s=password('%s')", c.RealNameFieldName, c.RolesFieldName, c.TableName, c.LoginFieldName, login, c.PassFieldName, passd)
	row := utils.DB.QueryRow(qstr)
	//row := utils.DB.QueryRow("select ?,? from ? where ?='?' and ?=password(?)", c.RealNameFieldName, c.RolesFieldName, c.TableName, c.LoginFieldName, login, c.PassFieldName, passd)
	err := row.Scan(&u.RealName, &u.Tags)

	if err != nil {
		// log.Println(err.Error())
		return &u, errors.New("Auth Failed")
	}
	u.Logined = true
	u.Login = login
	return &u, nil
}

// Serve starting grpcd service
func Serve(host, port string) {

	if Rest {
		go restRun()
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// creds, err := credentials.NewServerTLSFromFile(certFilePath, certKeyPath)

	// if err != nil {
	// 	log.Fatalf("could not load TLS keys: %s", err)
	// }
	// opts := []grpc.ServerOption{grpc.Creds(creds)}

	s := grpc.NewServer()
	RegisterGrpcdServer(s, &Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	fmt.Printf("Server is running on %s:%s \n", host, port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
