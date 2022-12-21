package hubspot

type Tickets interface {
	Get(ticketID string) (TicketsResponse, error)
}
type tickets struct {
	client
}

// Tickets constructor
func (c client) Tickets() Tickets {
	return &tickets{
		client: c,
	}
}

// Get Ticket
func (c *tickets) Get(ticketID string) (TicketsResponse, error) {
	r := TicketsResponse{}
	params := []string{
		"properties=content",
		"properties=subject",
		"properties=hubspot_owner_id",
	}
	err := c.client.request("GET", "/crm/v4/objects/tickets/"+ticketID, nil, &r, params)
	return r, err
}
