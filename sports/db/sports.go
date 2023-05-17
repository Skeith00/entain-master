package db

import (
	"database/sql"
	"errors"
	"fmt"
	"git.neds.sh/matty/entain/sports/proto/sports"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"sync"
	"time"
)

// Static slice with all the possible fields to use to sort. For this case, we have decided to limit it to 4 fields
var orderByFields = []string{Id, Name, Sport, Location, AdvertisedStartTime}

// EventsRepo provides repository access to sport_events.
type EventsRepo interface {
	// Init will initialise our sports repository.
	Init() error

	// List will return a list of events.
	List(filter *sports.EventsFilter) ([]*sports.Event, error)
}

type eventsRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewEventsRepo creates a new events repository.
func NewEventsRepo(db *sql.DB) EventsRepo {
	return &eventsRepo{db: db}
}

// Init prepares the event repository dummy data.
func (r *eventsRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy sports.
		err = r.seed()
	})

	return err
}

func (r *eventsRepo) List(filter *sports.EventsFilter) ([]*sports.Event, error) {
	return r.executeFilter(filter)
}

func (r *eventsRepo) executeFilter(filter *sports.EventsFilter) ([]*sports.Event, error) {
	var (
		err   error
		query string
		args  []interface{}
	)
	query = getSportQueries()[eventsList]

	query, args, err = r.applyFilter(query, filter)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanEvents(rows)
}

func (r *eventsRepo) applyFilter(query string, filter *sports.EventsFilter) (string, []interface{}, error) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args, nil
	}

	// Filters searches containing event name when name parameter is included in the request
	if filter.Name != nil {
		clauses = append(clauses, "name LIKE ?")
		args = append(args, "%"+*filter.Name+"%")
	}

	// Filters searches by sport name when sport parameter is included in the request
	if filter.Sport != nil {
		clauses = append(clauses, "sport = ?")
		args = append(args, filter.Sport)
	}

	// Sorting event by the field provided in the request
	orderByField, err := r.determineOrderByField(filter.OrderBy)
	if err != nil {
		return "", nil, err
	}

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
func (r *eventsRepo) determineOrderByField(orderBy *string) (string, error) {
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

func (r *eventsRepo) scanEvents(
	rows *sql.Rows,
) ([]*sports.Event, error) {
	var events []*sports.Event

	for rows.Next() {
		var event sports.Event
		var advertisedStart time.Time

		if err := rows.Scan(&event.Id, &event.Name, &event.Sport, &event.Location, &advertisedStart); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		// status field in event is calculated in every returned row based on AdvertisedStartTime
		event.AdvertisedStartTime = ts
		if ptypes.TimestampNow().Seconds > ts.Seconds {
			event.Status = "CLOSED"
		} else {
			event.Status = "OPEN"
		}

		events = append(events, &event)
	}

	return events, nil
}
