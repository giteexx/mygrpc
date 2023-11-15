package logic

import (
	"context"

	"github.com/giteexx/mygrpc/api"
	"github.com/giteexx/mygrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBalanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBalanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBalanceLogic {
	return &GetBalanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBalanceLogic) GetBalance(in *api.WayRequest) (*api.BalanceReply, error) {
	// todo: add your logic here and delete this line

	return &api.BalanceReply{}, nil
}
