package router

import (
	"database/sql"
	"net/http"

	_ "github.com/Meruya-Technology/go-boilerplate/docs"
	cfg "github.com/Meruya-Technology/go-boilerplate/lib/common/config"
	ctr "github.com/Meruya-Technology/go-boilerplate/lib/presentation/controllers"
	mdr "github.com/Meruya-Technology/go-boilerplate/lib/presentation/middlewares"
	ech "github.com/labstack/echo/v4"
	ecs "github.com/swaggo/echo-swagger"
)

type RouteHandler struct {
	Config   cfg.Config
	Database sql.DB
}

func (r RouteHandler) Handle() *ech.Echo {
	echoServer := ech.New()

	/// Health check
	echoServer.GET("/healtz", func(c ech.Context) error {
		return c.String(http.StatusOK, "Server is Healthy")
	})

	/// Swagger
	echoServer.GET("/swagger/*", ecs.WrapHandler)

	/// V1
	routerV1 := echoServer.Group("/api/auth")

	//  PATH /client/create
	clientController := ctr.ClientController{Config: r.Config, Database: r.Database}
	clientMiddleware := mdr.ClientMiddleware{Config: r.Config, Database: r.Database, ClientController: clientController}
	routerV1.POST("/client/create", clientController.Create, clientMiddleware.CheckClient)
	routerV1.POST("/client/check", clientController.Check)

	//  PATH /client/create
	UserController := ctr.UserController{Config: r.Config, Database: r.Database}
	routerV1.POST("/user/login", UserController.Login)
	routerV1.POST("/user/register", UserController.Register)

	//  PATH /access/refresh
	AccessController := ctr.AccessController{Config: r.Config, Database: r.Database}
	routerV1.POST("/access/refresh", AccessController.Refresh)

	return echoServer
}
