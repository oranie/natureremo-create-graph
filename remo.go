package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

type Item struct {
	Id         string  `json:"id"`
	Updated_at string  `json:"updated_at"`
	Value      float64 `json:"value"`
}

func GenarateSensorData(deviceData Device) map[string]map[string]*dynamodb.AttributeValue {
	allSensorData := map[string]map[string]*dynamodb.AttributeValue{}

	now := time.Now()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := now.In(jst).Format("2006-01-02T15:04:05Z")

	temperature := Item{
		Id:         deviceData.Id + "_Te",
		Updated_at: nowJST,
		Value:      deviceData.NewestEvents.Temperature.Value,
	}
	te, err := dynamodbattribute.MarshalMap(temperature)
	if err != nil {
		fmt.Println("Got error set sensorData:")
		fmt.Println(err.Error())
	}
	allSensorData["temperature"] = te
	fmt.Println(allSensorData["temperature"])

	humidity := Item{
		Id:         deviceData.Id + "_Hu",
		Updated_at: nowJST,
		Value:      deviceData.NewestEvents.Humidity.Value,
	}

	hu, err := dynamodbattribute.MarshalMap(humidity)
	if err != nil {
		fmt.Println("Got error set sensorData:")
		fmt.Println(err.Error())
	}
	allSensorData["humidity"] = hu
	fmt.Println(allSensorData["humidity"])

	illumination := Item{
		Id:         deviceData.Id + "_Il",
		Updated_at: nowJST,
		Value:      deviceData.NewestEvents.Illumination.Value,
	}
	il, err := dynamodbattribute.MarshalMap(illumination)
	allSensorData["illumination"] = il
	fmt.Println(allSensorData["illumination"])

	if err != nil {
		fmt.Println("Got error set sensorData:")
		fmt.Println(err.Error())
	}

	movement := Item{
		Id:         deviceData.Id + "_Mo",
		Updated_at: nowJST,
		Value:      deviceData.NewestEvents.Movement.Value,
	}
	mo, err := dynamodbattribute.MarshalMap(movement)
	allSensorData["movement"] = mo
	fmt.Println(allSensorData["movement"])

	fmt.Println(allSensorData)

	return allSensorData
}
