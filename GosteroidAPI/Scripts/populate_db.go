package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Distance represents the distance data
type Distance struct {
	Date     string  `json:"date"`
	Distance float64 `json:"distance"`
}

// Asteroid represents the structure of the asteroid data
type Asteroid struct {
	Name          string     `json:"name"`
	Diameter      float64    `json:"diameter"`
	DiscoveryDate string     `json:"discovery_date"`
	Observations  string     `json:"observations,omitempty"`
	Distances     []Distance `json:"distances,omitempty"`
}

func main() {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Replace with your actual JWT token
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTc3NTg2NzQsImlzcyI6Ikdvc3Rlcm9pZEFQSSJ9.HCpd_LwQVmVpbw5snWWfjt9YfBzeMivPzC77B2HicfE"

	// Skip the header
	for i, record := range records[1:] {
		if len(record) < 4 {
			log.Printf("Skipping invalid record at line %d: %v\n", i+2, record)
			continue
		}

		diameter, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Printf("Invalid diameter at line %d: %v\n", i+2, record)
			continue
		}

		asteroid := Asteroid{
			Name:          record[0],
			Diameter:      diameter,
			DiscoveryDate: record[2],
			Observations:  record[3],
			// Add other fields as necessary
		}

		// Serialize asteroid to JSON
		asteroidJSON, err := json.Marshal(asteroid)
		if err != nil {
			log.Printf("Error marshaling asteroid at line %d: %v\n", i+2, err)
			continue
		}

		// Make POST request to API
		req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/asteroides", bytes.NewBuffer(asteroidJSON))
		if err != nil {
			log.Printf("Error creating request at line %d: %v\n", i+2, err)
			continue
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+jwtToken)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error making request at line %d: %v\n", i+2, err)
			continue
		}

		body, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			log.Printf("Error response at line %d: %s\n", i+2, body)
		} else {
			fmt.Printf("Successfully added asteroid at line %d: %s\n", i+2, asteroid.Name)
		}
	}
}
