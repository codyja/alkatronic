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
	"os"
	"strings"
	"sync"
)

const (
	apiUrl = "https://alkatronic.focustronic.com"
)

// FocustronicClient is the client implementation that calls Focustronic
type FocustronicClient struct {
	c           *http.Client
	baseURL     *url.URL
	accessToken string
	m           sync.Mutex
}

// NewFocustronicClient returns a new Client
func NewFocustronicClient() (*FocustronicClient, error) {
	parsedURL, err := url.Parse(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse baseURL: %w", err)
	}

	h := &http.Client{}
	c := &FocustronicClient{
		c:       h,
		baseURL: parsedURL,
	}

	return c, nil
}

// BaseURL returns the Client's configured API URL
func (c *FocustronicClient) BaseURL() *url.URL {
	u, _ := url.Parse(c.baseURL.String())
	return u
}

// AccessToken returns the API Access Token stored in the Client
func (c *FocustronicClient) AccessToken() string {
	return c.accessToken
}

// SetAccessToken sets c.accessToken to use an existing token read from home directory
func (c *FocustronicClient) SetAccessToken(token string) {
	c.accessToken = token
}

// Do executes an HTTP request with the Client's HTTP client.
// Do automatically adds the Token as a cookie to each request
func (c *FocustronicClient) Do(req *http.Request) (*http.Response, error) {
	if c.accessToken != "" {
		req.Header.Set("User-Agent", "Focustronic_GO_API/1.0")

		cookie := http.Cookie{Name: "token", Value: c.accessToken}
		req.AddCookie(&cookie)
	}

	return c.c.Do(req)
}

// DoRequest executes an HTTP request and decodes the response
func (c *FocustronicClient) DoRequest(req *http.Request) (*http.Response, error) {
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

// Login uses the provided username and password to auth against the Focustronic site API
// and retrieves a single token. The token is stored internal to the Client. The username/password
// are encoded and sent as form data.
func (c *FocustronicClient) Login(username, password string) {
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

	log.Printf("Starting authentication against Focustronic's API")
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

// Authenticate performs initial authentication or reauthentication as needed
func (c *FocustronicClient) Authenticate(username string, password string) {
	// Read in home directory to read and write token file to
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting home directory: %s", err)
	}

	// Location to store token inside file
	tokenFileLocation := fmt.Sprintf("%s/.focustronic", home)

	// Read existing token from token storage file
	existingToken, err := ioutil.ReadFile(tokenFileLocation)
	if err != nil {
		log.Printf("Checking for existing token file: %s", err)
	}

	// Authenticate if there's no existing token file, otherwise call SetAccessToken to reuse existing token on subsequent calls
	if string(existingToken) == "" {
		log.Printf("No existing token file found, logging into Focustronic API.")
		c.Login(username, password)

		// write token locally
		tokenBytes := []byte(fmt.Sprintf("%s",c.AccessToken()))
		err = ioutil.WriteFile(tokenFileLocation, tokenBytes, 0600)
		if err != nil {
			log.Fatalf("Error writting token file: %s", err)
		}
		log.Printf("Saved access token to following location for reuse: %s", tokenFileLocation)
	} else {
		log.Printf("Using existing access token read from %s", tokenFileLocation)
		c.SetAccessToken(string(existingToken))
	}

}

////GetAllDevices returns all the user's devices
//func (c *FocustronicClient) GetAllDevices() (*Devices, error) {
//	p, _ := url.Parse("/api/v2/users/self/devices")
//
//	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
//	if err != nil {
//		fmt.Println("error in request") // TODO: cleanup!!!!!!!!!!
//		return nil, err
//	}
//
//	// Add token to query string
//	q := req.URL.Query()
//	q.Add("token", c.accessToken)
//	req.URL.RawQuery = q.Encode()
//
//	resp, err := c.DoRequest(req)
//	if err != nil {
//		fmt.Errorf("error making http call: %w", err)
//	}
//
//    var d *Devices
//	err = json.NewDecoder(resp.Body).Decode(&d)
//	if err != nil {
//		fmt.Errorf("error decoding response: %w", err)
//	}
//
//	return d, err
//
//}

////ListDevices returns all the user's devices in a simple ID: Name format
//func (c *FocustronicClient) ListDevices(devices *Devices) (d map[int]string) {
//
//	var deviceMap = map[int]string{}
//
//	for _, data := range devices.Data {
//		for _, device := range data.Devices {
//			fmt.Println(device.FriendlyName, device.Id)
//			deviceMap[device.Id] = device.FriendlyName
//		}
//	}
//
//	return deviceMap
//}

//// GetDevices calls the /devices endpoint and returns the user's own devices
//func (c *FocustronicClient) GetDevices() (*Devices, error) {
//	p, _ := url.Parse("/users/devices")
//
//	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
//	if err != nil {
//		fmt.Println("error in request")
//		return nil, err
//	}
//	// Add token to query string to get devices
//	q := req.URL.Query()
//	q.Add("token", c.accessToken)
//	req.URL.RawQuery = q.Encode()
//
//	resp, err := c.DoRequest(req)
//	if err != nil {
//		return nil, fmt.Errorf("error making http call: %w", err)
//	}
//
//	var d *Devices
//	err = json.NewDecoder(resp.Body).Decode(&d)
//	if err != nil {
//		return nil, fmt.Errorf("GetDevices error decoding response: %w", err)
//	}
//
//	return d, err
//}

//// GetRecords calls the /records endpoint for the specified device and days
//func (c *FocustronicClient) GetRecords(deviceID int, days int) (*Records, error) {
//	p, _ := url.Parse(fmt.Sprintf("/users/devices/%d/records/%d", deviceID, days))
//
//	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
//	if err != nil {
//		fmt.Println("must be an errors at req")
//		return nil, err
//	}
//
//	resp, err := c.DoRequest(req)
//	if err != nil {
//		return nil, fmt.Errorf("error making http call: %w", err)
//	}
//
//	var r *Records
//	err = json.NewDecoder(resp.Body).Decode(&r)
//	if err != nil {
//		return nil, fmt.Errorf("GetRecords error decoding response: %w", err)
//	}
//
//	return r, err
//}

//// GetLatestResult calls the GetRecords func, iterates over the dates and returns the most recent Record
//func (c *FocustronicClient) GetLatestResult(deviceID int) (Record, error) {
//	records, err := c.GetRecords(deviceID, 7)
//	if err != nil {
//		log.Fatalf("error retrieving latest record: %s", err)
//	}
//
//	var dates []int64
//	for _, v := range records.Data {
//		dates = append(dates, v.CreateTime)
//	}
//
//	var latest int64 = 0
//	for _, v := range dates {
//		if latest < v {
//			latest = v
//		}
//	}
//
//	r := Record{}
//	for _, record := range records.Data {
//		if record.CreateTime == latest {
//			r = record
//		}
//	}
//
//	return r, err
//}
//
//// ConvertKh takes in the reported KH value and converts to dKh
//func ConvertKh(kh float64) float64 {
//	return kh / 100.0
//}
