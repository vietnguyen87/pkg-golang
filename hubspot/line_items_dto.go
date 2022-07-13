package hubspot

// LineItemsRequest object
type LineItemsRequest struct {
	Properties LineItemsProperty `json:"properties"`
}

// LineItemsProperty object
type LineItemsProperty struct {
	Name                      string `json:"name"`
	HsProductId               string `json:"hs_product_id"`
	HsRecurringBillingPeriod  string `json:"hs_recurring_billing_period"` // Term
	Recurringbillingfrequency string `json:"recurringbillingfrequency"`
	Quantity                  string `json:"quantity"`
	Price                     string `json:"price"` // UNIT price
	NumberOfUnits             string `json:"number_of_units"`
	TypeOfUnit                string `json:"type_of_unit"`
	Amount                    string `json:"amount"`                     // Net price
	HsACV                     string `json:"hs_acv"`                     // Annual contract value
	HsCostOfGoodsSold         string `json:"hs_cost_of_goods_sold"`      // Unit cost
	HsDiscountPercentage      string `json:"hs_discount_percentage"`     // Discount Percentage
	HsLineItemCurrencyCode    string `json:"hs_line_item_currency_code"` // Currency Code
	HsMarginTCV               string `json:"hs_margin_tcv"`              // Total contract value margin
	HsMRR                     string `json:"hs_mrr"`                     // Monthly recurring revenue
	HsPreDiscountAmount       string `json:"hs_pre_discount_amount"`     // Pre Discount Amount
	HsTCV                     string `json:"hs_tcv"`                     // Total contract value
	HsTotalDiscount           string `json:"hs_total_discount"`          // Calculated Total Discount
	SchoolYear                string `json:"school_year"`                // School Year
	Tax                       string `json:"tax"`                        // Tax
	HsSku                     string `json:"hs_sku"`                     // Sku
	SkuCode                   string `json:"sku_code"`                   // Insight SKU code
	SkuType                   string `json:"sku_type"`
	StartDate                 string `json:"start_date"`
	EndDate                   string `json:"end_date"`
	Subject                   string `json:"subject"`
	CourseId                  string `json:"course_id"`
	Discount                  string `json:"discount"`
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

// LineItemAssociateResponse object
type LineItemAssociateResponse struct {
	ErrorResponse
	Id           string            `json:"id"`
	Properties   LineItemsProperty `json:"properties"`
	CreatedAt    string            `json:"createdAt"`
	UpdatedAt    string            `json:"updatedAt"`
	Archived     bool              `json:"archived"`
	Associations Associations      `json:"associations"`
}

type LineItemAssociationDeal struct {
	Deals LineItemAssociationDealProperty `json:"deals"`
}
type LineItemAssociationDealProperty struct {
	Results []LineItemAssociationDealResult `json:"results"`
}
type LineItemAssociationDealResult struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}
