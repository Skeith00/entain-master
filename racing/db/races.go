package db

import (
	"database/sql"
	"errors"
	"fmt"
	"git.neds.sh/matty/entain/racing/proto/racing"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"sync"
	"time"
)

// Static slice with all the possible fields to use to sort. For this case, we have decided to limit it to 4 fields
var orderByFields = []string{Id, MeetingId, Name, AdvertisedStartTime}

// RacesRepo provides repository access to races.
type RacesRepo interface {
	// Init will initialise our races repository.
	Init() error

	// List will return a list of races.
	List(filter *racing.ListRacesRequestFilter) ([]*racing.Race, error)
}

type racesRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewRacesRepo creates a new races repository.
func NewRacesRepo(db *sql.DB) RacesRepo {
	return &racesRepo{db: db}
}

// Init prepares the race repository dummy data.
func (r *racesRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy races.
		err = r.seed()
	})

	return err
}

func (r *racesRepo) List(filter *racing.ListRacesRequestFilter) ([]*racing.Race, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getRaceQueries()[racesList]

	query, args, err = r.applyFilter(query, filter)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanRaces(rows)
}

func (r *racesRepo) applyFilter(query string, filter *racing.ListRacesRequestFilter) (string, []interface{}, error) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args, nil
	}

	if len(filter.MeetingIds) > 0 {
		clauses = append(clauses, "meeting_id IN ("+strings.Repeat("?,", len(filter.MeetingIds)-1)+"?)")

		for _, meetingID := range filter.MeetingIds {
			args = append(args, meetingID)
		}
	}

	// When visible parameter is included in the request, the visible filter is applied
	if filter.Visible != nil {
		clauses = append(clauses, "visible = ?")
		args = append(args, filter.Visible)
	}

	// Adding a field to the arguments to sort races
	orderByField, err := determineOrderByField(filter.OrderBy)
	if err != nil {
		return "", nil, err
	}
	//args = append(args, orderByField)

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}
	query += fmt.Sprintf(" ORDER BY %s", orderByField)
	return query, args, nil
}

/*
The orderBy field provided by the user is validated using this function.
The default advertised_start_time field will be returned if the provided field is invalid.
*/
func determineOrderByField(orderBy *string) (string, error) {
	if orderBy == nil {
		return AdvertisedStartTime, nil
	}
	// Finding the element in the slice of valid orderBy fields
	for _, element := range orderByFields {
		if strings.ToLower(element) == strings.ToLower(*orderBy) {
			return strings.ToLower(element), nil
		}
	}
	return "", errors.New("order_by field incorrect")
}

func (m *racesRepo) scanRaces(
	rows *sql.Rows,
) ([]*racing.Race, error) {
	var races []*racing.Race

	for rows.Next() {
		var race racing.Race
		var advertisedStart time.Time

		if err := rows.Scan(&race.Id, &race.MeetingId, &race.Name, &race.Number, &race.Visible, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		// status field in race is calculated in every returned row based on AdvertisedStartTime
		race.AdvertisedStartTime = ts
		if ptypes.TimestampNow().Seconds > ts.Seconds {
			race.Status = "CLOSED"
		} else {
			race.Status = "OPEN"
		}

		races = append(races, &race)
	}

	return races, nil
}
