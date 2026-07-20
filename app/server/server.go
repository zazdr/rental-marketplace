package server

import (
	"app/config"

	"github.com/labstack/echo/v5"
)

type Server struct {
	Echo   *echo.Echo
	Config *config.Config
}

func New(config *config.Config) *Server {
	return &Server{
		Echo:   echo.New(),
		Config: config,
	}
}

func (s *Server) Start() error {
	return s.Echo.Start(s.Config.Server.Host + ":" + s.Config.Server.Port)
}
