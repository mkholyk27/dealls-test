package configs

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var once sync.Once

type Server struct {
	conf *Configs
}

var server *echo.Echo

func NewServer(conf *Configs) Server {
	return Server{
		conf: conf,
	}
}

func (srv *Server) InitServer() *echo.Echo {
	once.Do(func() {
		server = echo.New()
	})

	return server
}

func (srv *Server) Start() {
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	server.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	server.Logger.Fatal(server.Start(":" + srv.conf.Get("PORT_SERVER")))
}
