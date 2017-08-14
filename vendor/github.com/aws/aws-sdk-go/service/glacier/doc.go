// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package glacier provides the client and types for making API
// requests to Amazon Glacier.
//
// Amazon Glacier is a storage solution for "cold data."
//
// Amazon Glacier is an extremely low-cost storage service that provides secure,
// durable, and easy-to-use storage for data backup and archival. With Amazon
// Glacier, customers can store their data cost effectively for months, years,
// or decades. Amazon Glacier also enables customers to offload the administrative
// burdens of operating and scaling storage to AWS, so they don't have to worry
// about capacity planning, hardware provisioning, data replication, hardware
// failure and recovery, or time-consuming hardware migrations.
//
// Amazon Glacier is a great storage choice when low storage cost is paramount,
// your data is rarely retrieved, and retrieval latency of several hours is
// acceptable. If your application requires fast or frequent access to your
// data, consider using Amazon S3. For more information, see Amazon Simple Storage
// Service (Amazon S3) (http://aws.amazon.com/s3/).
//
// You can store any kind of data in any format. There is no maximum limit on
// the total amount of data you can store in Amazon Glacier.
//
// If you are a first-time user of Amazon Glacier, we recommend that you begin
// by reading the following sections in the Amazon Glacier Developer Guide:
//
//    * What is Amazon Glacier (http://docs.aws.amazon.com/amazonglacier/latest/dev/introduction.html)
//    - This section of the Developer Guide describes the underlying data model,
//    the operations it supports, and the AWS SDKs that you can use to interact
//    with the service.
//
//    * Getting Started with Amazon Glacier (http://docs.aws.amazon.com/amazonglacier/latest/dev/amazon-glacier-getting-started.html)
//    - The Getting Started section walks you through the process of creating
//    a vault, uploading archives, creating jobs to download archives, retrieving
//    the job output, and deleting archives.
//
// See glacier package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/glacier/
//
// Using the Client
//
// To use the client for Amazon Glacier you will first need
// to create a new instance of it.
//
// When creating a client for an AWS service you'll first need to have a Session
// already created. The Session provides configuration that can be shared
// between multiple service clients. Additional configuration can be applied to
// the Session and service's client when they are constructed. The aws package's
// Config type contains several fields such as Region for the AWS Region the
// client should make API requests too. The optional Config value can be provided
// as the variadic argument for Sessions and client creation.
//
// Once the service's client is created you can use it to make API requests the
// AWS service. These clients are safe to use concurrently.
//
//   // Create a session to share configuration, and load external configuration.
//   sess := session.Must(session.NewSession())
//
//   // Create the service's client with the session.
//   svc := glacier.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Glacier client Glacier for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/glacier/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.AbortMultipartUpload(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("AbortMultipartUpload result:")
//   fmt.Println(result)
//
// Using the Client with Context
//
// The service's client also provides methods to make API requests with a Context
// value. This allows you to control the timeout, and cancellation of pending
// requests. These methods also take request Option as variadic parameter to apply
// additional configuration to the API request.
//
//   ctx := context.Background()
//
//   result, err := svc.AbortMultipartUploadWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package glacier
