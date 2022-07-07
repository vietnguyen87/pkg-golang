package hubspot

import "fmt"

// LineItems client
type lineItems struct {
	client
}

type LineItems interface {
	Get(lineItemId string) (LineItemResponse, error)
	Create(data LineItemsRequest) (LineItemsResponse, error)
	GetByIds(ids []string) (GetLineItemsByIdsResponse, error)
	Update(lineItem string, data LineItemsRequest) (LineItemsResponse, error)
	Delete(lineItemId string) error
	Associate(lineItemId, toObjectType, toObjectId, associationType string) (LineItemAssociateResponse, error)
}

// LineItems constructor (from Client)
func (c client) LineItems() LineItems {
	return &lineItems{
		client: c,
	}
}

// Get Line Items
func (l *lineItems) Get(lineItemId string) (LineItemResponse, error) {
	r := LineItemResponse{}
	params := []string{
		"properties=course_id",
		"properties=hs_sku",
		"properties=start_date",
		"properties=end_date",
		"properties=name",
		"properties=quantity",
		"properties=subject",
		"properties=sku_code",
	}
	err := l.client.request("GET", fmt.Sprintf("/crm/v3/objects/line_items/%v", lineItemId), nil, &r, params)
	return r, err
}

// Create new Line Items
func (l *lineItems) Create(data LineItemsRequest) (LineItemsResponse, error) {
	r := LineItemsResponse{}
	err := l.client.request("POST", "/crm/v3/objects/line_items", data, &r, nil)
	return r, err
}

func (l *lineItems) GetByIds(ids []string) (GetLineItemsByIdsResponse, error) {
	r := GetLineItemsByIdsResponse{}
	var inputs []GetByIdsLineItemsInput
	for _, v := range ids {
		inputs = append(inputs, GetByIdsLineItemsInput{Id: v})
	}
	data := GetByIdsLineItemsRequest{
		Inputs:     inputs,
		Properties: []string{"course_id", "hs_sku", "start_date", "end_date", "name", "quantity", "subject"},
	}
	err := l.client.request("POST", "/crm/v3/objects/line_items/batch/read", data, &r, nil)
	return r, err
}

// Update Line Items
func (l *lineItems) Update(lineItem string, data LineItemsRequest) (LineItemsResponse, error) {
	r := LineItemsResponse{}
	err := l.client.request("PATCH", "/crm/v3/objects/line_items/"+lineItem, data, &r, nil)
	return r, err
}

// Delete Deal
func (l *lineItems) Delete(lineItemId string) error {
	err := l.client.request("DELETE", "/crm/v3/objects/line_items/"+lineItemId, nil, nil, nil)
	return err
}

// Associate Create new Line Items
// toObjectType: DEAL
// associationType: line_item_to_deal
func (l *lineItems) Associate(lineItemId, toObjectType, toObjectId, associationType string) (LineItemAssociateResponse, error) {
	r := LineItemAssociateResponse{}
	err := l.client.request("PUT",
		fmt.Sprintf("/crm/v3/objects/line_items/%s/associations/%s/%s/%s",
			lineItemId,
			toObjectType,
			toObjectId,
			associationType,
		),
		nil,
		&r,
		nil,
	)
	return r, err
}
