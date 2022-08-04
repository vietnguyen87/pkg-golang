package hubspot

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type EngagementsTestSuite struct {
	suite.Suite
	client Client
}

func TestEngagementsTestSuite(t *testing.T) {
	suite.Run(t, new(EngagementsTestSuite))
}

func (c *EngagementsTestSuite) SetupTest() {
	c.client = NewClient(NewClientConfig(ApiHost, ApiKey))
}

func (c *EngagementsTestSuite) TestCreateEngagements() {
	data := EngagementsRequest{
		EngagementProperties{
			Engagement: Engagement{
				Active:    true,
				OwnerId:   1,
				Type:      "TASK",
				Timestamp: time.Now().UnixMilli(),
			},
			Associations: EngagementAssociations{
				DealIds: []int64{9619060375},
			},
			Metadata: Metadata{
				Body:          "This is the body of the task.",
				Subject:       "This is the subject like of the task",
				Status:        "NOT_STARTED",
				ForObjectType: "CONTACT",
			},
		},
	}
	c.Run("Test create new engagement successful", func() {
		c.SetupTest()
		engagement, err := c.client.Engagements().Create(data)
		c.Suite.NoError(err)
		c.Suite.Equal("This is the body of the task.", engagement.Metadata.Body)
	})
}
