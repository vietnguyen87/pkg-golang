package hubspot

// TicketsResponse object
type TicketsResponse struct {
	ErrorResponse
	Id         string                  `json:"id"`
	Properties TicketsResponseProperty `json:"properties"`
	CreatedAt  string                  `json:"createdAt"`
	UpdatedAt  string                  `json:"updatedAt"`
	Archived   bool                    `json:"archived"`
}

type TicketsResponseProperty struct {
	Content        string `json:"content"`
	HsObjectId     string `json:"hs_object_id"`
	HubspotOwnerId string `json:"hubspot_owner_id"`
	Subject        string `json:"subject"`
}
