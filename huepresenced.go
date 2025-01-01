package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var logging = false

func main() {
	bridgeIP := os.Getenv("HUEIP")
	username := os.Getenv("HUESERNAME")
	sensorID := ""

	flag.StringVar(&sensorID, "sensor", sensorID, "Hue sensor ID")
	flag.BoolVar(&logging, "log", logging, "Enable logging")
	flag.Parse()

	if sensorID == "" {
		log.Fatal("sensor must be defined")
	}

	for {
		presence, err := checkPresence(bridgeIP, username, sensorID)
		if err != nil && logging {
			log.Printf("Error checking presence: %v\n", err)
		}
		if presence {
			cmd := exec.Command("caffeinate", "-dimsut", "1")
			if err := cmd.Run(); err != nil && logging {
				log.Printf("Error running caffeinate: %v\n", err)
			}
		}
		time.Sleep(5 * time.Second) // Check every 5 seconds - fyi the bridge seems to reset flags afer about 9s
	}
}

type SensorResponse struct {
	State struct {
		Presence bool `json:"presence"`
	} `json:"state"`
}

func checkPresence(bridgeIP, username, sensorID string) (bool, error) {
	url := fmt.Sprintf("http://%s/api/%s/sensors/%s", bridgeIP, username, sensorID)
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var data SensorResponse
	if err := json.Unmarshal(body, &data); err != nil {
		if logging {
			log.Printf("error getting sensor data: %s", body)
		}
		return false, err
	}

	if logging {
		log.Printf("motion: %v\n", data.State.Presence)
	}

	return data.State.Presence, nil
}
