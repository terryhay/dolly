test:
	go test ./utils/... -cover
	go test ./argparser/... -cover
	go test ./generator/... -cover
	go test ./examples/... -cover

test-report:
	go test ./utils/... -coverprofile cover.out
	go tool cover -html=cover.out -o cover-utils.html

	go test ./argparser/... -coverprofile cover.out
	go tool cover -html=cover.out -o cover-argparser.html

	go test ./generator/... -coverprofile cover.out
	go tool cover -html=cover.out -o cover-generator.html

	go test ./examples/... -coverprofile cover.out
	go tool cover -html=cover.out -o cover-examples.html

lint:
	golangci-lint run utils/...
	golangci-lint run argparser/...
	golangci-lint run generator/...
	golangci-lint run examples/...

update:
	sh ./script.sh
