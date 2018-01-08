package interceptor

import (
	"github.com/Attsun1031/jobnetes/utils/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := authorize(ctx); err != nil {
		log.Logger.Error(err)
		return nil, err
	}
	r, err := handler(ctx, req)
	if err != nil {
		log.Logger.Error(err)
	}
	return r, err
}
