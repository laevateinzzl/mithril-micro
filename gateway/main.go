package main

import (
	"context"
	"flag"
	"net/http"

	vsrvpb "mithril-micro/vsrv/pb"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// func main() {
// 	// logger
// 	logger, err := zap.NewDevelopment()
// 	if err != nil {
// 		log.Fatalf("cannot create logger: %v", err)
// 	}

// 	c := context.Background()
// 	c, cancel := context.WithCancel(c)
// 	defer cancel()

// 	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
// 		runtime.MIMEWildcard,
// 		&runtime.JSONPb{
// 			MarshalOptions: protojson.MarshalOptions{
// 				UseEnumNumbers: true,
// 				UseProtoNames:  true,
// 			},
// 			UnmarshalOptions: protojson.UnmarshalOptions{
// 				DiscardUnknown: true,
// 			},
// 		},
// 	))

// 	serverConfig := []struct {
// 		name         string
// 		addr         string
// 		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
// 	}{
// 		{
// 			name:         "video",
// 			addr:         "localhost:5040",
// 			registerFunc: vsrvpb.RegisterVideoServiceHandlerFromEndpoint,
// 		},
// 		// {
// 		// 	name:         "user",
// 		// 	addr:         "localhost:8082",
// 		// 	registerFunc: userpb.RegisterUserServiceHandlerFromEndpoint,
// 		// },
// 	}

// 	for _, s := range serverConfig {
// 		err := s.registerFunc(
// 			c, mux, s.addr,
// 			[]grpc.DialOption{grpc.WithInsecure()},
// 		)
// 		if err != nil {
// 			logger.Sugar().Fatalf("cannot register service %s : %v", s.name, err)
// 		}
// 	}
// 	addr := ":8580"
// 	logger.Sugar().Infof("grpc gateway started at %s", addr)
// 	logger.Sugar().Fatal(http.ListenAndServe(addr, mux))
// }

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:5040", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := vsrvpb.RegisterVideoServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
