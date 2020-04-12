package postgres

import (
	"github.com/Aneg/calendar/internal/models"
	"github.com/jmoiron/sqlx"
	"time"
)

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{DB: db}
}

type EventRepository struct {
	DB *sqlx.DB
}

func (e EventRepository) AddEvent(event *models.Event) error {
	panic("implement me")
}

func (e EventRepository) DropEvent(uesrId uint, dt time.Time) error {
	panic("implement me")
}

func (e EventRepository) EditEvent(dt time.Time, event *models.Event) error {
	panic("implement me")
}

func (e EventRepository) GetEvent(userId uint, dt time.Time) (models.Event, error) {
	panic("implement me")
}

func (e EventRepository) All(userId uint) []models.Event {
	panic("implement me")
}
