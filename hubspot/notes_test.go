package hubspot

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type NotesTestSuite struct {
	suite.Suite
	client Client
}

func TestNotesTestSuite(t *testing.T) {
	suite.Run(t, new(NotesTestSuite))
}

func (c *NotesTestSuite) SetupTest() {
	c.client = NewClient(NewClientConfig(ApiHost, ApiKey))
}

func (c *NotesTestSuite) TearDownTest() {
}

func (c *NotesTestSuite) TestGetOne() {
	c.Run("Test get one with id: 22106518880", func() {
		c.SetupTest()
		id := "22106518880"
		r, err := c.client.Notes().GetOne(id)
		c.Suite.Equal(r.Id, id)
		c.Suite.NoError(err)
	})
}

func (c *NotesTestSuite) TestCreate() {
	c.Run("Test create", func() {
		c.SetupTest()
		notesRequest := NotesRequest{
			Properties: NotesProperties{
				Timestamp: time.Now().Format("2006-01-02T15:04:05Z"),
				Body:      "Test Create Note",
				OwnerId:   "0908090909",
			},
		}
		r, err := c.client.Notes().Create(notesRequest)
		c.Suite.NoError(err)
		fmt.Printf("Response: %v", r)
	})
}

func (c *NotesTestSuite) TestAssociation() {
	c.Run("Test Association", func() {
		c.SetupTest()
		r, err := c.client.Notes().Association("23272220714", "contact", "47351", "202")
		c.Suite.NoError(err)
		fmt.Printf("Response: %v", r)
	})
}
