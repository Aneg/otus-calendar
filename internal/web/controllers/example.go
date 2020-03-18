package controllers

import (
	"github.com/Aneg/calendar/internal"
	"github.com/Aneg/calendar/pkg/log"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewExample(c *internal.Config) *Example {
	return &Example{c: c}
}

type Example struct {
	c *internal.Config
}

func (c *Example) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if _, err := w.Write([]byte("Hi")); err != nil {
		log.Logger.Error(err.Error())
	}
}
