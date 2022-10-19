package service

import (
	"context"
	"dao"
	otgrom "github.com/eddycjy/opentracing-gorm"
	"global"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(otgrom.WithContext(svc.ctx, global.DBEngine))
	return svc
}
