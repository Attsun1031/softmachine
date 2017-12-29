package model

import (
	"encoding/json"
	"fmt"

	"github.com/Attsun1031/jobnetes/utils/log"
	"k8s.io/api/batch/v1"
)

type JobDefProvider interface {
	GetJobDef() *JobDef
}

func GetJobDefFromString(jobDefStr string) *JobDef {
	var err error
	rawJobDef := &_RawJobDef{}
	err = json.Unmarshal([]byte(jobDefStr), rawJobDef)
	if err != nil {
		log.Logger.Fatal(err)
	}
	jobDef := &JobDef{Name: rawJobDef.Name}

	tasks := make([]Task, len(rawJobDef.Tasks))
	for i, t := range rawJobDef.Tasks {
		decoded := t.(map[string]interface{})
		tp := decoded["type"].(string)
		name := decoded["name"].(string)
		next := decoded["next"]
		switch tp {
		case "kube-job":
			var b []byte
			b, err = json.Marshal(decoded["job"])
			if err != nil {
				log.Logger.Fatal(err)
			}
			j := &v1.Job{}
			err = json.Unmarshal(b, j)
			if err != nil {
				log.Logger.Fatal(err)
			}
			kjt := &KubeJobTask{Name: name, KubeJobSpec: *j}
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
