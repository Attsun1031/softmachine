package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClient(config *KubeConfig) *kubernetes.Clientset {
	if config.InCluster {
		return getInClusterClient()
	} else {
		return getOutOfClusterKubeClient(config.KubeMasterUrl, config.ConfigPath)
	}
}

func getInClusterClient() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func getOutOfClusterKubeClient(masterUrl string, configPath string) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, configPath)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
