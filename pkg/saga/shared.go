package saga

const CreateLicenseTaskQueue = "CREATE_LICENSE_TASK_QUEUE"

type TransferDetails struct {
	Amount      float32
	FromAccount string
	ToAccount   string
	ReferenceID string
}
