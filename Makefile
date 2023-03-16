test:
	go version
	go test ./tools/... -cover
	go test ./argparser/... -cover
	go test ./generator/... -cover
	go test ./man_style_help/... -cover
	go test ./examples/... -cover

test-report:
	go test ./tools/... -coverprofile cover.out
	go tool cover -html=cover.out -o cover-tools.html

	go test ./argparser/... -coverprofile cover.out
	go tool cover -html=cover.out -o cover-argparser.html

	go test ./generator/... -coverprofile cover.out
	go tool cover -html=cover.out -o cover-generator.html

	go test ./man_style_help/... -coverprofile cover.out
	go tool cover -html=cover.out -o cover-man-style-help.html

	go test ./examples/... -coverprofile cover.out
	go tool cover -html=cover.out -o cover-examples.html

lint:
	golangci-lint run tools/... --config=.golangci.yml
	golangci-lint run argparser/... --config=.golangci.yml
	golangci-lint run generator/... --config=.golangci.yml
	golangci-lint run man_style_help/... --config=.golangci.yml
	golangci-lint run examples/... --config=.golangci.yml

update:
	sh ./script.sh
