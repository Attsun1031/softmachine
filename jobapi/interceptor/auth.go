package interceptor

import (
	"github.com/Attsun1031/jobnetes/utils/config"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func authorize(ctx context.Context) error {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if len(md["username"]) > 0 && md["username"][0] == config.JobnetesConfig.JobApiConfig.Username &&
			len(md["password"]) > 0 && md["password"][0] == config.JobnetesConfig.JobApiConfig.Password {
			return nil
		}
		return status.Error(codes.Unauthenticated, "Authentication failed")
	}
	return status.Error(codes.InvalidArgument, "No metadata")
}
