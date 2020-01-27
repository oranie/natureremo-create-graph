package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

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
			{
				Put: &dynamodb.Put{
					TableName: aws.String(tableName),
					Item:      AllSensorData["movement"],
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
