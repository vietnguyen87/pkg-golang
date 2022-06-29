package hubspot

// ContactsHubspot Contacts interface
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
			"company",
			"email",
			"createdate",
			"lastmodifieddate",
			"classin_add_date",
			"classin_remove_date",
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
