package hubspot

import (
	"testing"
)

func TestProducts(t *testing.T) {

	//data := ProductsRequest{
	//	Properties: ProductsProperty{
	//		Description:       "Onboarding service for data product",
	//		HsCostOfGoodsSold: "600.00",
	//		HsSku:             "191902",
	//		Name:              "Implementation Service",
	//		Price:             "6000.00",
	//	},
	//}
	//c := NewClient(NewClientConfig())
	//// Create new line items
	//r, err := c.Products().Create(data)
	//
	//if err != nil {
	//	t.Error(err)
	//}
	//if r.ErrorResponse.Status == "error" {
	//	t.Error(r.ErrorResponse.Message)
	//}
	//
	//if r.Id != "" {
	//	// Get product by ID
	//	lineItem, err := c.Products().Get(r.Id)
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
	//	// Update product by id
	//	r, err = c.Products().Update(r.Id, data)
	//	if err != nil {
	//		t.Error(err)
	//	}
	//	if r.ErrorResponse.Status == "error" {
	//		t.Error(r.ErrorResponse.Message)
	//	}
	//}
	//
	//if r.Id != "" {
	//	// Delete product by Id
	//	err = c.Products().Delete(r.Id)
	//	if err != nil {
	//		t.Error(err)
	//	}
	//}
	//
	//t.Logf("%+v", r)

	//c := NewClient(NewClientConfig(ApiHost, ApiKey))
	//lineItem, err := c.Products().Get("1596106416")
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//properties := map[string]interface{}{
	//	"properties": map[string]interface{}{},
	//}
	//properties["properties"] = map[string]interface{}{
	//	"number_of_units": 1.3,
	//}
	//
	//r, err := c.Products().Update("1596106416", properties)
	//if err != nil {
	//	t.Error(err)
	//}
	//if r.ErrorResponse.Status == "error" {
	//	t.Error(r.ErrorResponse.Message)
	//}
	//t.Errorf("lineItem: %v", lineItem)
}
