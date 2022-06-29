package hubspot

import (
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type DealsTestSuite struct {
	suite.Suite
	client Client
}

func TestDealsTestSuite(t *testing.T) {
	suite.Run(t, new(ContactsTestSuite))
}

func (c *DealsTestSuite) SetupTest() {
	_ = os.Setenv("HUBSPOT_API_KEY", "12c3033c-718e-42ec-b68d-e88ae6ef5e29")
	c.client = NewClient(NewClientConfig())
}

func (c *DealsTestSuite) TestDeals() {
	data := DealsRequest{
		Properties: DealsProperty{
			Amount:    "1500.00",
			Closedate: "2019-12-07T16:50:06.678Z",
			Dealname:  "Custom data integrations",
			//Dealstage: "presentationscheduled",
			Pipeline: "default",
		},
	}
	id := ""
	c.Run("Test create new deal success", func() {
		c.SetupTest()
		r, err := c.client.Deals().Create(data)

		c.Suite.NotEqual(r.Id, "")
		c.Suite.Equal(err, nil)
		c.Suite.Equal(r.Properties.Dealname, data.Properties.Dealname)

		id = r.Id
	})

	c.Run("Test get deal success", func() {
		c.SetupTest()
		r, err := c.client.Deals().Get(id)

		c.Suite.Equal(r.Id, id)
		c.Suite.Equal(r.Properties.Dealname, data.Properties.Dealname)
		c.Suite.Equal(err, nil)
	})

	c.Run("Test update deal success", func() {
		c.SetupTest()
		data.Properties.Dealname = data.Properties.Dealname + " updated"
		r, err := c.client.Deals().Update(id, data)

		c.Suite.Equal(r.Properties.Dealname, data.Properties.Dealname)
		c.Suite.Equal(err, nil)
	})

	c.Run("Test delete deal success", func() {
		c.SetupTest()
		err := c.client.Deals().Delete(id)

		c.Suite.Equal(err, nil)
	})
}
