package routes

import "github.com/gin-gonic/gin"

var (
	ginApp = gin.Default()
)

func RoutesServer1() {
	r := gin.Default()
	ApplyRoutesServer1(ginApp)
	r.Run(":8081")
}

func ApplyRoutesServer1(r *gin.Engine) {
	GroupRoutesServer1(&r.RouterGroup)
}

func GroupRoutesServer1(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		v1.GET("/get")
	}
}
