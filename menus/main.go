package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var menuidRegex = regexp.MustCompile(`[0-9,-]`)
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

type menu struct {
	MENUID     string `json:"menuid"`
	Restaurant string `json:"restaurant"`
	Cuisine    string `json:"cuisine"`
}

func show(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	menuid := req.QueryStringParameters["menuid"]
	if !menuidRegex.MatchString(menuid) {
		return clientError(http.StatusBadRequest)
	}

	menu, err := getItem(menuid)
	if err != nil {
		return serverError(err)
	}
	if menu == nil {
		return clientError(http.StatusNotFound)
	}
	// The APIGatewayProxyResponse.Body field needs to be a string, so
	// we marshal the book record into JSON.
	js, err := json.Marshal(menu)
	if err != nil {
		return serverError(err)
	}

	// Return a response with a 200 OK status and the JSON book record
	// as the body.
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(js),
	}, nil
}

// Add a helper for handling errors. This logs any error to os.Stderr
// and returns a 500 Internal Server Error response that the AWS API
// Gateway understands.
func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

// Similarly add a helper for send responses relating to client errors.
func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func create(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if req.Headers["content-type"] != "application/json" && req.Headers["Content-Type"] != "application/json" {
		return clientError(http.StatusNotAcceptable)
	}

	mn := new(menu)
	err := json.Unmarshal([]byte(req.Body), mn)
	if err != nil {
		return clientError(http.StatusUnprocessableEntity)
	}

	if !menuidRegex.MatchString(mn.MENUID) {
		return clientError(http.StatusBadRequest)
	}
	if mn.Restaurant == "" || mn.Cuisine == "" {
		return clientError(http.StatusBadRequest)
	}

	err = putItem(mn)
	if err != nil {
		return serverError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Headers:    map[string]string{"Location": fmt.Sprintf("/books?isbn=%s", mn.MENUID)},
	}, nil
}

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return show(req)
	case "POST":
		return create(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func main() {
	lambda.Start(router)
}
