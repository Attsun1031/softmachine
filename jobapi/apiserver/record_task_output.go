package apiserver

import (
	"fmt"

	jobapi_pb "github.com/Attsun1031/jobnetes/jobapi/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *JobApiServerImpl) RecordTaskOutput(ctx context.Context, request *jobapi_pb.TaskOutputRecordRequest) (*jobapi_pb.TaskOutputRecordResponse, error) {
	conn := s.connect()
	tx := conn.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()

	te, err := s.TaskExecutionDao.FindById(uint(request.GetTaskId()), tx.Set("gorm:query_option", "FOR UPDATE"))
	if err != nil {
		tx.Rollback()
		msg := fmt.Sprintf("Failed to load task execution by db error. err=%v", err)
		return nil, status.Errorf(codes.Internal, msg)
	}

	te.Output = string(request.GetData())
	err = s.TaskExecutionDao.Update(te, tx)
	if err != nil {
		tx.Rollback()
		msg := fmt.Sprintf("Failed to save task execution by db error. err=%v", err)
		return nil, status.Errorf(codes.Internal, msg)
	}

	tx.Commit()
	return &jobapi_pb.TaskOutputRecordResponse{TaskId: request.GetTaskId()}, nil
}
