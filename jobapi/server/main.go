package main

import (
	"log"
	"net"

	"fmt"

	"encoding/json"

	jobapi_pb "github.com/Attsun1031/jobnetes/jobapi"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

type SampleInput struct {
	Name string
	Age  int
}

func (s *server) StartWorkflow(ctx context.Context, request *jobapi_pb.WorkflowStartRequest) (*jobapi_pb.WorkflowStartResponse, error) {
	b := request.GetInput()
	i := &SampleInput{}
	err := json.Unmarshal(b, i)
	if err != nil {
		println(err)
	}
	fmt.Printf("Receive request. id=%v name=%v input.Name=%v input.Age=%v\n", request.GetWorkflowId(), request.GetExecName(), i.Name, i.Age)
	return &jobapi_pb.WorkflowStartResponse{Id: 1}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	jobapi_pb.RegisterJobapiServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
