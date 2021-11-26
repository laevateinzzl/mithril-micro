package service

import "context"

// UsrvService describes the service.
type UsrvService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	GetUser(ctx context.Context, id int64) (string, error)
}

type UserService struct{}

func (s *UserService) GetUser(ctx context.Context, id int64) (string, error) {
	return "test", nil
}

type basicUsrvService struct{}

func (b *basicUsrvService) GetUser(ctx context.Context, id int64) (s0 string, e1 error) {
	// TODO implement the business logic of GetUser
	return s0, e1
}

// NewBasicUsrvService returns a naive, stateless implementation of UsrvService.
func NewBasicUsrvService() UsrvService {
	return &basicUsrvService{}
}

// New returns a UsrvService with all of the expected middleware wired in.
func New(middleware []Middleware) UsrvService {
	var svc UsrvService = NewBasicUsrvService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
