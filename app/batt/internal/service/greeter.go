package service

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2/component"
	"hang-king-game/app/batt/internal/biz"
	v1 "hang-king-game/app/user/api/helloworld/v1"
)

type GreeterService struct {
	component.Base
	v1.UnimplementedGreeterServer
	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {

	return &GreeterService{
		uc: uc,
	}
}

func (s *GreeterService) AfterInit() {
	fmt.Println("GreeterUsecase AfterInit")
}

func (s *GreeterService) SayHello(ctx context.Context, req *v1.HelloRequest) (*v1.HelloReply, error) {
	greeter, err := s.uc.Test1Greeter(ctx, &biz.Greeter{Id: int(req.Id)})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: greeter.Hello}, nil
}
