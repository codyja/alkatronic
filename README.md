## Purpose
This third party project contains a GO api client to connect to the Focustronic API. 
You can authenticate with your own credentials and retrieve the test results for your devices. You can retrieve data
for the Alkatronic, Mastertronic, and Dosetronic. Then you can use the data for other use cases such as monitoring, 
graphing, data analysis, and alerting. If using this project, please be respectful of the Focustronic API, thank you!

## Examples

1. Set credentials in your shell:
```
export FOCUSTRONIC_USERNAME='user here'
export FOCUSTRONIC_PASSWORD='password here'
```
or for Windows:
```
$env:FOCUSTRONIC_USERNAME = 'user here'
$env:FOCUSTRONIC_PASSWORD = 'password here'
```
2. Create sample application. For permanent usage, consider writing the token locally (eg. ~/.focustronic) and reusing
during future requests.
```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/codyja/focustronic/api"
)

func main() {
	username, ok := os.LookupEnv("FOCUSTRONIC_USERNAME")
	if !ok {
		log.Fatalf("FOCUSTRONIC_USERNAME not set")
	}
	password, ok := os.LookupEnv("FOCUSTRONIC_PASSWORD")
	if !ok {
		log.Fatalf("FOCUSTRONIC_PASSWORD not set")
	}

	// Initialize new Focustronic Client
	c, err := api.NewFocustronicClient()
	if err != nil {
		fmt.Errorf("error initializing new Focustronic Client")
	}

	c.Authenticate(username, password)

	// Get all Focustronic devices under account
	d, err := c.GetDevices()
	if err != nil {
		log.Fatalf("Error getting devices: %s", err)
	}

	// Loop over devices and then GetLatestResult
	for _, data := range d.Data {
		if data.Type == "alkatronic" {
			for _, device := range data.Devices {
				fmt.Printf("Device found: Name: %s, DeviceID: %d\n", device.FriendlyName, device.Id)
				record, err := c.GetAlkatronicLatestResult(device.Id)
				if err != nil {
					fmt.Errorf("Error getting alkatronic results: %s", err)
				}
				fmt.Printf("  Latest Alkalinity Record: %v, Time: %v\n", record.KhValue, time.Unix(record.CreateTime, 0))
			}
		}
		if data.Type == "dosetronic" {
			for _, device := range data.Devices {
				fmt.Printf("Device found: Name: %s, DeviceID: %d\n", device.FriendlyName, device.Id)

				records, err := c.GetDosetronicLatestRecords(device.Id)
				if err != nil {
					fmt.Errorf("Error getting dosetronic results: %s", err)
				}

				for pumpId, record := range records {
					fmt.Printf("  Latest Dosetronic Doseage: PumpID:%d, Volume: %v, Time: %v\n", pumpId, record.DoseVolume, time.Unix(record.RecordTime, 0))
				}

			}
		}
		if data.Type == "mastertronic" {
			for _, device := range data.Devices {
				fmt.Printf("Device found: Name: %s, DeviceID: %d\n", device.FriendlyName, device.Id)
				d, err := c.GetMastertronicDeviceDetails(device.Id)
				if err != nil {
					log.Println(fmt.Errorf("Error getting mastertronic results: %s", err))
				}

				for _, p := range d.Data.Parameters {
					record, err := c.GetLatestMastertronicRecord(device.Id, p.Parameter)
					if err != nil {
						fmt.Println(fmt.Errorf("  Error getting mastertronic results for parameter '%s': %s", p.Parameter, err))
						continue
					}
					fmt.Printf("  Latest Mastertronic Record: %s, Latest Value: %v, Time: %v\n", record.Parameter, record.Value, time.Unix(record.RecordTime, 0))
				}
			}
		}
	}
}
```   
3. Run app:
```
$ go run main.go
Device found: Name: Alkatronic, DeviceID: 001
  Latest Alkalinity Record: 6.53, Time: 2023-02-10 16:09:39 -0600 CST
Device found: Name: Dosetronic, DeviceID: 001
  Latest Dosetronic Doseage: PumpID:1, Volume: 3.9, Time: 2023-02-10 18:05:47 -0600 CST
  Latest Dosetronic Doseage: PumpID:2, Volume: 9.74, Time: 2023-02-10 18:11:00 -0600 CST
  Latest Dosetronic Doseage: PumpID:3, Volume: 1.49, Time: 2023-02-10 18:15:41 -0600 CST
  Latest Dosetronic Doseage: PumpID:4, Volume: 2.5, Time: 2023-02-10 15:20:41 -0600 CST
  Latest Dosetronic Doseage: PumpID:5, Volume: 2, Time: 2023-02-10 14:25:52 -0600 CST
Device found: Name: Mastertronic, DeviceID: 001
  Error getting mastertronic results for parameter 'no2': no records found in last 7 days
  Error getting mastertronic results for parameter 'no3': no records found in last 7 days
  Latest Mastertronic Record: po4, Latest Value: 0.24, Time: 2023-02-08 04:28:06 -0600 CST
  Error getting mastertronic results for parameter 'ca': no records found in last 7 days
  Latest Mastertronic Record: mg, Latest Value: 1496, Time: 2023-02-08 09:21:59 -0600 CST
  Latest Mastertronic Record: oli, Latest Value: 0, Time: 2023-02-09 03:03:56 -0600 CST
  Error getting mastertronic results for parameter 'dkh': no records found in last 7 days
  Error getting mastertronic results for parameter 'i': no records found in last 7 days
  Error getting mastertronic results for parameter 'fe': no records found in last 7 days
  Error getting mastertronic results for parameter 'nh4': no records found in last 7 days
```
