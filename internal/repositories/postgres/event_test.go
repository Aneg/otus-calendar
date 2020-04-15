package postgres

import (
	"github.com/Aneg/calendar/internal/config"
	"github.com/Aneg/calendar/internal/models"
	"github.com/Aneg/calendar/pkg/db"
	"github.com/jmoiron/sqlx"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var pgConn *sqlx.DB

func init() {
	conf, err := config.GetConfigFromFile("../../../configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	if pgConn, err = db.PostgresOpenConnection(conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBName); err != nil {
		log.Fatal(err)
	}
}

func TestEventRepository_AddEvent(t *testing.T) {
	e := EventRepository{
		DB: pgConn,
	}
	uid := rand.Int31()

	events, err := e.All(uid)
	if err != nil {
		t.Error(err)
	}

	for _, event := range events {
		if err := e.DropEvent(event.UserId, event.Id); err != nil {
			t.Error(err)
		}
	}

	type args struct {
		event models.Event
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "1",
			args:    struct{ event models.Event }{event: models.NewEvent(uid, time.Now(), time.Now(), "test")},
			wantErr: false,
		},
		{
			name:    "2",
			args:    struct{ event models.Event }{event: models.NewEvent(uid, time.Now(), time.Now().AddDate(0, 0, -1), "test")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if id, err := e.AddEvent(&tt.args.event); (err != nil || id == 0) != tt.wantErr {
				t.Errorf("AddEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventRepository_EditEvent(t *testing.T) {
	type args struct {
		event models.Event
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "1",
			args:    struct{ event models.Event }{event: models.NewEvent(rand.Int31(), time.Now(), time.Now().AddDate(0, 0, -1), "test")},
			wantErr: false,
		},
		{
			name:    "2",
			args:    struct{ event models.Event }{event: models.NewEvent(rand.Int31(), time.Now(), time.Now().AddDate(0, 0, -1), "test")},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			e := EventRepository{DB: pgConn}

			uid := rand.Int31()

			events, err := e.All(uid)
			if err != nil {
				t.Error(err)
			}

			for _, event := range events {
				if err := e.DropEvent(event.UserId, event.Id); err != nil {
					t.Error(err)
				}
			}

			if !tt.wantErr {
				if tt.args.event.Id, err = e.AddEvent(&tt.args.event); err != nil {
					t.Error(err)
				}
			}

			tt.args.event.Description = "test 2"
			if err := e.EditEvent(tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("EditEvent() error = %v, wantErr %v", err, tt.wantErr)
			}

			if event, err := e.GetEvent(tt.args.event.UserId, tt.args.event.Id); (err != nil || event.Description != tt.args.event.Description) != tt.wantErr {
				t.Errorf("GetEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventRepository_All(t *testing.T) {
	e := EventRepository{
		DB: pgConn,
	}
	uid := rand.Int31()

	events, err := e.All(uid)
	if err != nil {
		t.Error(err)
	}

	for _, event := range events {
		if err := e.DropEvent(event.UserId, event.Id); err != nil {
			t.Error(err)
		}
	}

	for i := 0; i < 10; i++ {
		_, err := e.AddEvent(&models.Event{
			UserId:       uid,
			DateTimeFrom: time.Now().Add(time.Duration(i) * time.Hour),
			DateTimeTo:   time.Now().Add(time.Duration(i) * time.Hour),
			Description:  "test " + strconv.Itoa(i),
		})
		if err != nil {
			t.Error(err)
		}
	}

	events, err = e.All(uid)
	if err != nil {
		t.Error(err)
	}
	if len(events) != 10 {
		t.Error("не верное число событий для пользователя")
	}

}

func TestEventRepository_DropEvent(t *testing.T) {
	e := EventRepository{
		DB: pgConn,
	}
	uid := rand.Int31()

	events, err := e.All(uid)
	if err != nil {
		t.Error(err)
	}

	for _, event := range events {
		if err := e.DropEvent(event.UserId, event.Id); err != nil {
			t.Error(err)
		}
	}

	for i := 0; i < 10; i++ {
		_, err := e.AddEvent(&models.Event{
			UserId:       uid,
			DateTimeFrom: time.Now().Add(time.Duration(i) * time.Hour),
			DateTimeTo:   time.Now().Add(time.Duration(i) * time.Hour),
			Description:  "test " + strconv.Itoa(i),
		})
		if err != nil {
			t.Error(err)
		}
	}

	events, err = e.All(uid)
	if err != nil {
		t.Error(err)
	}
	if len(events) != 10 {
		t.Error("не верное число событий для пользователя (должно быть 10): ", len(events))
	}

	for _, event := range events {
		if err := e.DropEvent(event.UserId, event.Id); err != nil {
			t.Error(err)
		}
	}

	events, err = e.All(uid)
	if err != nil {
		t.Error(err)
	}
	if len(events) != 0 {
		t.Error("не верное число событий для пользователя (должно быть 0): ", len(events))
	}
}
