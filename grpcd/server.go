package grpcd

import (
	"context"
	"errors"
	"fmt"

	"github.com/lajunta/myauthd/utils"
)

var (
	c = utils.CFG
)

// Server used in main package
type Server struct{}

// Auth for grpc client
func (s *Server) Auth(ctx context.Context, in *AuthRequest) (*AuthReply, error) {

	u := AuthReply{}
	login := utils.Filter(in.Login)
	//passd := utils.Crypt(in.Password)
	passd := utils.Filter(in.Password)

	qstr := fmt.Sprintf("select %s,%s from %s where %s='%s' and %s='%s'", c.RealNameFieldName, c.RolesFieldName, c.TableName, c.LoginFieldName, login, c.PassFieldName, passd)

	row := utils.DB.QueryRow(qstr)
	err := row.Scan(&u.RealName, &u.Tags)

	if err != nil {
		u.Logined = false
		return &u, errors.New("Auth Failed")
	}
	u.Logined = true
	return &u, nil
}
