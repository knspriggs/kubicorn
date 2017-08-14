// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package applicationdiscoveryservice provides the client and types for making API
// requests to AWS Application Discovery Service.
//
// AWS Application Discovery Service helps you plan application migration projects
// by automatically identifying servers, virtual machines (VMs), software, and
// software dependencies running in your on-premises data centers. Application
// Discovery Service also collects application performance data, which can help
// you assess the outcome of your migration. The data collected by Application
// Discovery Service is securely retained in an Amazon-hosted and managed database
// in the cloud. You can export the data as a CSV or XML file into your preferred
// visualization tool or cloud-migration solution to plan your migration. For
// more information, see the Application Discovery Service FAQ (http://aws.amazon.com/application-discovery/faqs/).
//
// Application Discovery Service offers two modes of operation.
//
//    * Agentless discovery mode is recommended for environments that use VMware
//    vCenter Server. This mode doesn't require you to install an agent on each
//    host. Agentless discovery gathers server information regardless of the
//    operating systems, which minimizes the time required for initial on-premises
//    infrastructure assessment. Agentless discovery doesn't collect information
//    about software and software dependencies. It also doesn't work in non-VMware
//    environments. We recommend that you use agent-based discovery for non-VMware
//    environments and if you want to collect information about software and
//    software dependencies. You can also run agent-based and agentless discovery
//    simultaneously. Use agentless discovery to quickly complete the initial
//    infrastructure assessment and then install agents on select hosts to gather
//    information about software and software dependencies.
//
//    * Agent-based discovery mode collects a richer set of data than agentless
//    discovery by using Amazon software, the AWS Application Discovery Agent,
//    which you install on one or more hosts in your data center. The agent
//    captures infrastructure and application information, including an inventory
//    of installed software applications, system and process performance, resource
//    utilization, and network dependencies between workloads. The information
//    collected by agents is secured at rest and in transit to the Application
//    Discovery Service database in the cloud.
//
// Application Discovery Service integrates with application discovery solutions
// from AWS Partner Network (APN) partners. Third-party application discovery
// tools can query Application Discovery Service and write to the Application
// Discovery Service database using a public API. You can then import the data
// into either a visualization tool or cloud-migration solution.
//
// Application Discovery Service doesn't gather sensitive information. All data
// is handled according to the AWS Privacy Policy (http://aws.amazon.com/privacy/).
// You can operate Application Discovery Service using offline mode to inspect
// collected data before it is shared with the service.
//
// Your AWS account must be granted access to Application Discovery Service,
// a process called whitelisting. This is true for AWS partners and customers
// alike. To request access, sign up for AWS Application Discovery Service here
// (http://aws.amazon.com/application-discovery/preview/). We send you information
// about how to get started.
//
// This API reference provides descriptions, syntax, and usage examples for
// each of the actions and data types for Application Discovery Service. The
// topic for each action shows the API request parameters and the response.
// Alternatively, you can use one of the AWS SDKs to access an API that is tailored
// to the programming language or platform that you're using. For more information,
// see AWS SDKs (http://aws.amazon.com/tools/#SDKs).
//
// This guide is intended for use with the AWS Application Discovery Service
// User Guide (http://docs.aws.amazon.com/application-discovery/latest/userguide/).
//
// See https://docs.aws.amazon.com/goto/WebAPI/discovery-2015-11-01 for more information on this service.
//
// See applicationdiscoveryservice package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/applicationdiscoveryservice/
//
// Using the Client
//
// To use the client for AWS Application Discovery Service you will first need
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
//   svc := applicationdiscoveryservice.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Application Discovery Service client ApplicationDiscoveryService for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/applicationdiscoveryservice/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.AssociateConfigurationItemsToApplication(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("AssociateConfigurationItemsToApplication result:")
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
//   result, err := svc.AssociateConfigurationItemsToApplicationWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package applicationdiscoveryservice
