package grpc

import (
	"context"
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
	err := s.Calendar.DropEvent(req.UserId, req.EventId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &calendar.SuccessResponse{Success: true}, nil
}

func (s *CalendarServer) EditEvent(ctx context.Context, req *calendar.EditEventRequest) (*calendar.EventResponse, error) {
	event := models.NewEvent(
		req.NewEvent.UserId,
		time.Unix(req.NewEvent.Datetimefrom.GetSeconds(), int64(req.NewEvent.Datetimefrom.GetNanos())),
		time.Unix(req.NewEvent.Datetimeto.GetSeconds(), int64(req.NewEvent.Datetimeto.GetNanos())),
		req.NewEvent.Description,
	)
	if err := s.Calendar.EditEvent(&event); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dtf, err := ptypes.TimestampProto(event.DateTimeFrom)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	dtt, err := ptypes.TimestampProto(event.DateTimeTo)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &calendar.EventResponse{Success: true, Event: &calendar.Event{
		Id:           event.Id,
		UserId:       event.UserId,
		Datetimefrom: dtf,
		Datetimeto:   dtt,
		Description:  event.Description,
	}}, nil
}

func (s *CalendarServer) GetEvent(ctx context.Context, req *calendar.GetEventRequest) (*calendar.EventResponse, error) {
	event, err := s.Calendar.GetEvent(req.UserId, req.EventId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	dtf, err := ptypes.TimestampProto(event.DateTimeFrom)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	dtt, err := ptypes.TimestampProto(event.DateTimeTo)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &calendar.EventResponse{Success: true, Event: &calendar.Event{
		Id:           event.Id,
		UserId:       event.UserId,
		Datetimefrom: dtf,
		Datetimeto:   dtt,
		Description:  event.Description,
	}}, nil
}

func (s *CalendarServer) AllEvents(ctx context.Context, req *calendar.AllEventsRequest) (*calendar.AllEventsResponse, error) {
	events, err := s.Calendar.All(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	responseEvents := make([]*calendar.Event, 0, len(events))
	for _, event := range events {
		dtf, err := ptypes.TimestampProto(event.DateTimeFrom)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		dtt, err := ptypes.TimestampProto(event.DateTimeTo)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		responseEvents = append(responseEvents, &calendar.Event{
			Id:           event.Id,
			UserId:       event.UserId,
			Datetimefrom: dtf,
			Datetimeto:   dtt,
			Description:  event.Description,
		})
	}
	return &calendar.AllEventsResponse{Events: responseEvents}, nil
}

func (s *CalendarServer) AddEvent(ctx context.Context, req *calendar.AddEventRequest) (*calendar.EventResponse, error) {
	// TODO: Проверить корректность конвертации времени
	event := models.NewEvent(req.Event.UserId, time.Unix(
		req.Event.Datetimefrom.GetSeconds(),
		int64(req.Event.Datetimefrom.GetNanos())),
		time.Unix(req.Event.Datetimeto.GetSeconds(),
			int64(req.Event.Datetimeto.GetNanos())),
		req.Event.Description,
	)
	err := s.Calendar.AddEvent(&event)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	req.Event.Id = event.Id
	return &calendar.EventResponse{Success: true, Event: req.Event}, nil
}
