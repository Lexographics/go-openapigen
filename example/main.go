package main

import (
	"reflect"

	"github.com/Lexographics/go-openapigen"
	"github.com/Lexographics/go-openapigen/example/models"
	"github.com/Lexographics/go-openapigen/example/requests"
)

func main() {

	gen := openapigen.New("Example API", "1.0.0")

	gen.AddGroup("Users", "/users.*")
	gen.AddRoute("/users/{id}", "GET", reflect.TypeOf(requests.GetUser{}), openapigen.DefaultTag)
	gen.AddRoute("/users/{id}", "PUT", reflect.TypeOf(requests.EditUser{}), openapigen.DefaultTag)
	gen.AddRoute("/users", "POST", reflect.TypeOf(models.User{}), openapigen.DefaultTag)

	err := gen.WriteToFile("openapi.json")
	if err != nil {
		panic(err)
	}
}
