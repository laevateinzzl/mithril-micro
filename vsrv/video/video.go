package video

import (
	"context"
	vsrpb "mithril-micro/vsrv/api/gen"

	"go.uber.org/zap"
)

type VideoService struct {
	vsrpb.UnimplementedVideoServiceServer
	Logger *zap.Logger
}

func (s *VideoService) CreateVideo(ctx context.Context, req *vsrpb.CreateVideoReq) (*vsrpb.Video, error) {
	return &vsrpb.Video{
		VideoId:   1,
		Title:     "test",
		Summary:   "test",
		CreatedAt: 0,
		Poster:    "",
		Url:       "",
	}, nil
}

func (s *VideoService) GetVideo(ctx context.Context, req *vsrpb.GetVideoReq) (*vsrpb.Video, error) {
	return &vsrpb.Video{
		VideoId:   req.VideoId,
		Title:     "test",
		Summary:   "test",
		CreatedAt: 0,
		Poster:    "",
		Url:       "",
	}, nil
}
