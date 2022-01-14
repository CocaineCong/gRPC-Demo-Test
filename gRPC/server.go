package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "micro/proto"
	"net"
)

type UserInfoService struct {

}

var u = UserInfoService{}

// 实现服务端需要实现的端口
func (service *UserInfoService)GetUserInfo(ctx context.Context,req *pb.UserRequest)(resp *pb.UserResponse,err error) {
	name := req.Name
	// 在数据库查询用户信息
	if name == "zs" {
		resp = &pb.UserResponse{
			Id:   1,
			Name: name,
			Age:  19,
			// 切片字段
			Hobby :[]string{"FanOne","FanOneTwo"},
		}
	}
	err = nil
	return
}

func main() {
	// 1. 监听
	address := "127.0.0.1:8080"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("监听异常 ",err)
		return
	}
	fmt.Println("监听成功")
	// 2. 实例化rpc
	s := grpc.NewServer()
	// 3. 在gRPC上注册微服务
	// 第一个类型是服务，第二个类型是接口的变量
	pb.RegisterUserInfoServiceServer(s,&u)
	// 4. 启动gRPC服务端
	_ = s.Serve(lis)
}

