package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	InvokeNumber int `json:"InvokeNumber"`
}

func Handler(event Event) {
	fmt.Println("Hello Sfn!!")
	fmt.Println("InvokeNumber: ", event.InvokeNumber)
}

func main() {
	lambda.Start(Handler)
}
