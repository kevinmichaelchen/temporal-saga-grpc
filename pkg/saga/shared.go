package saga

import "time"

// CreateLicenseTaskQueue - Queue name for workflow tasks.
const CreateLicenseTaskQueue = "CREATE_LICENSE_TASK_QUEUE"

// CreateLicenseInputArgs - Arguments to kick off the Temporal workflow.
type CreateLicenseInputArgs struct {
	Org     Org
	Profile Profile
	License License
}

type Org struct {
	ID   string
	Name string
}

type Profile struct {
	ID       string
	FullName string
	OrgID    string
}

type License struct {
	ID     string
	Start  time.Time
	End    time.Time
	UserID string
}
