package main

import (
	"blog-admin-service/conf"
	"blog-admin-service/server"
)

// @title QB blog Swagger API
// @version 1.0
// @description This is a sample cms server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host api.kyuuu.be
// @BasePath /v1

func main() {

	conf.Init()
	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
