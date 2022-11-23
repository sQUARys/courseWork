package routers

import (
	"courseWork/app/controller"
	"github.com/gorilla/mux"
)

type Router struct {
	Router     *mux.Router
	Controller controller.Controller
}

func New(controller *controller.Controller) *Router {
	r := mux.NewRouter()
	return &Router{
		Controller: *controller,
		Router:     r,
	}
}

func (r *Router) SetRoutes() {
	r.Router.HandleFunc("/sorts", r.Controller.GetSorts).Methods("Get")
	r.Router.HandleFunc("/sorts", r.Controller.SendUserChoice).Methods("Post")
}
