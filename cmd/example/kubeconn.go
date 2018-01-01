package main

import (
	"fmt"

	"github.com/Attsun1031/jobnetes/kubernetes"
	"github.com/Attsun1031/jobnetes/utils/config"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes2 "k8s.io/client-go/kubernetes"
)

func main() {
	config.InitConfig()
	cli := kubernetes.GetClient(config.JobnetesConfig.KubernetesConfig)
	//hello(cli)

	//kubeJobSpec := `{"spec": {"template": {"spec": {"containers": [{"name": "pi", "image": "perl", "command": ["perl",  "-Mbignum=bpi", "-wle", "print bpi(2000)"]}], "restartPolicy": "Never"}}}, "metadata": {"name": "pi-5"}}`
	//j := &batchv1.Job{}
	//json.Unmarshal([]byte(kubeJobSpec), j)
	//_, err := cli.BatchV1().Jobs("default").Create(j)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	j, err := cli.BatchV1().Jobs("default").Get("pi-1-1-2017-12-29-12-45-29-99", v1.GetOptions{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(j)
	}
}

func hello(cli *kubernetes2.Clientset) {
	pods, err := cli.CoreV1().Pods("default").List(v1.ListOptions{})
	fmt.Printf("Pods count: %d\n", len(pods.Items))
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
		fmt.Println(pod.Labels)
		fmt.Println(pod.Spec)
	}
	if err != nil {
		fmt.Println(err.Error())
	}
}
