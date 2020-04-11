package grpc

import (
	"context"
	"fmt"
	"github.com/Aneg/calendar/internal/models"
	"github.com/Aneg/calendar/internal/repositories/memory"
	calendar2 "github.com/Aneg/calendar/internal/services/calendar"
	calendar "github.com/Aneg/calendar/pkg/api"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"testing"
	"time"
)

func getListener() net.Listener {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	return lis
}

func TestCalendarServer_AddEvent(t *testing.T) {
	service := calendar2.NewCalendarService(memory.NewCalendarMap())
	server := initServer(&CalendarServer{Calendar: service})
	go server.Serve(getListener())
	defer server.Stop()

	client, cc := initClient()
	defer cc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	grade, err := client.AddEvent(ctx, &calendar.AddEventRequest{Event: &calendar.Event{
		UserId:      1,
		Duration:    3,
		Datetime:    ptypes.TimestampNow(),
		Description: "test",
	}})
	if err != nil {
		handlerError(err, t)
	} else {
		if !grade.Success {
			t.Error("Success is false")
		}
	}

	if calendarLen := len(service.All(1)); calendarLen != 1 {
		t.Error("calendar len ", calendarLen)
	}
}

func TestCalendarServer_AllEvent(t *testing.T) {
	calendarRepository := memory.NewCalendarMap()
	e := models.NewEvent(1, time.Now(), 2, "test 1")
	calendarRepository.AddEvent(&e)
	e = models.NewEvent(1, time.Now().AddDate(0, 0, 2), 2, "test 2")
	calendarRepository.AddEvent(&e)
	e = models.NewEvent(2, time.Now(), 2, "test 3")
	calendarRepository.AddEvent(&e)

	service := calendar2.NewCalendarService(calendarRepository)
	server := initServer(&CalendarServer{Calendar: service})
	go server.Serve(getListener())
	defer server.Stop()

	client, cc := initClient()
	defer cc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	grade, err := client.AllEvents(ctx, &calendar.AllEventsRequest{UserId: 1})
	if err != nil {
		handlerError(err, t)
	} else {
		if len(grade.Events) != 2 {
			t.Error("events count not equal 2")
		}
	}

	if calendarLen := len(calendarRepository.All(1)); calendarLen != 2 {
		t.Error("calendar len ", calendarLen)
	}
}

func TestCalendarServer_DropEvent(t *testing.T) {
	calendarRepository := memory.NewCalendarMap()
	eventForDrop := models.NewEvent(1, time.Now(), 2, "test 1")
	calendarRepository.AddEvent(&eventForDrop)
	e := models.NewEvent(1, time.Now().AddDate(0, 0, 2), 2, "test 2")
	calendarRepository.AddEvent(&e)
	e = models.NewEvent(2, time.Now(), 2, "test 3")
	calendarRepository.AddEvent(&e)

	service := calendar2.NewCalendarService(calendarRepository)
	server := initServer(&CalendarServer{Calendar: service})
	go server.Serve(getListener())
	defer server.Stop()

	client, cc := initClient()
	defer cc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()
	dt, _ := ptypes.TimestampProto(eventForDrop.GetDateTime())
	grade, err := client.DropEvent(ctx, &calendar.DropEventRequest{UserId: 1, Datetime: dt})
	if err != nil {
		handlerError(err, t)
	} else {
		if !grade.Success {
			t.Error("events count not equal 2")
		}
	}

	if calendarLen := len(calendarRepository.All(1)); calendarLen != 1 {
		t.Error("calendar len ", calendarLen)
	}
	if e = calendarRepository.All(1)[0]; e.GetDateTime() == eventForDrop.GetDateTime() {
		t.Error("event not drop")
	}
}

func TestCalendarServer_EditEvent(t *testing.T) {
	calendarRepository := memory.NewCalendarMap()
	eventForEdit := models.NewEvent(1, time.Now(), 2, "test 1")
	calendarRepository.AddEvent(&eventForEdit)
	e := models.NewEvent(1, time.Now().AddDate(0, 0, 2), 2, "test 2")
	calendarRepository.AddEvent(&e)
	e = models.NewEvent(2, time.Now(), 2, "test 3")
	calendarRepository.AddEvent(&e)
	service := calendar2.NewCalendarService(calendarRepository)
	server := initServer(&CalendarServer{Calendar: service})

	go server.Serve(getListener())
	defer server.Stop()

	client, cc := initClient()
	defer cc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()
	dt, _ := ptypes.TimestampProto(eventForEdit.GetDateTime())
	newEvent := calendar.Event{
		UserId:      1,
		Duration:    5,
		Datetime:    dt,
		Description: "new test 2",
	}
	grade, err := client.EditEvent(ctx, &calendar.EditEventRequest{OldDatetime: dt, NewEvent: &newEvent})
	if err != nil {
		handlerError(err, t)
	} else {
		if !grade.Success {
			t.Error("events not edit")
		}
	}

	curEvent, _ := calendarRepository.GetEvent(1, eventForEdit.GetDateTime())
	if curEvent.GetDuration() != int(newEvent.GetDuration()) || curEvent.GetDescription() != newEvent.GetDescription() {
		t.Error("events not edit")
	}
}

func TestCalendarServer_GetEvent(t *testing.T) {
	calendarRepository := memory.NewCalendarMap()
	eventForGet := models.NewEvent(1, time.Now(), 2, "test 1")
	calendarRepository.AddEvent(&eventForGet)
	e := models.NewEvent(1, time.Now().AddDate(0, 0, 2), 2, "test 2")
	calendarRepository.AddEvent(&e)
	e = models.NewEvent(2, time.Now(), 2, "test 3")
	calendarRepository.AddEvent(&e)
	service := calendar2.NewCalendarService(calendarRepository)
	server := initServer(&CalendarServer{Calendar: service})

	go server.Serve(getListener())
	defer server.Stop()

	client, cc := initClient()
	defer cc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	dt, _ := ptypes.TimestampProto(eventForGet.GetDateTime())
	grade, err := client.GetEvent(ctx, &calendar.GetEventRequest{UserId: 1, Datetime: dt})
	if err != nil {
		handlerError(err, t)
	} else {
		if grade.Event.Description != eventForGet.GetDescription() || grade.Event.GetDuration() != int32(eventForGet.GetDuration()) {
			t.Error("event not equal")
		}
	}
}

func initClient() (calendar.CalendarClient, *grpc.ClientConn) {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	return calendar.NewCalendarClient(cc), cc
}

func initServer(calendarServer calendar.CalendarServer) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)

	calendar.RegisterCalendarServer(server, calendarServer)
	return server
}

func handlerError(err error, t *testing.T) {
	statusErr, ok := status.FromError(err)
	if ok {
		if statusErr.Code() == codes.DeadlineExceeded {
			fmt.Println("Deadline exceeded!")
		} else {
			t.Errorf("undexpected error %s", statusErr.Message())
		}
	} else {
		t.Errorf("Error while calling RPC CheckHomework: %v", err)
	}
}
