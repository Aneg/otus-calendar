package repositories

import (
	"errors"
	"github.com/Aneg/calendar/internal/models"
	"sort"
	"time"
)

func NewCalendarMap() *CalendarMap {
	return &CalendarMap{
		eventMap:  make(map[int]map[time.Month]map[int]map[int]*models.Event, 10),
		eventList: make([]models.Event, 0, 10),
	}
}

type CalendarMap struct {
	eventMap  map[int]map[time.Month]map[int]map[int]*models.Event
	eventList []models.Event
}

func (c *CalendarMap) AddEvent(event *models.Event) error {
	if !c.TimeIsFree(event) {
		return errors.New("время занято")
	}
	dt := event.GetDateTime()
	for h := 0; h < event.Duration; h++ {
		if _, ok := c.eventMap[dt.Year()]; !ok {
			c.eventMap[dt.Year()] = map[time.Month]map[int]map[int]*models.Event{}
		}
		if _, ok := c.eventMap[dt.Year()][dt.Month()]; !ok {
			c.eventMap[dt.Year()][dt.Month()] = map[int]map[int]*models.Event{}
		}
		if _, ok := c.eventMap[dt.Year()][dt.Month()][dt.Day()]; !ok {
			c.eventMap[dt.Year()][dt.Month()][dt.Day()] = map[int]*models.Event{}
		}
		if _, ok := c.eventMap[dt.Year()][dt.Month()][dt.Day()][dt.Hour()]; !ok {
			c.eventMap[dt.Year()][dt.Month()][dt.Day()][dt.Hour()] = event
		}

		dt = dt.Add(time.Hour)
	}
	c.eventList = append(c.eventList, *event)
	return nil
}

func (c *CalendarMap) TimeIsFree(event *models.Event) bool {
	dt := event.GetDateTime()
	for h := 0; h < event.Duration; h++ {
		if _, ok := c.eventMap[dt.Year()][dt.Month()][dt.Day()][dt.Hour()]; ok {
			return false
		}
		dt.Add(time.Hour * time.Duration(1))
	}
	return true
}

func (c *CalendarMap) CanEdit(oldEvent *models.Event, newEvent *models.Event) bool {
	dt := newEvent.GetDateTime()
	for h := 0; h < newEvent.Duration; h++ {
		if eOld, ok := c.eventMap[dt.Year()][dt.Month()][dt.Day()][dt.Hour()+h]; ok && eOld.GetDescription() != oldEvent.GetDescription() {
			return false
		}
	}
	return true
}

func (c *CalendarMap) DropEvent(dt time.Time) error {
	if e, ok := c.eventMap[dt.Year()][dt.Month()][dt.Day()][dt.Hour()]; !ok {
		return errors.New("на данное время ничего не запланировано")
	} else {
		for h := 0; h < e.Duration; h++ {
			if _, ok := c.eventMap[dt.Year()][dt.Month()][dt.Day()][dt.Hour()]; ok {
				// todo: реализовать рекурсивное удаление пустых мап
				delete(c.eventMap[dt.Year()][dt.Month()][dt.Day()], dt.Hour())
				if len(c.eventMap[dt.Year()][dt.Month()][dt.Day()]) == 0 {
					delete(c.eventMap[dt.Year()][dt.Month()], dt.Day())
					if len(c.eventMap[dt.Year()][dt.Month()]) == 0 {
						delete(c.eventMap[dt.Year()], dt.Month())
						if len(c.eventMap[dt.Year()]) == 0 {
							delete(c.eventMap, dt.Year())
						}
					}
				}
			}
			dt = dt.Add(time.Hour)
		}
		for i, event := range c.eventList {
			if event.Duration == e.Duration && event.GetDateTime() == e.GetDateTime() && event.GetDescription() == e.GetDescription() {
				copy(c.eventList[i:], c.eventList[i+1:])
				c.eventList = c.eventList[:len(c.eventList)-1]
			}
		}
	}
	return nil
}

func (c *CalendarMap) EditEvent(dt time.Time, event *models.Event) error {
	// todo: внедрить мьютексы
	oldEvent, ok := c.eventMap[dt.Year()][dt.Month()][dt.Day()][dt.Hour()]
	if !ok {
		return errors.New("на данное время ничего не запланировано")
	}

	if c.CanEdit(oldEvent, event) {
		if err := c.DropEvent(dt); err != nil {
			return err
		}
		if err := c.AddEvent(event); err != nil {
			return err
		}

		return nil
	}
	return errors.New("время занято другими событиями")
}

func (c *CalendarMap) GetEvent(dt time.Time) (models.Event, error) {
	if e, ok := c.eventMap[dt.Year()][dt.Month()][dt.Day()][dt.Hour()]; !ok {
		return models.Event{}, errors.New("на данное время ничего не запланировано")
	} else {
		return *e, nil
	}
}

func (c *CalendarMap) All() []models.Event {
	sort.Slice(c.eventList, func(i, j int) bool {
		return c.eventList[j].GetDateTime().Sub(c.eventList[i].GetDateTime()).Hours() > 0
	})
	// единственная проблема - ссылки на события и значит его можно поменять
	return c.eventList
}
