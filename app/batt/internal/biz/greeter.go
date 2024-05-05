package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
	Id    int
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Test1(context.Context, *Greeter) (*Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GreeterUsecase) Test1Greeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Test1(ctx, g)
}
