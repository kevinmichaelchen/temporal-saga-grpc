package saga

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func CreateLicense(ctx workflow.Context, args CreateLicenseInputArgs) (err error) {
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Second * 15,
		MaximumAttempts:    5,
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failures by default, this is just an example.
		RetryPolicy: retryPolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	// This nil pointer is deliberately nil!
	// From the docs:
	// https://pkg.go.dev/go.temporal.io/sdk/workflow#ExecuteActivity
	// To call an activity that is a member of a structure use the function
	// reference with nil receiver.
	var ctrl *Controller

	// Step 1: Create Org
	err = workflow.ExecuteActivity(ctx, ctrl.CreateOrg, args).Get(ctx, nil)
	if err != nil {
		return err
	}

	// Step 2: Create Profile
	err = workflow.ExecuteActivity(ctx, ctrl.CreateProfile, args).Get(ctx, nil)
	if err != nil {
		return err
	}

	// Step 3: Create License
	err = workflow.ExecuteActivity(ctx, ctrl.CreateLicense, args).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
