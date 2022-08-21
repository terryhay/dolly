generate:
	go build -o ../bin/generators/gen_dolly ../../internal/generator
	../bin/generators/gen_dolly -c ../../examples/example/dolly_config/arg_config.yaml -o ./

test:
	go test ./... -cover

run_example_sequence_one:
	go run main.go -sl one two three -il 1 2 3 -fl 1.0 2 0.0003
