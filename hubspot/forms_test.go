package hubspot

import (
	"github.com/stretchr/testify/suite"
	"net/url"
	"testing"
)

type FormsTestSuite struct {
	suite.Suite
	client Client
}

func TestFormsTestSuite(t *testing.T) {
	suite.Run(t, new(FormsTestSuite))
}

func (c *FormsTestSuite) SetupTest() {
	c.client = NewClient(NewClientConfig(ApiHost, ApiKey))
}

func (c *FormsTestSuite) TearDownTest() {
}

func (c *NotesTestSuite) TestSubmit() {
	c.Run("Test Submit", func() {
		c.SetupTest()
		data := url.Values{}
		data.Set("firstname", "Viet Nguyễn")
		data.Set("phone", "0919882581")
		data.Set("grade", "7")
		data.Set("interest_on_subjects", "Anh Văn")
		data.Set("note", "Viet Test Form")
		data.Set("mkt_campaign_engagement", "Affiliate program")
		data.Set("hs_analytics_source_data_1", "Dinos")
		data.Set("hs_analytics_source_data_2", "Dinos")
		statusCode, err := c.client.Forms().Submit(data, "21066554", "481328ea-bc6c-4d09-bf51-ff4ba3ec6675")
		c.Suite.NoError(err)
		c.Suite.Equal(1, statusCode)
	})
}
