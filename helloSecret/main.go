package main

import (
	"context"
	"fmt"
	"hello-secret/internal/dagger"
	"strings"
)

// A Dagger module for saying hello to the world
// with a secret greeting ;) Original version:
// github.com/shykes/daggerverse/helloWorld
type HelloSecret struct {
	Greeting *dagger.Secret
	Name     string
}

func (m *HelloSecret) MyFunction(ctx context.Context, stringArg string) (*dagger.Container, error) {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg}).Sync(ctx)
}

// Change the greeting
func (hello *HelloSecret) WithGreeting(greeting *dagger.Secret) *HelloSecret {
	hello.Greeting = greeting
	return hello
}

// Change the name
func (hello *HelloSecret) WithName(name string) *HelloSecret {
	hello.Name = name
	return hello
}

// Say hello to the world!
func (hello *HelloSecret) Message() string {
	var (
		greeting = hello.Greeting
		name     = hello.Name
	)
	if greeting == nil {
		greeting = dag.SetSecret("sh", "Secret Hello")
	}
	if name == "" {
		name = "World"
	}
	ctx := context.Background()
	pt, err := greeting.Plaintext(ctx)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s, %s!", pt, name)
}

// SHOUT HELLO TO THE WORLD!
func (hello *HelloSecret) Shout() string {
	return strings.ToUpper(hello.Message() + "!!!!!!")
}
