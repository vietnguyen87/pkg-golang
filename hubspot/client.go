package hubspot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"gitlab.marathon.edu.vn/pkg/go/hubspot/builder"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	ApiHost         = "https://api.hubapi.com"
	ApiForm         = "https://forms.hubspot.com"
	ApiKey          = "12c3033c-718e-42ec-b68d-e88ae6ef5e29"
	GroupPathHeader = "X-Group-Path"
)

type RequestOption struct {
	Header    map[string]string
	GroupPath string
}

// ClientConfig object used for client creation
type ClientConfig struct {
	APIHost    string
	APIKey     string
	OAuthToken string
	HTTPClient *http.Client
}

// NewClientConfig constructs a ClientConfig object with the environment variables set as default
func NewClientConfig(apiHost, apiKey string) ClientConfig {
	if apiHost == "" {
		apiHost = ApiHost
	}
	if apiKey == "" {
		apiKey = ApiKey
	}
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			//Dial: (&net.Dialer{
			//	Timeout: c.config.DialTimeout,
			//}).Dial,
			IdleConnTimeout:     5 * time.Second,
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}
	r := ClientConfig{
		APIHost:    apiHost,
		APIKey:     apiKey,
		OAuthToken: "",
		HTTPClient: httpClient,
	}
	return r
}

// Client object
type client struct {
	config ClientConfig
}

type Client interface {
	Contacts() Contacts
	Deals() Deals
	LineItems() LineItems
	Products() Products
	Notes() Notes
	Forms() Forms
	Engagements() Engagements
}

// NewClient constructor
func NewClient(config ClientConfig) Client {
	return &client{
		config: config,
	}
}

// addAPIKey adds HUBSPOT_API_KEY param to a given URL.
func (c *client) addAPIKey(u string) (string, error) {
	if c.config.APIKey != "" {
		uri, err := url.Parse(u)
		if err != nil {
			return u, err
		}
		q := uri.Query()
		q.Set("hapikey", c.config.APIKey)
		uri.RawQuery = q.Encode()
		u = uri.String()
	}

	return u, nil
}

// Request executes any HubSpot API method using the current client configuration
func (c *client) request(method, endpoint string, data, response interface{}, params []string) error {
	// Build URL
	uri, err := c.buildUri(endpoint, params)
	if err != nil {
		return fmt.Errorf("hubspot.Client.go.Request(): buildUri error: %v", err)
	}
	var req *http.Request
	// Build body payload data
	bodyPayload, _ := c.buildBodyRequest(data)

	req, err = http.NewRequest(method, uri, bodyPayload)
	// Create new request (maybe with no bodyPayload)

	if err != nil {
		return fmt.Errorf("hubspot.Client.go.Request(): http.NewRequest(): %v", err)
	}
	// Headers
	req.Header.Add("Content-Type", "application/json")
	// OAuth authentication
	if c.config.APIKey == "" && c.config.OAuthToken != "" {
		req.Header.Add("Authorization", "Bearer "+c.config.OAuthToken)
	}

	// Execute and read response body
	resp, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("hubspot.Client.go.Request(): c.config.HTTPClient.Do(): %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("hubspot.Client.go.Request(): ioutil.ReadAll(): %v", err)
	}
	//log := logger.CToL(context.Background(), "request")
	// Get data?
	if response != nil {
		err = json.Unmarshal(body, &response)
		if err != nil {
			return fmt.Errorf("hubspot.Client.go.Request(): json.Unmarshal(): %v \n%s", err, string(body))
		}
	}
	// Return HTTP errors
	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 204 {
		return fmt.Errorf("HubSpot API error: %d - %s \n%s", resp.StatusCode, resp.Status, string(body))
	}
	return nil
}

// requestForm executes any Form HubSpot API method using the current client configuration
func (c *client) submitForm(endpoint string, data url.Values, response interface{}) (statusCode int, err error) {
	// Build URL
	u, err := url.Parse(ApiForm)
	if err != nil {
		return statusCode, fmt.Errorf("url.Parse(): %v", err)
	}
	u.Path = path.Join(u.Path, endpoint)
	// API Key authentication
	uri := u.String()
	if c.config.APIKey != "" {
		uri, err = c.addAPIKey(uri)
		if err != nil {
			return statusCode, fmt.Errorf("c.addAPIKey(): %v", err)
		}
	}
	if err != nil {
		return statusCode, fmt.Errorf("hubspot.Client.go.Request(): buildUri error: %v", err)
	}
	stringData := data.Encode()
	reqOpts := []RequestOption{
		{
			Header: map[string]string{
				"Content-Type":   binding.MIMEPOSTForm,
				"Content-Length": strconv.Itoa(len(stringData)),
			},
		},
	}
	header := c.getRequestHeader(reqOpts...)
	req, err := builder.NewRequestBuilder().
		WithMethod(http.MethodPost).
		WithURL(uri).
		WithHeaders(header).
		WithFormBody(binding.MIMEPOSTForm, data).
		BuildForm()
	if err != nil {
		return 0, err
	}
	return c.Do(req, response)
}

func (c *client) buildUri(endpoint string, params []string) (uri string, err error) {
	// Construct endpoint URL
	u, err := url.Parse(c.config.APIHost)
	if err != nil {
		return uri, fmt.Errorf("url.Parse(): %v", err)
	}
	u.Path = path.Join(u.Path, endpoint)
	// API Key authentication
	uri = u.String()
	if c.config.APIKey != "" {
		uri, err = c.addAPIKey(uri)
		if err != nil {
			return uri, fmt.Errorf("c.addAPIKey(): %v", err)
		}
	}
	if !strings.Contains(uri, "?") {
		uri += "?temp="
	}
	for _, v := range params {
		uri = uri + fmt.Sprintf("&%v", v)
	}
	return uri, nil
}

func (c *client) buildBodyRequest(data interface{}) (body io.Reader, err error) {
	if data == nil {
		return nil, fmt.Errorf("hubspot.Client.go.Request() with nil data")
	}
	// Encode data to JSON
	dataEncoded, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal(): %v", err)
	}
	return bytes.NewBuffer(dataEncoded), nil
}

func (c *client) Do(request *http.Request, target interface{}) (int, error) {
	rsp, err := c.config.HTTPClient.Do(request)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = rsp.Body.Close()
	}()

	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return 0, err
	}

	if len(bodyBytes) == 0 {
		return rsp.StatusCode, nil
	}

	return rsp.StatusCode, json.Unmarshal(bodyBytes, target)
}

func (c *client) getRequestHeader(reqOpts ...RequestOption) map[string]string {
	if len(reqOpts) == 0 {
		return nil
	}
	reqOpt := reqOpts[0]
	header := reqOpt.Header
	if header == nil {
		header = make(map[string]string)
	}
	if reqOpt.GroupPath != "" {
		header[GroupPathHeader] = reqOpt.GroupPath
	}
	return header
}
