package errs

import "google.golang.org/grpc/codes"

var (
	ErrInvalidArgument = &Error{
		Code:     "invalid_argument",
		Message:  "invalid argument",
		GRPCCode: codes.InvalidArgument,
	}
	ErrInternal = &Error{
		Code:     "internal_error",
		Message:  "internal server error",
		GRPCCode: codes.Internal,
	}
	ErrPermissionDenied = &Error{
		Code:     "permission_denied",
		Message:  "permission denied",
		GRPCCode: codes.PermissionDenied,
	}
)
