package svc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/giteexx/mygrpc/internal/config"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
func (s *ServiceContext) Login(ctx context.Context, api, username, password string) (token string, err error) {
	tokenVar, err := gcache.GetOrSetFuncLock(ctx, fmt.Sprintf("one_%s", username), func(ctx context.Context) (value interface{}, err error) {
		url := fmt.Sprintf("%s/api/user/apiLogin?username=%s&password=%s", api, username, password)
		body := g.Client().Timeout(10*time.Second).GetContent(ctx, url)
		if len(body) == 0 || body == "" {
			return nil, gerror.New("获取TOKEN失败" + url)
		}
		if j, err := gjson.DecodeToJson(body); err != nil {

			return nil, err
		} else {
			if j.Get("code").Int() == 200 {
				return j.Get("result.token").String(), nil
			} else {
				return nil, gerror.New(j.Get("msg").String())
			}
		}

	}, time.Second*3600)
	if err != nil {
		return
	}
	token = tokenVar.String()
	return
}

func (s *ServiceContext) TokenErrWarp(ctx context.Context, username, msg string) {
	if strings.Contains(msg, "token") || strings.Contains(msg, "登录") {
		gcache.Remove(ctx, fmt.Sprintf("one_%s", username)) //清除
	}
}
