package logic

import (
	"context"

	"github.com/giteexx/mygrpc/api"
	"github.com/giteexx/mygrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RreleaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRreleaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RreleaseLogic {
	return &RreleaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RreleaseLogic) Rrelease(in *api.WayRequest) (*api.EmptyReply, error) {
	// todo: add your logic here and delete this line

	return &api.EmptyReply{}, nil
}
