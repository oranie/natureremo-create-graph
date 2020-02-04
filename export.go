package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	ExportDDBtoJson()
}

func ExportDDBtoJson() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	svc := dynamodb.New(sess)
	var tableName = "NatureRemo"

	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	var result []interface{}
	err = svc.ScanPages(params,
		func(page *dynamodb.ScanOutput, lastPage bool) bool {
			for i := range page.Items {
				result = append(result, page.Items[i])
			}
			return true
		})
	fmt.Println(result)
	if err != nil {
		fmt.Println("scan error :", err)
	}
	//var item Item
	/*
		result[0].(map[string]*dynamodb.AttributeValue)
		if err := json.Unmarshal(result[0].(map[string]*dynamodb.AttributeValue), &item); err != nil {
			log.Fatal(err)
		}*/

}
