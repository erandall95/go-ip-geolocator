package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/erandall95/goipgeolocator"
)

func main() {
	// search both ipv4 and ipv6 addresses
	ips := []string{"8.8.8.8", "2001:4860:4860::8888"}
	geoData, err := goipgeolocator.FetchIPGeolocation(ips)
	if err != nil {
		log.Fatalf("Error fetching geolocation: %v", err)
	}

	// Convert the geolocation data to JSON
	jsonData, err := json.MarshalIndent(geoData, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// Print the JSON data
	fmt.Println(string(jsonData))
}

/* Expected Reponse

[
  {
    "country": "United States",
    "country_code": "US",
    "country_code3": "USA",
    "continent_code": "NA",
    "region": "",
    "ip": "8.8.8.8",
    "organization_name": "GOOGLE",
    "asn": 15169,
    "organization": "AS15169 GOOGLE",
    "timezone": "America/Chicago",
    "accuracy": 1000,
    "latitude": "37.751",
    "city": "",
    "longitude": "-97.822",
    "area_code": "0"
  },
  {
    "country": "",
    "country_code": "",
    "country_code3": "",
    "continent_code": "",
    "region": "",
    "ip": "2001:4860:4860:: 8888",
    "organization_name": "Unknown",
    "asn": 64512,
    "organization": "AS64512 Unknown",
    "timezone": "",
    "accuracy": 0,
    "latitude": "nil",
    "city": "",
    "longitude": "nil",
    "area_code": "0"
  }
]

*/
