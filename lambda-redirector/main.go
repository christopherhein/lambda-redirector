package main

import (
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	pathsplit := strings.Split(request.Path, "/")
	pathsplit = deleteEmpty(pathsplit)

	if len(pathsplit) == 0 {
		pathsplit = []string{os.Getenv("DEFAULTREPO")}
	}

	additionalPath := ""
	if len(pathsplit) > 1 {
		path := []string{"blob/master"}
		path = append(path, pathsplit[0:]...)

		additionalPath = strings.Join(path, "/")
	}

	return events.APIGatewayProxyResponse{
		Body:       "",
		StatusCode: 301,
		Headers: map[string]string{
			"Location": strings.Join([]string{"https:/", os.Getenv("CODEPATH"), pathsplit[0], additionalPath}, "/"),
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
