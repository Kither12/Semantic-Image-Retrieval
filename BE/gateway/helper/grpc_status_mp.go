package helper

// grpcToHTTPStatus maps gRPC status codes to HTTP status codes.
func GrpcToHTTPStatus(grpcStatus int) int {
	grpcToHTTPMapping := map[int]int{
		0:  200, // OK
		1:  499, // Cancelled (non-standard)
		2:  500, // Unknown
		3:  400, // Invalid Argument
		4:  504, // Deadline Exceeded
		5:  404, // Not Found
		6:  409, // Already Exists
		7:  403, // Permission Denied
		8:  429, // Resource Exhausted
		9:  412, // Failed Precondition
		10: 409, // Aborted
		11: 416, // Out of Range
		12: 501, // Unimplemented
		13: 500, // Internal
		14: 503, // Unavailable
		15: 500, // Data Loss
		16: 401, // Unauthenticated
	}

	httpStatus, ok := grpcToHTTPMapping[grpcStatus]
	if !ok {
		// Default to 500 Internal Server Error if not found
		return 500
	}
	return httpStatus
}
