package app_test

import (
	"testing"

	"github.com/eugeniaguerrero/temporal-hello-world/app"
	"github.com/google/go-cmp/cmp"
)

func Test_Greet(t *testing.T) {
	gs := app.NewGreetingService()
	got := gs.Greet("Bob")
	want := "Hello, Bob!"
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("Greet() mismatch (-want +got):\n%s", diff)
	}
}
