package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("us-east-1"))

func getItem(menuid string) (*menu, error) {
	// Prepare the input for the query.
	input := &dynamodb.GetItemInput{
		TableName: aws.String("Menus"),
		Key: map[string]*dynamodb.AttributeValue{
			"MENUID": {
				S: aws.String(menuid),
			},
		},
	}

	// Retrieve the item from DynamoDB. If no matching item is found
	// return nil.
	result, err := db.GetItem(input)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	// The result.Item object returned has the underlying type
	// map[string]*AttributeValue. We can use the UnmarshalMap helper
	// to parse this straight into the fields of a struct. Note:
	// UnmarshalListOfMaps also exists if you are working with multiple
	// items.
	mn := new(menu)
	err = dynamodbattribute.UnmarshalMap(result.Item, mn)
	if err != nil {
		return nil, err
	}

	return mn, nil
}

func putItem(mn *menu) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String("Menus"),
		Item: map[string]*dynamodb.AttributeValue{
			"MENUID": {
				S: aws.String(mn.MENUID),
			},
			"Restaurant": {
				S: aws.String(mn.Restaurant),
			},
			"Cuisine": {
				S: aws.String(mn.Cuisine),
			},
		},
	}

	_, err := db.PutItem(input)
	return err
}
