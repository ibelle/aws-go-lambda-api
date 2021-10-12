package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type menu struct {
	MENUID     string `json:"menuid"`
	Restaurant string `json:"restaurant"`
	Cuisine    string `json:"cusisine"`
}

func show() (*menu, error) {
	mn, err := getItem("978-1420931693")
	if err != nil {
		return nil, err
	}

	return mn, nil
}

func main() {
	lambda.Start(show)
}
