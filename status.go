package smsprovider

type StatusCode string

const (
	StatusPending       = StatusCode("PENDING")
	StatusEnroute       = StatusCode("ENROUTE")
	StatusScheduled     = StatusCode("SCHEDULED")
	StatusDeleted       = StatusCode("DELETED")
	StatusAccepted      = StatusCode("ACCEPTED")
	StatusUnknown       = StatusCode("UNKNOWN")
	StatusSkipped       = StatusCode("SKIPPED")
	StatusUndeliverable = StatusCode("UNDELIVERABLE")
	StatusDelivered     = StatusCode("DELIVERED")
	StatusExpired       = StatusCode("EXPIRED")
	StatusRejected      = StatusCode("REJECTED")
)

type ReasonCode string

const (
	ReasonX = ReasonCode("X")
)

type ErrorCode string

const (
	ErrorNotImplemented = ErrorCode("NOT_IMPLEMENTED")
	ErrorNotSupported   = ErrorCode("NOT_SUPPORTED")
)

var (
	ResponseNotImplemented = Response{
		Error: &Error{
			Code: ErrorNotImplemented,
		},
	}

	ResponseNotSupported = Response{
		Error: &Error{
			Code: ErrorNotSupported,
		},
	}
)
