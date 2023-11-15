package logic

import (
	"context"

	"github.com/giteexx/mygrpc/internal/svc"
	"github.com/giteexx/mygrpc/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type PhoneFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPhoneFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PhoneFeedLogic {
	return &PhoneFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PhoneFeedLogic) PhoneFeed(in *api.WayRequest) (*api.EmptyReply, error) {
	// todo: add your logic here and delete this line

	return &api.EmptyReply{}, nil
}
