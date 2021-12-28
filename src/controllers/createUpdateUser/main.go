package main

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/src/golang/src/types"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user types.User
	err := json.Unmarshal([]byte(req.Body), &user)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	formatedBodyResponse, _ := json.Marshal(user)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(formatedBodyResponse),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
