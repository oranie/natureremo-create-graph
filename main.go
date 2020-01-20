package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

type Event struct {
	Value     float64 `json:"val"`
	CreatedAt string  `json:"created_at"`
}

type NewestEvents struct {
	Humidity     Event `json:"hu"`
	Illumination Event `json:"il"`
	Movement     Event `json:"mo"`
	Temperature  Event `json:"te"`
}

type User struct {
	Id        string `json:"id"`
	Nickname  string `json:"nickname"`
	Superuser bool   `json:"superuser"`
}

type Device struct {
	Name              string       `json:"name"`
	Id                string       `json:"id"`
	CreatedAt         string       `json:"created_at"`
	UpdatedAt         string       `json:"updated_at"`
	MacAddress        string       `json:"mac_address"`
	SerialNumber      string       `json:"serial_number"`
	FirmwareVersion   string       `json:"firmware_version"`
	TemperatureOffset int          `json:"temperature_offset"`
	HumidityOffset    int          `json:"humidity_offset"`
	Users             []User       `json:"users"`
	NewestEvents      NewestEvents `json:"newest_events"`
}

type Env struct {
	Token string `envconfig:"REMO_API_TOKEN" default:"test"`
}

func main() {
	env := GetEnvValue()

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.nature.global/1/devices", nil)
	req.Header.Add("Authorization", env.Token)

	resp, err := client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr := string(body)
	var devices []Device
	err = json.Unmarshal([]byte(jsonStr), &devices)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Get json response data : ", jsonStr)
	fmt.Println("Json parse : ", devices[0].Name)
	res := PutDeviceData(devices[0])
	fmt.Println(res)
}

func GetEnvValue() Env {
	var env Env
	err := envconfig.Process("", &env)
	fmt.Println(env)
	if err != nil {
		log.Fatal(err.Error())
	}
	return env
}

func createGraphData() {

}
