package main

import (
	"fmt"
	"net"

	"github.com/TykTechnologies/tyk/coprocess"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var d coprocess.DispatcherServer

const (
	port = 5000
)

type dispatcherImpl struct{}

func (d *dispatcherImpl) Dispatch(ctx context.Context, obj *coprocess.Object) (*coprocess.Object, error) {
	return nil, nil
}

func (d *dispatcherImpl) DispatchEvent(ctx context.Context, event *coprocess.Event) (*coprocess.EventReply, error) {
	return nil, nil
}

func main() {
	fmt.Println("sample-plugin starts.")

	d := new(dispatcherImpl)

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		panic(err)
	}

	coprocess.RegisterDispatcherServer(grpcServer, d)
	grpcServer.Serve(lis)
}
