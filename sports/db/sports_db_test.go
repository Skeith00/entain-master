package db

import (
	"database/sql"
	"git.neds.sh/matty/entain/sports/proto/sports"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

/*
EventsRepo test class with different filters
*/

func TestEventsRepoNameFilter_List(t *testing.T) {
	eventsRepo := createRepo(t)

	// Set up a request to pass to the repo to find all events
	filterAll := &sports.EventsFilter{}
	// Fetching all events
	eventsResponse, err := eventsRepo.List(filterAll)
	assert.NoError(t, err)
	assert.NotNilf(t, eventsResponse, "Events should not be nil.")
	assert.Equalf(t, 100, len(eventsResponse), "There should be a total of 100 events in DB.")

	// extracting a partial name of an event name to test the wildcard search
	var eventName string
	for i := 0; i < 100; i++ {
		event := eventsResponse[i].Name
		if len(event) > 2 {
			// Get substring from index 1 to the second-to-last index
			eventName = event[1 : len(event)-2]
		}
	}

	// Set up a filter with an event name
	filterEventName := &sports.EventsFilter{
		Name: &eventName,
	}
	eventsResponseBySport, err := eventsRepo.List(filterEventName)
	assert.NoError(t, err)
	assert.NotNilf(t, eventsResponseBySport, "Events should not be nil.")
	for _, event := range eventsResponseBySport {
		strings.Contains(event.Name, eventName)
		assert.Truef(t, strings.Contains(event.Name, eventName), "Event %d is not %s event.", event.Id, eventName)
	}
}

func TestEventsRepoSportFilter_List(t *testing.T) {
	eventsRepo := createRepo(t)

	// Set up a request to pass to the repo to find all events
	filterAll := &sports.EventsFilter{}
	// Fetching all events
	eventsResponse, err := eventsRepo.List(filterAll)
	assert.NoError(t, err)
	assert.NotNilf(t, eventsResponse, "Events should not be nil.")
	assert.Equalf(t, 100, len(eventsResponse), "There should be a total of 100 events in DB.")

	// using the first one to do search based on sport name
	sport := &eventsResponse[0].Sport
	// Set up a request with a sport name
	filter := &sports.EventsFilter{
		Sport: sport,
	}

	events, err := eventsRepo.List(filter)
	assert.NoError(t, err)
	assert.NotNilf(t, events, "Events by sport should not be nil.")
	for _, event := range events {
		assert.EqualValuesf(t, *sport, event.Sport, "Event %d is not a %s event.", event.Id, *sport)
	}
}

func TestEventsRepoSortByAdvertisedStartTimeDefault_List(t *testing.T) {
	eventsRepo := createRepo(t)
	// Set up a filter to pass to the List method
	filter := &sports.EventsFilter{}
	events, err := eventsRepo.List(filter)
	assert.NoError(t, err)
	assert.NotNilf(t, events, "Events should not be nil.")
	assert.Equalf(t, 100, len(events), "There should be a total of 100 events in DB.")

	for i := 0; i < len(events)-1; i++ {
		previousElement := events[i]
		element := events[i+1]
		assert.Truef(t, element.AdvertisedStartTime.Nanos >= previousElement.AdvertisedStartTime.Nanos, "Results are not sorted by AdvertisedStartTime")
	}
}

func TestEventsRepoSortByMeetingId_List(t *testing.T) {
	eventsRepo := createRepo(t)
	orderField := "sport"
	// Set up a  filter with OrderBy meeting_id
	filter := &sports.EventsFilter{
		OrderBy: &orderField,
	}
	events, err := eventsRepo.List(filter)
	assert.NoError(t, err)
	assert.NotNilf(t, events, "Events should not be nil.")
	assert.Equalf(t, 100, len(events), "There should be a total of 100 events in DB.")

	for i := 0; i < len(events)-1; i++ {
		previousElement := events[i]
		element := events[i+1]
		assert.Truef(t, element.Sport >= previousElement.Sport, "Results are not sorted by meeting_id")
	}
}

func TestEventsRepoSortByMeetingIdInCaps_List(t *testing.T) {
	eventsRepo := createRepo(t)
	orderField := "SPORT"
	// Set up a request with OrderBy sport
	filter := &sports.EventsFilter{
		OrderBy: &orderField,
	}
	events, err := eventsRepo.List(filter)
	assert.NoError(t, err)
	assert.NotNilf(t, events, "Events should not be nil.")
	assert.Equalf(t, 100, len(events), "There should be a total of 100 events in DB.")

	for i := 0; i < len(events)-1; i++ {
		previousElement := events[i]
		element := events[i+1]
		assert.Truef(t, element.Sport >= previousElement.Sport, "Results are not sorted by sport")
	}
}

func TestEventsRepoSortByName_List(t *testing.T) {
	eventsRepo := createRepo(t)
	orderField := "name"
	// Set up a  filter with OrderBy name
	filter := &sports.EventsFilter{
		OrderBy: &orderField,
	}
	events, err := eventsRepo.List(filter)
	assert.NoError(t, err)
	assert.NotNilf(t, events, "Events should not be nil.")
	assert.Equalf(t, 100, len(events), "There should be a total of 100 events in DB.")

	for i := 0; i < len(events)-1; i++ {
		previousElement := events[i]
		element := events[i+1]
		assert.Truef(t, element.Name >= previousElement.Name, "Results are not sorted by name")
	}
}

func TestEventsRepoSortByAdvertisedStartTime_List(t *testing.T) {
	eventsRepo := createRepo(t)
	// Set up a  filter with OrderBy AdvertisedStartTime
	orderField := "advertised_start_time"
	// Set up a  filter with OrderBy name
	filter := &sports.EventsFilter{
		OrderBy: &orderField,
	}
	events, err := eventsRepo.List(filter)
	assert.NoError(t, err)
	assert.NotNilf(t, events, "Events should not be nil.")
	assert.Equalf(t, 100, len(events), "There should be a total of 100 events in DB.")

	for i := 0; i < len(events)-1; i++ {
		previousElement := events[i]
		element := events[i+1]
		assert.Truef(t, element.AdvertisedStartTime.Nanos >= previousElement.AdvertisedStartTime.Nanos, "Results are not sorted by AdvertisedStartTime")
	}
}

func TestEventsRepoSortByNoneExistentField_List(t *testing.T) {
	eventsRepo := createRepo(t)
	// Set up a  filter with an invalid OrderBy parameter
	orderField := "nonexistent"
	filter := &sports.EventsFilter{
		OrderBy: &orderField,
	}
	_, err := eventsRepo.List(filter)
	assert.Error(t, err)
}

func TestEventsRepoStatusField_List(t *testing.T) {
	eventsRepo := createRepo(t)
	// Set up a filter to pass to the List method
	filter := &sports.EventsFilter{}
	events, err := eventsRepo.List(filter)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(events), "There should be a total of 100 events in DB.")

	// Validating that all events with CLOSED status have a past date and OPEN status in a future date
	for _, event := range events {
		switch event.Status {
		case "CLOSED":
			assert.Truef(t, ptypes.TimestampNow().Seconds > event.AdvertisedStartTime.Seconds, "Event %d should not be CLOSED.", event.Id)
		case "OPEN":
			assert.Truef(t, ptypes.TimestampNow().Seconds <= event.AdvertisedStartTime.Seconds, "Event %d should not be OPEN.", event.Id)
		default:
			assert.Failf(t, "There should only be OPEN and CLOSED status. %s is invalid.", event.Status)
		}
	}
}

func TestEventsRepoAll_List(t *testing.T) {
	eventsRepo := createRepo(t)
	// Set up a filter to pass to the List method
	filter := &sports.EventsFilter{}
	events, err := eventsRepo.List(filter)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(events), "There should be a total of 100 events in DB.")
}

func createRepo(t *testing.T) EventsRepo {
	sportsDB, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)

	repo := NewEventsRepo(sportsDB)

	err = repo.Init()
	assert.NoError(t, err)

	return repo
}
