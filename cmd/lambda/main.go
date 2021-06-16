package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	handlers "github.com/bernljung/lambda-golang"
)

func main() {
	lambda.Start(handlers.HandlerFunc)
}
