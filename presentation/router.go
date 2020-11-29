package presentation

import (
	"math/rand"
	"net/http"

	"github.com/KendoCross/kendoDDD/application"
	"github.com/KendoCross/kendoDDD/infrastructure/ddd"
	"github.com/KendoCross/kendoDDD/presentation/api"
	"github.com/KendoCross/kendoDDD/presentation/api/k8s"
	"github.com/KendoCross/kendoDDD/presentation/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//表现层主要的职责在于，表现形式的多样化。

func InitRouter() *gin.Engine {

	if viper.GetString("APP_MODE") == "prod" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}

	r := gin.Default()
	// 自定义error handler
	r.Use(middleware.Errors())
	//r.Use(md.GinCors())

	// 请求记录写入日志文件
	//r.Use(logger.RequestHandler(logger.Logger), gin.Recovery())
	r.GET("/", healthHand)

	testonly := r.Group("testonly")
	testonly.POST("", ddd.PreCommandDeals(application.TestOnlyCmd))

	files := r.Group("files")
	files.GET("/:id", api.GetFile)
	files.POST("", api.AddFile)

	kube := r.Group("kube")
	//pod := k8s.Group("pod")
	kube.GET("pods", k8s.GetPods)

	return r
}

var pongStr = []string{"Hello World", "你好，世界", "こんにちは世界", "Hallo Welt", "Привет, мир", "Bonjour le monde", "Hei maailma", "Saluton mondo", "salve Orbis Terrarum", "Сайн уу", "Nyob zoo lub ntiaj teb"}

//健康检查API
func healthHand(c *gin.Context) {
	// router := httprouter.New()
	// router.GET("", nil)

	index := rand.Intn(10)
	c.JSON(http.StatusOK, pongStr[index])
}
