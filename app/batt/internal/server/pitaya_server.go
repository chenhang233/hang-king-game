package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/cluster"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/groups"
	"github.com/topfreegames/pitaya/v2/route"
	v1 "hang-king-game/app/batt/api/auth/v1"
	"hang-king-game/app/batt/internal/conf"
	"hang-king-game/app/batt/internal/model"
	"hang-king-game/app/batt/internal/service"
	"strings"
	"time"
)

type RegisterComponent struct {
	GroupName string
	Component component.Component
	Options   []component.Option
}

func WsComponents(gService *service.GreeterService,
) []*RegisterComponent {
	var registerComponent []*RegisterComponent
	registerComponent = append(registerComponent, &RegisterComponent{
		GroupName: model.GameGroupName,
		Component: gService,
		Options:   []component.Option{component.WithName(model.GameGroupName), component.WithNameFunc(strings.ToLower)},
	})
	return registerComponent
}

func NewServer(c *conf.Server_Websocket, logger log.Logger) pitaya.Pitaya {
	appConf := config.NewDefaultPitayaConfig()
	appConf.Buffer.Handler.LocalProcess = int(c.GetLocalProcess())
	appConf.Heartbeat.Interval = c.GetInterval().AsDuration()
	appConf.Buffer.Agent.Messages = int(c.GetMessages())
	appConf.Handler.Messages.Compression = c.GetCompression()
	builder := pitaya.NewDefaultBuilder(true, c.ServerType, pitaya.Standalone, map[string]string{}, *appConf)
	builder.AddAcceptor(acceptor.NewWSAcceptor(c.GetAddr()))
	d := config.MemoryGroupConfig{TickDuration: 30 * time.Second}
	builder.Groups = groups.NewMemoryGroupService(d)
	middlewareServer := NewWebsocketMiddleware()
	builder.HandlerHooks.BeforeHandler.PushBack(middlewareServer.RequestBefore)
	builder.HandlerHooks.AfterHandler.PushBack(middlewareServer.RequestAfter)
	app := builder.Build()
	middlewareServer.SetAPPInstance(app)
	err := app.AddRoute("test", func(ctx context.Context, route *route.Route, payload []byte, servers map[string]*cluster.Server) (*cluster.Server, error) {
		fmt.Println(route, string(payload), servers)
		return nil, nil
	})
	if err != nil {
		panic(err)
	}
	return app
}

type WebsocketMiddleware struct {
	component.Base
	app pitaya.Pitaya
}

func NewWebsocketMiddleware() *WebsocketMiddleware {
	return &WebsocketMiddleware{}
}

// RequestBefore 请求之前处理
func (w *WebsocketMiddleware) RequestBefore(ctx context.Context, in any) (context.Context, interface{}, error) {
	fmt.Println("--------------RequestBefore---------------------START")
	if w.app == nil {
		return ctx, in, nil
	}

	session := w.app.GetSessionFromCtx(ctx)
	// 鉴权方式：
	if !session.HasKey(model.SessionUserIdKey) {
		if val, ok := in.(*v1.AuthLoginRequest); ok {
			// 类型断言成功，可以继续处理
			if (*val).GetToken() == "" {
				return nil, nil, pitaya.Error(errors.New("UNKNOWN_ERROR"), "CH-000")
			}
		} else {
			// 类型断言失败，处理错误
			_ = session.Kick(ctx)
			return nil, nil, pitaya.Error(errors.New("NO_PERMISSION"), "CH-000")
		}

	}
	_ = session.Set(model.SessionIpTokenKey, session.RemoteAddr().String())
	_ = session.Set("IE:EXEC:START", time.Now())
	return ctx, in, nil
}

func (w *WebsocketMiddleware) RequestAfter(ctx context.Context, resp interface{}, err error) (interface{}, error) {
	session := w.app.GetSessionFromCtx(ctx)
	elapsed := time.Since(session.Get(model.SessionExecStart).(time.Time))
	fmt.Println("Program run time : ", elapsed)
	return resp, err
}

// SetAPPInstance 设置app实例
func (w *WebsocketMiddleware) SetAPPInstance(app pitaya.Pitaya) {
	w.app = app
}
