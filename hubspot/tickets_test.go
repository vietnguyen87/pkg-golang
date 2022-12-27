package hubspot

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TicketsTestSuite struct {
	suite.Suite
	client Client
}

func TestTicketsTestSuite(t *testing.T) {
	suite.Run(t, new(TicketsTestSuite))
}

func (c *TicketsTestSuite) SetupTest() {
	//c.client = NewClient(NewClientConfig(ApiHost, "6af298d2-be72-4246-94a6-5550063e251d"))
	c.client = NewClient(NewClientConfig(ApiHost, ApiKey))
}

func (c *TicketsTestSuite) TestGet() {
	c.Run("Test get tickets successful", func() {
		ticket, _ := c.client.Tickets().Get("1328410192")
		c.Suite.Equal(ticket.Id, "1328410192")
		c.Suite.Equal(ticket.Properties.HubspotOwnerId, "187299234")
	})
}

func (c *TicketsTestSuite) TestGetAssociate() {
	c.Run("Test get associate successful", func() {
		contact, _ := c.client.Tickets().GetAssociate("1328410192", "contact")
		c.Suite.Equal(contact.Results[0].ToObjectId, 83401)
	})
}
