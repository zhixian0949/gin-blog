package service

import (
	"context"

	"github.com/zhixian0949/gin-blog/global"
	"github.com/zhixian0949/gin-blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
