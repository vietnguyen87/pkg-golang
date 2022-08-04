package hubspot

type Engagements interface {
	Create(data EngagementsRequest) (EngagementsResponse, error)
}

// Engagements client
type engagements struct {
	client
}

func (c client) Engagements() Engagements {
	return &engagements{
		client: c,
	}
}

// Create new Engagement
func (n *engagements) Create(data EngagementsRequest) (note EngagementsResponse, err error) {
	err = n.client.request("POST", "/engagements/v1/engagements", data, &note, nil)
	return note, err
}
