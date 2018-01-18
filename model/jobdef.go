package model

import (
	"encoding/json"
	"fmt"

	"github.com/Attsun1031/jobnetes/utils/log"
	"k8s.io/api/batch/v1"
)

const JobTypeKube = "kube-job"
const JobTypeParallel = "parallel-job"
const JobTypeChoice = "choice-job"

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
	tasks := decodeTasks(rawJobDef.Tasks)
	return &JobDef{Tasks: tasks}
}

func decodeTasks(rawTasks []interface{}) []Task {
	tasks := make([]Task, len(rawTasks))
	for i, t := range rawTasks {
		var err error
		decoded := t.(map[string]interface{})
		tp := decoded["type"].(string)
		name := decoded["name"].(string)
		next := decoded["next"]
		switch tp {
		case JobTypeKube:
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
		case JobTypeParallel:
			pt := &ParallelTask{Name: name}
			if next != nil {
				pt.NextTaskName = next.(string)
			}
			taskSets := decoded["task-sets"].([]interface{})
			pt.TaskSets = make([][]Task, len(taskSets))
			for i, tasksInSet := range taskSets {
				pt.TaskSets[i] = decodeTasks(tasksInSet.([]interface{}))
			}
			tasks[i] = pt
		case JobTypeChoice:
			log.Logger.Warn("TO BE IMPLEMENTED")
		default:
			log.Logger.Error(fmt.Sprintf("Unknown job type %s", tp))
		}
	}
	return tasks
}

type _RawJobDef struct {
	Name  string
	Tasks []interface{}
}

type JobDef struct {
	Tasks []Task
}

// Get first job of this definition
func (jobDef *JobDef) GetStartTask() Task {
	return jobDef.Tasks[0]
}

// Get te's task definition
func (jobDef *JobDef) GetCurrentTask(te *TaskExecution) Task {
	return filterTasks(jobDef.Tasks, te.TaskName)
}

// Get next job
// Next job is decided by tasks' status which belong to the execution
func (jobDef *JobDef) GetNextTask(te *TaskExecution) Task {
	current := jobDef.GetCurrentTask(te)
	nextName := current.GetNextTaskName()
	if nextName == "" {
		return nil
	}
	nextTask := filterTasks(jobDef.Tasks, nextName)
	if nextTask == nil {
		log.Logger.Fatalf("Next task %v not found in job definition.", nextName)
	}
	return nextTask
}

func filterTasks(tasks []Task, taskName string) Task {
	for _, t := range tasks {
		if t.GetName() == taskName {
			return t
		}
		if t.GetJobType() == JobTypeParallel {
			for _, ts := range t.(*ParallelTask).TaskSets {
				it := filterTasks(ts, taskName)
				if it != nil {
					return it
				}
			}
		}
	}
	return nil
}
