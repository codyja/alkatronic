package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// GetRecords calls the /records endpoint for the specified device and days
func (c *FocustronicClient) GetRecords(deviceType string, deviceID int, days string) (*Records, error) {
	p, _ := url.Parse(fmt.Sprintf("/api/v2/devices/%s/%d/data/test-records", deviceType, deviceID))

	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
	if err != nil {
		fmt.Println("must be an errors at req")
		return nil, err
	}

	// Add token to query string
	q := req.URL.Query()
	q.Add("day", days)
	req.URL.RawQuery = q.Encode()

	log.Printf("Pulling data from: %v", req.URL)
	//fmt.Println(req.URL)

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error making http call: %w", err)
	}

	var r *Records
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("GetRecords error decoding response: %w", err)
	}

	for k, v := range r.Data {
		r.Data[k].KhValue = ConvertValue(v.KhValue)
		r.Data[k].SolutionAdded = ConvertValue(v.SolutionAdded)
		if v.Parameter == "no3" ||  v.Parameter == "po4" || v.Parameter == "dkh" {
			r.Data[k].Value = ConvertValue(v.Value)
		}
	}

	return r, err
}

func (c *FocustronicClient) GetDosetronicRecords(deviceID int, days string) (*DosetronicRecords, error) {
	p, _ := url.Parse(fmt.Sprintf("/api/v2/devices/dosetronic/%d/data/dose-records", deviceID))

	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
	if err != nil {
		fmt.Println("must be an errors at req")
		return nil, err
	}

	// Add day and token to query string
	q := req.URL.Query()
	q.Add("day", days)
	q.Add("token", c.accessToken)
	req.URL.RawQuery = q.Encode()

	log.Printf("Pulling data from: %v", req.URL)

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error making http call: %w", err)
	}

	var r *DosetronicRecords
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("GetRecords error decoding response: %w", err)
	}

	//for k, v := range r.Data {
	//	r.Data[k].KhValue = ConvertValue(v.KhValue)
	//	r.Data[k].SolutionAdded = ConvertValue(v.SolutionAdded)
	//	if v.Parameter == "no3" ||  v.Parameter == "po4" || v.Parameter == "dkh" {
	//		r.Data[k].Value = ConvertValue(v.Value)
	//	}
	//}

	return r, err
}

//// GetLatestResult calls the GetRecords func, iterates over the dates and returns the most recent Record
//func (c *FocustronicClient) GetLatestResult(deviceType string, deviceID int) (Record, error) {
//	records, err := c.GetRecords(deviceType, deviceID, "7")
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

// ConvertValue takes in the reported KH value and converts to dKh
func ConvertValue(v float64) float64 {
	return v / 100.0
}
