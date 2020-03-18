package controllers

import (
	"github.com/Aneg/calendar/internal"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExample_Index(t *testing.T) {
	conf := internal.Config{
		HttpListen: "",
		LogFile:    "",
		LogLevel:   "",
	}

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
