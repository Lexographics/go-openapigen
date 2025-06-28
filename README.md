# go-openapigen

A Go library for generating OpenAPI 3.0 specifications from Go types and route definitions. It is designed to make it easy to document your Go APIs and generate OpenAPI-compliant JSON files directly from your codebase.

## Features

- **Automatic OpenAPI 3.0 Spec Generation**: Generate OpenAPI specifications from Go structs and route definitions.
- **Custom Type and Enums**: Register custom types and enums to be used in schema generation.
- **Group and Tagging**: Group endpoints using regex patterns for better organization in the OpenAPI specification.
- **File Output**: Write the generated OpenAPI specification directly to a file.

## Installation

Add the package to your Go project:

```sh
go get github.com/Lexographics/go-openapigen
```

## Usage

Here's a minimal example of how to use `go-openapigen` to generate an OpenAPI spec for a simple user API:

```go
package main

import (
	"reflect"
  "time"

	"github.com/Lexographics/go-openapigen"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetUser struct {
	ID string `param:"id"`
}

type EditUser struct {
	ID        string    `param:"id"`
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}

func main() {
	gen := openapigen.New("Example API", "1.0.0")

	gen.AddGroup("Users", "/users.*")
	gen.AddRoute("/users/{id}", "GET", reflect.TypeOf(GetUser{}), openapigen.DefaultTag)
	gen.AddRoute("/users/{id}", "PUT", reflect.TypeOf(EditUser{}), openapigen.DefaultTag)
	gen.AddRoute("/users", "POST", reflect.TypeOf(User{}), openapigen.DefaultTag)

	err := gen.WriteToFile("openapi.json")
	if err != nil {
		// handle error
	}
}
```

### Output

Running the above example will generate an `openapi.json` file in the current directory, containing the OpenAPI 3.0 specification for your API.

## API Overview

- `openapigen.New(title, version)`: Create a new generator instance.
- `AddGroup(group, regex)`: Group endpoints by regex pattern.
- `AddRoute(path, method, inputType, tag)`: Register a route with its HTTP method, input struct type, and optional tag.
- `RegisterEnum(type, values)`: Register an enum type and its possible values.
- `RegisterCustomType(type, schema)`: Register a custom type with a specific OpenAPI schema.
- `WriteToFile(filename)`: Write the generated OpenAPI spec to a file.

## Contributing

Contributions are welcome! Please open issues or pull requests for bug fixes, features, or documentation improvements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
