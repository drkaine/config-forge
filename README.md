# Config-forge

This project is a Go application designed to generate configuration files for diverse tools, like nginx or apache. The application allows users to specify various configuration parameters such as server name, document root path etc., and automatically generates corresponding configuration files for diverse tools, like nginx or apache.

## Prerequisites

Before you can run this application, make sure you have Go installed on your system. You can find installation instructions for Go on the [official Go website](https://golang.org/doc/install).

## Installation

To install the application, you can clone this Git repository to your local machine:

```bash
git clone https://github.com/your-user/server-config-generator.git
cd config-forge
go install
```

## Executable

You can use the build with :
```
./config-forge
```

And for rebuild :
```
go build
```

## Run

You can run the application with :
```
go run main.go
```

For run the tests :
```
go test
```

And the linter :
```
golangci-lint run
```

### Arguments

You can run the application with :
```
go run main.go -c toolName fileName serverName documentRoot
```
Or :
```
./config-forge --cli toolName fileName serverName documentRoot
```
By example :
```
./config-forge -c apache test test.test dev/test/public
```

## Contributions

Contributions to this project are welcome! If you'd like to improve the application, please submit a pull request with your changes.

## License

This project is licensed under the MIT License.
