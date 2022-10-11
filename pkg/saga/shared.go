package saga

const CreateLicenseTaskQueue = "CREATE_LICENSE_TASK_QUEUE"

type CreateLicenseInputArgs struct {
	OrgName     string
	ProfileName string
	LicenseName string
}
