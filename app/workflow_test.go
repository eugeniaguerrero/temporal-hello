// workflow_test.go
package app

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/testsuite"
)

type MockActivities struct {
	ComposeGreetingFunc func(ctx context.Context, name string) (string, error)
}

func (m *MockActivities) ComposeGreeting(ctx context.Context, name string) (string, error) {
	return m.ComposeGreetingFunc(ctx, name)
}

func TestGreetingWorkflow(t *testing.T) {
	ts := testsuite.WorkflowTestSuite{}
	env := ts.NewTestWorkflowEnvironment()

	mockActivities := &MockActivities{
		ComposeGreetingFunc: func(ctx context.Context, name string) (string, error) {
			return "Hello, " + name + "!", nil
		},
	}

	workflow := NewWorkflow(mockActivities)

	env.RegisterWorkflow(workflow.GreetingWorkflow)
	env.RegisterActivity(mockActivities.ComposeGreeting)

	env.ExecuteWorkflow(workflow.GreetingWorkflow, "Temporal")

	assert.True(t, env.IsWorkflowCompleted())
	assert.NoError(t, env.GetWorkflowError())

	var result string
	err := env.GetWorkflowResult(&result)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, Temporal!", result)
}
