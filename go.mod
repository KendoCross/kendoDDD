module github.com/KendoCross/kendoDDD

go 1.15

require (
	github.com/astaxie/beego v1.12.3
	github.com/ghodss/yaml v1.0.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.2.0
	github.com/go-redis/redis/v7 v7.4.0
	github.com/google/uuid v1.1.2
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/looplab/eventhorizon v0.8.0
	github.com/spf13/viper v1.7.1
	github.com/tidwall/gjson v1.6.4
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.8
	helm.sh/helm/v3 v3.4.2
	k8s.io/api v0.19.4
	k8s.io/apimachinery v0.19.4
	k8s.io/client-go v0.19.4
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920 // indirect
	rsc.io/letsencrypt v0.0.3 // indirect
)

// replace (
// 	golang.org/x/sync => golang.org/x/sync v0.0.0-20181108010431-42b317875d0f
// 	golang.org/x/sys => golang.org/x/sys v0.0.0-20190209173611-3b5209105503
// 	golang.org/x/tools => golang.org/x/tools v0.0.0-20190313210603-aa82965741a9
// 	k8s.io/api => k8s.io/api v0.0.0-20191016110246-af539daaa43a
// 	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191004115701-31ade1b30762
// )
