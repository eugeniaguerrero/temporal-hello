package main

import (
	"log"

	"github.com/eugeniaguerrero/temporal-hello-world/app"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	// Register workflow and activity implementations with the worker
	greetingService := app.NewGreetingService()
	activities := app.NewActivities(greetingService)
	workflow := app.NewWorkflow(activities)

	w := worker.New(c, app.GreetingTaskQueue, worker.Options{})
	w.RegisterWorkflow(workflow.GreetingWorkflow)
	w.RegisterActivity(activities.ComposeGreeting)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
