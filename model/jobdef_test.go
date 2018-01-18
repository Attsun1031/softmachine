package model

import (
	"reflect"
	"testing"
)

func TestGetJobDefFromString(t *testing.T) {
	str := `{"name":"test1","tasks":[{"type":"kube-job","name":"task-1","next":"task-2"}]}`
	jobDef := GetJobDefFromString(str)
	expected := &JobDef{
		Tasks: []Task{
			&KubeJobTask{Name: "task-1", NextTaskName: "task-2"},
		},
	}
	if !reflect.DeepEqual(expected, jobDef) {
		t.FailNow()
	}
}

func TestGetJobDefFromStringWithParallel(t *testing.T) {
	str := `
	{
		"name":"test1",
		"tasks":[
			{
				"type":"parallel-job",
				"name":"para-task-1",
				"task-sets":[
					[
						{"type":"kube-job","name":"kube-task-1","next":"kube-task-2"},
						{"type":"kube-job","name":"kube-task-2"}
					],
					[
						{"type":"kube-job","name":"kube-task-3","next":"kube-task-4"},
						{"type":"kube-job","name":"kube-task-4"}
					],
					[
						{"type":"kube-job","name":"kube-task-5","next":"para-task-2"},
						{
							"type":"parallel-job",
							"name":"para-task-2",
							"task-sets":[
								[
									{"type":"kube-job","name":"kube-task-6","next":"kube-task-7"},
									{"type":"kube-job","name":"kube-task-7"}
								]
							]
						}
					]
				],
				"next":"task-3"
			}
		]
	}`
	jobDef := GetJobDefFromString(str)
	expected := &JobDef{
		Tasks: []Task{
			&ParallelTask{
				Name:         "para-task-1",
				NextTaskName: "task-3",
				TaskSets: [][]Task{
					{
						&KubeJobTask{Name: "kube-task-1", NextTaskName: "kube-task-2"},
						&KubeJobTask{Name: "kube-task-2"},
					},
					{
						&KubeJobTask{Name: "kube-task-3", NextTaskName: "kube-task-4"},
						&KubeJobTask{Name: "kube-task-4"},
					},
					{
						&KubeJobTask{Name: "kube-task-5", NextTaskName: "para-task-2"},
						&ParallelTask{
							Name: "para-task-2",
							TaskSets: [][]Task{
								{
									&KubeJobTask{Name: "kube-task-6", NextTaskName: "kube-task-7"},
									&KubeJobTask{Name: "kube-task-7"},
								},
							},
						},
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(expected, jobDef) {
		t.FailNow()
	}
}

func TestJobDef_GetStartTask(t *testing.T) {
	startTask := &KubeJobTask{Name: "task-1", NextTaskName: "task-2"}
	jobDef := &JobDef{
		Tasks: []Task{
			startTask,
			&KubeJobTask{Name: "task-2", NextTaskName: "task-3"},
		},
	}
	if !reflect.DeepEqual(startTask, jobDef.GetStartTask()) {
		t.FailNow()
	}
}
