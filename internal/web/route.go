package web

import (
	"github.com/Aneg/calendar/internal/config"
	"github.com/Aneg/calendar/internal/web/controllers"
	"github.com/julienschmidt/httprouter"
)

var Router *httprouter.Router

func Init(conf *config.Config) {
	ExampleController := controllers.NewExample(conf)

	Router = httprouter.New()
	{
		Router.GET("/hello", ExampleController.Hello)
	}
}
