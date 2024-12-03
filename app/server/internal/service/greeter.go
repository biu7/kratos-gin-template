package service

import (
	"github.com/biu7/gokit/ginutils/response"
	"github.com/gin-gonic/gin"
	"kratos-gin-template/app/shared/request"

	v1 "kratos-gin-template/api/helloworld/v1"
	"kratos-gin-template/app/server/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) SayHello(c *gin.Context) {
	var args v1.HelloRequest
	if err := request.Bind(c, &args); err != nil {
		response.Fail(c, err)
		return
	}

	g, err := s.uc.CreateGreeter(c, &biz.Greeter{Hello: args.Name})
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, &v1.HelloReply{Message: "Hello " + g.Hello})
}
