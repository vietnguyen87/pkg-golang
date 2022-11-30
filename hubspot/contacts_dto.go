package hubspot

// ContactsRequest object
type ContactsRequest struct {
	Properties ContactsRequestProperty `json:"properties"`
}

type ContactsRequestProperty struct {
	Company               string `json:"company,omitempty" example:"Biglytics"`
	Email                 string `json:"email,omitempty" example:"bcooper@biglytics.net"`
	Firstname             string `json:"firstname,omitempty" example:"Bryan"`
	Lastname              string `json:"lastname,omitempty" example:"Cooper"`
	Phone                 string `json:"phone,omitempty" example:"(877) 929-0687"`
	LDBAccount            string `json:"ldb_account,omitempty" example:"(877) 929-0687"`
	Website               string `json:"website,omitempty" example:"biglytics.net"`
	ClassinAccountId      string `json:"classin_account_id,omitempty" example:"biglytics.net"`
	ClassinVirtualAccount string `json:"classin_virtual_account,omitempty" example:"biglytics.net"`
	ClassinPassword       string `json:"classin_password,omitempty" example:"biglytics.net"`
	SubjectTobeRenewed    string `json:"subject_to_be_renewed,omitempty" example:"biglytics.net"`
	Grade                 string `json:"grade,omitempty" example:"biglytics.net"`
	GradeNew              string `json:"grade_new,omitempty" example:"biglytics.net"`
	InterestOnSubjects    string `json:"interest_on_subjects,omitempty" example:"biglytics.net"`
	HistoricalSmsSend     string `json:"historical_sms_send,omitempty" example:"biglytics.net"`
	UtmCampaign           string `json:"utm_campaign,omitempty" example:"biglytics.net"`
	UtmContent            string `json:"utm_content,omitempty" example:"biglytics.net"`
	UtmMedium             string `json:"utm_medium,omitempty" example:"biglytics.net"`
	UtmSource             string `json:"utm_source,omitempty" example:"biglytics.net"`
	UtmSourcePage         string `json:"utm_source_page,omitempty" example:"biglytics.net"`
	UtmTerm               string `json:"utm_term,omitempty" example:"biglytics.net"`
	DigitalTrackingStg    string `json:"digital_tracking,omitempty" example:"biglytics.net"`
	DigitalTrackingProd   string `json:"doi_chi_vy_dat_ten,omitempty" example:"biglytics.net"`
	HubspotOwnerId        string `json:"hubspot_owner_id,omitempty"  example:"biglytics.net"`
}

// ContactsResponse object
type ContactsResponse struct {
	ErrorResponse
	Id                    string                        `json:"id"`
	Properties            ContactsResponseProperty      `json:"properties"`
	PropertiesWithHistory ContactsPropertiesWithHistory `json:"propertiesWithHistory"`
	CreatedAt             string                        `json:"createdAt"`
	UpdatedAt             string                        `json:"updatedAt"`
	Archived              bool                          `json:"archived"`
}

type ContactsPropertiesWithHistory struct {
	ECSource                  []PropertiesResponse `json:"ec_source"`
	ECStatus                  []PropertiesResponse `json:"ec_status"`
	Lifecyclestage            []PropertiesResponse `json:"lifecyclestage"`
	RecentConversionEventName []PropertiesResponse `json:"recent_conversion_event_name"`
	HSLatestSource            []PropertiesResponse `json:"hs_latest_source"`
	HSLatestSourceData1       []PropertiesResponse `json:"hs_latest_source_data_1"`
	HSLatestSourceData2       []PropertiesResponse `json:"hs_latest_source_data_2"`
	HSOwnerIds                []PropertiesResponse `json:"hubspot_owner_id"`
	HSOwnerAssignedDate       []PropertiesResponse `json:"hubspot_owner_assigneddate"`
}

type PropertiesResponse struct {
	Value      string `json:"value"`
	Timestamp  string `json:"timestamp"`
	SourceType string `json:"sourceType"`
	SourceId   string `json:"sourceId"`
}

type ContactsResponseProperty struct {
	Company               string `json:"company" example:"Biglytics"`
	Createdate            string `json:"createdate" example:"2019-10-30T03:30:17.883"`
	Email                 string `json:"email" example:"bcooper@biglytics.net"`
	Firstname             string `json:"firstname" example:"Bryan"`
	Lastname              string `json:"lastname" example:"Cooper"`
	Lastmodifieddate      string `json:"lastmodifieddate" example:"2019-12-07T16:50:06.678Z"`
	Phone                 string `json:"phone" example:"(877) 929-0687"`
	Website               string `json:"website" example:"biglytics.net"`
	ClassinAccountId      string `json:"classin_account_id"`
	ClassinAccount        string `json:"classin_account"`
	ClassinVirtualAccount string `json:"classin_virtual_account"`
	ClassinPassword       string `json:"classin_password"`
	ClassinAddDate        string `json:"classin_add_date"`
	ClassinRemoveDate     string `json:"classin_remove_date"`
	TypeOfUser            string `json:"type_of_user"`
	Grade                 string `json:"grade"`
	HsLatestSource        string `json:"hs_latest_source"`
	HsLatestSource1       string `json:"hs_latest_source_1"`
	HsLatestSource2       string `json:"hs_latest_source_2"`
	HubspotOwnerId        string `json:"hubspot_owner_id"`
	LDBAccount            string `json:"ldb_account"` // Phone sign in on ldb
}

