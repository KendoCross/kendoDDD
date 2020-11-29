package testcmd

import (
	"path/filepath"

	"github.com/KendoCross/kendoDDD/domain/services/helm"
	"github.com/KendoCross/kendoDDD/infrastructure/bus"
	"github.com/spf13/viper"
	"k8s.io/client-go/util/homedir"
)

//单例聚合根的特殊实例，ID为NIL
var SingletonAggTestOnly *testOnlyAggregate

var helmClient *helm.HelmAction

// 仓储（能封装成工作单元更好）

func StartUp() (err error) {
	bus.SetDealer(SingletonAggTestOnly, TestOnlyCmdType)
	k8sMode := viper.GetString("K8S_MODE")
	kubeconfig := ""
	if k8sMode == "Outer" {
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = filepath.Join(home, ".kube", "config")
		}
	}

	helmClient, err = helm.NewHelmAction(kubeconfig, "")

	return
}
