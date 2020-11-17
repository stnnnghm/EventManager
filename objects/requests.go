package objects

import (
	"encoding/json"
	"net/http"
)

// MaxListLimit max listing
const MaxListLimit = 200

// GetRequest for retrieving single Event
type GetRequest struct {
	ID string `json:"id"`
}

// ListRequest for retrieving list of Events
type ListRequest struct {
	Limit int `json:"limit"`
	After string `json:"after"`
	// optional name matching
	Name string `json:"name"`
}

// CreateRequest for creating a new Event
type CreateRequest struct {
	Event *Event `json:"event"`
}

// UpdateDetailsRequest to update existing Event
type UpdateDetailsRequest struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Website string `json:"website"`
	Address string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

// CancelRequest to cancel Event
type CancelRequest struct {
	ID string `json:"id"`
}

// RescheduleRequest to reschedule an Event
type RescheduleRequest struct {
	ID string `json:"id"`
	NewSlot *TimeSlot `json:"new_slot"`
}

// DeleteRequest to delete an event
type DeleteRequest struct {
	ID string `json:"id"`
}

// EventResponseWrapper response of any Event req
type EventResponseWrapper struct {
	Event *Event `json:"event,omitempty"`
	Events []*Event `json:"events,omitempty"`
	Code int `json:"-"`
}

// JSON convert EventResponseWrapper in json
func (e *EventResponseWrapper) Json() []byte {
	if e == nil {
		return []byte("{}")
	}
	res, _ := json.Marshal(e)
	return res
}

// StatusCode return the status code
func (e *EventResponseWrapper) StatusCode() int {
	if e == nil || e.Code == 0 {
		return http.StatusOK
	}
	return e.Code
}