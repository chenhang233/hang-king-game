package data

import (
	"context"
	"encoding/json"
	"hang-king-game/app/user/internal/model"

	"hang-king-game/app/user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Test1(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	p := &model.Product{}
	result := r.data.GetDB(model.DBTest).Table(p.TableName()).Select("*").Where("id = ?", g.Id).Find(p)
	if err := result.Error; err != nil {
		r.log.WithContext(ctx).Errorf("Save error err:%v", err)
		return nil, err
	}
	m, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	g.Hello = string(m)
	return g, nil
}
