package main

import (
	"log"

	jobapi_pb "github.com/Attsun1031/jobnetes/jobapi"
	"github.com/gin-gonic/gin/json"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:50051"
)

type SampleInput struct {
	Name string
	Age  int
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := jobapi_pb.NewJobapiClient(conn)

	// Contact the server and print out its response.
	i := &SampleInput{Name: "Jon", Age: 20}
	b, err := json.Marshal(i)
	if err != nil {
		println(err)
		return
	}
	md := metadata.Pairs("username", "jobnetesadmin", "password", "test")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.StartWorkflow(ctx, &jobapi_pb.WorkflowStartRequest{
		WorkflowId: 2,
		ExecName:   "test-seq-2",
		Input:      b,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Result: %d", r.Id)
}
