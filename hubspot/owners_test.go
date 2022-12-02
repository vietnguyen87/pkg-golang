package hubspot

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type OwnersTestSuite struct {
	suite.Suite
	client Client
}

func TestOwnersTestSuite(t *testing.T) {
	suite.Run(t, new(OwnersTestSuite))
}

func (c *OwnersTestSuite) SetupTest() {
	//c.client = NewClient(NewClientConfig(ApiHost, "6af298d2-be72-4246-94a6-5550063e251d"))
	c.client = NewClient(NewClientConfig(ApiHost, ApiKey))
}

func (c *OwnersTestSuite) TestContacts() {
	id := "124023325"

	c.Run("Test get owner successful", func() {
		c.SetupTest()
		owner, _ := c.client.Owners().Get(id)
		c.Suite.Equal(owner.Id, id)
		c.Suite.NotEqual(owner.Firstname, "")
	})

	c.Run("Test get owner not found", func() {
		c.SetupTest()
		_, err := c.client.Owners().Get("10003451111")
		c.Suite.NotEqual(err, nil)
	})

}
