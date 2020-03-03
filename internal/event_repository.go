package internal

import (
	"github.com/Aneg/calendar/internal/models"
	"time"
)

type CalendarRepository interface {
	AddEvent(event *models.Event) error
	DropEvent(dt time.Time) error
	EditEvent(dt time.Time, event *models.Event) error
	GetEvent(dt time.Time) (models.Event, error)
	All() []models.Event
}
