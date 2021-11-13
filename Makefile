all: build deploy

build:
	GOARCH=amd64 GOOS=linux go build -o handler handler.go
	zip function.zip handler

deploy:
	aws lambda update-function-code --function-name invoke-lambda-from-sfn --zip-file fileb://function.zip