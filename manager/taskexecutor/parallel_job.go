package taskexecutor

import (
	"fmt"

	"time"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/jinzhu/gorm"
)

type ParallelJobTaskExecutor struct {
	Task                *model.ParallelTask
	TaskExecutionDao    dao.TaskExecutionDao
	TaskExecutorFactory Factory
}

func (executor *ParallelJobTaskExecutor) Execute(we *model.WorkflowExecution, db *gorm.DB, input string, parentId uint, prevId uint) (*model.TaskExecution, error) {
	task := executor.Task
	log.Logger.Infof("Requesting parallel task. ExecutionName=%v Type=%v ParentId=%v PrevId=%v", task.Name, task.GetJobType(), parentId, prevId)

	// create execution record
	startedAt := time.Now()
	te := &model.TaskExecution{
		WorkflowExecution:     we,
		ParentTaskExecutionID: parentId,
		PrevTaskExecutionID:   prevId,
		TaskName:              task.Name,
		TaskType:              task.GetJobType(),
		StartedAt:             &startedAt,
		Status:                model.TaskRunning,
		Input:                 input,
		Output:                "{}",
	}
	err := executor.TaskExecutionDao.Update(te, db)
	if err != nil {
		return nil, err
	}

	name := fmt.Sprintf(
		"%s-%d-%d-%s",
		task.Name,
		we.ID,
		te.ID,
		time.Now().Format("2006-01-02-15-04-05-99"))
	te.ExecutionName = name
	err = executor.TaskExecutionDao.Update(te, db)
	if err != nil {
		return nil, err
	}

	executors := make([]TaskExecutor, len(task.TaskSets))
	for _, taskSet := range task.TaskSets {
		startTask := taskSet[0]
		executor, err := executor.TaskExecutorFactory.GetTaskExecutor(startTask)
		if err != nil {
			return nil, err
		}
		executors = append(executors, executor)
	}

	for _, executor := range executors {
		if executor == nil {
			continue
		}
		_, err = executor.Execute(we, db, input, te.ID, prevId)
		if err != nil {
			return nil, err
		}
	}
	return te, nil
}
