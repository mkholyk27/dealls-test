package routes

import (
	"dating-app/controllers"
	"dating-app/middlewares"
	"dating-app/repositories"

	"github.com/labstack/echo/v4"
)

type Route struct {
	*echo.Echo
}

func New(srv *echo.Echo) Route {
	return Route{srv}
}

func (route *Route) GenerateRoutes() {
	// init repositories
	loginRepository := repositories.NewLoginRepository()
	signupRepository := repositories.NewSignupRepository()

	// init controllers
	loginController := controllers.NewLoginController(loginRepository)
	homeController := controllers.NewHomeController()
	signupController := controllers.NewSignupController(signupRepository)

	// middlewares
	route.Use(middlewares.MiddlewareJWTAuthorization)

	// routes
	route.POST("/login", loginController.Login)
	route.POST("/register", signupController.Register)
	route.GET("/", homeController.Index)
}
