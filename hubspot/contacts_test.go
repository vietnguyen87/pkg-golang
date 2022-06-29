package hubspot

import (
	"github.com/leonelquinteros/gorand"
	"github.com/stretchr/testify/suite"
	"os"
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
	_ = os.Setenv("HUBSPOT_API_KEY", "12c3033c-718e-42ec-b68d-e88ae6ef5e29")
	c.client = NewClient(NewClientConfig())
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
