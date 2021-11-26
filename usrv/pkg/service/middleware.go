package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UsrvService) UsrvService

type loggingMiddleware struct {
	logger log.Logger
	next   UsrvService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UsrvService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UsrvService) UsrvService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) GetUser(ctx context.Context, id int64) (s0 string, e1 error) {
	defer func() {
		l.logger.Log("method", "GetUser", "id", id, "s0", s0, "e1", e1)
	}()
	return l.next.GetUser(ctx, id)
}
