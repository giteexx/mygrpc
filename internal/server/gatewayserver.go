package server

import (
	"context"
	"fmt"
	"time"

	"errors"

	"github.com/giteexx/mygrpc/api"
	"github.com/giteexx/mygrpc/internal/svc"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type GatewayServer struct {
	svcCtx *svc.ServiceContext
	api.UnimplementedGatewayServer
}

func NewGatewayServer(svcCtx *svc.ServiceContext) *GatewayServer {
	return &GatewayServer{
		svcCtx: svcCtx,
	}
}

func (s *GatewayServer) GetBalance(ctx context.Context, in *api.WayRequest) (*api.BalanceReply, error) {
	//url:= fmt.
	return &api.BalanceReply{}, gerror.New("暂时未提供此服务")
}

func (s *GatewayServer) UsePhone(ctx context.Context, in *api.WayRequest) (*api.PhoneReply, error) {
	return &api.PhoneReply{}, gerror.New("暂时未提供此服务")
}

func (s *GatewayServer) GetPhone(ctx context.Context, in *api.WayRequest) (res *api.PhoneReply, err error) {
	token, err := s.svcCtx.Login(ctx, in.ChannelInfo.ApiUrl, in.ChannelInfo.Username, in.ChannelInfo.Password)
	if err != nil {
		return
	}
	url := fmt.Sprintf("%s/api/phone/getPhone?productId=%s&username=%s&token=%s", in.ChannelInfo.ApiUrl, in.ChannelInfo.EnVal, in.ChannelInfo.Username, token)
	body := g.Client().Timeout(10*time.Second).GetContent(ctx, url)
	if len(body) == 0 || body == "" {
		return nil, gerror.New("接口访问返回空")
	}
	if j, err := gjson.DecodeToJson(body); err != nil {
		return nil, err
	} else {
		if j.Get("code").Int() == 200 {
			return &api.PhoneReply{
				Phone: j.Get("result.phones").String(),
			}, nil
		} else {
			s.svcCtx.TokenErrWarp(ctx, in.ChannelInfo.Username, body)
			return nil, gerror.New(j.Get("msg").String())
		}
	}
}

func (s *GatewayServer) GetCode(ctx context.Context, in *api.WayRequest) (*api.CodeReply, error) {
	token, err := s.svcCtx.Login(ctx, in.ChannelInfo.ApiUrl, in.ChannelInfo.Username, in.ChannelInfo.Password)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	url := fmt.Sprintf("%s/api/phone/getCode?productId=%s&phone=%s&username=%s&token=%s", in.ChannelInfo.ApiUrl, in.ChannelInfo.EnVal, in.OrderField.Phone, in.ChannelInfo.Username, token)
	body := g.Client().Timeout(10*time.Second).GetContent(ctx, url)
	if len(body) == 0 || body == "" {
		return nil, errors.New("接口返回为空")
	}
	if j, err := gjson.DecodeToJson(body); err != nil {
		return nil, err
	} else {
		if j.Get("code").Int() == 200 {
			if j.Get("result.status").Int() == 1 {
				codeIn := in
				codeIn.ApiField.Codefeed = 1
				go s.CodeFeed(gctx.New(), codeIn)
				return &api.CodeReply{
					Code: j.Get("result.code").String(),
				}, nil
			} else {
				return nil, gerror.New("暂时未获取到验证码")
			}

		} else {
			s.svcCtx.TokenErrWarp(ctx, in.ChannelInfo.Username, body)
			return nil, gerror.New(j.Get("msg").String())
		}
	}
}

func (s *GatewayServer) CodeFeed(ctx context.Context, in *api.WayRequest) (*api.EmptyReply, error) {
	token, err := s.svcCtx.Login(ctx, in.ChannelInfo.ApiUrl, in.ChannelInfo.Username, in.ChannelInfo.Password)
	if err != nil {
		return nil, nil
	}
	url := fmt.Sprintf("%s/api/phone/reportResult?productId=%s&username=%s&token=%s&result=%d&phone=%s", in.ChannelInfo.ApiUrl, in.ChannelInfo.EnVal, in.ChannelInfo.Username, token, in.ApiField.Codefeed, in.OrderField.Phone)
	body := g.Client().Timeout(10*time.Second).GetContent(ctx, url)
	if len(body) == 0 || body == "" {
		return nil, gerror.New("接口访问返回空")
	}
	if j, err := gjson.DecodeToJson(body); err != nil {
		return nil, err
	} else {
		if j.Get("code").Int() == 200 {
			return nil, nil
		} else {
			s.svcCtx.TokenErrWarp(ctx, in.ChannelInfo.Username, body)
			return nil, gerror.New(j.Get("msg").String())
		}
	}
}

func (s *GatewayServer) PhoneFeed(ctx context.Context, in *api.WayRequest) (*api.EmptyReply, error) {
	return &api.EmptyReply{}, nil
}

func (s *GatewayServer) Rrelease(ctx context.Context, in *api.WayRequest) (*api.EmptyReply, error) {
	//newin := in
	//newin.ApiField.Codefeed = 2
	return &api.EmptyReply{}, nil
}
