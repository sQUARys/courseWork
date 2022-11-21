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
	//r.Router.HandleFunc("/mail/{email}", r.Controller.MailHandler).Methods("Post")
	//
	r.Router.HandleFunc("/sorts", r.Controller.GetSorts).Methods("Get")
	r.Router.HandleFunc("/sorts", r.Controller.SendUserChoice).Methods("Post")

	//r.Router.HandleFunc("/mail/got-emails/{email}", r.Controller.GetMailsByEmail).Methods("Get")
	//r.Router.HandleFunc("/mail/message/{message-id:[0-9]+}", r.Controller.GetMailById).Methods("Get")
}
