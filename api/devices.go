package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)


// ListDevices returns all the user's devices in a simple ID: Name format
func (c *FocustronicClient) ListDevices(devices *Devices) (d map[int]string) {

	var deviceMap = map[int]string{}

	for _, data := range devices.Data {
		for _, device := range data.Devices {
			deviceMap[device.Id] = device.FriendlyName
		}
	}

	return deviceMap
}

func (c *FocustronicClient) GetMastertronicDeviceDetails(id int) (*MastertronicDevice, error) {
	p, _ := url.Parse(fmt.Sprintf("/api/v2/devices/mastertronic/%d", id))

	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	q := req.URL.Query()
	q.Add("token", c.accessToken)
	req.URL.RawQuery = q.Encode()

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error making http call: %w", err)
	}
	defer resp.Body.Close()

	var d *MastertronicDevice
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return d, nil
}

func (c *FocustronicClient) GetDevices() (*Devices, error) {
	p, _ := url.Parse("/api/v2/users/self/devices")

	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	q := req.URL.Query()
	q.Add("token", c.accessToken)
	req.URL.RawQuery = q.Encode()

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error making http call: %w", err)
	}
	defer resp.Body.Close()

	var d *Devices
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return d, nil
}
