package hubspot

import "fmt"

// Contacts ContactsHubspot interface
type Contacts interface {
	Get(contactID string) (ContactsResponse, error)
	Create(data ContactsRequest) (ContactsResponse, error)
	Search(body SearchContactRequest) (SearchContactResponse, error)
	GetByIds(ids []string) (GetContactsByIdsResponse, error)
	GetByEmail(email string) (ContactsResponse, error)
	Update(contactID string, data ContactsRequest) error
	UpdateByEmail(email string, data ContactsRequest) error
	CreateOrUpdate(email string, data ContactsRequest) (CreateOrUpdateContactResponse, error)
	Delete(contactID string) error
	Associate(contactId, toObjectType, toObjectId, associationType string) (ContactAssociateResponse, error)
	GetList(after, limit string) (GetListContactResponse, error)
}
type contacts struct {
	client
}

// Contacts constructor
func (c client) Contacts() Contacts {
	return &contacts{
		client: c,
	}
}

// Get List Contact
func (c *contacts) GetList(after, limit string) (GetListContactResponse, error) {
	r := GetListContactResponse{}
	params := []string{
		"propertiesWithHistory=ec_source",
		"propertiesWithHistory=ec_status",
		"propertiesWithHistory=lifecyclestage",
		"propertiesWithHistory=recent_conversion_event_name",
		"propertiesWithHistory=hs_latest_source",
		"propertiesWithHistory=hs_latest_source_data_1",
		"propertiesWithHistory=hs_latest_source_data_2",
		"propertiesWithHistory=hubspot_owner_id",
		"after=" + after,
		"limit=" + limit,
	}
	err := c.client.request("GET", "/crm/v3/objects/contacts/", nil, &r, params)
	return r, err
}

// Get Contact
func (c *contacts) Get(contactID string) (ContactsResponse, error) {
	r := ContactsResponse{}
	params := []string{
		"properties=firstname",
		"properties=lastname",
		"properties=email",
		"properties=phone",
		"properties=classin_account_id",
		"properties=classin_account",
		"properties=classin_virtual_account",
		"properties=classin_password",
		"properties=classin_add_date",
		"properties=classin_remove_date",
		"properties=type_of_user",
		"properties=hubspot_owner_id",
		"propertiesWithHistory=hubspot_owner_id",
		"propertiesWithHistory=hubspot_owner_assigneddate",
	}
	err := c.client.request("GET", "/crm/v3/objects/contacts/"+contactID, nil, &r, params)
	return r, err
}

// Search a Contact
func (c *contacts) Search(body SearchContactRequest) (SearchContactResponse, error) {
	r := SearchContactResponse{}
	err := c.client.request("POST", "/crm/v3/objects/contacts/search", body, &r, nil)
	return r, err
}

// GetByIds Get multi contact by ids
func (c *contacts) GetByIds(ids []string) (GetContactsByIdsResponse, error) {
	r := GetContactsByIdsResponse{}
	var inputs []GetByIdsContactsInput
	for _, v := range ids {
		inputs = append(inputs, GetByIdsContactsInput{Id: v})
	}
	data := GetByIdsContactsRequest{
		Inputs: inputs,
		Properties: []string{
			"firstname",
			"lastname",
			"phone",
			"classin_account_id",
			"classin_account",
			"classin_password",
			"classin_virtual_account",
			"company",
			"email",
			"createdate",
			"lastmodifieddate",
			"classin_add_date",
			"classin_remove_date",
			"hubspot_owner_id",
		},
	}
	err := c.client.request("POST", "/crm/v3/objects/contacts/batch/read", data, &r, nil)
	return r, err
}

// GetByEmail a Contact [deprecated]
func (c *contacts) GetByEmail(email string) (ContactsResponse, error) {
	r := ContactsResponse{}
	err := c.client.request("GET", "/contacts/v1/contact/email/"+email+"/profile", nil, &r, nil)
	return r, err
}

// Create new Contact [deprecated]
func (c *contacts) Create(data ContactsRequest) (ContactsResponse, error) {
	r := ContactsResponse{}
	err := c.client.request("POST", "/crm/v3/objects/contacts", data, &r, nil)
	return r, err
}

// Update Contact
func (c *contacts) Update(contactID string, data ContactsRequest) error {
	return c.client.request("PATCH", "/crm/v3/objects/contacts/"+contactID, data, nil, nil)
}

// UpdateByEmail a Contact [deprecated]
func (c *contacts) UpdateByEmail(email string, data ContactsRequest) error {
	return c.client.request("POST", "/contacts/v1/contact/email/"+email+"/profile", data, nil, nil)
}

// CreateOrUpdate a Contact [deprecated]
func (c *contacts) CreateOrUpdate(email string, data ContactsRequest) (CreateOrUpdateContactResponse, error) {
	r := CreateOrUpdateContactResponse{}
	err := c.client.request("POST", "/contacts/v1/contact/createOrUpdate/email/"+email, data, &r, nil)
	return r, err
}

// Delete Contact
func (c *contacts) Delete(contactID string) error {
	return c.client.request("DELETE", "/crm/v3/objects/contacts/"+contactID, nil, nil, nil)
}

// Associate Create new Line Items
// toObjectType: DEAL
// associationType: line_item_to_deal
func (c *contacts) Associate(contactId, toObjectType, toObjectId, associationType string) (ContactAssociateResponse, error) {
	r := ContactAssociateResponse{}
	err := c.client.request("PUT",
		fmt.Sprintf("/crm/v3/objects/contacts/%s/associations/%s/%s/%s",
			contactId,
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
