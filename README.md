# Nico Robin

Robin is a URL fuzzer project written in Go, designed to help test and identify vulnerabilities in web applications. The project generates a large number of requests with varying inputs to a target web application to detect common vulnerabilities such as SQL injection, cross-site scripting (XSS), and other common web application vulnerabilities.

## Features

- Generates a large number of requests to test web applications
- Identifies input validation errors, SQL injection, cross-site scripting (XSS), and other common web application vulnerabilities
- Configurable request parameters such as query strings and form fields
- Advanced capabilities for efficient generation and management of requests

## Run it yourself

To use Robin, simply clone the project and run it using the `go run` command:

```
$ git clone https://github.com/raissonsouto/Robin.git && cd Robin
$ go run main.go
```

By default, Robin generates requests to a target URL specified in the main.go file. You can modify the URL and other parameters in the code to suit your needs.

## Requirements

Robin requires Go version 1.16 or later.