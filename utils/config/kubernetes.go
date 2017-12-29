package config

import (
	"github.com/spf13/viper"
)

type KubernetesConfig struct {
	InCluster    bool
	ConfigPath   string
	MasterUrl    string
	JobNamespace string
}

func LoadKubernetesConfig() *KubernetesConfig {
	return &KubernetesConfig{
		InCluster:    viper.GetBool("k8s.in-cluster"),
		ConfigPath:   viper.GetString("k8s.config-path"),
		MasterUrl:    viper.GetString("k8s.master-url"),
		JobNamespace: viper.GetString("k8s.job-namespace"),
	}
}
