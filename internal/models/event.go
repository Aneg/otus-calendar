package models

import "time"

func NewEvent(DateTime time.Time, Duration int, Description string) Event {
	return Event{
		dateTime:    DateTime,
		Duration:    Duration,
		description: Description,
	}
}

type Event struct {
	dateTime    time.Time
	Duration    int
	description string
}

func (e *Event) GetDateTime() time.Time {
	return e.dateTime
}

func (e *Event) GetDuration() int {
	return e.Duration
}

func (e *Event) GetDescription() string {
	return e.description
}

func (e *Event) SetDateTimeAndDuration(dt time.Time, duration int) error {
	// можно ли переопределить время и продолжительность?
	// если да, то дропаем и добавляем его.
	return nil
}

func (e *Event) SetDescription(disc string) {
	e.description = disc
}
