package logic

import (
	"context"

	"github.com/giteexx/mygrpc/internal/svc"
	"github.com/giteexx/mygrpc/pb/api"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPhoneLogic {
	return &GetPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPhoneLogic) GetPhone(in *api.WayRequest) (*api.PhoneReply, error) {
	// todo: add your logic here and delete this line

	return &api.PhoneReply{}, nil
}
