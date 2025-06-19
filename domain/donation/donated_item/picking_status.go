package donated_item

import "fmt"

const (
	PickingStatusPick    = "pick"
	PickingStatusDeliver = "deliver"
	PickingStatusBoth    = "both"
)

var (
	PickingStatuses = []PickingStatus{
		{Status: PickingStatusPick},
		{Status: PickingStatusDeliver},
		{Status: PickingStatusBoth},
	}
)

type PickingStatus struct {
	Status string
}

func NewPickingStatus(status string) (PickingStatus, error) {
	if !isValidPickingStatus(status) {
		return PickingStatus{}, fmt.Errorf("invalid picking status: %s", status)
	}
	return PickingStatus{
		Status: status,
	}, nil
}

func NewPickingStatusFromSchema(status string) PickingStatus {
	return PickingStatus{
		Status: status,
	}
}

func isValidPickingStatus(status string) bool {
	for _, s := range PickingStatuses {
		if s.Status == status {
			return true
		}
	}
	return false
}
