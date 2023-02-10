package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
)

// GetRecords calls the /records endpoint for the specified device and days
func (c *FocustronicClient) GetAlkatronicRecords(deviceID int, days int) (*AlkatronicRecords, error) {
	p, _ := url.Parse(fmt.Sprintf("/api/v2/devices/alkatronic/%d/data/test-records", deviceID))

	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
	if err != nil {
		fmt.Println("must be an errors at req")
		return nil, err
	}

	// Add token to query string
	q := req.URL.Query()
	q.Add("day", strconv.Itoa(days))
	req.URL.RawQuery = q.Encode()

	//log.Printf("Pulling data from: %v", req.URL)
	//fmt.Println(req.URL)

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error making http call: %w", err)
	}

	var r *AlkatronicRecords
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("GetRecords error decoding response: %w", err)
	}

	for k, v := range r.Data {
		r.Data[k].KhValue = ConvertValue(v.KhValue)
		r.Data[k].SolutionAdded = ConvertValue(v.SolutionAdded)
		// if v.Parameter == "no3" || v.Parameter == "po4" || v.Parameter == "dkh" {
		// 	r.Data[k].Value = ConvertValue(v.Value)
		// }
	}

	return r, err
}

func (c *FocustronicClient) GetAlkatronicLatestResult(deviceID int) (*AlkatronicRecord, error) {
	records, err := c.GetAlkatronicRecords(deviceID, 7)
	if err != nil {
		return nil, err
	}

	if len(records.Data) == 0 {
		return nil, fmt.Errorf("no records found")
	}

	sort.Slice(records.Data, func(i, j int) bool {
		return records.Data[i].CreateTime > records.Data[j].CreateTime
	})

	// return &records.Data[0], nil
	latest := &AlkatronicRecord{
		// RecordID:   records.Data[0].RecordID,
		// DeviceID:   records.Data[0].DeviceID,
		KhValue:    records.Data[0].KhValue,
		RecordTime: records.Data[0].RecordTime,
		CreateTime: records.Data[0].CreateTime,
	}

	return latest, nil

}

// ConvertValue takes in the reported KH value and converts to dKh
func ConvertValue(v float64) float64 {
	return v / 100.0
}

// Will provide all test results unless a specific parameter is provided
func (c *FocustronicClient) GetMastertronicRecords(days, deviceID int, parameter string) (*MastertronicRecords, error) {
	p, _ := url.Parse(fmt.Sprintf("/api/v2/devices/mastertronic/%d/data/test-records", deviceID))

	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
	if err != nil {
		fmt.Println("must be an errors at req")
		return nil, err
	}

	// Add token to query string
	q := req.URL.Query()
	q.Add("day", strconv.Itoa(days))
	q.Add("token", c.accessToken)
	if parameter != "" {
		q.Add("parameter", parameter)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error making http call: %w", err)
	}

	var r *MastertronicRecords
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	for k, v := range r.Data {
		if v.Parameter == "no3" || v.Parameter == "po4" || v.Parameter == "dkh" {
			r.Data[k].Value = ConvertValue(v.Value)
		}
	}

	return r, err
}

func (c *FocustronicClient) GetLatestMastertronicRecord(deviceID int, parameter string) (*MastertronicRecord, error) {
	records, err := c.GetMastertronicRecords(7, deviceID, parameter)
	if err != nil {
		return nil, fmt.Errorf("error getting records: %w", err)
	}

	if len(records.Data) == 0 {
		return nil, fmt.Errorf("no records found in last 7 days")
	}

	sort.Slice(records.Data, func(i, j int) bool {
		return records.Data[i].RecordTime > records.Data[j].RecordTime
	})

	return &records.Data[0], nil

}

func (c *FocustronicClient) GetLatestMastertronicRecordId(records *MastertronicRecords) (int, error) {
	if len(records.Data) == 0 {
		return 0, fmt.Errorf("no records found")
	}

	sort.Slice(records.Data, func(i, j int) bool {
		return records.Data[i].RecordTime > records.Data[j].RecordTime
	})

	return records.Data[0].ID, nil

}

func (c *FocustronicClient) GetDosetronicRecords(deviceID, days int) (*DosetronicRecords, error) {
	p, _ := url.Parse(fmt.Sprintf("/api/v2/devices/dosetronic/%d/data/dose-records", deviceID))

	req, err := http.NewRequest(http.MethodGet, c.baseURL.ResolveReference(p).String(), nil)
	if err != nil {
		fmt.Println("must be an errors at req")
		return nil, err
	}

	// Add day and token to query string
	q := req.URL.Query()
	q.Add("day", strconv.Itoa(days))
	q.Add("token", c.accessToken)
	req.URL.RawQuery = q.Encode()

	//log.Printf("Pulling data from: %v", req.URL)

	resp, err := c.DoRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error making http call: %w", err)
	}

	var r *DosetronicRecords
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("GetRecords error decoding response: %w", err)
	}

	for k, v := range r.Data {
		for k1, v1 := range v {
			r.Data[k][k1].DoseVolume = ConvertValue(v1.DoseVolume)
		}
	}

	return r, err
}

func (c *FocustronicClient) GetDosetronicLatestRecords(deviceID int) (map[int]DosetronicRecord, error) {
	records, err := c.GetDosetronicRecords(deviceID, 7)
	if err != nil {
		return nil, fmt.Errorf("error getting records: %w", err)
	}

	latestRecords := make(map[int]DosetronicRecord)

	for _, record := range records.Data {
		for _, pumpRecord := range record {
			// Check if the pump_id already exists in the map
			if currentRecord, ok := latestRecords[pumpRecord.PumpID]; ok {
				// If it does, check if the record_time for the current record is greater than the record_time for the existing record
				if pumpRecord.RecordTime > currentRecord.RecordTime {
					// If it is, update the map with the new record
					latestRecords[pumpRecord.PumpID] = pumpRecord
				}
			} else {
				// If the pump_id does not exist in the map, add it
				latestRecords[pumpRecord.PumpID] = pumpRecord
			}
		}

	}

	return latestRecords, nil
}
