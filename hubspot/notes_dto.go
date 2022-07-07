package hubspot

type NotesRequest struct {
	Properties NotesProperties `json:"properties"`
}

type NotesResponse struct {
	ErrorResponse
	Id           string          `json:"id"`
	Properties   NotesProperties `json:"properties"`
	CreatedAt    string          `json:"createdAt"`
	UpdatedAt    string          `json:"updatedAt"`
	Archived     bool            `json:"archived"`
	Associations Associations    `json:"associations"`
}

type NotesProperties struct {
	Timestamp     string `json:"hs_timestamp"`
	Body          string `json:"hs_note_body"`
	OwnerId       string `json:"hubspot_owner_id"`
	AttachmentIds string `json:"hs_attachment_ids,omitempty"`
}
