package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	Id         string  `json:"id"`
	Updated_at string  `json:"updatedt_at"`
	Value      float64 `json:"value"`
}

type AllSensortData struct {
	SensorData map[string]map[string]*dynamodb.AttributeValue
}

//Device データをDynamoDBに書き込む
func PutDeviceData(deviceData Device) (res *dynamodb.PutItemOutput) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)

	svc := dynamodb.New(sess)
	av, err := dynamodbattribute.MarshalMap(deviceData)
	input := &dynamodb.PutItemInput{
		Item:                   av,
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String("NatureRemo"),
	}
	fmt.Println("dynamodb mapping : ", av)

	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
	}

	fmt.Println("start time :", time.Now())
	AllSensorData := GenarateSensorData(deviceData)

	var tableName = "NatureRemo"

	_, err = svc.TransactWriteItems(&dynamodb.TransactWriteItemsInput{
		TransactItems: []*dynamodb.TransactWriteItem{
			{
				Put: &dynamodb.Put{
					TableName: aws.String(tableName),
					Item:      AllSensorData["temperature"],
				},
			},
			{
				Put: &dynamodb.Put{
					TableName: aws.String(tableName),
					Item:      AllSensorData["humidity"],
				},
			},
			{
				Put: &dynamodb.Put{
					TableName: aws.String(tableName),
					Item:      AllSensorData["illumination"],
				},
			},
		},
	})

	if err != nil {
		fmt.Println("Got error calling TransactWriteItems:")
		fmt.Println(err.Error())
	}

	response, err := svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
	}

	return response
}
func GenarateSensorData(deviceData Device) map[string]map[string]*dynamodb.AttributeValue {
	allSensorData := map[string]map[string]*dynamodb.AttributeValue{}

	temperature := Item{
		Id:         deviceData.Id + "_Te",
		Updated_at: deviceData.NewestEvents.Temperature.CreatedAt,
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
		Updated_at: deviceData.NewestEvents.Humidity.CreatedAt,
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
		Updated_at: deviceData.NewestEvents.Illumination.CreatedAt,
		Value:      deviceData.NewestEvents.Illumination.Value,
	}
	il, err := dynamodbattribute.MarshalMap(illumination)
	allSensorData["illumination"] = il
	fmt.Println(allSensorData["illumination"])

	if err != nil {
		fmt.Println("Got error set sensorData:")
		fmt.Println(err.Error())
	}

	fmt.Println(allSensorData)

	return allSensorData
}
