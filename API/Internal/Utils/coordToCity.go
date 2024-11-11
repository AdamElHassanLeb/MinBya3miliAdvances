package Utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Struct to parse the response
type MapsCoResponse struct {
	DisplayName string `json:"display_name"`
	Address     struct {
		City    string `json:"city"`
		Village string `json:"village"`
		Town    string `json:"town"`
		Suburb  string `json:"suburb"`
		State   string `json:"state"`
		Country string `json:"country"`
	} `json:"address"`
}

func ReverseGeocode(lat, lon float64) (string, string, error) {
	//fmt.Println(lat, lon)
	// Create a ticker that ticks once per second to limit requests
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	// Wait for the next tick (ensures 1 request per second)
	<-ticker.C

	// Your API key
	apiKey := "67306e2d9fb92945904915kom9878b8" // Replace with your actual API key

	// Construct the request URL with the API key
	reqURL := fmt.Sprintf("https://geocode.maps.co/reverse?lat=%f&lon=%f&api_key=%s", lat, lon, apiKey)

	//fmt.Println(reqURL)
	// Send the HTTP GET request
	resp, err := http.Get(reqURL)
	if err != nil {
		return "", "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("error reading response: %w", err)
	}

	// Parse the JSON response
	var response MapsCoResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", "", fmt.Errorf("error parsing JSON: %w", err)
	}

	//fmt.Println(string(body))

	// Determine the city field
	city := response.Address.City
	if city == "" {
		if response.Address.Village != "" {
			city = response.Address.Village
		} else if response.Address.Town != "" {
			city = response.Address.Town
		} else if response.Address.Suburb != "" {
			city = response.Address.Suburb
		} else if response.Address.State != "" {
			city = response.Address.State
		}
	}

	// Return the city and country
	return city, response.Address.Country, nil
}
