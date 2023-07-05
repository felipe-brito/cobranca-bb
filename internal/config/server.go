package config

import (
	"context"
	"fmt"
	"github.com/felipe-brito/cobranca-bb/internal/global"
	router "github.com/felipe-brito/cobranca-bb/internal/http"
	"net/http"
)

type Server struct {
}

func (s Server) Init(_ context.Context) error {

	r := router.NewRouter()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", global.AppConfig.Server.Port),
		Handler: r.GetRouter(),
	}

	return server.ListenAndServe()
}

func (s Server) Stop(_ context.Context) error {
	return nil
}
