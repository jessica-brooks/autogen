package controllers

import "github.com/future/automate-similar-text-generation/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/titles", middlewares.SetMiddlewareJSON(s.GetTitles)).Methods("GET")
	s.Router.HandleFunc("/feedback", middlewares.SetMiddlewareJSON(s.SendFeedback)).Methods("POST")
}
