package kubernetes

import (
	"os"

	"github.com/Attsun1031/jobnetes/utils/consts"
)

type KubeConfig struct {
	InCluster     bool
	ConfigPath    string
	KubeMasterUrl string
}

func LoadKubeConfig() *KubeConfig {
	c := &KubeConfig{
		InCluster: true,
	}
	inCluster := os.Getenv(consts.KubeInCluster)
	if inCluster == "0" {
		c.InCluster = false
	}
	kubeConfigPath := os.Getenv(consts.KubeConfigPath)
	if kubeConfigPath != "" {
		c.ConfigPath = kubeConfigPath
	}
	kubeMasterUrl := os.Getenv(consts.KubeMasterUrl)
	if kubeMasterUrl != "" {
		c.KubeMasterUrl = kubeMasterUrl
	}
	return c
}
