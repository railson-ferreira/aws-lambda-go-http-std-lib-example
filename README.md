# aws-lambda-go-http-std-lib-example

This is a simple implementation of a REST API in Go using standard library, ready to run on AWS Lambda.

It has been tested on **Custom Runtime - Amazon Linux 2 - arm64**

```shell
make
# or
GOARCH=arm64 GOOS=linux go build -tags lambda.norpc -o ./bin/bootstrap
```

