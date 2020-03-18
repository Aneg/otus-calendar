package web

import (
	"github.com/Aneg/calendar/internal"
	"github.com/Aneg/calendar/internal/web/controllers"
	"github.com/julienschmidt/httprouter"
)

var Router *httprouter.Router

func Init(conf *internal.Config) {
	ExampleController := controllers.NewExample(conf)

	Router = httprouter.New()
	{
		Router.GET("/example", ExampleController.Index)
	}
}
