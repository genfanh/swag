package main

import (
	"github.com/genfanh/swag/example/object-map-example/controller"
	_ "github.com/genfanh/swag/example/object-map-example/docs"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/genfanh/files"
	ginSwagger "github.com/genfanh/gin-swagger"
)

//	@title			Swagger Map Example API
//	@version		1.0
//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

func main() {
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		test := v1.Group("/map")
		{
			test.GET("", c.GetMap)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
