package kubernetes

import (
	"github.com/Attsun1031/jobnetes/utils/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClient(config *config.KubernetesConfig) *kubernetes.Clientset {
	if config.InCluster {
		return getInClusterClient()
	} else {
		return getOutOfClusterKubeClient(config.MasterUrl, config.ConfigPath)
	}
}

func getInClusterClient() *kubernetes.Clientset {
	c, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	cli, err := kubernetes.NewForConfig(c)
	if err != nil {
		panic(err.Error())
	}
	return cli
}

func getOutOfClusterKubeClient(masterUrl string, configPath string) *kubernetes.Clientset {
	c, err := clientcmd.BuildConfigFromFlags(masterUrl, configPath)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	cli, err := kubernetes.NewForConfig(c)
	if err != nil {
		panic(err.Error())
	}
	return cli
}
