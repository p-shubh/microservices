package routes

import "github.com/gin-gonic/gin"

var (
	ginApp = gin.Default()
)

func RoutesServer2() {
	r := gin.Default()
	ApplyRoutesServer2(ginApp)
	r.Run(":8082")
}

func ApplyRoutesServer2(r *gin.Engine) {
	GroupRoutesServer2(&r.RouterGroup)
}

func GroupRoutesServer2(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/get")
	}
}
