package saga

// CreateLicenseTaskQueue - Queue name for workflow tasks.
const CreateLicenseTaskQueue = "CREATE_LICENSE_TASK_QUEUE"

// CreateLicenseInputArgs - Arguments to kick off the Temporal workflow.
type CreateLicenseInputArgs struct {
	OrgName     string
	ProfileName string
	LicenseName string
}
