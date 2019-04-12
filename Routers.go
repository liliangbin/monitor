/**
  * @author liliangbin  dumpling1520@gmail.com
  * @date 2019/4/12  15:03
  *
  **/

package main

import (
	"github.com/gorilla/mux"
	"monitor/model"
	"monitor/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}
	return router
}

var routes = model.Routes{
	model.Router{
		"Index",
		"GET",
		"/",
		handlers.TestInfo,
	},
}
