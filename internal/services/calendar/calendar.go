package calendar

import (
	"github.com/Aneg/calendar/internal/models"
	"github.com/Aneg/calendar/internal/repositories"
)

func NewCalendarService(repository repositories.CalendarRepository) *CalendarService {
	return &CalendarService{repository: repository}
}

type CalendarService struct {
	repository repositories.CalendarRepository
}

func (c CalendarService) AddEvent(event *models.Event) (err error) {
	event.Id, err = c.repository.AddEvent(event)
	return err
}

func (c CalendarService) DropEvent(uesrId int32, id int32) error {
	return c.repository.DropEvent(uesrId, id)
}

func (c CalendarService) EditEvent(event *models.Event) error {
	return c.repository.EditEvent(*event)
}

func (c CalendarService) GetEvent(userId int32, id int32) (models.Event, error) {
	return c.repository.GetEvent(userId, id)
}

func (c CalendarService) All(userId int32) ([]models.Event, error) {
	return c.repository.All(userId)
}
