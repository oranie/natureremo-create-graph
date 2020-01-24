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

	temperature := Item{
		Id:         deviceData.Id + "_Te",
		Updated_at: deviceData.NewestEvents.Temperature.CreatedAt,
		Value:      deviceData.NewestEvents.Temperature.Value,
	}
	te, err := dynamodbattribute.MarshalMap(temperature)
	fmt.Println(te)

	humidity := Item{
		Id:         deviceData.Id + "_Hu",
		Updated_at: deviceData.NewestEvents.Humidity.CreatedAt,
		Value:      deviceData.NewestEvents.Humidity.Value,
	}

	hu, err := dynamodbattribute.MarshalMap(humidity)
	fmt.Println(hu)

	illumination := Item{
		Id:         deviceData.Id + "_Il",
		Updated_at: deviceData.NewestEvents.Illumination.CreatedAt,
		Value:      deviceData.NewestEvents.Illumination.Value,
	}
	il, err := dynamodbattribute.MarshalMap(illumination)
	fmt.Println(il)

	var tableName = "NatureRemo"

	_, err = svc.TransactWriteItems(&dynamodb.TransactWriteItemsInput{
		TransactItems: []*dynamodb.TransactWriteItem{
			{
				Put: &dynamodb.Put{
					TableName: aws.String(tableName),
					Item:      te,
				},
			},
			{
				Put: &dynamodb.Put{
					TableName: aws.String(tableName),
					Item:      hu,
				},
			},
			{
				Put: &dynamodb.Put{
					TableName: aws.String(tableName),
					Item:      il,
				},
			},
		},
	})

	if err != nil {
		fmt.Println("Got error calling TransactWriteItems:")
		fmt.Println(err.Error())
	}

	ressponse, err := svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
	}

	return ressponse
}
