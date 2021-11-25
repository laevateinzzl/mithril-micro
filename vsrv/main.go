package main

import (
	"log"
	"mithril-micro/shared/server"
	vsrvpb "mithril-micro/vsrv/api/gen"
	"mithril-micro/vsrv/video"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	// RunGRPCServer
	logger.Sugar().Fatal(
		server.RunGRPCServer(&server.GRPCConfig{
			Name:   "video",
			Addr:   ":8581",
			Logger: logger,
			RegisterFunc: func(s *grpc.Server) {
				vsrvpb.RegisterVideoServiceServer(s, &video.VideoService{
					Logger: logger,
				})
			},
		}),
	)
}
