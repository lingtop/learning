package custom_error

const errorCodeBase = 0

const (
	// original codes
	UnknownError         int32 = errorCodeBase + 1
	InvalidJSONString    int32 = errorCodeBase + 2
	InputValidationError int32 = errorCodeBase + 3

	// proprietary codes
	DatabaseError int32 = errorCodeBase + 8
)
