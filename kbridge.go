package main


import (
	"context"
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"l0calh0st.cn/k8s-bridge/pkg/controller"
	"l0calh0st.cn/k8s-bridge/pkg/controller/kube-resource"
)

var (
	masterUrl = flag.String("masterUrl", "","")
	kubeConfig = flag.String("kubeConfig", "", "")
)

var (
	kubeClientSet *kubernetes.Clientset
)

func main() {
	if err := initializeClientSets();err != nil{
		return
	}
	ctx,cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	ksvController := kube_resource.NewKubeResourceController(kubeClientSet)
	go runController(ctx, ksvController)
}


func runController(ctx context.Context, controller controller.Controller){

}

func initializeClientSets()error{
	var err error
	cfg,err := clientcmd.BuildConfigFromFlags(*masterUrl, *kubeConfig)
	if err != nil{
		return err
	}
	kubeClientSet,err = kubernetes.NewForConfig(cfg)
	if err != nil{
		return err
	}
	return nil

}