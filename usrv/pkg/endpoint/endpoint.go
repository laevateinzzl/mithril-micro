package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "mithril-micro/usrv/pkg/service"
)

// GetUserRequest collects the request parameters for the GetUser method.
type GetUserRequest struct {
	Id int64 `json:"id"`
}

// GetUserResponse collects the response parameters for the GetUser method.
type GetUserResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeGetUserEndpoint returns an endpoint that invokes GetUser on the service.
func MakeGetUserEndpoint(s service.UsrvService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserRequest)
		s0, e1 := s.GetUser(ctx, req.Id)
		return GetUserResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetUserResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// GetUser implements Service. Primarily useful in a client.
func (e Endpoints) GetUser(ctx context.Context, id int64) (s0 string, e1 error) {
	request := GetUserRequest{Id: id}
	response, err := e.GetUserEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetUserResponse).S0, response.(GetUserResponse).E1
}
