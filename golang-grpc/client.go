package main

import (
	"context"
	"math/rand"
	"time"
	// "fmt"
	// "net"

	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"

	"snippet/common"
	"snippet/grpc/proto"
)

func main() {
	//common.Logger.Debug("hello")
	var err error
	var conn *grpc.ClientConn
	conn, err = grpc.Dial("127.0.0.1:15001", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()),
	)
	common.Must(err)

	var client = proto.NewOrderServiceClient(conn)
	var req = proto.UserOrderFindReq{
		UserId: 0,
	}

	// var user *proto.User
	for {
		req.UserId = rand.Int63n(10000)
		_, err = client.FindUserOrder(context.Background(), &req)
		if err != nil {
			common.Logger.Error(err)
		}
		time.Sleep(time.Second)
	}

	// common.Logger.Debug(common.Pretty(user.Id))
	// common.Logger.Debug(common.Pretty(user.Name))
	// common.Logger.Debug(common.Pretty(user.Age))

}
