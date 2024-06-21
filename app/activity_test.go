package app_test

import (
	"context"
	"testing"

	"github.com/eugeniaguerrero/temporal-hello-world/app"
	"github.com/google/go-cmp/cmp"
)

type MockGreet struct {
	GreetFunc func(name string) string
}

func (m *MockGreet) Greet(name string) string {
	return m.GreetFunc(name)
}

func Test_ComposeGreeting(t *testing.T) {
	gs := &MockGreet{
		GreetFunc: func(name string) string {
			return "Hello, " + name + "!"
		},
	}
	a := app.NewActivities(gs)

	want := "Hello, Sara!"
	got, err := a.ComposeGreeting(context.Background(), "Sara")
	if err != nil {
		t.Errorf("Expected no error, got: %s", err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Greet() mismatch (-want +got):\n%s", diff)
	}
}
