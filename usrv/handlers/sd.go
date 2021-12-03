package handlers

import (
	"context"
	"io"
	"os"
	"time"

	pb "mithril-micro/usrv/pb"
	"mithril-micro/usrv/svc"
	c "mithril-micro/usrv/svc/client/grpc"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

var etcdAddr []string = []string{"127.0.0.1:12379", "127.0.0.1:22379", "127.0.0.1:32379"}

var (
	prefix   = "/usrvpb"
	instance = "127.0.0.1" + os.Getenv("GRPC_ADDR")
)

func RegisterToServer() *etcdv3.Registrar {

	options := etcdv3.ClientOptions{}

	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddr, options)

	if err != nil {
		panic(err)
	}

	registar := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   prefix + os.Getenv("GRPC_ADDR"),
		Value: instance,
		TTL:   &etcdv3.TTLOption{},
	}, log.NewNopLogger())

	return registar
}

type UserAgent struct {
	Instancerm *etcdv3.Instancer
	Logger     log.Logger
}

func NewUserAgent() *UserAgent {

	options := etcdv3.ClientOptions{
		DialTimeout:   0,
		DialKeepAlive: 0,
	}

	logger := log.NewNopLogger()

	client, err := etcdv3.NewClient(context.Background(), etcdAddr, options)
	if err != nil {
		panic(err)
	}
	instancerm, err := etcdv3.NewInstancer(client, prefix, logger)

	if err != nil {
		panic(err)
	}
	return &UserAgent{
		Instancerm: instancerm,
		Logger:     logger,
	}
}

func (v *UserAgent) UserAgentClient() pb.UserServiceServer {
	var (
		retryMax     = 3
		retryTimeout = 10 * time.Second
	)

	var endpoints svc.Endpoints

	{
		factory := v.factoryFor(svc.MakeCreateUserEndpoint)
		endpointer := sd.NewEndpointer(v.Instancerm, factory, v.Logger)
		balancer := lb.NewRoundRobin(endpointer)
		retery := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.CreateUserEndpoint = retery
	}
	{
		factory := v.factoryFor(svc.MakeGetUserEndpoint)
		endpointer := sd.NewEndpointer(v.Instancerm, factory, v.Logger)
		balancer := lb.NewRoundRobin(endpointer)
		retery := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.GetUserEndpoint = retery
	}
	{
		factory := v.factoryFor(svc.MakeLoginEndpoint)
		endpointer := sd.NewEndpointer(v.Instancerm, factory, v.Logger)
		balancer := lb.NewRoundRobin(endpointer)
		retery := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.LoginEndpoint = retery
	}
	{
		factory := v.factoryFor(svc.MakeUpdateUserEndpoint)
		endpointer := sd.NewEndpointer(v.Instancerm, factory, v.Logger)
		balancer := lb.NewRoundRobin(endpointer)
		retery := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.UpdateUserEndpoint = retery
	}

	return endpoints
}

func (v *UserAgent) factoryFor(makeEndpoints func(pb.UserServiceServer) endpoint.Endpoint) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		conn, err := grpc.Dial(instance, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		srv, err := c.New(conn)
		if err != nil {
			panic(err)
		}
		endpoints := makeEndpoints(srv)
		return endpoints, conn, nil
	}
}
