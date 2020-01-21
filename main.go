package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

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
