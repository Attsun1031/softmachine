package taskpoller

import (
	"time"

	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/jinzhu/gorm"
)

type ParallelJobTaskPoller struct {
	TaskExecutionDao dao.TaskExecutionDao
	TaskDef          *model.ParallelTask
}

func (poller *ParallelJobTaskPoller) Poll(te *model.TaskExecution, db *gorm.DB) (bool, error) {
	childTasks, err := poller.TaskExecutionDao.FindChildTasks(te.ID, db)
	if err != nil {
		return false, err
	}

	expectedChildCount := 0
	for _, ts := range poller.TaskDef.TaskSets {
		expectedChildCount += len(ts)
	}

	hasFailedTask := false
	actualCompletedChildCount := 0
	for _, t := range childTasks {
		if t.IsCompleted() {
			actualCompletedChildCount += 1
			if t.IsFailed() {
				hasFailedTask = true
			}
		}
	}

	// check if all child task completed
	if expectedChildCount == actualCompletedChildCount {
		endedAt := time.Now()
		if hasFailedTask {
			te.Status = model.TaskFailed
			te.MarkFailed(&endedAt, "Some child job failed.", "")
		} else {
			te.Status = model.TaskSuccess
			te.MarkSuccess(&endedAt)
		}
		err = poller.TaskExecutionDao.Update(te, db)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}
