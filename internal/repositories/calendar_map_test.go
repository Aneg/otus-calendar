package repositories

import (
	"github.com/Aneg/calendar/internal/models"
	"testing"
	"time"
)

func TestCalendarMap_AddEvent(t *testing.T) {
	c := NewCalendarMap()

	dt := time.Now()
	d := time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
	var e = models.NewEvent(d, 10, "test event")
	if err := c.AddEvent(&e); err != nil {
		t.Error(err)
	}
	for i := 0; i < e.Duration; i++ {
		if _, ok := c.eventMap[d.Year()][d.Month()][d.Day()][d.Hour()+i]; !ok {
			t.Error("Не найден добавленый эллемент")
		}
	}

	if _, ok := c.eventMap[d.Year()][d.Month()][d.Day()][d.Hour()+10]; ok {
		t.Error("Найдено неожиданный событие")
	}

	d = time.Date(dt.Year(), dt.Month(), dt.Day(), 2, 0, 0, 0, dt.Location())
	e = models.NewEvent(d, 5, "test event")
	if err := c.AddEvent(&e); err == nil {
		t.Error("не хватает ошибки")
	}

	for i := 0; i <= 5; i++ {
		if _, ok := c.eventMap[d.Year()][d.Month()][d.Day()][d.Day()+i]; !ok {
			t.Error()
		}
	}
}

func TestCalendarMap_DropEvent(t *testing.T) {
	c := NewCalendarMap()

	dt := time.Now()
	d := time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
	var e = models.NewEvent(d, 10, "test event")
	if err := c.AddEvent(&e); err != nil {
		t.Error(err)
	}
	for i := 0; i < e.Duration; i++ {
		if _, ok := c.eventMap[d.Year()][d.Month()][d.Day()][d.Hour()+i]; !ok {
			t.Error("Не найден добавленый эллемент")
		}
	}

	d = time.Date(dt.Year(), dt.Month(), dt.Day(), 2, 0, 0, 0, dt.Location())
	e = models.NewEvent(d, 5, "test event")
	if err := c.AddEvent(&e); err == nil {
		t.Error("не хватает ошибки")
	}
	if err := c.DropEvent(e.GetDateTime()); err != nil {
		t.Error(err)
	}

	for i := 0; i < e.Duration; i++ {
		if _, ok := c.eventMap[d.Year()][d.Month()][d.Day()][d.Hour()+i]; ok {
			t.Error("Найден удалённый эллемент")
		}
	}
}

func TestCalendarMap_EditEvent(t *testing.T) {
	c := NewCalendarMap()

	dt := time.Now()
	d := time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
	var e = models.NewEvent(d, 10, "test event")
	if err := c.AddEvent(&e); err != nil {
		t.Error(err)
	}
	for i := 0; i < e.Duration; i++ {
		if _, ok := c.eventMap[d.Year()][d.Month()][d.Day()][d.Hour()+i]; !ok {
			t.Error("Не найден добавленый эллемент")
		}
	}
	if _, ok := c.eventMap[d.Year()][d.Month()][d.Day()][d.Hour()+10]; ok {
		t.Error("Найдено неожиданный событие")
	}

	dNew := time.Date(dt.Year(), dt.Month(), dt.Day(), 1, 0, 0, 0, dt.Location())
	var eNew = models.NewEvent(dNew, 11, "test event")

	if err := c.EditEvent(d, &eNew); err != nil {
		t.Error(err)
	}

	if _, ok := c.eventMap[d.Year()][d.Month()][d.Day()][0]; ok {
		t.Error("Найдено неожиданный событие")
	}

	for i := 0; i < eNew.Duration; i++ {
		if _, ok := c.eventMap[dNew.Year()][dNew.Month()][dNew.Day()][dNew.Hour()+i]; !ok {
			t.Error("Не найден добавленый эллемент")
		}
	}

	if _, ok := c.eventMap[d.Year()][d.Month()][d.Day()][eNew.Duration+1]; ok {
		t.Error("Найдено неожиданный событие")
	}
}

func TestCalendarMap_EditEvent_Error(t *testing.T) {
	c := NewCalendarMap()

	dt := time.Now()
	d := time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
	var e = models.NewEvent(d, 10, "test event")
	if err := c.AddEvent(&e); err != nil {
		t.Error(err)
	}

	d2 := time.Date(dt.Year(), dt.Month(), dt.Day(), 11, 0, 0, 0, dt.Location())
	var e2 = models.NewEvent(d2, 10, "test event 2")
	if err := c.AddEvent(&e2); err != nil {
		t.Error(err)
	}

	dNew := time.Date(dt.Year(), dt.Month(), dt.Day(), 3, 0, 0, 0, dt.Location())
	var eNew = models.NewEvent(dNew, 11, "test event 3")

	if err := c.EditEvent(d, &eNew); err == nil {
		t.Error("Нехватает ошибки")
	}
}

func TestCalendarMap_GetEvent(t *testing.T) {
	c := NewCalendarMap()

	dt := time.Now()
	d := time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
	var e = models.NewEvent(d, 2, "test event")
	if err := c.AddEvent(&e); err != nil {
		t.Error(err)
	}

	d2 := time.Date(dt.Year(), dt.Month(), dt.Day(), 3, 0, 0, 0, dt.Location())
	e2 := models.NewEvent(d2, 3, "test event 2")
	if err := c.AddEvent(&e2); err != nil {
		t.Error(err)
	}

	d = time.Date(dt.Year(), dt.Month(), dt.Day(), 4, 0, 0, 0, dt.Location())
	event, err := c.GetEvent(d)
	if err != nil {
		t.Error(err)
	}
	if event.GetDescription() != "test event 2" {
		t.Error("не верное событие")
	}
}

func TestCalendarMap_All(t *testing.T) {
	c := NewCalendarMap()

	dt := time.Now()
	d := time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
	var e = models.NewEvent(d, 2, "test event")
	if err := c.AddEvent(&e); err != nil {
		t.Error(err)
	}

	d2 := time.Date(dt.Year(), dt.Month(), dt.Day(), 3, 0, 0, 0, dt.Location())
	e2 := models.NewEvent(d2, 3, "test event 2")
	if err := c.AddEvent(&e2); err != nil {
		t.Error(err)
	}

	if len(c.All()) != 2 {
		t.Error("в репозитории не все события")
	}
}
