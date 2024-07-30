# config Package

## Overview

The `config` package provides a set of functions for working with JSON configuration files. It offers methods to retrieve different types of data (string, int, float, bool, etc.) from a JSON key. This is useful for applications that need to read and manage configurations stored in JSON format.

## Installation

To install the `config` package, use the following command:

```bash
go get github.com/yourusername/config
```

## Usage

Here’s a basic example of how to use the `config` package:

```go
package main

import (
    "fmt"
    "log"
    "github.com/yourusername/config"
)

func main() {
    err := config.LoadConfigFile("config.json")
    if err != nil {
        log.Fatalf("Failed to open config file: %v", err)
    }

    appName := config.GetString("app.name")
    port := config.GetInt("server.port")
    debugMode := config.GetBool("debug")

    fmt.Printf("App Name: %s\n", appName)
    fmt.Printf("Port: %d\n", port)
    fmt.Printf("Debug Mode: %t\n", debugMode)
}
```

## Functions

The `config` package provides several functions to retrieve data from a JSON configuration file.

### `GetString`

```go
func GetString(key string) string
```

Gets a string JSON node from a JSON key.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The string value associated with the JSON key.

### `GetInt64`

```go
func GetInt64(key string) int64
```

Gets an integer 64 JSON node from a JSON key.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The int64 value associated with the JSON key.

### `GetInt`

```go
func GetInt(key string) int
```

Gets an integer JSON node from a JSON key.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The int value associated with the JSON key.

### `GetFloat64`

```go
func GetFloat64(key string) float64
```

Gets a float 64 JSON node from a JSON key.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The float64 value associated with the JSON key.

### `GetBool`

```go
func GetBool(key string) bool
```

Gets a boolean JSON node from a JSON key.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The boolean value associated with the JSON key.

### `Get`

```go
func Get(key string) interface{}
```

Gets a JSON node from a JSON key.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The value associated with the JSON key as an `interface{}`.

### `GetArrayMaps`

```go
func GetArrayMaps(key string) []map[string]interface{}
```

Gets an array JSON node from a JSON key. Returns a JSON array as a map.

- **key**: The JSON key (e.g., `foo.bar`).
- **Returns**: The JSON array as a slice of maps.

### `LoadConfigFile`

```go
func LoadConfigFile(name string) error
```

Opens a file that contains a JSON object.

- **name**: The name of the JSON file.
- **Returns**: An error if the file cannot be opened.

## Example

Here’s an example configuration file (`config.json`):

```json
{
    "app": {
        "name": "MyApp"
    },
    "server": {
        "port": 8080
    },
    "debug": true
}
```

Using the configuration file:

```go
package main

import (
    "fmt"
    "log"
    "github.com/jgolang/config"
)

func main() {
    err := config.LoadConfigFile("config.json")
    if err != nil {
        log.Fatalf("Failed to open config file: %v", err)
    }

    appName := config.GetString("app.name")
    port := config.GetInt("server.port")
    debugMode := config.GetBool("debug")

    fmt.Printf("App Name: %s\n", appName)
    fmt.Printf("Port: %d\n", port)
    fmt.Printf("Debug Mode: %t\n", debugMode)
}
```

This example demonstrates how to open a configuration file and retrieve different types of data from it using the `config` package.

## Contributing

If you have suggestions for how We could be improved, or want to report a bug, open an issue! We'd love all and any contributions.

For more, check out the [Contributing Guide](CONTRIBUTING.md).

## License

This project is licensed under the [MIT License](LICENSE).

## Support

If you find this repository helpful and would like to support its development, consider making a donation. Your contributions will help ensure the continued improvement and maintenance of this repository.

Thank you for your support!

[![ko-fi](https://www.ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/josuegiron)