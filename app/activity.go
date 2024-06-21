package app

import (
	"context"
)

type ActivityInterface interface {
	ComposeGreeting(ctx context.Context, name string) (string, error)
}

type Activities struct {
	greetingService GreetingServiceInterface
}

// NewActivities returns a new Activities instance
func NewActivities(service GreetingServiceInterface) ActivityInterface {
	return &Activities{greetingService: service}
}

// ComposeGreeting implements the ActivityInterface
func (a *Activities) ComposeGreeting(ctx context.Context, name string) (string, error) {
	return a.greetingService.Greet(name), nil
}
