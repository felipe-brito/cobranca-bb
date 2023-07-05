package http

import (
	"fmt"
	"github.com/felipe-brito/cobranca-bb/internal/http/swagger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/cors"
	"net/http"
)

type ChiRouter struct {
	router  chi.Router
	routers []Router
	manager []Router
}

func NewRouter() ChiRouter {
	cRouter := ChiRouter{
		router: chi.NewRouter(),
	}

	cRouter.init()
	cRouter.configuration()

	return cRouter
}

func (c *ChiRouter) GetRouter() chi.Router {
	return c.router
}

func (c *ChiRouter) cors() func(next http.Handler) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposedHeaders:     []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type"},
		MaxAge:             300,
		AllowCredentials:   true,
		OptionsPassthrough: false,
	}).Handler
}

func (c *ChiRouter) init() {
	c.routers = []Router{
		GetComponents().Ticket,
	}

	c.manager = []Router{
		GetComponents().Management,
	}
}

func (c *ChiRouter) configuration() {

	c.router.Use(c.cors())
	c.router.Route("/api/v1", func(r chi.Router) {
		// TODO fazer um middleware de recover
		r.Group(func(r chi.Router) {
			r.Use(render.SetContentType(render.ContentTypeJSON))
			for _, v := range c.routers {
				r.Route(v.Path(), v.Router)
			}
		})
	})

	c.router.Group(func(r chi.Router) {
		for _, v := range c.manager {
			r.Route(v.Path(), v.Router)
			fmt.Println(v.Path())
		}
	})

	swag := swagger.Swagger{}

	c.router.Get("/swagger/*", swag.Config())
	c.router.Get("/swagger.json", swag.Read())

}
