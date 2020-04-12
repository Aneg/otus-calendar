package postgres

import (
	"github.com/Aneg/calendar/internal/models"
	"github.com/jmoiron/sqlx"
	"reflect"
	"testing"
	"time"
)

func TestEventRepository_AddEvent(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		event *models.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EventRepository{
				DB: tt.fields.DB,
			}
			if err := e.AddEvent(tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("AddEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventRepository_All(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		userId uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []models.Event
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EventRepository{
				DB: tt.fields.DB,
			}
			if got := e.All(tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEventRepository_DropEvent(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		uesrId uint
		dt     time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EventRepository{
				DB: tt.fields.DB,
			}
			if err := e.DropEvent(tt.args.uesrId, tt.args.dt); (err != nil) != tt.wantErr {
				t.Errorf("DropEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventRepository_EditEvent(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		dt    time.Time
		event *models.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EventRepository{
				DB: tt.fields.DB,
			}
			if err := e.EditEvent(tt.args.dt, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("EditEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEventRepository_GetEvent(t *testing.T) {
	type fields struct {
		DB *sqlx.DB
	}
	type args struct {
		userId uint
		dt     time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Event
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := EventRepository{
				DB: tt.fields.DB,
			}
			got, err := e.GetEvent(tt.args.userId, tt.args.dt)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEventRepository(t *testing.T) {
	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want *EventRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEventRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEventRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
