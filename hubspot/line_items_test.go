package hubspot

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type LineItemsTestSuite struct {
	suite.Suite
	client Client
}

func TestLineItemsTestSuite(t *testing.T) {
	suite.Run(t, new(ContactsTestSuite))
}

func (c *LineItemsTestSuite) SetupTest() {
	c.client = NewClient(NewClientConfig(ApiHost, ApiKey))
}

func (c *LineItemsTestSuite) TestLineItems() {
	dataProduct := ProductsRequest{
		Properties: ProductsProperty{
			Description:       "Onboarding service for data product",
			HsCostOfGoodsSold: "600.00",
			HsSku:             "191902",
			Name:              "Implementation Service",
			Price:             "6000.00",
		},
	}
	// Create new line items
	rProduct, err := c.client.Products().Create(dataProduct)

	if err != nil {
		c.Suite.Error(err)
	}
	data := LineItemsRequest{
		Properties: LineItemsProperty{
			Name:                      "1 year implementation consultation",
			HsProductId:               rProduct.Id,
			Recurringbillingfrequency: "monthly",
			Quantity:                  "2",
			Price:                     "6000.00",
		},
	}
	id := ""
	c.Run("Test create new line item success", func() {
		c.SetupTest()
		r, err := c.client.LineItems().Create(data)

		c.Suite.NotEqual(r.Id, "")
		c.Suite.Equal(err, nil)
		c.Suite.Equal(r.Properties.Name, data.Properties.Name)

		id = r.Id
	})

	c.Run("Test update line item success", func() {
		c.SetupTest()
		data.Properties.Name = data.Properties.Name + " updated"

		r, err := c.client.LineItems().Update(id, data)

		c.Suite.NotEqual(r.Id, "")
		c.Suite.NotEqual(r.Properties.Name, data.Properties.Name)
		c.Suite.Equal(err, nil)

		id = r.Id
	})

	c.Run("Test delete line item success", func() {
		c.SetupTest()
		err := c.client.LineItems().Delete(id)

		c.Suite.Equal(err, nil)
	})
	//c := NewClient(NewClientConfig())
	//// Create new line items
	//rProduct, err := c.Products().Create(dataProduct)
	//
	//data := LineItemsRequest{
	//	Properties: LineItemsProperty{
	//		Name:        "1 year implementation consultation",
	//		HsProductId: rProduct.Id,
	//		//HsRecurringBillingPeriod: "24",
	//		Recurringbillingfrequency: "monthly",
	//		Quantity:                  "2",
	//		Price:                     "6000.00",
	//	},
	//}
	//if r.ErrorResponse.Status == "error" {
	//	t.Error(r.ErrorResponse.Message)
	//}
	//
	//if r.Id != "" {
	//	// Get Deal by ID
	//	lineItem, err := c.LineItems().Get(r.Id)
	//	if err != nil {
	//		t.Error(err)
	//	}
	//	if lineItem.ErrorResponse.Status == "error" {
	//		t.Error(r.ErrorResponse.Message)
	//	}
	//}
	//
	//data.Properties.Name = data.Properties.Name + " updated"
	//
	//if r.Id != "" {
	//	// Update line item by id
	//	r, err = c.LineItems().Update(r.Id, data)
	//	if err != nil {
	//		t.Error(err)
	//	}
	//	if r.ErrorResponse.Status == "error" {
	//		t.Error(r.ErrorResponse.Message)
	//	}
	//}
	//
	//if r.Id != "" {
	//	// Delete line items by Id
	//	err = c.LineItems().Delete(r.Id)
	//	if err != nil {
	//		t.Error(err)
	//	}
	//}
	//
	//t.Logf("%+v", r)
	//
	//// Clear product after testing
	//if rProduct.Id != "" {
	//	// Delete product by Id
	//	err = c.Products().Delete(rProduct.Id)
	//	if err != nil {
	//		t.Error(err)
	//	}
	//}

}
