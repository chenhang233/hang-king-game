package service

import (
	"context"
	v1 "hang-king-game/app/user/api/helloworld/v1"
	"hang-king-game/app/user/internal/biz"
)

type GreeterService struct {
	v1.UnimplementedGreeterServer
	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {

	return &GreeterService{
		uc: uc,
	}
}

func (s *GreeterService) SayHello(ctx context.Context, req *v1.HelloRequest) (*v1.HelloReply, error) {
	greeter, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Id: int(req.Id)})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: greeter.Hello}, nil
}
