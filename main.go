package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
)

// HealthCheck will contain the data related to endpoints status
type HealthCheck struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Status string `json:"status"`
	Error  error  `json:"error"`
}

func callHealthcheckAPI(url string, health *[]HealthCheck) error {
	client := &http.Client{}
	resp, err := client.Get(url)

	if err != nil {
		fmt.Printf("Error: Failed to execute the HTTP request. %s\n", err)
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Errror: Unable to read the body. %s\n", err)
		return err
	}

	if resp.StatusCode != 200 {
		return err
	}

	err = json.Unmarshal(body, &health)
	if err != nil {
		// fmt.Printf("Errror: Unable to unmarshal the body. %s\n", err)
		return err
	}

	return nil

}

func createTable() error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "URL", "Status Code"})

	// convert from struct to 2D array
	data := [][]string{}
	for _, element := range h {
		data = append(data, []string{
			element.Name,
			element.URL,
			element.Status,
		})
	}

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output

	return nil
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

var h []HealthCheck

func main() {

	endpoint := getenv("HEALTHCHECK_ENDPOINT", "http://127.0.0.1:3333/healthcheck")
	callHealthcheckAPI(endpoint, &h)

	createTable()

}
