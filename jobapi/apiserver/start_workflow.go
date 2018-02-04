package apiserver

import (
	"fmt"

	jobapi_pb "github.com/Attsun1031/jobnetes/jobapi/proto"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *JobApiServerImpl) StartWorkflow(ctx context.Context, request *jobapi_pb.WorkflowStartRequest) (*jobapi_pb.WorkflowStartResponse, error) {
	conn := s.connect()
	defer conn.Close()

	wfId := uint(request.GetWorkflowId())
	wf, err := s.WorkflowDao.FindById(wfId, conn)
	if err != nil {
		msg := fmt.Sprintf("Failed to load workflow by db error. err=%v", err)
		return nil, status.Errorf(codes.Internal, msg)
	}

	in := string(request.GetInput())
	exec := &model.WorkflowExecution{
		WorkflowID: wfId,
		Name:       request.GetExecName(),
		Definition: wf.Definition,
		Input:      in,
	}
	err = s.WorkflowExecutionDao.Create(exec, conn)
	if err != nil {
		msg := fmt.Sprintf("Failed to create workflow execution. req=%v err=%v", request, err)
		return nil, status.Errorf(codes.InvalidArgument, msg)
	}
	log.Logger.Infof("Accept start workflow request. execId=%v", exec.ID)
	return &jobapi_pb.WorkflowStartResponse{Id: uint64(exec.ID)}, nil
}
