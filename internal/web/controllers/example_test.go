package controllers

import (
	"github.com/Aneg/calendar/internal"
	"github.com/Aneg/calendar/pkg/log"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExample_Index(t *testing.T) {

	log.Logger = zap.NewExample()
	conf := internal.Config{}

	c := NewExample(&conf)

	w := httptest.NewRecorder()
	r := http.Request{}
	c.Index(w, &r, []httprouter.Param{})

	if w.Code != 200 {
		t.Error("status is not ok: ", w.Code)
	}

	if result, _ := ioutil.ReadAll(w.Body); string(result) != "Hi" {
		t.Error("body is not `Hi`")
	}
}
