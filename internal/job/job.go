package job

type Status string

const (
	StatusPending Status = "pending"
	StatusRunning Status = "running"
	StatusDone    Status = "done"
	StatusFailed  Status = "failed"
)

type Job struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status Status `json:"status"`
}

func New(id, name string) Job {
	return Job{
		ID:     id,
		Name:   name,
		Status: StatusPending,
	}
}
