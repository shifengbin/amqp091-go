package wire

import "errors"

var (
	ErrUnexpectedMethod    = errors.New("Bad protocol: Received out of order method")
	ErrUnknownMethod       = errors.New("Unknown wire method id")
	ErrUnknownClass        = errors.New("Unknown wire class id")
	ErrUnknownFrameType    = errors.New("Bad frame: unknown type")
	ErrBadFrameSize        = errors.New("Bad frame: invalid size")
	ErrBadFrameTermination = errors.New("Bad frame: invalid terminator")
)

// XXX(ST) discuss where best to put these types, rather in buffer.go?
type Unit struct{}

type Decimal struct {
	Scale uint8
	Value uint32
}

type Timestamp uint64

type Table map[string]interface{}
