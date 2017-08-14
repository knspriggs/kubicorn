// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package dynamodb provides the client and types for making API
// requests to Amazon DynamoDB.
//
// Amazon DynamoDB is a fully managed NoSQL database service that provides fast
// and predictable performance with seamless scalability. DynamoDB lets you
// offload the administrative burdens of operating and scaling a distributed
// database, so that you don't have to worry about hardware provisioning, setup
// and configuration, replication, software patching, or cluster scaling.
//
// With DynamoDB, you can create database tables that can store and retrieve
// any amount of data, and serve any level of request traffic. You can scale
// up or scale down your tables' throughput capacity without downtime or performance
// degradation, and use the AWS Management Console to monitor resource utilization
// and performance metrics.
//
// DynamoDB automatically spreads the data and traffic for your tables over
// a sufficient number of servers to handle your throughput and storage requirements,
// while maintaining consistent and fast performance. All of your data is stored
// on solid state disks (SSDs) and automatically replicated across multiple
// Availability Zones in an AWS region, providing built-in high availability
// and data durability.
//
// See https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10 for more information on this service.
//
// See dynamodb package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/
//
// Using the Client
//
// To use the client for Amazon DynamoDB you will first need
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
//   svc := dynamodb.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon DynamoDB client DynamoDB for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.BatchGetItem(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("BatchGetItem result:")
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
//   result, err := svc.BatchGetItemWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package dynamodb
