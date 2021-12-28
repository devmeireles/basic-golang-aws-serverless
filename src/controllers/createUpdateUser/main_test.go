package main

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/src/golang/src/helpers"
	"github.com/src/golang/src/types"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	mockedUser := &types.User{
		Name:     "Gabriel",
		Email:    "dev.meireles@gmail.com",
		UserName: "devmeireles",
	}

	t.Run("It should create a user", func(t *testing.T) {
		formatedRequestBody, _ := json.Marshal(mockedUser)

		body := events.APIGatewayProxyRequest{
			Body: string(formatedRequestBody),
		}

		response, _ := Handler(body)

		formatedBodyResponse := &types.User{}
		json.Unmarshal([]byte(response.Body), formatedBodyResponse)

		mockedKeys := helpers.HandleTestKey(mockedUser)
		responseKeys := helpers.HandleTestKey(formatedBodyResponse)

		assert.Equal(t, 201, response.StatusCode)

		assert.True(t, reflect.DeepEqual(mockedKeys, responseKeys))
	})

	t.Run("It shouldn't create a user because the body is empty", func(t *testing.T) {

		body := events.APIGatewayProxyRequest{
			Body: "",
		}

		response, _ := Handler(body)

		formatedBodyResponse := &types.User{}
		json.Unmarshal([]byte(response.Body), formatedBodyResponse)

		assert.Equal(t, 500, response.StatusCode)

	})
}
