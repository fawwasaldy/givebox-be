package identity

import "github.com/google/uuid"

var (
	NilID = uuid.Nil
)

type ID struct {
	ID uuid.UUID
}

func NewID(id string) ID {
	return ID{
		ID: uuid.MustParse(id),
	}
}

func NewIDFromSchema(id uuid.UUID) ID {
	return ID{
		ID: id,
	}
}

func (i ID) String() string {
	return i.ID.String()
}
