package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

type Workflow struct {
	activities ActivityInterface
}

type WorkflowInterface interface {
	GreetingWorkflow(ctx workflow.Context, name string) (string, error)
}

// NewWorkflow returns a new Workflow instance
func NewWorkflow(activities ActivityInterface) WorkflowInterface {
	return &Workflow{activities: activities}
}

// GreetingWorkflow implements the WorkflowInterface
func (w *Workflow) GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	err := workflow.ExecuteActivity(ctx, w.activities.ComposeGreeting, name).Get(ctx, &result)
	if err != nil {
		return "", err
	}
	return result, nil
}
