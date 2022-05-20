package smsprovider

type StatusCode string

const (
	StatusNotImplemented = StatusCode("NOT_IMPLEMENTED")
	StatusPending        = StatusCode("PENDING")
	StatusEnroute        = StatusCode("ENROUTE")
	StatusScheduled      = StatusCode("SCHEDULED")
	StatusDeleted        = StatusCode("DELETED")
	StatusAccepted       = StatusCode("ACCEPTED")
	StatusUnknown        = StatusCode("UNKNOWN")
	StatusSkipped        = StatusCode("SKIPPED")
	StatusUndeliverable  = StatusCode("UNDELIVERABLE")
	StatusDelivered      = StatusCode("DELIVERED")
	StatusExpired        = StatusCode("EXPIRED")
	StatusRejected       = StatusCode("REJECTED")
)

type StatusReason string

const (
	ReasonX = StatusReason("")
)
