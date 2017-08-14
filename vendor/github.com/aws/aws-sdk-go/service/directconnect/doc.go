// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package directconnect provides the client and types for making API
// requests to AWS Direct Connect.
//
// AWS Direct Connect links your internal network to an AWS Direct Connect location
// over a standard 1 gigabit or 10 gigabit Ethernet fiber-optic cable. One end
// of the cable is connected to your router, the other to an AWS Direct Connect
// router. With this connection in place, you can create virtual interfaces
// directly to the AWS cloud (for example, to Amazon Elastic Compute Cloud (Amazon
// EC2) and Amazon Simple Storage Service (Amazon S3)) and to Amazon Virtual
// Private Cloud (Amazon VPC), bypassing Internet service providers in your
// network path. An AWS Direct Connect location provides access to AWS in the
// region it is associated with, as well as access to other US regions. For
// example, you can provision a single connection to any AWS Direct Connect
// location in the US and use it to access public AWS services in all US Regions
// and AWS GovCloud (US).
//
// See https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25 for more information on this service.
//
// See directconnect package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/directconnect/
//
// Using the Client
//
// To use the client for AWS Direct Connect you will first need
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
//   svc := directconnect.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Direct Connect client DirectConnect for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/directconnect/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.AllocateConnectionOnInterconnect(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("AllocateConnectionOnInterconnect result:")
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
//   result, err := svc.AllocateConnectionOnInterconnectWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package directconnect
