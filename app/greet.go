package app

type GreetingServiceInterface interface {
	Greet(name string) string
}

type GreetingService struct{}

// NewGreetingService returns a new GreetingService instance
func NewGreetingService() GreetingServiceInterface {
	return &GreetingService{}
}

// Greet implements the GreetingServiceInterface
func (s *GreetingService) Greet(name string) string {
	return "Hello, " + name + "!"
}
