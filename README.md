# My Umbrella task implementation

This project aims to solve the problem outlined in the [test task](https://github.com/spatecon/echo-middleware-assessment/blob/master/docs/testtask.pdf). The project is built using [Go](https://golang.org) and [Echo](https://echo.labstack.com).

## Getting Started

To get started with the project, clone this repository and run `go run cmd/days_left/main.go`.

## Usage

To make a new request, use the following command:

```
http http://127.0.0.1:1223/status

594 days left until 1 Jan 2025
```

To make a new request with admin's User-Role, use the following command:

```
http http://127.0.0.1:1223/status

594 days left until 1 Jan 2025

# IN CONSOLE
2023/05/18 18:57:57 red button user detected
```

## License

This project's source code can be shared and used for educational purposes.
