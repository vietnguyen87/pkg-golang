package hubspot

import "fmt"

// Products client
type products struct {
	client
}

type Products interface {
	Get(productId string) (ProductsResponse, error)
	Create(data ProductsRequest) (ProductsResponse, error)
	Update(productId string, data ProductsRequest) (ProductsResponse, error)
	Delete(productId string) error
}

// Products constructor (from Client)
func (c client) Products() Products {
	return &products{
		client: c,
	}
}

// Get products
func (p *products) Get(productId string) (ProductsResponse, error) {
	r := ProductsResponse{}
	err := p.client.request("GET", fmt.Sprintf("/crm/v3/objects/products/%v", productId), nil, &r, nil)
	return r, err
}

// Create new products
func (p *products) Create(data ProductsRequest) (ProductsResponse, error) {
	r := ProductsResponse{}
	err := p.client.request("POST", "/crm/v3/objects/products", data, &r, nil)
	return r, err
}

// Update product
func (p *products) Update(productId string, data ProductsRequest) (ProductsResponse, error) {
	r := ProductsResponse{}
	err := p.client.request("PATCH", "/crm/v3/objects/products/"+productId, data, &r, nil)
	return r, err
}

// Delete product
func (p *products) Delete(productId string) error {
	err := p.client.request("DELETE", "/crm/v3/objects/products/"+productId, nil, nil, nil)
	return err
}
