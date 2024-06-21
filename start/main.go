package main

import (
	"context"
	"fmt"
	"log"

	"github.com/eugeniaguerrero/temporal-hello-world/app"

	"go.temporal.io/sdk/client"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Start a workflow execution
	options := client.StartWorkflowOptions{
		ID:        "greeting_workflow",
		TaskQueue: app.GreetingTaskQueue,
	}

	greetingService := app.NewGreetingService()
	workflow := app.NewWorkflow(app.NewActivities(greetingService))
	name := "World"
	we, err := c.ExecuteWorkflow(context.Background(), options, workflow.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("Unable to complete Workflow", err)
	}

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get Workflow result", err)
	}

	printResults(result, we.GetID(), we.GetRunID())
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
