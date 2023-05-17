package service

import (
	"database/sql"
	"git.neds.sh/matty/entain/sports/db"
	"git.neds.sh/matty/entain/sports/proto/sports"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

/*
SportsService test class with different requests
*/

func TestSportsServiceNameFilter_List(t *testing.T) {
	sportsService := createService(t)

	// Set up a request to pass to the service to find all events
	listEventsRequest := &sports.ListEventsRequest{Filter: &sports.EventsFilter{}}
	// Fetching all events
	listEventsResponse, err := sportsService.ListEvents(nil, listEventsRequest)
	assert.NoError(t, err)
	assert.NotNilf(t, listEventsResponse, "ListEventsResponse should not be nil.")
	assert.Equalf(t, 100, len(listEventsResponse.Events), "There should be a total of 100 events in DB.")

	// extracting a partial name of an event name to test the wildcard search
	var eventName string
	for i := 0; i < 100; i++ {
		event := listEventsResponse.Events[i].Name
		if len(event) > 2 {
			// Get substring from index 5 to 15
			eventName = event[1 : len(event)-2]
		}
	}

	// Set up a request with an event name
	filter := &sports.EventsFilter{
		Name: &eventName,
	}
	listEventsRequestBySport := &sports.ListEventsRequest{Filter: filter}
	listEventsResponseBySport, err := sportsService.ListEvents(nil, listEventsRequestBySport)
	assert.NoError(t, err)
	assert.NotNilf(t, listEventsResponse, "ListEventsResponse should not be nil.")
	for _, event := range listEventsResponseBySport.Events {
		strings.Contains(event.Name, eventName)
		assert.Truef(t, strings.Contains(event.Name, eventName), "Event %d is not %s event.", event.Id, eventName)
	}
}

func TestSportsServiceSportFilter_List(t *testing.T) {
	sportsService := createService(t)

	// Set up a request to pass to the service
	listEventsRequest := &sports.ListEventsRequest{Filter: &sports.EventsFilter{}}
	// Fetching all events
	listEventsResponse, err := sportsService.ListEvents(nil, listEventsRequest)
	assert.NoError(t, err)
	assert.NotNilf(t, listEventsResponse, "ListEventsResponse should not be nil.")
	assert.Equalf(t, 100, len(listEventsResponse.Events), "There should be a total of 100 events in DB.")

	// using the first one to do search based on sport name
	sport := &listEventsResponse.Events[0].Sport
	// Set up a request with a sport name
	filter := &sports.EventsFilter{
		Sport: sport,
	}
	listEventsRequestBySport := &sports.ListEventsRequest{Filter: filter}
	listEventsResponseBySport, err := sportsService.ListEvents(nil, listEventsRequestBySport)
	assert.NoError(t, err)
	assert.NotNilf(t, listEventsResponseBySport, "ListEventsResponse by sport should not be nil.")
	for _, event := range listEventsResponseBySport.Events {
		assert.EqualValuesf(t, *sport, event.Sport, "Event %d is not a %s event.", event.Id, *sport)
	}
}

func TestSportsServiceSortByAdvertisedStartTimeDefault_List(t *testing.T) {
	sportsService := createService(t)
	// Set up a request to pass to the service
	filter := &sports.EventsFilter{}
	listEventsRequest := &sports.ListEventsRequest{Filter: filter}
	listEventsResponse, err := sportsService.ListEvents(nil, listEventsRequest)
	assert.NoError(t, err)
	assert.NotNilf(t, listEventsResponse, "ListEventsResponse should not be nil.")
	assert.Equalf(t, 100, len(listEventsResponse.Events), "There should be a total of 100 events in DB.")

	for i := 0; i < len(listEventsResponse.Events)-1; i++ {
		previousElement := listEventsResponse.Events[i]
		element := listEventsResponse.Events[i+1]
		assert.Truef(t, element.AdvertisedStartTime.Nanos >= previousElement.AdvertisedStartTime.Nanos, "Results are not sorted by AdvertisedStartTime")
	}
}

func TestSportsServiceSortByMeetingId_List(t *testing.T) {
	sportsService := createService(t)
	orderField := "sport"
	// Set up a  request with OrderBy sport
	filter := &sports.EventsFilter{
		OrderBy: &orderField,
	}
	listEventsRequest := &sports.ListEventsRequest{Filter: filter}
	listEventsResponse, err := sportsService.ListEvents(nil, listEventsRequest)
	assert.NoError(t, err)
	assert.NotNilf(t, listEventsResponse, "ListEventsResponse should not be nil.")
	assert.Equalf(t, 100, len(listEventsResponse.Events), "There should be a total of 100 events in DB.")

	for i := 0; i < len(listEventsResponse.Events)-1; i++ {
		previousElement := listEventsResponse.Events[i]
		element := listEventsResponse.Events[i+1]
		assert.Truef(t, element.Sport >= previousElement.Sport, "Results are not sorted by sport")
	}
}

func TestSportsServiceSortBySportInCaps_List(t *testing.T) {
	sportsService := createService(t)
	orderField := "SPORT"
	// Set up a request with OrderBy sport
	filter := &sports.EventsFilter{
		OrderBy: &orderField,
	}
	listEventsRequest := &sports.ListEventsRequest{Filter: filter}
	listEventsResponse, err := sportsService.ListEvents(nil, listEventsRequest)
	assert.NoError(t, err)
	assert.NotNilf(t, listEventsResponse, "ListEventsResponse should not be nil.")
	assert.Equalf(t, 100, len(listEventsResponse.Events), "There should be a total of 100 events in DB.")

	for i := 0; i < len(listEventsResponse.Events)-1; i++ {
		previousElement := listEventsResponse.Events[i]
		element := listEventsResponse.Events[i+1]
		assert.Truef(t, element.Sport >= previousElement.Sport, "Results are not sorted by sport")
	}
}

func TestSportsServiceSortByName_List(t *testing.T) {
	sportsService := createService(t)
	orderField := "name"
	// Set up a request with OrderBy name
	filter := &sports.EventsFilter{
		OrderBy: &orderField,
	}
	listEventsRequest := &sports.ListEventsRequest{Filter: filter}
	listEventsResponse, err := sportsService.ListEvents(nil, listEventsRequest)
	assert.NoError(t, err)
	assert.NotNilf(t, listEventsResponse, "ListEventsResponse should not be nil.")
	assert.Equalf(t, 100, len(listEventsResponse.Events), "There should be a total of 100 events in DB.")

	for i := 0; i < len(listEventsResponse.Events)-1; i++ {
		previousElement := listEventsResponse.Events[i]
		element := listEventsResponse.Events[i+1]
		assert.Truef(t, element.Name >= previousElement.Name, "Results are not sorted by name")
	}
}

func TestSportsServiceSortByAdvertisedStartTime_List(t *testing.T) {
	sportsService := createService(t)
	// Set up a  filter with OrderBy AdvertisedStartTime
	orderField := "advertised_start_time"
	filter := &sports.EventsFilter{
		OrderBy: &orderField,
	}
	listEventsRequest := &sports.ListEventsRequest{Filter: filter}
	listEventsResponse, err := sportsService.ListEvents(nil, listEventsRequest)
	assert.NoError(t, err)
	assert.NotNilf(t, listEventsResponse, "ListEventsResponse should not be nil.")
	assert.Equalf(t, 100, len(listEventsResponse.Events), "There should be a total of 100 events in DB.")

	for i := 0; i < len(listEventsResponse.Events)-1; i++ {
		previousElement := listEventsResponse.Events[i]
		element := listEventsResponse.Events[i+1]
		assert.Truef(t, element.AdvertisedStartTime.Nanos >= previousElement.AdvertisedStartTime.Nanos, "Results are not sorted by AdvertisedStartTime")
	}
}

func TestSportsServiceSortByNoneExistentField_List(t *testing.T) {
	sportsService := createService(t)
	// Set up a request with an invalid OrderBy parameter to pass to the service
	orderField := "nonexistent"
	listEventsRequest := &sports.ListEventsRequest{Filter: &sports.EventsFilter{OrderBy: &orderField}}
	_, err := sportsService.ListEvents(nil, listEventsRequest)
	assert.Error(t, err)
}

func TestSportsServiceStatusField_List(t *testing.T) {
	sportsService := createService(t)
	// Set up a request to pass to the service
	listEventsRequest := &sports.ListEventsRequest{Filter: &sports.EventsFilter{}}
	listEventsResponse, err := sportsService.ListEvents(nil, listEventsRequest)
	assert.NoError(t, err)
	assert.NotNilf(t, listEventsResponse, "ListEventsResponse should not be nil.")
	assert.Equalf(t, 100, len(listEventsResponse.Events), "There should be a total of 100 events in DB.")
	events := listEventsResponse.Events
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

func TestSportsServiceAll_List(t *testing.T) {
	sportsService := createService(t)
	// Set up a request to pass to the service
	listEventsRequest := &sports.ListEventsRequest{Filter: &sports.EventsFilter{}}
	listEventsResponse, err := sportsService.ListEvents(nil, listEventsRequest)
	assert.NoError(t, err)
	assert.NotNilf(t, listEventsResponse, "ListEventsResponse should not be nil.")
	assert.Equalf(t, 100, len(listEventsResponse.Events), "There should be a total of 100 events in DB.")
}

func createService(t *testing.T) Sports {
	sportsDB, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)

	repo := db.NewEventsRepo(sportsDB)

	err = repo.Init()
	assert.NoError(t, err)

	return NewSportsService(
		repo,
	)
}
