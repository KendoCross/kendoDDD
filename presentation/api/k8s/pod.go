package k8s

import (
	"net/http"

	"github.com/KendoCross/kendoDDD/application/k8s"
	"github.com/gin-gonic/gin"
)

func GetPods(c *gin.Context) {

	var parm k8s.PodsReq
	rst, err := k8s.PodList(parm)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, rst)
}
