package repositories

import (
	"github.com/Aneg/calendar/internal/models"
)

type CalendarRepository interface {
	AddEvent(event *models.Event) (int32, error)
	DropEvent(uesrId int32, id int32) error
	EditEvent(event models.Event) error
	GetEvent(userId int32, id int32) (event models.Event, err error)
	All(userId int32) (events []models.Event, err error)
}
