package goipgeolocator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type IpGeoResp struct {
	Country          string `json:"country"`
	CountryCode      string `json:"country_code"`
	CountryCode3     string `json:"country_code3"`
	ContinentCode    string `json:"continent_code"`
	Region           string `json:"region"`
	IP               string `json:"ip"`
	OrganizationName string `json:"organization_name"`
	Asn              int    `json:"asn"`
	Organization     string `json:"organization"`
	Timezone         string `json:"timezone"`
	Accuracy         int    `json:"accuracy"`
	Latitude         string `json:"latitude"`
	City             string `json:"city"`
	Longitude        string `json:"longitude"`
	AreaCode         string `json:"area_code"`
}

// FetchIPGeolocation takes an array of IP addresses and returns their geolocation information
//
// Uses https://www.geojs.io to fet ip information
func FetchIPGeolocation(ips []string) ([]IpGeoResp, error) {
	// Client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Construct URL with all IPs chained by commas
	queryParam := strings.Join(ips, ",")
	url := fmt.Sprintf("https://get.geojs.io/v1/ip/geo.json?ip=%s", url.QueryEscape(queryParam))

	// Execute HTTP GET request
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Check for HTTP error codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status: %d", resp.StatusCode)
	}

	// Read and decode JSON response
	// The API needs to support returning an array of GeoResponse; adjust accordingly
	var geoResponses []IpGeoResp
	if err := json.NewDecoder(resp.Body).Decode(&geoResponses); err != nil {
		return nil, fmt.Errorf("JSON decode failed: %v", err)
	}

	return geoResponses, nil
}
