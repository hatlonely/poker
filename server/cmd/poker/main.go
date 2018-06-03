package main

import (
	"fmt"
	"net"

	"golang.org/x/net/context"

	echo "github.com/hatlonely/poker/server/api"
	"google.golang.org/grpc"
)

type EchoServerImp struct {
}

func (e *EchoServerImp) Echo(ctx context.Context, req *echo.EchoReq) (*echo.EchoRes, error) {
	fmt.Printf("message from client: %v\n", req.GetMsg())
	res := &echo.EchoRes{
		Msg: req.GetMsg(),
	}
	return res, nil
}

func main() {
	server := grpc.NewServer()
	echo.RegisterEchoServer(server, &EchoServerImp{})
	address, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(address); err != nil {
		panic(err)
	}

	fmt.Println("hello world")
}
