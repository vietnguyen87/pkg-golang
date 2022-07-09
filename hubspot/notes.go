package hubspot

import (
	"fmt"
)

type Notes interface {
	GetOne(noteID string) (NotesResponse, error)
	Create(data NotesRequest) (NotesResponse, error)
	Association(noteID, toObjectType, toObjectId, associationType string) (NotesResponse, error)
}

// Notes client
type notes struct {
	client
}

func (c client) Notes() Notes {
	return &notes{
		client: c,
	}
}

// GetOne Note
func (n *notes) GetOne(noteID string) (note NotesResponse, err error) {
	err = n.client.request("GET", fmt.Sprintf(
		"/crm/v3/objects/notes/%s", noteID), nil, &note, []string{
		"properties=hs_note_body",
		"associations=contacts",
		"archived=false",
	})
	return note, err
}

// Create new Deal
func (n *notes) Create(data NotesRequest) (note NotesResponse, err error) {
	err = n.client.request("POST", "/crm/v3/objects/notes", data, &note, nil)
	return note, err
}

// Association Note with ObjectType
// toObjectType: exp: contact
// associationType: exp: note_to_contact (22)
func (n *notes) Association(noteID, toObjectType, toObjectId, associationType string) (note NotesResponse, err error) {
	err = n.client.request("PUT", fmt.Sprintf(
		"/crm/v3/objects/notes/%s/associations/%s/%s/%s", noteID, toObjectType, toObjectId, associationType), nil, &note, []string{})
	return note, err
}
