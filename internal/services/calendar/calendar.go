package calendar

import (
	"github.com/Aneg/calendar/internal/models"
	"github.com/Aneg/calendar/internal/repositories"
	"time"
)

func NewCalendarService(repository repositories.CalendarRepository) *CalendarService {
	return &CalendarService{repository: repository}
}

type CalendarService struct {
	repository repositories.CalendarRepository
}

func (c CalendarService) AddEvent(event *models.Event) error {
	return c.repository.AddEvent(event)
}

func (c CalendarService) DropEvent(uesrId uint, dt time.Time) error {
	return c.repository.DropEvent(uesrId, dt)
}

func (c CalendarService) EditEvent(dt time.Time, event *models.Event) error {
	return c.repository.EditEvent(dt, event)
}

func (c CalendarService) GetEvent(userId uint, dt time.Time) (models.Event, error) {
	return c.repository.GetEvent(userId, dt)
}

func (c CalendarService) All(userId uint) []models.Event {
	return c.repository.All(userId)
}
