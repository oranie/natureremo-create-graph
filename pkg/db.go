package pkg

import (
	"fmt"
	"strconv"
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
		Id:         deviceData.Id + "_Te",
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
					Item: map[string]*dynamodb.AttributeValue{
						"id": {
							S: aws.String(fmt.Sprintf(temperature.Id)),
						},
						"updated_at": {
							S: aws.String(fmt.Sprintf(temperature.Updated_at)),
						},
						"value": {
							N: aws.String(fmt.Sprintf(strconv.FormatFloat(temperature.Value, 'f', 4, 64))),
						},
					},
				},
			},
			{
				Put: &dynamodb.Put{
					TableName: aws.String(tableName),
					Item: map[string]*dynamodb.AttributeValue{
						"id": {
							S: aws.String(fmt.Sprintf(humidity.Id)),
						},
						"updated_at": {
							S: aws.String(fmt.Sprintf(humidity.Updated_at)),
						},
						"value": {
							N: aws.String(fmt.Sprintf(strconv.FormatFloat(humidity.Value, 'f', 4, 64))),
						},
					},
				},
			},
			{
				Put: &dynamodb.Put{
					TableName: aws.String(tableName),
					Item: map[string]*dynamodb.AttributeValue{
						"id": {
							S: aws.String(fmt.Sprintf(illumination.Id)),
						},
						"updated_at": {
							S: aws.String(fmt.Sprintf(illumination.Updated_at)),
						},
						"value": {
							N: aws.String(fmt.Sprintf(strconv.FormatFloat(illumination.Value, 'f', 4, 64))),
						},
					},
				},
			},
		},
	})

	fmt.Println("response :", res)

	if err != nil {
		fmt.Println("Got error calling TransactWriteItems:")
		fmt.Println(err.Error())
	}

	res, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
	}

	return res
}
