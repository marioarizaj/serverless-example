package common

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func InternalServerError() events.APIGatewayProxyResponse {
	resp := map[string]interface{}{
		"detail": "Internal Server Error",
	}
	val, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(val),
	}
}
