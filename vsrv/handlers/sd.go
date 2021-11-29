package handlers

import (
	"context"
	"io"
	"time"

	pb "mithril-micro/vsrv/pb"
	"mithril-micro/vsrv/svc"
	c "mithril-micro/vsrv/svc/client/grpc"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

var etcdAddr []string = []string{"127.0.0.1:12379", "127.0.0.1:22379", "127.0.0.1:32379"}

func RegisterToServer() *etcdv3.Registrar {

	options := etcdv3.ClientOptions{}

	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddr, options)

	if err != nil {
		panic(err)
	}

	registar := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   "/vsrv",
		Value: "http://127.0.0.1:5040",
		TTL:   &etcdv3.TTLOption{},
	}, log.NewNopLogger())

	return registar
}

type VideoAgent struct {
	Instancerm *etcdv3.Instancer
	Logger     log.Logger
}

func NewVideoAgent(srv string, logger log.Logger) *VideoAgent {

	options := etcdv3.ClientOptions{
		DialTimeout:   0,
		DialKeepAlive: 0,
	}

	client, err := etcdv3.NewClient(context.Background(), etcdAddr, options)
	if err != nil {
		panic(err)
	}
	instancerm, err := etcdv3.NewInstancer(client, srv, logger)

	if err != nil {
		panic(err)
	}
	return &VideoAgent{
		Instancerm: instancerm,
		Logger:     logger,
	}
}

func (v *VideoAgent) VideoAgentClient() pb.VideoServiceServer {
	var (
		retryMax     = 3
		retryTimeout = 10 * time.Second
	)

	var endpoints svc.Endpoints

	{
		factory := v.factoryFor(svc.MakeCreateVideoEndpoint)
		endpointer := sd.NewEndpointer(v.Instancerm, factory, v.Logger)
		balancer := lb.NewRoundRobin(endpointer)
		retery := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.CreateVideoEndpoint = retery
	}
	{
		factory := v.factoryFor(svc.MakeGetVideoEndpoint)
		endpointer := sd.NewEndpointer(v.Instancerm, factory, v.Logger)
		balancer := lb.NewRoundRobin(endpointer)
		retery := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.CreateVideoEndpoint = retery
	}
	{
		factory := v.factoryFor(svc.MakeGetVideoListEndpoint)
		endpointer := sd.NewEndpointer(v.Instancerm, factory, v.Logger)
		balancer := lb.NewRoundRobin(endpointer)
		retery := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.CreateVideoEndpoint = retery
	}
	{
		factory := v.factoryFor(svc.MakeUpdateVideoEndpoint)
		endpointer := sd.NewEndpointer(v.Instancerm, factory, v.Logger)
		balancer := lb.NewRoundRobin(endpointer)
		retery := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.CreateVideoEndpoint = retery
	}
	{
		factory := v.factoryFor(svc.MakeDeleteVideoEndpoint)
		endpointer := sd.NewEndpointer(v.Instancerm, factory, v.Logger)
		balancer := lb.NewRoundRobin(endpointer)
		retery := lb.Retry(retryMax, retryTimeout, balancer)
		endpoints.CreateVideoEndpoint = retery
	}

	return endpoints
}

func (v *VideoAgent) factoryFor(makeEndpoints func(pb.VideoServiceServer) endpoint.Endpoint) sd.Factory {
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
		return endpoints, conn, err
	}
}
