package handlers

import (
	"context"

	pb "mithril-micro/vsrv/pb"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.VideoServiceServer {
	return videoserviceService{}
}

type videoserviceService struct{}

func (s videoserviceService) CreateVideo(ctx context.Context, in *pb.CreateVideoReq) (*pb.Video, error) {
	var resp pb.Video
	return &resp, nil
}

func (s videoserviceService) GetVideo(ctx context.Context, in *pb.GetVideoReq) (*pb.Video, error) {
	var resp pb.Video
	return &resp, nil
}

func (s videoserviceService) GetVideoList(ctx context.Context, in *pb.GetVideoListReq) (*pb.GetVideoListRes, error) {
	var resp pb.GetVideoListRes
	return &resp, nil
}

func (s videoserviceService) UpdateVideo(ctx context.Context, in *pb.UpdateVideoReq) (*pb.UpdateVideoRes, error) {
	var resp pb.UpdateVideoRes
	return &resp, nil
}

func (s videoserviceService) DeleteVideo(ctx context.Context, in *pb.DeleteVideoReq) (*pb.DeleteVideoRes, error) {
	var resp pb.DeleteVideoRes
	return &resp, nil
}
