package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	InvokeNumber int `json:"InvokeNumber"`
}

func Handler(ctx context.Context, event Event) {
	fmt.Println("Hello Sfn!!")
	fmt.Println("InvokeNumber: ", event.InvokeNumber)
}

func main() {
	lambda.Start(Handler)
}
