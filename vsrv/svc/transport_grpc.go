// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: fcd9ff140d
// Version Date: 2021-07-14T06:36:40Z

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "mithril-micro/vsrv/pb"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC VideoServiceServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.VideoServiceServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// videoservice

		createvideo: grpctransport.NewServer(
			endpoints.CreateVideoEndpoint,
			DecodeGRPCCreateVideoRequest,
			EncodeGRPCCreateVideoResponse,
			serverOptions...,
		),
		getvideo: grpctransport.NewServer(
			endpoints.GetVideoEndpoint,
			DecodeGRPCGetVideoRequest,
			EncodeGRPCGetVideoResponse,
			serverOptions...,
		),
		getvideolist: grpctransport.NewServer(
			endpoints.GetVideoListEndpoint,
			DecodeGRPCGetVideoListRequest,
			EncodeGRPCGetVideoListResponse,
			serverOptions...,
		),
		updatevideo: grpctransport.NewServer(
			endpoints.UpdateVideoEndpoint,
			DecodeGRPCUpdateVideoRequest,
			EncodeGRPCUpdateVideoResponse,
			serverOptions...,
		),
		deletevideo: grpctransport.NewServer(
			endpoints.DeleteVideoEndpoint,
			DecodeGRPCDeleteVideoRequest,
			EncodeGRPCDeleteVideoResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the VideoServiceServer interface
type grpcServer struct {
	createvideo  grpctransport.Handler
	getvideo     grpctransport.Handler
	getvideolist grpctransport.Handler
	updatevideo  grpctransport.Handler
	deletevideo  grpctransport.Handler
}

// Methods for grpcServer to implement VideoServiceServer interface

func (s *grpcServer) CreateVideo(ctx context.Context, req *pb.CreateVideoReq) (*pb.Video, error) {
	_, rep, err := s.createvideo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Video), nil
}

func (s *grpcServer) GetVideo(ctx context.Context, req *pb.GetVideoReq) (*pb.Video, error) {
	_, rep, err := s.getvideo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.Video), nil
}

func (s *grpcServer) GetVideoList(ctx context.Context, req *pb.GetVideoListReq) (*pb.GetVideoListRes, error) {
	_, rep, err := s.getvideolist.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetVideoListRes), nil
}

func (s *grpcServer) UpdateVideo(ctx context.Context, req *pb.UpdateVideoReq) (*pb.UpdateVideoRes, error) {
	_, rep, err := s.updatevideo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateVideoRes), nil
}

func (s *grpcServer) DeleteVideo(ctx context.Context, req *pb.DeleteVideoReq) (*pb.DeleteVideoRes, error) {
	_, rep, err := s.deletevideo.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteVideoRes), nil
}

// Server Decode

// DecodeGRPCCreateVideoRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC createvideo request to a user-domain createvideo request. Primarily useful in a server.
func DecodeGRPCCreateVideoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CreateVideoReq)
	return req, nil
}

// DecodeGRPCGetVideoRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getvideo request to a user-domain getvideo request. Primarily useful in a server.
func DecodeGRPCGetVideoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetVideoReq)
	return req, nil
}

// DecodeGRPCGetVideoListRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getvideolist request to a user-domain getvideolist request. Primarily useful in a server.
func DecodeGRPCGetVideoListRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetVideoListReq)
	return req, nil
}

// DecodeGRPCUpdateVideoRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC updatevideo request to a user-domain updatevideo request. Primarily useful in a server.
func DecodeGRPCUpdateVideoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UpdateVideoReq)
	return req, nil
}

// DecodeGRPCDeleteVideoRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC deletevideo request to a user-domain deletevideo request. Primarily useful in a server.
func DecodeGRPCDeleteVideoRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.DeleteVideoReq)
	return req, nil
}

// Server Encode

// EncodeGRPCCreateVideoResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain createvideo response to a gRPC createvideo reply. Primarily useful in a server.
func EncodeGRPCCreateVideoResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.Video)
	return resp, nil
}

// EncodeGRPCGetVideoResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getvideo response to a gRPC getvideo reply. Primarily useful in a server.
func EncodeGRPCGetVideoResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.Video)
	return resp, nil
}

// EncodeGRPCGetVideoListResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getvideolist response to a gRPC getvideolist reply. Primarily useful in a server.
func EncodeGRPCGetVideoListResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetVideoListRes)
	return resp, nil
}

// EncodeGRPCUpdateVideoResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain updatevideo response to a gRPC updatevideo reply. Primarily useful in a server.
func EncodeGRPCUpdateVideoResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.UpdateVideoRes)
	return resp, nil
}

// EncodeGRPCDeleteVideoResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain deletevideo response to a gRPC deletevideo reply. Primarily useful in a server.
func EncodeGRPCDeleteVideoResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.DeleteVideoRes)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
