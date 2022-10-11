package saga

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	// Mock activity implementation
	testDetails := CreateLicenseInputArgs{
		Amount:      1.00,
		FromAccount: "001-001",
		ToAccount:   "002-002",
		ReferenceID: "1234",
	}
	var ctrl *Controller
	env.OnActivity(ctrl.CreateOrg, mock.Anything, testDetails).Return(nil)
	env.OnActivity(ctrl.CreateProfile, mock.Anything, testDetails).Return(nil)
	env.ExecuteWorkflow(CreateLicense, testDetails)
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
}
