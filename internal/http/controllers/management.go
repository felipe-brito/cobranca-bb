package controllers

import (
	"github.com/felipe-brito/cobranca-bb/internal/http/handlers"
	"github.com/felipe-brito/cobranca-bb/internal/http/model/response"
	"github.com/felipe-brito/cobranca-bb/internal/services"
	"github.com/felipe-brito/cobranca-bb/internal/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Management struct {
	responseHandlers handlers.ResponseHandler
	manager          services.Manager
}

func NewManagement(manager services.Manager, response handlers.ResponseHandler) Management {
	return Management{responseHandlers: response, manager: manager}
}

func (m Management) Router(r chi.Router) {
	r.Get("/health", m.health)

}
func (m Management) Path() string {
	return "/manager"
}

// @Title			checkHealth
// @Tags        	Internal
// @Summary      	Check health
// @Description 	Check the API, and it's components health
// @Description 	**Errors http codes response**
// @Description 	HTTP | Description | Code | Note
// @Description 	----- | ----- | ----- | -----
// @Description 	500 | Internal server error | 500 | N/A
// @Success			200		{object}	response.ManagerHealth
// @Failure			500		"Internal Server Error"
// @Accept       	json
// @Produce      	json
// @Router       	/manager/health [get]
func (m Management) health(w http.ResponseWriter, r *http.Request) {

	health := m.manager.Health()

	if health.Status != utils.UP {
		// TODO internal error
		m.responseHandlers.InternalServerError(w, r)
		return
	}

	m.responseHandlers.OK(w, response.ConvertTo(health))
}
