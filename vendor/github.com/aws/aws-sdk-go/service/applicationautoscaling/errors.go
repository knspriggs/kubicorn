// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationautoscaling

const (

	// ErrCodeConcurrentUpdateException for service response error code
	// "ConcurrentUpdateException".
	//
	// Concurrent updates caused an exception, for example, if you request an update
	// to an Application Auto Scaling resource that already has a pending update.
	ErrCodeConcurrentUpdateException = "ConcurrentUpdateException"

	// ErrCodeFailedResourceAccessException for service response error code
	// "FailedResourceAccessException".
	//
	// Failed access to resources caused an exception. This exception is thrown
	// when Application Auto Scaling is unable to retrieve the alarms associated
	// with a scaling policy due to a client error, for example, if the role ARN
	// specified for a scalable target does not have permission to call the CloudWatch
	// DescribeAlarms (http://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarms.html)
	// API operation on behalf of your account.
	ErrCodeFailedResourceAccessException = "FailedResourceAccessException"

	// ErrCodeInternalServiceException for service response error code
	// "InternalServiceException".
	//
	// The service encountered an internal error.
	ErrCodeInternalServiceException = "InternalServiceException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The next token supplied was invalid.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Your account exceeded a limit. This exception is thrown when a per-account
	// resource limit is exceeded. For more information, see Application Auto Scaling
	// Limits (http://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_as-app).
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeObjectNotFoundException for service response error code
	// "ObjectNotFoundException".
	//
	// The specified object could not be found. For any Put or Register API operation,
	// which depends on the existence of a scalable target, this exception is thrown
	// if the scalable target with the specified service namespace, resource ID,
	// and scalable dimension does not exist. For any Delete or Deregister API operation,
	// this exception is thrown if the resource that is to be deleted or deregistered
	// cannot be found.
	ErrCodeObjectNotFoundException = "ObjectNotFoundException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// An exception was thrown for a validation issue. Review the available parameters
	// for the API request.
	ErrCodeValidationException = "ValidationException"
)
