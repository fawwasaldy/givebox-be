package donated_item

import "fmt"

const (
	StatusOpened   = "opened"
	StatusAccepted = "accepted"
	StatusRejected = "rejected"
)

var (
	Statuses = []Status{
		{Status: StatusOpened},
		{Status: StatusAccepted},
		{Status: StatusRejected},
	}
)

type Status struct {
	Status string
}

func NewStatus(status string) (Status, error) {
	if !isValidStatus(status) {
		return Status{}, fmt.Errorf("invalid status: %s", status)
	}
	return Status{
		Status: status,
	}, nil
}

func NewStatusFromSchema(status string) Status {
	return Status{
		Status: status,
	}
}

func isValidStatus(status string) bool {
	for _, s := range Statuses {
		if s.Status == status {
			return true
		}
	}
	return false
}
