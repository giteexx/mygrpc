package logic

import (
	"context"

	"github.com/giteexx/mygrpc/internal/svc"
	"github.com/giteexx/mygrpc/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReleaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReleaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReleaseLogic {
	return &ReleaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReleaseLogic) Release(in *api.WayRequest) (*api.EmptyReply, error) {
	// todo: add your logic here and delete this line

	return &api.EmptyReply{}, nil
}
