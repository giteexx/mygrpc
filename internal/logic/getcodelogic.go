package logic

import (
	"context"

	"github.com/giteexx/mygrpc/api"
	"github.com/giteexx/mygrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCodeLogic {
	return &GetCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCodeLogic) GetCode(in *api.WayRequest) (*api.CodeReply, error) {
	// todo: add your logic here and delete this line

	return &api.CodeReply{}, nil
}
