package db

import "fmt"

const (
	eventsList = "list"
)

// DB constants with table sports field names
const (
	Id                  = "id"
	Name                = "name"
	Sport               = "sport"
	Location            = "location"
	AdvertisedStartTime = "advertised_start_time"
)

func getSportQueries() map[string]string {
	return map[string]string{
		eventsList: fmt.Sprintf(
			"SELECT %s, %s, %s, %s, %s FROM sport_events",
			Id,
			Name,
			Sport,
			Location,
			AdvertisedStartTime),
	}
}
