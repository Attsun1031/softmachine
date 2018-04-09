package taskpoller

import (
	"github.com/Attsun1031/jobnetes/dao"
	"github.com/Attsun1031/jobnetes/model"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/jinzhu/gorm"
	"k8s.io/api/batch/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type KubeJobTaskPoller struct {
	Client           kubernetes.Interface
	TaskExecutionDao dao.TaskExecutionDao
}

func (poller *KubeJobTaskPoller) Poll(te *model.TaskExecution, db *gorm.DB) (bool, error) {
	result, err := poller.Client.
		BatchV1().
		Jobs(config.JobnetesConfig.KubernetesConfig.JobNamespace).
		Get(te.ExecutionName, v1meta.GetOptions{})

	if err != nil {
		return false, err
	}
	if len(result.Status.Conditions) == 0 {
		return false, nil
	}

	lastCondition := result.Status.Conditions[0]

	switch lastCondition.Type {
	case v1.JobComplete:
		te.MarkSuccess(&result.Status.CompletionTime.Time)
	default:
		te.MarkFailed(&lastCondition.LastProbeTime.Time, lastCondition.Reason, lastCondition.Message)
	}
	err = poller.TaskExecutionDao.Update(te, db)
	if err != nil {
		return false, err
	}
	return true, nil
}
