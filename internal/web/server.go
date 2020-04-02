package web

import (
	"context"
	"github.com/Aneg/calendar/internal"
	"github.com/Aneg/calendar/internal/models"
	calendar "github.com/Aneg/calendar/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type CalendarServer struct {
	CalendarRepository internal.CalendarRepository
}

func (s *CalendarServer) DropEvent(ctx context.Context, req *calendar.DropEventRequest) (*calendar.SuccessResponse, error) {
	err := s.CalendarRepository.DropEvent(uint(req.UserId), time.Unix(req.Datetime.GetSeconds(), int64(req.Datetime.GetNanos())))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &calendar.SuccessResponse{Success: true}, nil
}

func (s *CalendarServer) EditEvent(ctx context.Context, req *calendar.EditEventRequest) (*calendar.SuccessResponse, error) {
	event := models.NewEvent(uint(req.NewEvent.UserId), time.Unix(req.NewEvent.Datetime.GetSeconds(), int64(req.NewEvent.Datetime.GetNanos())), int(req.NewEvent.Duration), req.NewEvent.Description)
	err := s.CalendarRepository.EditEvent(time.Unix(req.OldDatetime.GetSeconds(), int64(req.OldDatetime.GetNanos())), &event)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &calendar.SuccessResponse{Success: true}, nil
}

func (s *CalendarServer) GetEvent(ctx context.Context, req *calendar.GetEventRequest) (*calendar.GetEventResponse, error) {
	event, err := s.CalendarRepository.GetEvent(uint(req.UserId), time.Unix(req.Datetime.GetSeconds(), int64(req.Datetime.GetNanos())))
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &calendar.GetEventResponse{Event: &calendar.Event{
		UserId:   uint32(event.UserId),
		Duration: int32(event.GetDuration()),
		Datetime: &timestamp.Timestamp{
			Seconds: int64(event.GetDateTime().Second()),
			Nanos:   int32(event.GetDateTime().Nanosecond()),
		},
		Description: event.GetDescription(),
	}}, nil
}

func (s *CalendarServer) AllEvent(ctx context.Context, req *calendar.AllRequest) (*calendar.AllResponse, error) {
	// TODO: Простоестить поведение при запросе для неизвестного пользователя
	events := s.CalendarRepository.All(uint(req.UserId))
	responseEvents := make([]*calendar.Event, 0, len(events))
	for _, event := range events {
		responseEvents = append(responseEvents, &calendar.Event{
			UserId:   uint32(event.UserId),
			Duration: int32(event.GetDuration()),
			Datetime: &timestamp.Timestamp{
				Seconds: int64(event.GetDateTime().Second()),
				Nanos:   int32(event.GetDateTime().Nanosecond()),
			},
			Description: event.GetDescription(),
		})
	}
	return &calendar.AllResponse{Events: responseEvents}, nil
}

func (s *CalendarServer) AddEvent(ctx context.Context, req *calendar.AddEventRequest) (*calendar.SuccessResponse, error) {
	// TODO: Проверить корректность конвертации времени
	event := models.NewEvent(uint(req.Event.UserId), time.Unix(req.Event.Datetime.GetSeconds(), int64(req.Event.Datetime.GetNanos())), int(req.Event.Duration), req.Event.Description)
	err := s.CalendarRepository.AddEvent(&event)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &calendar.SuccessResponse{Success: true}, nil
}
