package apiserver

import (
	"fmt"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/dao/db"
	"github.com/Attsun1031/jobnetes/di"
	jobapi_pb "github.com/Attsun1031/jobnetes/jobapi"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	WorkflowDao          dao.WorkflowDao
	WorkflowExecutionDao dao.WorkflowExecutionDao
}

func connect() *gorm.DB {
	d := db.Connect(config.JobnetesConfig.DbConfig)
	d.SetLogger(log.Logger)
	return d
}

func (s *server) StartWorkflow(ctx context.Context, request *jobapi_pb.WorkflowStartRequest) (*jobapi_pb.WorkflowStartResponse, error) {
	conn := connect()
	wfId := uint(request.GetWorkflowId())
	wf := s.WorkflowDao.FindById(wfId, conn)
	in := string(request.GetInput())
	exec := &model.WorkflowExecution{
		WorkflowID: wfId,
		Name:       request.GetExecName(),
		Definition: wf.Definition,
		Input:      in,
	}
	err := s.WorkflowExecutionDao.Create(exec, conn)
	if err != nil {
		msg := fmt.Sprintf("Failed to create workflow execution. req=%v err=%v", request, err)
		return nil, status.Errorf(codes.InvalidArgument, msg)
	}
	return &jobapi_pb.WorkflowStartResponse{Id: int64(exec.ID)}, nil
}

func MakeJobApiServer() jobapi_pb.JobapiServer {
	return &server{
		WorkflowDao:          di.InjectWorkflowDao(),
		WorkflowExecutionDao: di.InjectWorkflowExecutionDao(),
	}
}
