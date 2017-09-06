# Uptime Gopher

Uptime Gopher is a simple commmand line tool for checking the availability of the given web resources.

It checks every given host by doing a web request, ping, port check and keyword check (checks used per host are user-defined).

It displays the results on the standard output but it can be configured to log the results to a MySQL database (with or without displaying them on the standard output).

## Installation

You will need the Go compiler installed. Git clone this repository, `cd` to it and execute `go build`. If you encounter any missing packages, `go get` them and then try to build the program again.

## Configuration

To configure the program, copy `config.json.example` to `config.json` and fill the missing values.

If you do not want to log the results to the database, set `log-to-db` to `false` (`log-to-stdout` should be then `true`). If you want to log to the database, create an MySQL database and a MySQL user with ALL privileges on the new database and fill the missing MySQL values.

To define the hosts you want to check, copy `hosts.json.example` to `hosts.json` and add your hosts. 

The config file and the hosts config file are very self-explanatory.

## Usage
Go to the build directory and execute the `./uptime-gopher` binary.
