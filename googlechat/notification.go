package googlechat

import (
	"context"
	"gitlab.marathon.edu.vn/pkg/go/xhttp"
)

func SendMessageTemplate(ctx context.Context, chatGroupUrl string, data interface{}) {
	client := xhttp.NewClient()
	reqOptions := xhttp.RequestOption{
		Header: map[string]string{
			"Content-Type": "application/json; charset=UTF-8",
		},
	}
	_, _ = client.PostJSON(ctx, chatGroupUrl, data, nil, reqOptions)

}
