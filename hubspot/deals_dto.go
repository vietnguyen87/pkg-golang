package hubspot

// DealsRequest object
type DealsRequest struct {
	Properties DealsProperty `json:"properties"`
}

// DealsProperty object
type DealsProperty struct {
	Amount                 string `json:"amount,omitempty"`
	Closedate              string `json:"closedate,omitempty"`
	Dealname               string `json:"dealname,omitempty"`
	Dealstage              string `json:"dealstage,omitempty"`
	HubspotOwnerId         string `json:"hubspot_owner_id,omitempty"`
	Pipeline               string `json:"pipeline,omitempty"`
	AffiliateSName         string `json:"affiliate_s_name,omitempty"`
	QualifiedForCommission string `json:"qualified_for_commission,omitempty"`
	B2bDealId              string `json:"b2b_deal_id,omitempty"`
	ApprovedCommission     string `json:"approved_commission,omitempty"`
}

// DealsResponse object
type DealsResponse struct {
	ErrorResponse
	Id           string         `json:"id"`
	Properties   DealProperties `json:"properties"`
	CreatedAt    string         `json:"createdAt"`
	UpdatedAt    string         `json:"updatedAt"`
	Archived     bool           `json:"archived"`
	Associations Associations   `json:"associations"`
}

// DealProperties object
type DealProperties struct {
	Amount             string `json:"amount"`
	Closedate          string `json:"closedate"`
	Createdate         string `json:"createdate"`
	Dealname           string `json:"dealname"`
	Dealstage          string `json:"dealstage"`
	HsLastmodifieddate string `json:"hs_lastmodifieddate"`
	HubspotOwnerID     string `json:"hubspot_owner_id"`
	Pipeline           string `json:"pipeline"`
}
