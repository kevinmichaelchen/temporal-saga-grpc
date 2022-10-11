package saga

import (
	"time"

	"go.uber.org/multierr"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func TransferMoney(ctx workflow.Context, transferDetails TransferDetails) (err error) {
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    3,
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

	//////////////
	// WITHDRAW //
	//////////////

	err = workflow.ExecuteActivity(ctx, ctrl.Withdraw, transferDetails).Get(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			errCompensation := workflow.ExecuteActivity(ctx, ctrl.WithdrawCompensation, transferDetails).Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()

	/////////////
	// DEPOSIT //
	/////////////

	err = workflow.ExecuteActivity(ctx, ctrl.Deposit, transferDetails).Get(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			errCompensation := workflow.ExecuteActivity(ctx, ctrl.DepositCompensation, transferDetails).Get(ctx, nil)
			err = multierr.Append(err, errCompensation)
		}

		// uncomment to have time to shut down worker to simulate worker rolling update and ensure that compensation sequence preserves after restart
		// workflow.Sleep(ctx, 10*time.Second)
	}()

	// Trigger an error to simulate failure
	err = workflow.ExecuteActivity(ctx, ctrl.StepWithError, transferDetails).Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
