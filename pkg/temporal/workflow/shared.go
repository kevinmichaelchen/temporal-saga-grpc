package workflow

import "time"

// TaskQueue - Queue name for workflow tasks.
const TaskQueue = "CREATE_LICENSE_TASK_QUEUE"

// InputArgs - Arguments to kick off the Temporal workflow.
type InputArgs struct {
	Org     Org
	Profile Profile
	License License
}

// Org - Encapsulates all the necessary info to create an Org.
type Org struct {
	ID   string
	Name string
}

// Profile - Encapsulates all the necessary info to create an Profile.
type Profile struct {
	ID       string
	FullName string
	OrgID    string
}

// License - Encapsulates all the necessary info to create an License.
type License struct {
	ID     string
	Start  time.Time
	End    time.Time
	UserID string
}
