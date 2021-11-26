package grpc

import (
	"context"
	"errors"
	usersrvproto "mithril-micro/usrv/api/gen"
	endpoint1 "mithril-micro/usrv/pkg/endpoint"
	service "mithril-micro/usrv/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.UsrvService, error) {
	var getUserEndpoint endpoint.Endpoint
	{
		getUserEndpoint = grpc1.NewClient(conn, "pb.Usrv", "GetUser", encodeGetUserRequest, decodeGetUserResponse, usersrvproto.User{}, options["GetUser"]...).Endpoint()
	}

	return endpoint1.Endpoints{GetUserEndpoint: getUserEndpoint}, nil
}

// encodeGetUserRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain GetUser request to a gRPC request.
func encodeGetUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, errors.New("'Usrv' Encoder is not impelemented")
}

// decodeGetUserResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeGetUserResponse(_ context.Context, reply interface{}) (interface{}, error) {
	return nil, errors.New("'Usrv' Decoder is not impelemented")
}
