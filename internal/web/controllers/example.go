package controllers

import (
	"github.com/Aneg/calendar/internal/config"
	"github.com/Aneg/calendar/pkg/log"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewExample(c *config.Config) *Example {
	return &Example{c: c}
}

type Example struct {
	c *config.Config
}

func (c *Example) Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Logger.Info("Request: Example.Hello")
	if _, err := w.Write([]byte("Hi")); err != nil {
		log.Logger.Error(err.Error())
	}
}
