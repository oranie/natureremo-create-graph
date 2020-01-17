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
	Val       float64 `json:"val"`
	CreatedAt string  `json:"created_at"`
}

type NewestEvents struct {
	Hu Event `json:"hu"`
	Il Event `json:"il"`
	Mo Event `json:"mo"`
	Te Event `json:"te"`
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
	Token string `envconfig:"TOKEN" default:"test"`
}

func main() {
	env := GetEnvValue()
	fmt.Println(env.Token)

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.nature.global/1/devices", nil)
	req.Header.Add("Authorization", env.Token)
	fmt.Println(req)

	resp, err := client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr := string(body)
	fmt.Println("Body string : ", jsonStr, err)

	var devices []Device
	err = json.Unmarshal([]byte(jsonStr), &devices)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("test json data", jsonStr)
	fmt.Println("Json parse : ", devices[0].NewestEvents.Hu.Val)

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
