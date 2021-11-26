package video

import (
	"context"
	vsrvpb "mithril-micro/vsrv/api/gen"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetVideoEndpoint(svc VideoService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*vsrvpb.GetVideoReq)
		v, err := svc.GetVideo(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}
