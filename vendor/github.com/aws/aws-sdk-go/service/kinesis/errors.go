// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

const (

	// ErrCodeExpiredIteratorException for service response error code
	// "ExpiredIteratorException".
	//
	// The provided iterator exceeds the maximum age allowed.
	ErrCodeExpiredIteratorException = "ExpiredIteratorException"

	// ErrCodeInvalidArgumentException for service response error code
	// "InvalidArgumentException".
	//
	// A specified parameter exceeds its restrictions, is not supported, or can't
	// be used. For more information, see the returned message.
	ErrCodeInvalidArgumentException = "InvalidArgumentException"

	// ErrCodeKMSAccessDeniedException for service response error code
	// "KMSAccessDeniedException".
	//
	// The ciphertext references a key that doesn't exist or that you don't have
	// access to.
	ErrCodeKMSAccessDeniedException = "KMSAccessDeniedException"

	// ErrCodeKMSDisabledException for service response error code
	// "KMSDisabledException".
	//
	// The request was rejected because the specified CMK isn't enabled.
	ErrCodeKMSDisabledException = "KMSDisabledException"

	// ErrCodeKMSInvalidStateException for service response error code
	// "KMSInvalidStateException".
	//
	// The request was rejected because the state of the specified resource isn't
	// valid for this request. For more information, see How Key State Affects Use
	// of a Customer Master Key (http://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
	// in the AWS Key Management Service Developer Guide.
	ErrCodeKMSInvalidStateException = "KMSInvalidStateException"

	// ErrCodeKMSNotFoundException for service response error code
	// "KMSNotFoundException".
	//
	// The request was rejected because the specified entity or resource couldn't
	// be found.
	ErrCodeKMSNotFoundException = "KMSNotFoundException"

	// ErrCodeKMSOptInRequired for service response error code
	// "KMSOptInRequired".
	//
	// The AWS access key ID needs a subscription for the service.
	ErrCodeKMSOptInRequired = "KMSOptInRequired"

	// ErrCodeKMSThrottlingException for service response error code
	// "KMSThrottlingException".
	//
	// The request was denied due to request throttling. For more information about
	// throttling, see Limits (http://docs.aws.amazon.com/kms/latest/developerguide/limits.html#requests-per-second)
	// in the AWS Key Management Service Developer Guide.
	ErrCodeKMSThrottlingException = "KMSThrottlingException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The requested resource exceeds the maximum number allowed, or the number
	// of concurrent stream requests exceeds the maximum number allowed (5).
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeProvisionedThroughputExceededException for service response error code
	// "ProvisionedThroughputExceededException".
	//
	// The request rate for the stream is too high, or the requested data is too
	// large for the available throughput. Reduce the frequency or size of your
	// requests. For more information, see Streams Limits (http://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html)
	// in the Amazon Kinesis Streams Developer Guide, and Error Retries and Exponential
	// Backoff in AWS (http://docs.aws.amazon.com/general/latest/gr/api-retries.html)
	// in the AWS General Reference.
	ErrCodeProvisionedThroughputExceededException = "ProvisionedThroughputExceededException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The resource is not available for this operation. For successful operation,
	// the resource needs to be in the ACTIVE state.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The requested resource could not be found. The stream might not be specified
	// correctly.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)
