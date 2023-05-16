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
	if err != nil {
		return
	}

	for _, race := range races {
		assert.Truef(t, race.Visible, "Race %d is not visible.", race.Id)
	}

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
