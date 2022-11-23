package service

import (
	"context"

	"starfission_go_api/global"
	"starfission_go_api/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	//svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	svc.dao = dao.New(global.DBEngine)
	//svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}

func (svc *Service) Ctx() context.Context {
	return svc.ctx
}
