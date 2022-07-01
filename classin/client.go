package classin

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	Host                  = "https://api.eeo.cn"
	DefaultPasswordSuffix = "@Marathon"
	Country               = "0065"
	SID1                  = "37216036"
	Secret1               = "6G91jkn4"
	SID2                  = "37333936"
	Secret2               = "6rDUcTJu"
)

type ClientConfig struct {
	ApiHost       string
	SIDSchool1    string
	SecretSchool1 string
	SIDSchool2    string
	SecretSchool2 string
	HTTPClient    *http.Client
}

func NewClientConfig(apiHost, sidSchool1, secretSchool1, sidSchool2, secretSchool2 string) ClientConfig {
	if apiHost == "" {
		apiHost = Host
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
		SIDSchool1:    sidSchool1,
		SecretSchool1: secretSchool1,
		SIDSchool2:    sidSchool2,
		SecretSchool2: secretSchool2,
		ApiHost:       apiHost,
		HTTPClient:    httpClient,
	}
	return r
}

// Client object
type client struct {
	config ClientConfig
}

type Client interface {
	Students() Students
	GetMessage(code int) string
}

// NewClient constructor
func NewClient(config ClientConfig) Client {
	return &client{
		config: config,
	}
}

// Request executes any Classin API method using the current client configuration
func (c client) request(method, endpoint string, data map[string]string, response interface{}, school int) error {
	// Init request object
	var req *http.Request
	timeStamp := strconv.Itoa(int(math.Round(float64(time.Now().Unix()))))
	v := url.Values{}
	if school == 2 {
		safeKeyHash := md5.Sum([]byte(c.config.SecretSchool2 + timeStamp))
		safeKey := strings.ToLower(hex.EncodeToString(safeKeyHash[:]))
		v.Add("SID", c.config.SIDSchool2)
		v.Add("timeStamp", timeStamp)
		v.Add("safeKey", safeKey)
	} else {
		safeKeyHash := md5.Sum([]byte(c.config.SecretSchool1 + timeStamp))
		safeKey := strings.ToLower(hex.EncodeToString(safeKeyHash[:]))
		v.Add("SID", c.config.SIDSchool1)
		v.Add("timeStamp", timeStamp)
		v.Add("safeKey", safeKey)
	}
	// Create form value
	for key, val := range data {
		v.Add(key, val)
	}

	req, err := http.NewRequest(method, c.config.ApiHost+endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return fmt.Errorf("classin.Client.go.Request(): http.NewRequest(): %v", err)
	}
	// Headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// Execute and read response body
	fmt.Println(req.Form.Encode())
	resp, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("classin.Client.go.Request(): Do(req): %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("classin.Client.go.Request(): ioutil.ReadAll(): %v", err)
	}
	// Get data?
	if response != nil {
		err = json.Unmarshal(body, &response)
		if err != nil {
			return fmt.Errorf("classin.Client.go.Request(): json.Unmarshal(): %v \n%s", err, string(body))
		}
	}

	// Return HTTP errors
	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 204 {
		return fmt.Errorf("Classin API error: %d - %s \n%s", resp.StatusCode, resp.Status, string(body))
	}

	//Done!
	return nil
}
