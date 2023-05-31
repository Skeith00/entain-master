package db

import "fmt"

const (
	racesList = "list"
)

// DB constants with table race field names
const (
	Id                  = "id"
	MeetingId           = "meeting_id"
	Name                = "name"
	Number              = "number"
	Visible             = "visible"
	AdvertisedStartTime = "advertised_start_time"
)

func getRaceQueries() map[string]string {
	return map[string]string{
		racesList: fmt.Sprintf(
			"SELECT %s, %s, %s, %s, %s, %s FROM races",
			Id,
			MeetingId,
			Name,
			Number,
			Visible,
			AdvertisedStartTime),
	}
}
