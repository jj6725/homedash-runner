package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type sensorData struct {
	Humidity    float32
	Temperature float32
}

func convertToF(t float32) float32 {
	return float32((t*9)/5 + 32)
}

func fetchData(host string) (sensorData, error) {
	path := fmt.Sprintf("http://%s:6725/data", host)
	resp, err := http.Get(path)
	if err != nil {
		return sensorData{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return sensorData{}, err
	}

	var data sensorData
	json.Unmarshal(body, &data)
	return data, nil
}

func parseAndWrite(host string) string {
	time.Sleep(1 * time.Second)
	data, err := fetchData(host)
	if err != nil {
		return fmt.Sprintf("Error fetching from %s", host)
	}

	return fmt.Sprintf("[%s] H: %.2f T: %.2fºF", host, data.Humidity, convertToF(data.Temperature))
}

func main() {
	for {
		go fmt.Println(parseAndWrite("pi1"))
		go fmt.Println(parseAndWrite("pi4"))
	}
}
