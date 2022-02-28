package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"

	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"snippet/common"
	"snippet/grpc/proto"
)

type OrderServiceImpl struct {
	us proto.UserServiceClient
}

func NewOrderServiceImpl() *OrderServiceImpl {

	// dial up to user service
	var conn *grpc.ClientConn
	var err error
	conn, err = grpc.Dial("127.0.0.1:15000", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()),
	)
	common.Must(err)

	var client = proto.NewUserServiceClient(conn)
	return &OrderServiceImpl{us: client}
}

func (this *OrderServiceImpl) FindOrder(ctx context.Context, req *proto.OrderFindReq) (
	rv *proto.Order, err error) {
	common.Logger.Debug("id: ", req.Id)
	if req.Id > int64(1000) {
		err = status.Error(codes.NotFound, "order not found")
	} else {
		rv = &proto.Order{
			Id:      req.Id,
			Product: "prod name",
			UserId:  233,
		}
	}
	return
}

func (this *OrderServiceImpl) FindUserOrder(
	ctx context.Context, req *proto.UserOrderFindReq) (rv *proto.Orders, err error) {
	common.Logger.Debug("user_id: ", req.UserId)

	// var user *proto.User
	var ureq = proto.UserFindReq{
		Id: req.UserId,
	}
	var user *proto.User
	user, err = this.us.FindUser(ctx, &ureq)
	if err != nil {
		common.Logger.Error(err)
		var rs = status.Convert(err)
		err = status.Error(rs.Code(), rs.Message())
		return
	}

	var oid = rand.Int63n(3000000)
	var orders = []*proto.Order{}
	for i := 0; i < rand.Intn(10); i++ {
		oid++
		orders = append(orders, &proto.Order{
			Id:      oid,
			Product: "prod name",
			UserId:  user.Id,
		})
		rv = &proto.Orders{
			Orders: orders,
		}
	}
	return
}

func main() {
	//common.Logger.Debug("hello")
	var err error

	var server = grpc.NewServer(grpc.UnaryInterceptor(
		apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery())))

	proto.RegisterOrderServiceServer(server, NewOrderServiceImpl())
	var port = 15001
	var l net.Listener
	l, err = net.Listen("tcp", fmt.Sprintf(":%d", port))
	common.Must(err)
	common.Logger.Debug("server listen at :", port)
	err = server.Serve(l)
	common.Must(err)
}
