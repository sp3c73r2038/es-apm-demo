package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmgrpc"
	"go.elastic.co/apm/module/apmhttp"
	"golang.org/x/net/context/ctxhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"snippet/common"
	"snippet/grpc/proto"
)

var tracingClient = apmhttp.WrapClient(http.DefaultClient)

type UserServiceImpl struct {
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}

func (this *UserServiceImpl) FindUser(ctx context.Context, req *proto.UserFindReq) (
	rv *proto.User, err error) {
	common.Logger.Debug(common.Pretty(req.Id))
	if req.Id > int64(2000) {
		err = status.Error(codes.Internal, "random fatal error")
		common.Logger.Panic("random fatal error")
	} else if req.Id > int64(1000) {
		err = status.Error(codes.NotFound, "user not found")
	} else {
		var resp *http.Response
		resp, err = ctxhttp.Get(ctx, tracingClient, "http://127.0.0.1")
		if err != nil {
			apm.CaptureError(ctx, err).Send()
			err = status.Error(codes.Internal, "http client error")
			return
		}
		io.ReadAll(resp.Body)

		rv = &proto.User{
			Id:   req.Id,
			Name: fmt.Sprintf("Name %d", req.Id),
			Age:  int32(req.Id) - 100,
		}
	}
	return
}

func main() {
	//common.Logger.Debug("hello")
	var err error

	var server = grpc.NewServer(grpc.UnaryInterceptor(
		apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery())))

	proto.RegisterUserServiceServer(server, NewUserServiceImpl())
	var port = 15000
	var l net.Listener
	l, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	common.Must(err)
	common.Logger.Debug("server listen at :15000")
	err = server.Serve(l)
	common.Must(err)
}
