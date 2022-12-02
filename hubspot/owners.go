package hubspot

type Owners interface {
	Get(ownerID string) (OwnersResponse, error)
	GetList(after, limit string) (GetListOwnerResponse, error)
}

const (
	OwnerPatch string = "/crm/v3/owners/"
)

type owners struct {
	client
}

func (c client) Owners() Owners {
	return &owners{
		client: c,
	}
}

func (c *owners) Get(ownerID string) (OwnersResponse, error) {
	response := OwnersResponse{}
	params := []string{
		"idProperty=id",
	}
	err := c.client.request("GET", OwnerPatch+ownerID, nil, &response, params)
	return response, err
}

func (c *owners) GetList(after, limit string) (GetListOwnerResponse, error) {
	response := GetListOwnerResponse{}
	params := []string{
		"after=" + after,
		"limit=" + limit,
	}
	err := c.client.request("GET", OwnerPatch, nil, &response, params)
	return response, err
}
