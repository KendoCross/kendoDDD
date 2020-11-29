package k8s_info

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type podEntity struct{

}

func newpodEnByOV()*podEntity{
return new(podEntity)
}

func (en *podEntity) PodList() (pods *v1.PodList, err error) {
	pods, err = clientset.CoreV1().Pods("").List(context.Background(),metav1.ListOptions{})
	return
}
