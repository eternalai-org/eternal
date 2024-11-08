package zip_hf_model_to_light_house

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test_downloadZipFileFromLightHouseConcurrentNew(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 24*time.Hour)
	defer cancel() //

	_ = ctx

	info := &HFModelInLightHouse{}
	url := "https://gateway.lighthouse.storage/ipfs/bafkreihwdqea37ii7s63klrmn2imh6nw2lm4tgyqwqzxalytiwhp6scruy"

	// Send HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	// Ensure response body is closed after reading
	defer response.Body.Close()

	// Check if the response status is OK
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", response.StatusCode)
		return
	}

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	// Unmarshal JSON into the Post struct
	err = json.Unmarshal(body, &info)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	downloadZipFileFromLightHouseConcurrentNew(info, "./test")
}
