package main

import (
	"ccwe/frekwencja/librus"
	"context"
	"encoding/base64"
	"encoding/json"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
)

func HandleLambdaEvent(ctx context.Context, event events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	form, _ := base64.StdEncoding.DecodeString(event.Body)
	parsedBody, _ := url.ParseQuery(string(form))

	data, err := librus.Init(
		parsedBody.Get("username"),
		parsedBody.Get("password"),
	)
	if err != nil {
		if e, ok := err.(*librus.AppError); ok {
			return makeResponse(e.Status, e.Error()), nil
		}
	}

	err = data.LoadEndpoints("Subjects", "Lessons", "Attendances")
	if err != nil {
		if e, ok := err.(*librus.AppError); ok {
			return makeResponse(e.Status, e.Error()), nil
		}
	}

	target := data.MakeLessonSubjectMap()
	data.GetAttendance(target)

	var nodes []*librus.CalculatedNode
	for _, node := range target {
		nodes = append(nodes, node)
	}

	result, err := json.Marshal(&nodes)
	if err != nil {
		return makeResponse(500, err.Error()), nil
	}

	return makeResponse(200, string(result)), nil
}

func makeResponse(statusCode int, result string) *events.APIGatewayV2HTTPResponse {
	return &events.APIGatewayV2HTTPResponse{
		Body:            result,
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Headers:         map[string]string{"content-type": "application/json"},
	}
}
