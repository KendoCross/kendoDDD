package trips

import (
	"path/filepath"

	eh "github.com/looplab/eventhorizon"
	"github.com/spf13/viper"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const TripsAggregateType = eh.AggregateType("AggregateType_Trips")

//单例聚合根的特殊实例，ID为NIL
var SingleTripsAgg *tripsAggregate

//K8S集群
var clientset *kubernetes.Clientset

// 仓储（能封装成工作单元更好）

func StartUp() (err error) {
	k8sMode := viper.GetString("K8S_MODE")
	kubeconfig := ""
	if k8sMode == "Outer" {
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
			config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
			if err != nil {
				panic(err.Error())
			}

			// creates the clientset
			clientset, err = kubernetes.NewForConfig(config)
			if err != nil {
				panic(err.Error())
			}
		}
	} else {
		// creates the in-cluster config
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		// creates the clientset
		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
	}

	return
}
