package postgres

import (
	"errors"
	"github.com/Aneg/calendar/internal/models"
	"github.com/jmoiron/sqlx"
)

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{DB: db}
}

type EventRepository struct {
	DB *sqlx.DB
}

func (e EventRepository) AddEvent(event *models.Event) (int32, error) {
	if ok, err := e.CanAddEvent(event); err != nil {
		return 0, err
	} else if !ok {
		return 0, errors.New("событие на данное время уже существует")
	}

	_, err := e.DB.Exec(`INSERT INTO events (user_id, date_time_from, date_time_to, description) VALUES ($1,$2,$3,$4)`,
		event.UserId, event.DateTimeFrom, event.DateTimeTo, event.Description)

	if err != nil {
		return 0, err
	}

	newEvent := models.Event{}
	if err = e.DB.Get(&newEvent, `Select * from events where user_id = $1 and date_time_from = $2 and date_time_to = $3`, event.UserId, event.DateTimeFrom, event.DateTimeTo); err != nil {
		return 0, err
	}
	//error = LastInsertId is not supported by this driver, wantErr false
	//return r.LastInsertId()
	return newEvent.Id, err
}

func (e EventRepository) CanAddEvent(event *models.Event) (bool, error) {
	canInsert := struct {
		Count int `db:"c"`
	}{}
	err := e.DB.Get(&canInsert, "Select count(*) as c from events where user_id = $1 and id != $2 and ((date_time_from <= $3 and date_time_to >= $3) or (date_time_from <= $4 and date_time_to >= $4))",
		event.UserId, event.Id, event.DateTimeFrom, event.DateTimeTo)
	if err != nil {
		return false, err
	}
	return canInsert.Count == 0, err
}

func (e EventRepository) DropEvent(uesrId int32, id int32) error {
	_, err := e.DB.Exec(`DELETE FROM events WHERE events.id = $1 AND events.user_id = $2`, id, uesrId)
	return err
}

func (e EventRepository) EditEvent(event models.Event) error {
	if ok, err := e.CanAddEvent(&event); err != nil {
		return err
	} else if !ok {
		return errors.New("событие на данное время уже существует")
	}

	r, err := e.DB.Exec(`UPDATE events SET date_time_from=$1, date_time_to=$2, description=$3 WHERE user_id=$4 and id=$5`,
		event.DateTimeFrom, event.DateTimeTo, event.Description, event.UserId, event.Id)
	if err != nil {
		return err
	}

	if countUpdate, err := r.RowsAffected(); err != nil {
		return err
	} else if countUpdate == 0 {
		return errors.New("ни одна запись не была обновлена")
	}
	return nil
}

func (e EventRepository) GetEvent(userId int32, id int32) (event models.Event, err error) {
	err = e.DB.Get(&event, `Select * from events where user_id = $1 and id = $2`, userId, id)
	return
}

func (e EventRepository) All(userId int32) (events []models.Event, err error) {
	err = e.DB.Select(&events, `Select * from events where user_id = $1`, userId)
	return
}
