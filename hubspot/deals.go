package hubspot

import "fmt"

type Deals interface {
	Get(dealID string) (DealsResponse, error)
	Create(data DealsRequest) (DealsResponse, error)
	Update(dealID string, data DealsRequest) (DealsResponse, error)
	Delete(dealID string) error
}

// Deals client
type deals struct {
	client
}

func (c client) Deals() Deals {
	return &deals{
		client: c,
	}
}

// Get Deal
func (d *deals) Get(dealID string) (DealsResponse, error) {
	r := DealsResponse{}
	err := d.client.request("GET", fmt.Sprintf(
		"/crm/v3/objects/deals/%s", dealID), nil, &r, []string{
		"associations=contacts",
		"associations=line_items",
		"archived=false",
	})
	return r, err
}

// Create new Deal
func (d *deals) Create(data DealsRequest) (DealsResponse, error) {
	r := DealsResponse{}
	err := d.client.request("POST", "/crm/v3/objects/deals", data, &r, nil)
	return r, err
}

// Update Deal
func (d *deals) Update(dealID string, data DealsRequest) (DealsResponse, error) {
	r := DealsResponse{}
	err := d.client.request("PATCH", "/crm/v3/objects/deals/"+dealID, data, &r, nil)
	return r, err
}

// Delete Deal
func (d *deals) Delete(dealID string) error {
	err := d.client.request("DELETE", "/crm/v3/objects/deals/"+dealID, nil, nil, nil)
	return err
}