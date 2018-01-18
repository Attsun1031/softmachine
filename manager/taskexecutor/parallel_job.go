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

func (executor *ParallelJobTaskExecutor) Execute(we *model.WorkflowExecution, db *gorm.DB, input string, parentID uint) error {
	task := executor.Task
	log.Logger.Infof("Requesting parallel task. ExecutionName=%s Type=%s ParentID", task.Name, task.GetJobType(), parentID)

	// create execution record
	startedAt := time.Now()
	te := &model.TaskExecution{
		WorkflowExecution:     we,
		ParentTaskExecutionID: parentID,
		TaskName:              task.Name,
		StartedAt:             &startedAt,
		Status:                model.TaskRunning,
		Input:                 input,
		Output:                "{}",
	}
	err := executor.TaskExecutionDao.Update(te, db)
	if err != nil {
		return err
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
		return err
	}

	executors := make([]TaskExecutor, len(task.TaskSets))
	for _, taskSet := range task.TaskSets {
		startTask := taskSet[0]
		executor, err := executor.TaskExecutorFactory.GetTaskExecutor(startTask)
		if err != nil {
			return err
		}
		executors = append(executors, executor)
	}

	for _, executor := range executors {
		if executor == nil {
			continue
		}
		err = executor.Execute(we, db, input, te.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
