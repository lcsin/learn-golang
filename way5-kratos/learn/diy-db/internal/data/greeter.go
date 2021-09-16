package data

import (
	"context"
	"diy-db/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	ID        int
	Nickname  string
	Password  string
	Telephone string
	Email     string
	Birth     string
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	var u []*User
	r.data.db.Select(&u, "select * from tb_user")

	for _, user := range u {
		r.data.log.Infof("userId:%d, username:%s", user.ID, user.Nickname)
	}
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}
