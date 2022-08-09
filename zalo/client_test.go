package zalo

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientTestSuite struct {
	suite.Suite
	client Client
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (c *ClientTestSuite) SetupTest() {
	c.client = NewClient(
		NewClientConfig("https://oauth.zaloapp.com/v4", "pqrLDU7Ufmt11SJULJMR", "2788475007218903087", nil),
	)
}

func (c *ClientTestSuite) TearDownTest() {
}

func (c *ClientTestSuite) TestRefreshAccessToken() {
	c.Run("Test RefreshAccessToken", func() {
		c.SetupTest()
		data, err := c.client.RefreshAccessToken(context.Background(), "P3NPaIFoh4b3Fup3ISojE0uOizOF-QuORH2dwrxwi1Wf0DNnVi_93aSBr_K0qyyG0JVQwpBduIPnCSBN3iEXFrCAzxTe_UrK4nJwiHF_rsTT8zQ73i7NFcjpy81uu8bAENlJvIdFc1igVVhDAA2sTcj5b8CPl-58Vd3EysoVfXnZCehEGHB461pFzNRNG0sqa89OJgGu6TBJs6Cio5iLrgIZS0YuLn6YY-OY9UXsKgfUMIOXFJme4zZ9Km")
		c.Suite.NoError(err)
		fmt.Println("data: ", data.RefreshToken, data.AccessToken)
	})
}
