package hubspot

import (
	"fmt"
	"net/url"
)

type Forms interface {
	Submit(data url.Values, portalId, formId string) (int, error)
}

type forms struct {
	client
}

// Forms constructor
func (c client) Forms() Forms {
	return &forms{
		client: c,
	}
}

// Submit a form
func (c *forms) Submit(data url.Values, portalId, formId string) (statusCode int, err error) {
	statusCode, err = c.client.submitForm(fmt.Sprintf("/uploads/form/v2/%s/%s", portalId, formId), data, nil)
	return statusCode, err
}
