package logic

import (
	"context"

	"github.com/giteexx/mygrpc/api"
	"github.com/giteexx/mygrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CodeFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCodeFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CodeFeedLogic {
	return &CodeFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CodeFeedLogic) CodeFeed(in *api.WayRequest) (*api.EmptyReply, error) {
	// todo: add your logic here and delete this line

	return &api.EmptyReply{}, nil
}
