package service

import (
	"database/sql"
	"git.neds.sh/matty/entain/racing/db"
	"git.neds.sh/matty/entain/racing/proto/racing"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
RacingService test class with different requests
*/

func TestRacingServiceFetchById_Get(t *testing.T) {
	racingService := createService(t)

	// Set up a request with visible true
	filter := &racing.GetRaceRequest{
		Id: 1,
	}
	race, err := racingService.GetRace(nil, filter)
	assert.NoError(t, err)
	assert.EqualValuesf(t, 1, race.Id, "Race should have id 1.")

}

func TestRacingServiceVisibleFilter_List(t *testing.T) {
	racingService := createService(t)

	visible := true
	// Set up a request with visible true
	filter := &racing.ListRacesRequestFilter{
		Visible: &visible,
	}
	listRacesRequest := &racing.ListRacesRequest{Filter: filter}
	listRacesResponse, err := racingService.ListRaces(nil, listRacesRequest)
	assert.NoError(t, err)
	for _, race := range listRacesResponse.Races {
		assert.Truef(t, race.Visible, "Race %d is not visible.", race.Id)
	}
}

func TestRacingServiceSortByAdvertisedStartTimeDefault_List(t *testing.T) {
	racingService := createService(t)
	// Set up a request to pass to the service
	filter := &racing.ListRacesRequestFilter{}
	listRacesRequest := &racing.ListRacesRequest{Filter: filter}
	listRacesResponse, err := racingService.ListRaces(nil, listRacesRequest)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(listRacesResponse.Races), "There should be a total of 100 races in DB.")

	for i := 0; i < len(listRacesResponse.Races)-1; i++ {
		previousElement := listRacesResponse.Races[i]
		element := listRacesResponse.Races[i+1]
		assert.Truef(t, element.AdvertisedStartTime.Nanos >= previousElement.AdvertisedStartTime.Nanos, "Results are not sorted by AdvertisedStartTime")
	}
}

func TestRacingServiceSortByMeetingId_List(t *testing.T) {
	racingService := createService(t)
	orderField := "meeting_id"
	// Set up a  request with OrderBy meeting_id
	filter := &racing.ListRacesRequestFilter{
		OrderBy: &orderField,
	}
	listRacesRequest := &racing.ListRacesRequest{Filter: filter}
	listRacesResponse, err := racingService.ListRaces(nil, listRacesRequest)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(listRacesResponse.Races), "There should be a total of 100 races in DB.")

	for i := 0; i < len(listRacesResponse.Races)-1; i++ {
		previousElement := listRacesResponse.Races[i]
		element := listRacesResponse.Races[i+1]
		assert.Truef(t, element.MeetingId >= previousElement.MeetingId, "Results are not sorted by meeting_id")
	}
}

func TestRacingServiceSortByMeetingIdInCaps_List(t *testing.T) {
	racingService := createService(t)
	orderField := "MEETING_ID"
	// Set up a request with OrderBy meeting_id
	filter := &racing.ListRacesRequestFilter{
		OrderBy: &orderField,
	}
	listRacesRequest := &racing.ListRacesRequest{Filter: filter}
	listRacesResponse, err := racingService.ListRaces(nil, listRacesRequest)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(listRacesResponse.Races), "There should be a total of 100 races in DB.")

	for i := 0; i < len(listRacesResponse.Races)-1; i++ {
		previousElement := listRacesResponse.Races[i]
		element := listRacesResponse.Races[i+1]
		assert.Truef(t, element.MeetingId >= previousElement.MeetingId, "Results are not sorted by meeting_id")
	}
}

func TestRacingServiceSortByName_List(t *testing.T) {
	racingService := createService(t)
	orderField := "name"
	// Set up a request with OrderBy name
	filter := &racing.ListRacesRequestFilter{
		OrderBy: &orderField,
	}
	listRacesRequest := &racing.ListRacesRequest{Filter: filter}
	listRacesResponse, err := racingService.ListRaces(nil, listRacesRequest)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(listRacesResponse.Races), "There should be a total of 100 races in DB.")

	for i := 0; i < len(listRacesResponse.Races)-1; i++ {
		previousElement := listRacesResponse.Races[i]
		element := listRacesResponse.Races[i+1]
		assert.Truef(t, element.Name >= previousElement.Name, "Results are not sorted by name")
	}
}

func TestRacingServiceSortByAdvertisedStartTime_List(t *testing.T) {
	racingService := createService(t)
	// Set up a  filter with OrderBy AdvertisedStartTime
	orderField := "advertised_start_time"
	filter := &racing.ListRacesRequestFilter{
		OrderBy: &orderField,
	}
	listRacesRequest := &racing.ListRacesRequest{Filter: filter}
	listRacesResponse, err := racingService.ListRaces(nil, listRacesRequest)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(listRacesResponse.Races), "There should be a total of 100 races in DB.")

	for i := 0; i < len(listRacesResponse.Races)-1; i++ {
		previousElement := listRacesResponse.Races[i]
		element := listRacesResponse.Races[i+1]
		assert.Truef(t, element.AdvertisedStartTime.Nanos >= previousElement.AdvertisedStartTime.Nanos, "Results are not sorted by AdvertisedStartTime")
	}
}

func TestRacingServiceSortByNoneExistentField_List(t *testing.T) {
	racingService := createService(t)
	// Set up a request with an invalid OrderBy parameter to pass to the service
	orderField := "nonexistent"
	listRacesRequest := &racing.ListRacesRequest{Filter: &racing.ListRacesRequestFilter{OrderBy: &orderField}}
	_, err := racingService.ListRaces(nil, listRacesRequest)
	assert.Error(t, err)
}

func TestRacingServiceStatusField_List(t *testing.T) {
	racingService := createService(t)
	// Set up a request to pass to the service
	listRacesRequest := &racing.ListRacesRequest{Filter: &racing.ListRacesRequestFilter{}}
	listRacesResponse, err := racingService.ListRaces(nil, listRacesRequest)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(listRacesResponse.Races), "There should be a total of 100 races in DB.")
	races := listRacesResponse.Races
	// Validating that all races with CLOSED status have a past date and OPEN status in a future date
	for _, race := range races {
		switch race.Status {
		case "CLOSED":
			assert.Truef(t, ptypes.TimestampNow().Seconds > race.AdvertisedStartTime.Seconds, "Race %d should not be CLOSED.", race.Id)
		case "OPEN":
			assert.Truef(t, ptypes.TimestampNow().Seconds <= race.AdvertisedStartTime.Seconds, "Race %d should not be OPEN.", race.Id)
		default:
			assert.Failf(t, "There should only be OPEN and CLOSED status. %s is invalid.", race.Status)
		}
	}
}

func TestRacingServiceAll_List(t *testing.T) {
	racingService := createService(t)
	// Set up a request to pass to the service
	listRacesRequest := &racing.ListRacesRequest{Filter: &racing.ListRacesRequestFilter{}}
	listRacesResponse, err := racingService.ListRaces(nil, listRacesRequest)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(listRacesResponse.Races), "There should be a total of 100 races in DB.")
}

func createService(t *testing.T) Racing {

	racingDB, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)

	repo := db.NewRacesRepo(racingDB)

	err = repo.Init()
	assert.NoError(t, err)

	return NewRacingService(
		repo,
	)
}
