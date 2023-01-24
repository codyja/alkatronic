package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

//GetAllDevices returns all the user's devices
func (c *FocustronicClient) GetAllDevices() (*Devices, error) {
	p, _ := url.Parse("/api/v2/users/self/devices")

	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
	if err != nil {
		fmt.Println("error in request") // TODO: cleanup!!!!!!!!!!
		return nil, err
	}

	// Add token to query string
	q := req.URL.Query()
	q.Add("token", c.accessToken)
	req.URL.RawQuery = q.Encode()

	resp, err := c.DoRequest(req)
	if err != nil {
		fmt.Errorf("error making http call: %w", err)
	}

	var d *Devices
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		fmt.Errorf("error decoding response: %w", err)
	}

	return d, err

}

//ListDevices returns all the user's devices in a simple ID: Name format
func (c *FocustronicClient) ListDevices(devices *Devices) (d map[int]string) {

	var deviceMap = map[int]string{}

	for _, data := range devices.Data {
		for _, device := range data.Devices {
			deviceMap[device.Id] = device.FriendlyName
		}
	}

	return deviceMap
}

// GetDevices calls the /devices endpoint and returns the user's own devices
func (c *FocustronicClient) GetDevices() (*Devices, error) {
	p, _ := url.Parse("/users/devices")

	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
	if err != nil {
		fmt.Println("error in request")
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