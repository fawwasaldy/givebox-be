package message

import "errors"

var (
	ErrorGetMessageById     = errors.New("failed to get message by id")
	ErrorInvalidMessageType = errors.New("failed to convert to message type")
)
