package grpc

import (
	"context"
	"fmt"
	"github.com/Aneg/calendar/internal/models"
	calendar2 "github.com/Aneg/calendar/internal/services/calendar"
	calendar "github.com/Aneg/calendar/pkg/api"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type CalendarServer struct {
	Calendar *calendar2.CalendarService
}

func (s *CalendarServer) DropEvent(ctx context.Context, req *calendar.DropEventRequest) (*calendar.SuccessResponse, error) {
	dt := time.Unix(req.Datetime.GetSeconds(), int64(req.Datetime.GetNanos()))
	fmt.Println(dt.String(), dt, dt.Second(), dt.Nanosecond())
	err := s.Calendar.DropEvent(uint(req.UserId), dt)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &calendar.SuccessResponse{Success: true}, nil
}

func (s *CalendarServer) EditEvent(ctx context.Context, req *calendar.EditEventRequest) (*calendar.SuccessResponse, error) {
	event := models.NewEvent(uint(req.NewEvent.UserId), time.Unix(req.NewEvent.Datetime.GetSeconds(), int64(req.NewEvent.Datetime.GetNanos())), int(req.NewEvent.Duration), req.NewEvent.Description)
	err := s.Calendar.EditEvent(time.Unix(req.OldDatetime.GetSeconds(), int64(req.OldDatetime.GetNanos())), &event)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &calendar.SuccessResponse{Success: true}, nil
}

func (s *CalendarServer) GetEvent(ctx context.Context, req *calendar.GetEventRequest) (*calendar.GetEventResponse, error) {
	event, err := s.Calendar.GetEvent(uint(req.UserId), time.Unix(req.Datetime.GetSeconds(), int64(req.Datetime.GetNanos())))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	dt, err := ptypes.TimestampProto(event.GetDateTime())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &calendar.GetEventResponse{Event: &calendar.Event{
		UserId:      uint32(event.UserId),
		Duration:    int32(event.GetDuration()),
		Datetime:    dt,
		Description: event.GetDescription(),
	}}, nil
}

func (s *CalendarServer) AllEvents(ctx context.Context, req *calendar.AllEventsRequest) (*calendar.AllEventsResponse, error) {
	// TODO: Простоестить поведение при запросе для неизвестного пользователя
	events := s.Calendar.All(uint(req.UserId))
	responseEvents := make([]*calendar.Event, 0, len(events))
	for _, event := range events {
		dt, err := ptypes.TimestampProto(event.GetDateTime())
		if err != nil {
			// Не уверен, что хорошее решение фэйлить запрос если одна дата кривая
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		responseEvents = append(responseEvents, &calendar.Event{
			UserId:      uint32(event.UserId),
			Duration:    int32(event.GetDuration()),
			Datetime:    dt,
			Description: event.GetDescription(),
		})
	}
	return &calendar.AllEventsResponse{Events: responseEvents}, nil
}

func (s *CalendarServer) AddEvent(ctx context.Context, req *calendar.AddEventRequest) (*calendar.SuccessResponse, error) {
	// TODO: Проверить корректность конвертации времени
	event := models.NewEvent(uint(req.Event.UserId), time.Unix(req.Event.Datetime.GetSeconds(), int64(req.Event.Datetime.GetNanos())), int(req.Event.Duration), req.Event.Description)
	err := s.Calendar.AddEvent(&event)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &calendar.SuccessResponse{Success: true}, nil
}
