package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const (
	apiUrl = "https://alkatronic.focustronic.com"
)

// Client is the API client for making calls to alkatronic's api
type Client interface {
	AccessToken() string
	Authenticate(username, password string) error
}

// AlkatronicClient is the client implementation that calls Alkatronic
type AlkatronicClient struct {
	c           *http.Client
	baseURL     *url.URL
	accessToken string
	m           sync.Mutex
}

// NewAlkatronicClient returns a new Client
func NewAlkatronicClient() (*AlkatronicClient, error) {
	parsedURL, err := url.Parse(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse baseURL: %w", err)
	}

	h := &http.Client{}
	c := &AlkatronicClient{
		c:       h,
		baseURL: parsedURL,
	}

	return c, nil
}

// BaseURL returns the Client's configured API URL
func (c *AlkatronicClient) BaseURL() *url.URL {
	u, _ := url.Parse(c.baseURL.String())
	return u
}

// AccessToken returns the API Access Token stored in the Client
func (c *AlkatronicClient) AccessToken() string {
	return c.accessToken
}

// SetAccessToken sets c.accessToken to use an existing token read from home directory
func (c *AlkatronicClient) SetAccessToken(token string) {
	c.accessToken = token
}

// Do executes an HTTP request with the Client's HTTP client.
// Do automatically adds the Token as a cookie to each request
func (c *AlkatronicClient) Do(req *http.Request) (*http.Response, error) {
	if c.accessToken != "" {
		req.Header.Set("User-Agent", "Alkatronic_GO_API/1.0")

		cookie := http.Cookie{Name: "token", Value: c.accessToken}
		req.AddCookie(&cookie)
	}

	return c.c.Do(req)
}

// DoRequest executes an HTTP request and decodes the response
func (c *AlkatronicClient) DoRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	var bodyBytes []byte
	if resp.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	var r Response
	_ = json.NewDecoder(resp.Body).Decode(&r)

	resp.Body.Close()
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	if r.Result != true {
		err = errors.New(r.Message)
		return nil, err
	}

	return resp, nil
}

// Authenticate uses the provided username and password to auth against the Alkatronic site API
// and retrieves a single token. The token is stored internal to the Client. The username/password
// are encoded and sent as form data.
func (c *AlkatronicClient) Authenticate(username, password string) {
	p, _ := url.Parse("/users/login")
	b := url.Values{
		"email":    {username},
		"password": {password},
	}
	req, err := http.NewRequest(http.MethodPost, c.baseURL.ResolveReference(p).String(), strings.NewReader(b.Encode()))
	if err != nil {
		fmt.Errorf("error building token generate request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	log.Printf("Starting authentication against Alkatronic's API")
	resp, err := c.DoRequest(req)
	if err != nil {
		log.Fatalf("Authentication failed: %s", err)
		return
	}

	var r Response
	_ = json.NewDecoder(resp.Body).Decode(&r)
	c.accessToken = r.Data

	log.Printf("Authentication has succeeded")

	return
}
// GetDevices calls the /devices endpoint and returns the user's own devices
func (c *AlkatronicClient) GetDevices() (*Devices, error) {
	p, _ := url.Parse("/users/devices")

	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
	if err != nil {
		fmt.Println("must be an errors at req")
		return nil, err
	}
	// Add token to query string to get devices
	q := req.URL.Query()
	q.Add("token", c.accessToken)
	req.URL.RawQuery = q.Encode()

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error making http call: %w", err)
	}

	var d *Devices
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return nil, fmt.Errorf("GetDevices error decoding response: %w", err)
	}

	return d, err
}

// GetRecords calls the /records endpoint for the specified device and days
func (c *AlkatronicClient) GetRecords(deviceID int, days int) (*Records, error) {
	p, _ := url.Parse(fmt.Sprintf("/users/devices/%d/records/%d", deviceID, days))

	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
	if err != nil {
		fmt.Println("must be an errors at req")
		return nil, err
	}

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error making http call: %w", err)
	}

	var r *Records
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("GetRecords error decoding response: %w", err)
	}

	return r, err
}

// GetLatestResult calls the GetRecords func, iterates over the dates and returns the most recent Record
func (c *AlkatronicClient) GetLatestResult(deviceID int) (Record, error) {
	records, err := c.GetRecords(deviceID, 7)

	var dates []int64
	for _, v := range records.Data {
		dates = append(dates, v.CreateTime)
	}

	var latest int64 = 0
	for _, v := range dates {
		if latest < v {
			latest = v
		}
	}

	r := Record{}
	for _, record := range records.Data {
		if record.CreateTime == latest {
			r = record
		}
	}

	return r, err
}

// ConvertKh takes in the reported KH value and converts to dKh
func ConvertKh(kh int) float32 {
	return float32(kh) / 100.0
}