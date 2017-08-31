# Uptime Gopher

Uptime Gopher is a simple commmand line tool for checking the availability of the given web resources.

It checks every given host (in the config.json file) by doing a web request, ping, port check and keyword check.

## Installation

You will need the Go compiler installed. Git clone this repository, `cd` to it and execute `go build`. If you encounter any missing packages, `go get` them and then try to build the program again.

## Configuration

Copy the `config.json.example` to `config.json`. Config file is very self-explanatory.

## Usage
Go to the build directory and execute the `./uptime-gopher` binary.
