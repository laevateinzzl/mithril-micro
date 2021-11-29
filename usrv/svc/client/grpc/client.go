// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 68ba2c3132
// Version Date: 2021-06-08T17:59:18Z

// Package grpc provides a gRPC client for the UserService service.
package grpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "mithril-micro/usrv/pb"
	"mithril-micro/usrv/svc"
)

// New returns an service backed by a gRPC client connection. It is the
// responsibility of the caller to dial, and later close, the connection.
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.UserServiceServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var createuserEndpoint endpoint.Endpoint
	{
		createuserEndpoint = grpctransport.NewClient(
			conn,
			"usrvpb.UserService",
			"CreateUser",
			EncodeGRPCCreateUserRequest,
			DecodeGRPCCreateUserResponse,
			pb.CreateUserRes{},
			clientOptions...,
		).Endpoint()
	}

	var getuserEndpoint endpoint.Endpoint
	{
		getuserEndpoint = grpctransport.NewClient(
			conn,
			"usrvpb.UserService",
			"GetUser",
			EncodeGRPCGetUserRequest,
			DecodeGRPCGetUserResponse,
			pb.User{},
			clientOptions...,
		).Endpoint()
	}

	var updateuserEndpoint endpoint.Endpoint
	{
		updateuserEndpoint = grpctransport.NewClient(
			conn,
			"usrvpb.UserService",
			"UpdateUser",
			EncodeGRPCUpdateUserRequest,
			DecodeGRPCUpdateUserResponse,
			pb.UpdateUserRes{},
			clientOptions...,
		).Endpoint()
	}

	var loginEndpoint endpoint.Endpoint
	{
		loginEndpoint = grpctransport.NewClient(
			conn,
			"usrvpb.UserService",
			"Login",
			EncodeGRPCLoginRequest,
			DecodeGRPCLoginResponse,
			pb.LoginRes{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		CreateUserEndpoint: createuserEndpoint,
		GetUserEndpoint:    getuserEndpoint,
		UpdateUserEndpoint: updateuserEndpoint,
		LoginEndpoint:      loginEndpoint,
	}, nil
}

// GRPC Client Decode

// DecodeGRPCCreateUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC createuser reply to a user-domain createuser response. Primarily useful in a client.
func DecodeGRPCCreateUserResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.CreateUserRes)
	return reply, nil
}

// DecodeGRPCGetUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC getuser reply to a user-domain getuser response. Primarily useful in a client.
func DecodeGRPCGetUserResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.User)
	return reply, nil
}

// DecodeGRPCUpdateUserResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC updateuser reply to a user-domain updateuser response. Primarily useful in a client.
func DecodeGRPCUpdateUserResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.UpdateUserRes)
	return reply, nil
}

// DecodeGRPCLoginResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC login reply to a user-domain login response. Primarily useful in a client.
func DecodeGRPCLoginResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.LoginRes)
	return reply, nil
}

// GRPC Client Encode

// EncodeGRPCCreateUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain createuser request to a gRPC createuser request. Primarily useful in a client.
func EncodeGRPCCreateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateUserReq)
	return req, nil
}

// EncodeGRPCGetUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain getuser request to a gRPC getuser request. Primarily useful in a client.
func EncodeGRPCGetUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetUserReq)
	return req, nil
}

// EncodeGRPCUpdateUserRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain updateuser request to a gRPC updateuser request. Primarily useful in a client.
func EncodeGRPCUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateUserReq)
	return req, nil
}

// EncodeGRPCLoginRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain login request to a gRPC login request. Primarily useful in a client.
func EncodeGRPCLoginRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.LoginReq)
	return req, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}