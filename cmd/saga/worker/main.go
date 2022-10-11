package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/kevinmichaelchen/temporal-saga-grpc/pkg/saga"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// TODO dial grpc connections to upstream services;
	//  use uber-go/fx;
	//  create a service struct w/ 3 grpc connections;
	//  register the struct itself which causes all its exported functions to become activities:
	//  from the docs:
	//  And activity struct is a structure with all its exported methods treated as activities.

	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, saga.CreateLicenseTaskQueue, worker.Options{})
	w.RegisterWorkflow(saga.TransferMoney)
	w.RegisterActivity(saga.Withdraw)
	w.RegisterActivity(saga.WithdrawCompensation)
	w.RegisterActivity(saga.Deposit)
	w.RegisterActivity(saga.DepositCompensation)
	w.RegisterActivity(saga.StepWithError)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
