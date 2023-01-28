<img src="http://www.focustronic.net/skin/frontend/wow/default/images/logo.gif" width="150" alt="Alkatronic">

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
	"github.com/codyja/focustronic/api"
	"log"
	"os"
	"time"
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
	for _, v := range d.Data {
		// Get the last test result
		l, err := c.GetLatestResult(v.DeviceID)
		if err != nil {
			log.Fatalf("Error getting last test result: %s", err)
		}
		
		fmt.Printf(
			"Latest test result: Device Name: %s, Record ID: %d, KH_Value: %.2f, Create Time: %s\n",
			v.FriendlyName,
			l.RecordID,
			api.ConvertKh(l.KhValue),
			time.Unix(l.CreateTime, 0).Format(time.RFC822Z))
	}
}
```   
3. Run app:
```
$ go run main.go
2021/02/28 15:32:34 Starting authentication against Alkatronic's API
2021/02/28 15:32:34 Authentication has succeeded
Latest test result: Device Name: Alkatronic, Record ID: 0000001, KH_Value: 7.72, Create Time: 28 Feb 21 15:08 -0600
```
