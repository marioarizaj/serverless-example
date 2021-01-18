package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/marioarizaj/serverless-example/api/common"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler() (events.APIGatewayProxyResponse, error) {
	db, err := common.ConnectToDB()
	if err != nil {
		return common.InternalServerError(), err
	}
	authors, err := db.GetAllAuthors()
	if err != nil {
		return common.InternalServerError(), err
	}
	authorsBts, err := json.Marshal(authors)
	if err != nil {
		return common.InternalServerError(), err
	}
	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(authorsBts),
		Headers: map[string]string{
			"Content-Type":           "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
