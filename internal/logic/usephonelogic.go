package logic

import (
	"context"

	"github.com/giteexx/mygrpc/internal/svc"
	"github.com/giteexx/mygrpc/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsePhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUsePhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsePhoneLogic {
	return &UsePhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UsePhoneLogic) UsePhone(in *api.WayRequest) (*api.PhoneReply, error) {
	// todo: add your logic here and delete this line

	return &api.PhoneReply{}, nil
}
