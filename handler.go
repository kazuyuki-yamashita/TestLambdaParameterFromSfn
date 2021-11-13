package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Id string `json:"InvokeNumber"`
}

func Handler(event Event){
	fmt.Println("Hello Sfn!!")
	fmt.Println("InvokeNumber %d", event.Id)
}

func main() {
	lambda.Start(Handler)
}
