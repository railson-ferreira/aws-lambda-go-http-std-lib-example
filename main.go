package main

import (
	"encoding/json"
	"github.com/a-h/awsapigatewayv2handler"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
	"time"
)

var httpLambdaHandler lambda.Handler

type Response struct {
	From    string
	Message string
}

func init() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(Response{From: "net/http", Message: time.Now().Format(time.UnixDate)})
	})

	httpLambdaHandler = awsapigatewayv2handler.NewLambdaHandler(http.DefaultServeMux)
}

func main() {
	lambda.Start(httpLambdaHandler)
}
