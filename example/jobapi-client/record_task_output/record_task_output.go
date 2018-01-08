package main

import (
	"log"

	"flag"
	"fmt"

	jobapi_pb "github.com/Attsun1031/jobnetes/jobapi"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	endpoint := flag.String("endpoint", "localhost:50051", "jobnetes-api-server endpoint URL")
	username := flag.String("username", "jobnetesadmin", "jobnetes-api-server auth user")
	password := flag.String("password", "test", "jobnetes-api-server auth password")
	taskId := flag.Uint64("task-id", 0, "task id")
	data := flag.String("data", `{"x":"string","y":1}`, "task output data")
	flag.Parse()
	fmt.Printf("Arguments: endpoint=%v, username=%v, password=%v, task-id=%v, data=%v\n", *endpoint, *username, *password, *taskId, *data)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := jobapi_pb.NewJobapiClient(conn)

	// Contact the server and print out its response.
	md := metadata.Pairs("username", *username, "password", *password)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	recordTaskOutput(c, ctx, *taskId, *data)
}

func recordTaskOutput(c jobapi_pb.JobapiClient, ctx context.Context, taskId uint64, data string) {
	request := &jobapi_pb.TaskOutputRecordRequest{
		TaskId: taskId,
		Data:   []byte(data),
	}
	log.Printf("Call RecordTaskOutput: %v", request)

	r2, err := c.RecordTaskOutput(ctx, request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Result: %d", r2.TaskId)
}
