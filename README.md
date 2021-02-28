<img src="https://alkatronic.focustronic.com/images/alkatronic_logo.png" width="300" alt="Alkatronic">

## Purpose
This project contains an GO api client to connect to the Focustronic Alkatronic api. 
You can authenticate with your own credentials and retrieve the test results for your devices. 
Then you can use the data for other use cases such as monitoring, graphing, and alerting.

## Examples

1. Set credentials in your shell:
```
export ALKATRONIC_USERNAME='user here'
export ALKATRONIC_PASSWORD='password here'
```
2. Create sample application. For permenant usage, consider writing the token locally (eg. ~/.alkatronic) and reusing
during future requests.
```go
package main

import (
	"fmt"
	"github.com/codyja/alkatronic/api"
	"log"
	"os"
	"time"
)

func main() {
	username, ok := os.LookupEnv("ALKATRONIC_USERNAME")
	if !ok {
		log.Fatalf("ALKATRONIC_USERNAME not set")
	}
	password, ok := os.LookupEnv("ALKATRONIC_PASSWORD")
	if !ok {
		log.Fatalf("ALKATRONIC_PASSWORD not set")
	}

	// Initialize new Alkatronic Client
	c, err := api.NewAlkatronicClient()
	if err != nil {
		fmt.Errorf("error initializing new Alkatronic Client")
	}

	c.Authenticate(username, password)

	// Get all alkatronic devices under account
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

