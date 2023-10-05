package routesServer2

import "github.com/gin-gonic/gin"

func RoutesServer2() {
	r := gin.Default()

	r.Run(":8082")
}
