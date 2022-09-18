test:
	go test ./... -cover

test-report:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out -o cover.html

lint:
	golangci-lint run