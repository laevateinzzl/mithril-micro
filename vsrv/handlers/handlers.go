package handlers

import (
	"context"
	"log"
	"time"

	usrv "mithril-micro/usrv/handlers"
	usrvpb "mithril-micro/usrv/pb"
	"mithril-micro/vsrv/dao"
	pb "mithril-micro/vsrv/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.VideoServiceServer {
	r := RegisterToServer()
	r.Register()
	return videoserviceService{}
}

type videoserviceService struct{}

func (s videoserviceService) CreateVideo(ctx context.Context, in *pb.CreateVideoReq) (*pb.Video, error) {
	var resp pb.Video = pb.Video{
		VideoId:   primitive.NewObjectID().Hex(),
		Title:     in.Title,
		Summary:   in.Summary,
		CreatedAt: time.Now().Unix(),
		Poster:    in.Poster,
		Url:       in.Url,
		User:      &pb.User{},
	}

	user := &pb.User{}

	req := usrvpb.GetUserReq{

		UserId: in.UserId,
	}
	i := usrv.NewUserAgent()

	res, err := i.UserAgentClient().GetUser(ctx, &req)

	if err != nil {
		log.Fatal(err)
	}

	user.Avatar = res.Avatar
	user.Nickname = res.Nickname

	resp.User = user

	c := dao.NewCl()

	_, err = c.InsertOne(ctx, resp)
	if err != nil {
		log.Fatal(err)
	}

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

func (s videoserviceService) GetVideo(ctx context.Context, in *pb.GetVideoReq) (*pb.Video, error) {
	var resp pb.Video
	return &resp, nil
}
