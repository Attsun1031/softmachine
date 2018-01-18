package main

import (
	"log"

	"flag"

	"fmt"

	jobapi_pb "github.com/Attsun1031/jobnetes/jobapi/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// Set up a connection to the server.
	endpoint := flag.String("endpoint", "localhost:50051", "jobnetes-api-server endpoint URL")
	username := flag.String("username", "jobnetesadmin", "jobnetes-api-server auth user")
	password := flag.String("password", "test", "jobnetes-api-server auth password")
	workflowId := flag.Uint64("workflow-id", 0, "workflow id")
	execName := flag.String("exec-name", "", "exec name")
	input := flag.String("input", `{"x":"string","y":1}`, "workflow input")
	flag.Parse()
	fmt.Printf("Arguments: endpoint=%v, username=%v, password=%v, workflow-id=%v, exec-name=%v, input=%v\n", *endpoint, *username, *password, *workflowId, *execName, *input)

	conn, err := grpc.Dial(*endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := jobapi_pb.NewJobapiClient(conn)

	// Contact the server and print out its response.
	md := metadata.Pairs("username", *username, "password", *password)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	startWorkflow(c, ctx, *workflowId, *execName, *input)
}

func startWorkflow(c jobapi_pb.JobapiClient, ctx context.Context, workflowId uint64, execName string, input string) {
	r, err := c.StartWorkflow(ctx, &jobapi_pb.WorkflowStartRequest{
		WorkflowId: workflowId,
		ExecName:   execName,
		Input:      []byte(input),
	})
	if err != nil {
		log.Fatalf("Failed to request start workflow: %v", err)
	}
	log.Printf("Result: %d", r.Id)
}
