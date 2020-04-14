package models

import "time"

func NewEvent(UserId int32, DateTimeFrom time.Time, DateTimeTo time.Time, Description string) Event {
	return Event{
		UserId:       UserId,
		DateTimeFrom: DateTimeFrom,
		DateTimeTo:   DateTimeTo,
		Description:  Description,
	}
}

type Event struct {
	Id           int32     `db:"id"`
	UserId       int32     `db:"user_id"`
	DateTimeFrom time.Time `db:"date_time_from"`
	DateTimeTo   time.Time `db:"date_time_to"`
	Description  string    `db:"description"`
}
