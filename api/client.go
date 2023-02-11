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
