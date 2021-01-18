package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/marioarizaj/serverless-example/api/common"
	"github.com/marioarizaj/serverless-example/models"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var newArticle models.Article
	err := json.Unmarshal([]byte(request.Body), &newArticle)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	db, err := common.ConnectToDB()
	if err != nil {
		return common.InternalServerError(), err
	}
	author, err := db.CreateArticle(&newArticle)
	if err != nil {
		return common.InternalServerError(), err
	}
	authorBts, err := json.Marshal(author)
	if err != nil {
		return common.InternalServerError(), err
	}
	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(authorBts),
		Headers: map[string]string{
			"Content-Type":           "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
