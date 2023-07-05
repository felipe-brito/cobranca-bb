package http

import "github.com/go-chi/chi/v5"

type Router interface {
	Router(r chi.Router)
	Path() string
}

type Component struct {
	Management Router
	Ticket     Router
}

func GetComponents() Component {
	return Component{
		Management: GetInstance().ManagerController(),
		Ticket:     GetInstance().ManagerController(),
	}
}
