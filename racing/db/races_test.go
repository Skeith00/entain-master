package db

import (
	"database/sql"
	"git.neds.sh/matty/entain/racing/proto/racing"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
RacesRepo test class with different filters in ListRacesRequestFilter
*/
func TestRacesRepoVisibleFilter_List(t *testing.T) {

	racesRepo := createRepo(t)

	visible := true
	// Set up a  filter with visibile true
	filter := &racing.ListRacesRequestFilter{
		Visible: &visible,
	}
	races, err := racesRepo.List(filter)
	assert.NoError(t, err)
	for _, race := range races {
		assert.Truef(t, race.Visible, "Race %d is not visible.", race.Id)
	}
}

func TestRacesRepoSortByAdvertisedStartTimeDefault_List(t *testing.T) {
	racesRepo := createRepo(t)
	// Set up a filter to pass to the List method
	filter := &racing.ListRacesRequestFilter{}
	races, err := racesRepo.List(filter)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(races), "There should be a total of 100 races in DB.")

	for i := 0; i < len(races)-1; i++ {
		previousElement := races[i]
		element := races[i+1]
		assert.Truef(t, element.AdvertisedStartTime.Nanos >= previousElement.AdvertisedStartTime.Nanos, "Results are not sorted by AdvertisedStartTime")
	}
}

func TestRacesRepoSortByMeetingId_List(t *testing.T) {

	racesRepo := createRepo(t)
	orderField := "meeting_id"
	// Set up a  filter with OrderBy meeting_id
	filter := &racing.ListRacesRequestFilter{
		OrderBy: &orderField,
	}
	races, err := racesRepo.List(filter)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(races), "There should be a total of 100 races in DB.")

	for i := 0; i < len(races)-1; i++ {
		previousElement := races[i]
		element := races[i+1]
		assert.Truef(t, element.MeetingId >= previousElement.MeetingId, "Results are not sorted by meeting_id")
	}
}

func TestRacesRepoSortByMeetingIdInCaps_List(t *testing.T) {

	racesRepo := createRepo(t)
	orderField := "MEETING_ID"
	// Set up a  filter with OrderBy meeting_id
	filter := &racing.ListRacesRequestFilter{
		OrderBy: &orderField,
	}
	races, err := racesRepo.List(filter)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(races), "There should be a total of 100 races in DB.")

	for i := 0; i < len(races)-1; i++ {
		previousElement := races[i]
		element := races[i+1]
		assert.Truef(t, element.MeetingId >= previousElement.MeetingId, "Results are not sorted by meeting_id")
	}
}

func TestRacesRepoSortByName_List(t *testing.T) {

	racesRepo := createRepo(t)
	orderField := "name"
	// Set up a  filter with OrderBy name
	filter := &racing.ListRacesRequestFilter{
		OrderBy: &orderField,
	}
	races, err := racesRepo.List(filter)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(races), "There should be a total of 100 races in DB.")

	for i := 0; i < len(races)-1; i++ {
		previousElement := races[i]
		element := races[i+1]
		assert.Truef(t, element.Name >= previousElement.Name, "Results are not sorted by name")
	}
}

func TestRacesRepoSortByAdvertisedStartTime_List(t *testing.T) {
	racesRepo := createRepo(t)
	// Set up a  filter with OrderBy AdvertisedStartTime
	orderField := "advertised_start_time"
	// Set up a  filter with OrderBy name
	filter := &racing.ListRacesRequestFilter{
		OrderBy: &orderField,
	}
	races, err := racesRepo.List(filter)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(races), "There should be a total of 100 races in DB.")

	for i := 0; i < len(races)-1; i++ {
		previousElement := races[i]
		element := races[i+1]
		assert.Truef(t, element.AdvertisedStartTime.Nanos >= previousElement.AdvertisedStartTime.Nanos, "Results are not sorted by AdvertisedStartTime")
	}
}

func TestRacesRepoSortByNoneExistentField_List(t *testing.T) {
	racesRepo := createRepo(t)
	// Set up a  filter with an invalid OrderBy parameter
	orderField := "nonexistent"
	filter := &racing.ListRacesRequestFilter{
		OrderBy: &orderField,
	}
	_, err := racesRepo.List(filter)
	assert.Error(t, err)
}

func TestRacesRepoAll_List(t *testing.T) {
	racesRepo := createRepo(t)
	// Set up a filter to pass to the List method
	filter := &racing.ListRacesRequestFilter{}
	races, err := racesRepo.List(filter)
	assert.NoError(t, err)
	assert.Equalf(t, 100, len(races), "There should be a total of 100 races in DB.")
}

func createRepo(t *testing.T) RacesRepo {
	racingDB, err := sql.Open("sqlite3", ":memory:")
	assert.NoError(t, err)

	repo := NewRacesRepo(racingDB)
	err = repo.Init()
	assert.NoError(t, err)

	return repo
}
