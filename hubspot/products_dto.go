package hubspot

// ProductsRequest object
type ProductsRequest struct {
	Properties ProductsProperty `json:"properties"`
}

// ProductsProperty object
type ProductsProperty struct {
	Description              string `json:"description"`
	HsCostOfGoodsSold        string `json:"hs_cost_of_goods_sold"`
	HsRecurringBillingPeriod string `json:"hs_recurring_billing_period"`
	HsSku                    string `json:"hs_sku"`
	Name                     string `json:"name"`
	Price                    string `json:"price"`
}

// ProductsResponse object
type ProductsResponse struct {
	ErrorResponse
	Id         string           `json:"id"`
	Properties ProductsProperty `json:"properties"`
	CreatedAt  string           `json:"createdAt"`
	UpdatedAt  string           `json:"updatedAt"`
	Archived   bool             `json:"archived"`
}

// ProductResponse object
type ProductResponse struct {
	Properties ProductsProperty `json:"properties"`
}

// ProductsPropertyResponse object
type ProductsPropertyResponse struct {
	ProductsProperty
	Createdate         string `json:"createdate"`
	HsLastmodifieddate string `json:"hs_lastmodifieddate"`
}
