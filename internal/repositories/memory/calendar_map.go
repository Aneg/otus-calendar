package memory

import (
	"errors"
	"github.com/Aneg/calendar/internal/models"
	"math/rand"
)

func NewCalendarMap() *CalendarMap {
	return &CalendarMap{
		events: make(map[int32]map[int32]models.Event),
	}
}

type CalendarMap struct {
	events map[int32]map[int32]models.Event
}

func (c *CalendarMap) AddEvent(event *models.Event) (int32, error) {
	for _, e := range c.events[event.UserId] {
		if c.TimeIsNotFree(*event, e) {
			return 0, errors.New("время занято")
		}
	}
	event.Id = rand.Int31()
	if _, ok := c.events[event.UserId]; !ok {
		c.events[event.UserId] = make(map[int32]models.Event)
	}
	c.events[event.UserId][event.Id] = *event

	return event.Id, nil
}

func (c *CalendarMap) TimeIsNotFree(newEvent models.Event, oldEvent models.Event) bool {
	return (newEvent.DateTimeFrom.After(oldEvent.DateTimeFrom) && newEvent.DateTimeTo.Before(oldEvent.DateTimeFrom)) ||
		newEvent.DateTimeFrom.After(oldEvent.DateTimeTo) && newEvent.DateTimeTo.Before(oldEvent.DateTimeTo) ||
		(oldEvent.DateTimeFrom.After(newEvent.DateTimeFrom) && oldEvent.DateTimeTo.Before(newEvent.DateTimeFrom)) ||
		oldEvent.DateTimeFrom.After(newEvent.DateTimeTo) && oldEvent.DateTimeTo.Before(newEvent.DateTimeTo)
}

func (c *CalendarMap) DropEvent(uesrId int32, id int32) error {
	delete(c.events[uesrId], id)
	return nil
}

func (c *CalendarMap) EditEvent(event models.Event) error {
	for _, e := range c.events[event.UserId] {
		if e.Id != event.Id && c.TimeIsNotFree(event, e) {
			return errors.New("время занято другими событиями")
		}
	}
	if _, ok := c.events[event.UserId]; !ok {
		c.events[event.UserId] = make(map[int32]models.Event)
	}
	c.events[event.UserId][event.Id] = event
	return nil
}

func (c *CalendarMap) GetEvent(userId int32, id int32) (models.Event, error) {
	if e, ok := c.events[userId][id]; !ok {
		return e, errors.New("событие не найдено")
	} else {
		return e, nil
	}
}

func (c *CalendarMap) All(userId int32) (events []models.Event, err error) {
	if _, ok := c.events[userId]; !ok {
		return events, errors.New("у данного пользователя нету событий")
	}
	for _, e := range c.events[userId] {
		events = append(events, e)
	}
	return
}
