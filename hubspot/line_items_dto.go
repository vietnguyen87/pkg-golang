package hubspot

// LineItemsRequest object
type LineItemsRequest struct {
	Properties LineItemsProperty `json:"properties"`
}

// LineItemsProperty object
type LineItemsProperty struct {
	Name                      string `json:"name"`
	HsProductId               string `json:"hs_product_id"`
	HsRecurringBillingPeriod  string `json:"hs_recurring_billing_period"`
	Recurringbillingfrequency string `json:"recurringbillingfrequency"`
	Quantity                  string `json:"quantity"`
	Price                     string `json:"price"`
	HsSku                     string `json:"hs_sku"`
	SkuCode                   string `json:"sku_code"`
	StartDate                 string `json:"start_date"`
	EndDate                   string `json:"end_date"`
	Subject                   string `json:"subject"`
	CourseId                  string `json:"course_id"`
}

// LineItemsResponse object
type LineItemsResponse struct {
	ErrorResponse
	Id         string            `json:"id"`
	Properties LineItemsProperty `json:"properties"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	Archived   bool              `json:"archived"`
}

// LineItemResponse object
type LineItemResponse struct {
	ErrorResponse
	Id         string            `json:"id"`
	Properties LineItemsProperty `json:"properties"`
}

type GetByIdsLineItemsRequest struct {
	Properties []string                 `json:"properties"`
	Inputs     []GetByIdsLineItemsInput `json:"inputs"`
}
type GetByIdsLineItemsInput struct {
	Id string `json:"id"`
}
type GetLineItemsByIdsResponse struct {
	Status  string              `json:"status"`
	Results []LineItemsResponse `json:"results"`
}