// AssociatedCompany object
type AssociatedCompany struct {
	CompanyID  int                       `json:"company-id"`
	PortalID   int                       `json:"portal-id"`
	Properties AssociatedCompanyProperty `json:"properties"`
}
type ObjectValue struct {
	Value string `json:"value"`
}
type AssociatedCompanyProperty struct {
	HsNumOpenDeals               ObjectValue `json:"hs_num_open_deals"`
	FirstContactCreatedate       ObjectValue `json:"first_contact_createdate"`
	Website                      ObjectValue `json:"website"`
	HsLastmodifieddate           ObjectValue `json:"hs_lastmodifieddate"`
	HsNumDecisionMakers          ObjectValue `json:"hs_num_decision_makers"`
	NumAssociatedContacts        ObjectValue `json:"num_associated_contacts"`
	NumConversionEvents          ObjectValue `json:"num_conversion_events"`
	Domain                       ObjectValue `json:"domain"`
	HsNumChildCompanies          ObjectValue `json:"hs_num_child_companies"`
	HsNumContactsWithBuyingRoles ObjectValue `json:"hs_num_contacts_with_buying_roles"`
	HsObjectId                   ObjectValue `json:"hs_object_id"`
	Createdate                   ObjectValue `json:"createdate"`
	HsNumBlockers                ObjectValue `json:"hs_num_blockers"`
}

// CreateOrUpdateContactResponse object
type CreateOrUpdateContactResponse struct {
	ErrorResponse
	Vid   int  `json:"vid"`
	IsNew bool `json:"isNew"`
}

// DeleteContactResponse object
type DeleteContactResponse struct {
	ErrorResponse
	Vid     int    `json:"vid"`
	Deleted bool   `json:"deleted"`
	Reason  string `json:"reason"`
}

// IdentityProfile response object
type IdentityProfile struct {
	Identities []struct {
		Timestamp int64  `json:"timestamp"`
		Type      string `json:"type"`
		Value     string `json:"value"`
	} `json:"identities"`
	Vid int `json:"vid"`
}

// ContactProperty response object
type ContactProperty struct {
	Value    string `json:"value"`
	Versions []struct {
		Value       string      `json:"value"`
		Timestamp   int64       `json:"timestamp"`
		SourceType  string      `json:"source-type"`
		SourceID    interface{} `json:"source-id"`
		SourceLabel interface{} `json:"source-label"`
		Selected    bool        `json:"selected"`
	} `json:"versions"`
}

type SearchContactRequest struct {
	FilterGroups []SearchContactFilterGroups `json:"filterGroups,omitempty"`
	Sorts        []string                    `json:"sorts,omitempty"`
	Query        string                      `json:"query,omitempty"`
	Properties   []string                    `json:"properties,omitempty"`
	Limit        int                         `json:"limit,omitempty"`
	After        int                         `json:"after,omitempty"`
}

type SearchContactFilterGroups struct {
	Filters []SearchContactFilter `json:"filters,omitempty"`
}

type SearchContactFilter struct {
	Value        string   `json:"value,omitempty"`
	Values       []string `json:"values,omitempty"`
	PropertyName string   `json:"propertyName,omitempty"`
	Operator     string   `json:"operator,omitempty"`
}

type GetListContactResponse struct {
	Results []ContactsResponse `json:"results"`
	Paging  ContactPagination  `json:"paging"`
}

type SearchContactResponse struct {
	Total   int                `json:"total"`
	Results []ContactsResponse `json:"results"`
	Paging  ContactPagination  `json:"paging"`
}

type ContactPagination struct {
	Next ContactPaginationNext `json:"next"`
}
type ContactPaginationNext struct {
	After string `json:"after"`
	Link  string `json:"link"`
}
type GetContactsByIdsResponse struct {
	Status  string             `json:"status"`
	Results []ContactsResponse `json:"results"`
}
type GetByIdsContactsRequest struct {
	Properties []string                `json:"properties"`
	Inputs     []GetByIdsContactsInput `json:"inputs"`
}
type GetByIdsContactsInput struct {
	Id string `json:"id"`
}
type ContactAssociateProperties struct {
	Createdate       string `json:"createdate,omitempty"`
	HsObjectId       string `json:"hs_object_id,omitempty"`
	Lastmodifieddate string `json:"lastmodifieddate,omitempty"`
}
type ContactAssociateResponse struct {
	ErrorResponse
	Id         string                     `json:"id"`
	Properties ContactAssociateProperties `json:"properties"`
	CreatedAt  string                     `json:"createdAt"`
	UpdatedAt  string                     `json:"updatedAt"`
	Archived   bool                       `json:"archived"`
}
