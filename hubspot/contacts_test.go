package hubspot

import (
	"github.com/leonelquinteros/gorand"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

type ContactsTestSuite struct {
	suite.Suite
	client Client
}

func TestContactsTestSuite(t *testing.T) {
	suite.Run(t, new(ContactsTestSuite))
}

func (c *ContactsTestSuite) SetupTest() {
	//c.client = NewClient(NewClientConfig(ApiHost, "6af298d2-be72-4246-94a6-5550063e251d"))
	c.client = NewClient(NewClientConfig(ApiHost, ApiKey))
}

func (c *ContactsTestSuite) TestContacts() {

	tEmailUser, err := gorand.GetAlphaNumString(8)
	if err != nil {
		c.Suite.Error(err)
	}
	tCompanyName, err := gorand.GetAlphaNumString(9)
	if err != nil {
		c.Suite.Error(err)
	}
	tEmailHost, err := gorand.GetAlphaNumString(6)
	if err != nil {
		c.Suite.Error(err)
	}
	tPhone, err := gorand.GetNumString(10)
	if err != nil {
		c.Suite.Error(err)
	}
	tEmail := tEmailUser + "@" + tEmailHost + ".com"
	if err != nil {
		c.Suite.Error(err)
	}
	id := ""
	data := ContactsRequest{
		Properties: ContactsRequestProperty{
			Company:   tCompanyName,
			Email:     tEmail,
			Firstname: tEmailUser,
			Lastname:  tEmailUser,
			Phone:     tPhone,
			Website:   strings.ToLower(tCompanyName) + ".net",
		},
	}
	c.Run("Test create new contact successful", func() {
		c.SetupTest()
		contact, _ := c.client.Contacts().Create(data)
		c.Suite.NotEqual(contact.Id, "")
		c.Suite.NotEqual(contact.Properties.Firstname, "")
		id = contact.Id
	})
	c.Run("Test get contact successful", func() {
		c.SetupTest()
		contact, _ := c.client.Contacts().Get(id)
		c.Suite.Equal(contact.Id, id)
		c.Suite.NotEqual(contact.Properties.Firstname, "")
	})
	c.Run("Test update contact successful", func() {
		c.SetupTest()
		data.Properties.Company = data.Properties.Company + " updated"
		err := c.client.Contacts().Update(id, data)
		c.Suite.Equal(err, nil)
	})
	c.Run("Test delete contact successful", func() {
		c.SetupTest()
		err := c.client.Contacts().Delete(id)
		c.Suite.Equal(err, nil)
	})
	c.Run("Test get contact not found", func() {
		c.SetupTest()
		_, err := c.client.Contacts().Get("10003451")
		c.Suite.NotEqual(err, nil)
	})

}

//[IN, NOT_HAS_PROPERTY, LT, EQ, GT, NOT_IN, GTE, CONTAINS_TOKEN, HAS_PROPERTY, LTE, NOT_CONTAINS_TOKEN, BETWEEN, NEQ]
func (c *ContactsTestSuite) TestSearch() {
	c.Run("Test Search contact by phone", func() {
		c.SetupTest()
		contacts, err := c.client.Contacts().Search(SearchContactRequest{
			FilterGroups: []SearchContactFilterGroups{
				{
					Filters: []SearchContactFilter{
						{
							Value:        "",
							PropertyName: "classin_virtual_account",
							Operator:     "HAS_PROPERTY",
						},
						{
							Values:       []string{"opportunity", "salesqualifiedlead", "customer"},
							PropertyName: "lifecyclestage",
							Operator:     "IN",
						},
						{
							Value:        "",
							PropertyName: "invited_to_be_affiliate",
							Operator:     "NOT_HAS_PROPERTY",
						},
					},
				},
			},
			Properties: []string{"firstname", "lastname", "phone", "grade"},
			Limit:      10,
			After:      0,
		})
		c.Suite.NoError(err)
		c.Suite.Equal("0333993985", contacts.Results[0].Properties.Phone)
	})
}

func (c *ContactsTestSuite) TestUpdate() {
	c.Run("Test update contact successful", func() {
		c.SetupTest()
		data := ContactsRequest{}
		data.Properties.HubspotOwnerId = "124023325"
		err := c.client.Contacts().Update("83401", data)
		c.Suite.NoError(err)
	})
}

func (c *ContactsTestSuite) TestGet() {
	c.Run("Test update contact successful", func() {
		contact, _ := c.client.Contacts().Get("83401")
		c.Suite.Equal(contact.Id, "83401")
		c.Suite.NotEqual(contact.Properties.Firstname, "")
	})
}
