package main

import (
	"time"

	"os"
	"os/signal"

	"syscall"

	"github.com/Attsun1031/jobnetes/controller"
	"github.com/Attsun1031/jobnetes/kubernetes"
	clientset "github.com/Attsun1031/jobnetes/pkg/client/clientset/versioned"
	informers "github.com/Attsun1031/jobnetes/pkg/client/informers/externalversions"
	"github.com/Attsun1031/jobnetes/utils/config"
	"github.com/Attsun1031/jobnetes/utils/log"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config.InitConfig()
	log.SetupLogger(config.JobnetesConfig.LogConfig)
	kubeClient := kubernetes.GetClient(config.JobnetesConfig.KubernetesConfig)
	jobnetesClient := getOutClusterCrdClient()
	jobnetesInformerFactory := informers.NewSharedInformerFactory(jobnetesClient, time.Second*30)
	ctlr := controller.NewController(kubeClient, jobnetesClient, jobnetesInformerFactory)

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := SetupSignalHandler()
	go jobnetesInformerFactory.Start(stopCh)

	if err := ctlr.Run(2, stopCh); err != nil {
		log.Logger.Fatalf("Error running controller: %s", err.Error())
	}
}

func getOutClusterCrdClient() *clientset.Clientset {
	c, err := clientcmd.BuildConfigFromFlags(config.JobnetesConfig.KubernetesConfig.MasterUrl, config.JobnetesConfig.KubernetesConfig.ConfigPath)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	cli, err := clientset.NewForConfig(c)
	if err != nil {
		panic(err.Error())
	}
	return cli
}

func getInClusterCrdClient() *clientset.Clientset {
	c, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	cli, err := clientset.NewForConfig(c)
	if err != nil {
		panic(err.Error())
	}
	return cli
}

var onlyOneSignalHandler = make(chan struct{})
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
