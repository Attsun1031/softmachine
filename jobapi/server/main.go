package main

import (
	"net"

	"fmt"

	jobapi_pb "github.com/Attsun1031/jobnetes/jobapi"
	"github.com/Attsun1031/jobnetes/jobapi/server/apiserver"
	"github.com/Attsun1031/jobnetes/jobapi/server/interceptor"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/Attsun1031/jobnetes/utils/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.InitConfig()
	log.SetupLogger(config.JobnetesConfig.LogConfig)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.JobnetesConfig.JobApiConfig.Port))
	if err != nil {
		log.Logger.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.UnaryInterceptor),
	)
	jobapi_pb.RegisterJobapiServer(s, apiserver.MakeJobApiServer())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Logger.Fatalf("failed to serve: %v", err)
	}
}
