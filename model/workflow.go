package model

import (
	"encoding/json"
	"time"

	"fmt"

	"github.com/Attsun1031/jobnetes/utils/log"
)

type Workflow struct {
	ID         uint   `gorm:"primary_key"`
	Name       string `gorm:"not null"`
	Definition string `gorm:"type:json;not null"`
	User       *User  `gorm:"not null"`
	UserID     uint   `gorm:"not null"`
	CreatedAt  *time.Time
	UpdatedAt  *time.Time

	jobDef *JobDef `gorm:"-"`
}

type JobDefProvider interface {
	GetJobDef() *JobDef
}

func (workflow *Workflow) GetJobDef() *JobDef {
	if workflow.jobDef != nil {
		return workflow.jobDef
	}

	rawJobDef := &_RawJobDef{}
	json.Unmarshal([]byte(workflow.Definition), rawJobDef)
	jobDef := &JobDef{Name: rawJobDef.Name}

	tasks := make([]Task, len(rawJobDef.Tasks))
	for i, t := range rawJobDef.Tasks {
		decoded := t.(map[string]interface{})
		tp := decoded["type"].(string)
		name := decoded["name"].(string)
		next := decoded["next"]
		switch tp {
		case "kube-job":
			kjt := &KubeJobTask{Name: name}
			if next != nil {
				kjt.NextTaskName = next.(string)
			}
			tasks[i] = kjt
		case "parallel":
			log.Logger.Warn("TO BE IMPLEMENTED")
		case "choice":
			log.Logger.Warn("TO BE IMPLEMENTED")
		default:
			log.Logger.Error(fmt.Sprintf("Unknown job type %s", tp))
		}
	}
	jobDef.Tasks = tasks
	workflow.jobDef = jobDef
	return jobDef
}

type _RawJobDef struct {
	Name  string
	Tasks []interface{}
}

type JobDef struct {
	Name  string
	Tasks []Task
}

// Get first job of this definition
func (jobDef *JobDef) GetStartTask() Task {
	return jobDef.Tasks[0]
}

// Get next job
// Next job is decided by tasks' status which belong to the execution
func (jobDef *JobDef) GetNextTask(execution *WorkflowExecution) Task {
	return nil
}
