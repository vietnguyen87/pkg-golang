package hubspot

// DealsRequest object
type DealsRequest struct {
	Properties DealsProperty `json:"properties"`
}

// DealsProperty object
type DealsProperty struct {
	Amount                 string `json:"amount"`
	Closedate              string `json:"closedate"`
	Dealname               string `json:"dealname"`
	Dealstage              string `json:"dealstage"`
	HubspotOwnerId         string `json:"hubspot_owner_id"`
	Pipeline               string `json:"pipeline"`
	AffiliateSName         string `json:"affiliate_s_name"`
	QualifiedForCommission string `json:"qualified_for_commission"`
	B2bDealId              string `json:"b2b_deal_id"`
	ApprovedCommission     string `json:"approved_commission"`
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
