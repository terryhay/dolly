# dolly
The module for work with command line arguments.

The main advantage *dolly* is printing perfect formatted help information and checking input argument data.

## Setup

Import *dolly* module into your project:

`go get github.com/terryhay/dolly`

Build *dolly* generator into some directory where you are keeping bin files for your project:

`go build -o ./bin/gen_dolly github.com/terryhay/dolly/internal/generator`

The *dolly* generator will be built into *./bin* directory and will be named *gen_dolly*.

OK, the next step is to write *arg_config.yaml* file with description of your application commands, flags and other data for help info. For more information (you need it I guess) you can research arg_config.yaml files in example directories of dolly module. Good luck.

If your *arg_tools_config.yaml* is in *./config* directory you can generate

`./bin/gen_dolly -c ./config/arg_tools_config.yaml -o ./`

The *dolly* generator will create *./dolly* directory with *arg_tools.go* file. The file will contain *Parse* method.

Besides you can add strings

`go build -o ./bin/gen_dolly github.com/terryhay/dolly/internal/generator`

`./bin/gen_dolly -c ./config/arg_tools_config.yaml -o ./`

into your *Makefile* and use `make generate` command like it is done in examples.