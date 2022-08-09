package zalo

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"gitlab.marathon.edu.vn/pkg/go/xhttp"
	"gitlab.marathon.edu.vn/pkg/go/zalo/schema"
	"net/url"
	"strconv"
)

type Client interface {
	RefreshAccessToken(ctx context.Context, refreshToken string) (schema.AccessTokenResp, error)
}

// Client object
type client struct {
	config ClientConfig
}

// ClientConfig object used for client creation
type ClientConfig struct {
	APIHost   string
	SecretKey string
	AppId     string
	client    xhttp.Client
}

func NewClientConfig(apiHost, secretKey, appId string, options []xhttp.Option) ClientConfig {
	client := xhttp.NewClient(options...)
	return ClientConfig{
		APIHost:   apiHost,
		SecretKey: secretKey,
		AppId:     appId,
		client:    client,
	}
}

// NewClient constructor
func NewClient(config ClientConfig) Client {
	return &client{
		config: config,
	}
}

func (c *client) RefreshAccessToken(ctx context.Context, refreshToken string) (response schema.AccessTokenResp, err error) {
	uri := fmt.Sprintf("%v/%v", c.config.APIHost, "oa/access_token")
	params := url.Values{}
	params.Set("app_id", c.config.AppId)
	params.Set("refresh_token", refreshToken)
	params.Set("grant_type", "refresh_token")

	reqOptions := []xhttp.RequestOption{
		{
			Header: map[string]string{
				"secret_key":     c.config.SecretKey,
				"Content-Type":   binding.MIMEPOSTForm,
				"Content-Length": strconv.Itoa(len(params.Encode())),
			},
		},
	}
	_, err = c.config.client.PostForm(ctx, uri, params, &response, reqOptions...)
	return response, err
}
