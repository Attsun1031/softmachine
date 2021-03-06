package main

import (
	"net"

	"fmt"

	"github.com/Attsun1031/jobnetes/di"
	"github.com/Attsun1031/jobnetes/jobapi/apiserver"
	"github.com/Attsun1031/jobnetes/jobapi/interceptor"
	jobapi_pb "github.com/Attsun1031/jobnetes/jobapi/proto"
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
	jobapi_pb.RegisterJobapiServer(s, makeJobApiServer())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Logger.Infof("Start jobapi server. port: %v", config.JobnetesConfig.JobApiConfig.Port)
	if err := s.Serve(lis); err != nil {
		log.Logger.Fatalf("failed to serve: %v", err)
	}
}

func makeJobApiServer() jobapi_pb.JobapiServer {
	return &apiserver.JobApiServerImpl{
		WorkflowDao:          di.InjectWorkflowDao(),
		WorkflowExecutionDao: di.InjectWorkflowExecutionDao(),
		TaskExecutionDao:     di.InjectTaskExecutionDao(),
	}
}
