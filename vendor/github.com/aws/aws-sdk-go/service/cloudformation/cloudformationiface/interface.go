// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudformationiface provides an interface to enable mocking the AWS CloudFormation service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudformationiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// CloudFormationAPI provides an interface to enable mocking the
// cloudformation.CloudFormation service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CloudFormation.
//    func myFunc(svc cloudformationiface.CloudFormationAPI) bool {
//        // Make svc.CancelUpdateStack request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloudformation.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudFormationClient struct {
//        cloudformationiface.CloudFormationAPI
//    }
//    func (m *mockCloudFormationClient) CancelUpdateStack(input *cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudFormationClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CloudFormationAPI interface {
	CancelUpdateStack(*cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error)
	CancelUpdateStackWithContext(aws.Context, *cloudformation.CancelUpdateStackInput, ...request.Option) (*cloudformation.CancelUpdateStackOutput, error)
	CancelUpdateStackRequest(*cloudformation.CancelUpdateStackInput) (*request.Request, *cloudformation.CancelUpdateStackOutput)

	ContinueUpdateRollback(*cloudformation.ContinueUpdateRollbackInput) (*cloudformation.ContinueUpdateRollbackOutput, error)
	ContinueUpdateRollbackWithContext(aws.Context, *cloudformation.ContinueUpdateRollbackInput, ...request.Option) (*cloudformation.ContinueUpdateRollbackOutput, error)
	ContinueUpdateRollbackRequest(*cloudformation.ContinueUpdateRollbackInput) (*request.Request, *cloudformation.ContinueUpdateRollbackOutput)

	CreateChangeSet(*cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error)
	CreateChangeSetWithContext(aws.Context, *cloudformation.CreateChangeSetInput, ...request.Option) (*cloudformation.CreateChangeSetOutput, error)
	CreateChangeSetRequest(*cloudformation.CreateChangeSetInput) (*request.Request, *cloudformation.CreateChangeSetOutput)

	CreateStack(*cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error)
	CreateStackWithContext(aws.Context, *cloudformation.CreateStackInput, ...request.Option) (*cloudformation.CreateStackOutput, error)
	CreateStackRequest(*cloudformation.CreateStackInput) (*request.Request, *cloudformation.CreateStackOutput)

	CreateStackInstances(*cloudformation.CreateStackInstancesInput) (*cloudformation.CreateStackInstancesOutput, error)
	CreateStackInstancesWithContext(aws.Context, *cloudformation.CreateStackInstancesInput, ...request.Option) (*cloudformation.CreateStackInstancesOutput, error)
	CreateStackInstancesRequest(*cloudformation.CreateStackInstancesInput) (*request.Request, *cloudformation.CreateStackInstancesOutput)

	CreateStackSet(*cloudformation.CreateStackSetInput) (*cloudformation.CreateStackSetOutput, error)
	CreateStackSetWithContext(aws.Context, *cloudformation.CreateStackSetInput, ...request.Option) (*cloudformation.CreateStackSetOutput, error)
	CreateStackSetRequest(*cloudformation.CreateStackSetInput) (*request.Request, *cloudformation.CreateStackSetOutput)

	DeleteChangeSet(*cloudformation.DeleteChangeSetInput) (*cloudformation.DeleteChangeSetOutput, error)
	DeleteChangeSetWithContext(aws.Context, *cloudformation.DeleteChangeSetInput, ...request.Option) (*cloudformation.DeleteChangeSetOutput, error)
	DeleteChangeSetRequest(*cloudformation.DeleteChangeSetInput) (*request.Request, *cloudformation.DeleteChangeSetOutput)

	DeleteStack(*cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error)
	DeleteStackWithContext(aws.Context, *cloudformation.DeleteStackInput, ...request.Option) (*cloudformation.DeleteStackOutput, error)
	DeleteStackRequest(*cloudformation.DeleteStackInput) (*request.Request, *cloudformation.DeleteStackOutput)

	DeleteStackInstances(*cloudformation.DeleteStackInstancesInput) (*cloudformation.DeleteStackInstancesOutput, error)
	DeleteStackInstancesWithContext(aws.Context, *cloudformation.DeleteStackInstancesInput, ...request.Option) (*cloudformation.DeleteStackInstancesOutput, error)
	DeleteStackInstancesRequest(*cloudformation.DeleteStackInstancesInput) (*request.Request, *cloudformation.DeleteStackInstancesOutput)

	DeleteStackSet(*cloudformation.DeleteStackSetInput) (*cloudformation.DeleteStackSetOutput, error)
	DeleteStackSetWithContext(aws.Context, *cloudformation.DeleteStackSetInput, ...request.Option) (*cloudformation.DeleteStackSetOutput, error)
	DeleteStackSetRequest(*cloudformation.DeleteStackSetInput) (*request.Request, *cloudformation.DeleteStackSetOutput)

	DescribeAccountLimits(*cloudformation.DescribeAccountLimitsInput) (*cloudformation.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsWithContext(aws.Context, *cloudformation.DescribeAccountLimitsInput, ...request.Option) (*cloudformation.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsRequest(*cloudformation.DescribeAccountLimitsInput) (*request.Request, *cloudformation.DescribeAccountLimitsOutput)

	DescribeChangeSet(*cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error)
	DescribeChangeSetWithContext(aws.Context, *cloudformation.DescribeChangeSetInput, ...request.Option) (*cloudformation.DescribeChangeSetOutput, error)
	DescribeChangeSetRequest(*cloudformation.DescribeChangeSetInput) (*request.Request, *cloudformation.DescribeChangeSetOutput)

	DescribeStackEvents(*cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error)
	DescribeStackEventsWithContext(aws.Context, *cloudformation.DescribeStackEventsInput, ...request.Option) (*cloudformation.DescribeStackEventsOutput, error)
	DescribeStackEventsRequest(*cloudformation.DescribeStackEventsInput) (*request.Request, *cloudformation.DescribeStackEventsOutput)

	DescribeStackEventsPages(*cloudformation.DescribeStackEventsInput, func(*cloudformation.DescribeStackEventsOutput, bool) bool) error
	DescribeStackEventsPagesWithContext(aws.Context, *cloudformation.DescribeStackEventsInput, func(*cloudformation.DescribeStackEventsOutput, bool) bool, ...request.Option) error

	DescribeStackInstance(*cloudformation.DescribeStackInstanceInput) (*cloudformation.DescribeStackInstanceOutput, error)
	DescribeStackInstanceWithContext(aws.Context, *cloudformation.DescribeStackInstanceInput, ...request.Option) (*cloudformation.DescribeStackInstanceOutput, error)
	DescribeStackInstanceRequest(*cloudformation.DescribeStackInstanceInput) (*request.Request, *cloudformation.DescribeStackInstanceOutput)

	DescribeStackResource(*cloudformation.DescribeStackResourceInput) (*cloudformation.DescribeStackResourceOutput, error)
	DescribeStackResourceWithContext(aws.Context, *cloudformation.DescribeStackResourceInput, ...request.Option) (*cloudformation.DescribeStackResourceOutput, error)
	DescribeStackResourceRequest(*cloudformation.DescribeStackResourceInput) (*request.Request, *cloudformation.DescribeStackResourceOutput)

	DescribeStackResources(*cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error)
	DescribeStackResourcesWithContext(aws.Context, *cloudformation.DescribeStackResourcesInput, ...request.Option) (*cloudformation.DescribeStackResourcesOutput, error)
	DescribeStackResourcesRequest(*cloudformation.DescribeStackResourcesInput) (*request.Request, *cloudformation.DescribeStackResourcesOutput)

	DescribeStackSet(*cloudformation.DescribeStackSetInput) (*cloudformation.DescribeStackSetOutput, error)
	DescribeStackSetWithContext(aws.Context, *cloudformation.DescribeStackSetInput, ...request.Option) (*cloudformation.DescribeStackSetOutput, error)
	DescribeStackSetRequest(*cloudformation.DescribeStackSetInput) (*request.Request, *cloudformation.DescribeStackSetOutput)

	DescribeStackSetOperation(*cloudformation.DescribeStackSetOperationInput) (*cloudformation.DescribeStackSetOperationOutput, error)
	DescribeStackSetOperationWithContext(aws.Context, *cloudformation.DescribeStackSetOperationInput, ...request.Option) (*cloudformation.DescribeStackSetOperationOutput, error)
	DescribeStackSetOperationRequest(*cloudformation.DescribeStackSetOperationInput) (*request.Request, *cloudformation.DescribeStackSetOperationOutput)

	DescribeStacks(*cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error)
	DescribeStacksWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...request.Option) (*cloudformation.DescribeStacksOutput, error)
	DescribeStacksRequest(*cloudformation.DescribeStacksInput) (*request.Request, *cloudformation.DescribeStacksOutput)

	DescribeStacksPages(*cloudformation.DescribeStacksInput, func(*cloudformation.DescribeStacksOutput, bool) bool) error
	DescribeStacksPagesWithContext(aws.Context, *cloudformation.DescribeStacksInput, func(*cloudformation.DescribeStacksOutput, bool) bool, ...request.Option) error

	EstimateTemplateCost(*cloudformation.EstimateTemplateCostInput) (*cloudformation.EstimateTemplateCostOutput, error)
	EstimateTemplateCostWithContext(aws.Context, *cloudformation.EstimateTemplateCostInput, ...request.Option) (*cloudformation.EstimateTemplateCostOutput, error)
	EstimateTemplateCostRequest(*cloudformation.EstimateTemplateCostInput) (*request.Request, *cloudformation.EstimateTemplateCostOutput)

	ExecuteChangeSet(*cloudformation.ExecuteChangeSetInput) (*cloudformation.ExecuteChangeSetOutput, error)
	ExecuteChangeSetWithContext(aws.Context, *cloudformation.ExecuteChangeSetInput, ...request.Option) (*cloudformation.ExecuteChangeSetOutput, error)
	ExecuteChangeSetRequest(*cloudformation.ExecuteChangeSetInput) (*request.Request, *cloudformation.ExecuteChangeSetOutput)

	GetStackPolicy(*cloudformation.GetStackPolicyInput) (*cloudformation.GetStackPolicyOutput, error)
	GetStackPolicyWithContext(aws.Context, *cloudformation.GetStackPolicyInput, ...request.Option) (*cloudformation.GetStackPolicyOutput, error)
	GetStackPolicyRequest(*cloudformation.GetStackPolicyInput) (*request.Request, *cloudformation.GetStackPolicyOutput)

	GetTemplate(*cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error)
	GetTemplateWithContext(aws.Context, *cloudformation.GetTemplateInput, ...request.Option) (*cloudformation.GetTemplateOutput, error)
	GetTemplateRequest(*cloudformation.GetTemplateInput) (*request.Request, *cloudformation.GetTemplateOutput)

	GetTemplateSummary(*cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error)
	GetTemplateSummaryWithContext(aws.Context, *cloudformation.GetTemplateSummaryInput, ...request.Option) (*cloudformation.GetTemplateSummaryOutput, error)
	GetTemplateSummaryRequest(*cloudformation.GetTemplateSummaryInput) (*request.Request, *cloudformation.GetTemplateSummaryOutput)

	ListChangeSets(*cloudformation.ListChangeSetsInput) (*cloudformation.ListChangeSetsOutput, error)
	ListChangeSetsWithContext(aws.Context, *cloudformation.ListChangeSetsInput, ...request.Option) (*cloudformation.ListChangeSetsOutput, error)
	ListChangeSetsRequest(*cloudformation.ListChangeSetsInput) (*request.Request, *cloudformation.ListChangeSetsOutput)

	ListExports(*cloudformation.ListExportsInput) (*cloudformation.ListExportsOutput, error)
	ListExportsWithContext(aws.Context, *cloudformation.ListExportsInput, ...request.Option) (*cloudformation.ListExportsOutput, error)
	ListExportsRequest(*cloudformation.ListExportsInput) (*request.Request, *cloudformation.ListExportsOutput)

	ListExportsPages(*cloudformation.ListExportsInput, func(*cloudformation.ListExportsOutput, bool) bool) error
	ListExportsPagesWithContext(aws.Context, *cloudformation.ListExportsInput, func(*cloudformation.ListExportsOutput, bool) bool, ...request.Option) error

	ListImports(*cloudformation.ListImportsInput) (*cloudformation.ListImportsOutput, error)
	ListImportsWithContext(aws.Context, *cloudformation.ListImportsInput, ...request.Option) (*cloudformation.ListImportsOutput, error)
	ListImportsRequest(*cloudformation.ListImportsInput) (*request.Request, *cloudformation.ListImportsOutput)

	ListImportsPages(*cloudformation.ListImportsInput, func(*cloudformation.ListImportsOutput, bool) bool) error
	ListImportsPagesWithContext(aws.Context, *cloudformation.ListImportsInput, func(*cloudformation.ListImportsOutput, bool) bool, ...request.Option) error

	ListStackInstances(*cloudformation.ListStackInstancesInput) (*cloudformation.ListStackInstancesOutput, error)
	ListStackInstancesWithContext(aws.Context, *cloudformation.ListStackInstancesInput, ...request.Option) (*cloudformation.ListStackInstancesOutput, error)
	ListStackInstancesRequest(*cloudformation.ListStackInstancesInput) (*request.Request, *cloudformation.ListStackInstancesOutput)

	ListStackResources(*cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error)
	ListStackResourcesWithContext(aws.Context, *cloudformation.ListStackResourcesInput, ...request.Option) (*cloudformation.ListStackResourcesOutput, error)
	ListStackResourcesRequest(*cloudformation.ListStackResourcesInput) (*request.Request, *cloudformation.ListStackResourcesOutput)

	ListStackResourcesPages(*cloudformation.ListStackResourcesInput, func(*cloudformation.ListStackResourcesOutput, bool) bool) error
	ListStackResourcesPagesWithContext(aws.Context, *cloudformation.ListStackResourcesInput, func(*cloudformation.ListStackResourcesOutput, bool) bool, ...request.Option) error

	ListStackSetOperationResults(*cloudformation.ListStackSetOperationResultsInput) (*cloudformation.ListStackSetOperationResultsOutput, error)
	ListStackSetOperationResultsWithContext(aws.Context, *cloudformation.ListStackSetOperationResultsInput, ...request.Option) (*cloudformation.ListStackSetOperationResultsOutput, error)
	ListStackSetOperationResultsRequest(*cloudformation.ListStackSetOperationResultsInput) (*request.Request, *cloudformation.ListStackSetOperationResultsOutput)

	ListStackSetOperations(*cloudformation.ListStackSetOperationsInput) (*cloudformation.ListStackSetOperationsOutput, error)
	ListStackSetOperationsWithContext(aws.Context, *cloudformation.ListStackSetOperationsInput, ...request.Option) (*cloudformation.ListStackSetOperationsOutput, error)
	ListStackSetOperationsRequest(*cloudformation.ListStackSetOperationsInput) (*request.Request, *cloudformation.ListStackSetOperationsOutput)

	ListStackSets(*cloudformation.ListStackSetsInput) (*cloudformation.ListStackSetsOutput, error)
	ListStackSetsWithContext(aws.Context, *cloudformation.ListStackSetsInput, ...request.Option) (*cloudformation.ListStackSetsOutput, error)
	ListStackSetsRequest(*cloudformation.ListStackSetsInput) (*request.Request, *cloudformation.ListStackSetsOutput)

	ListStacks(*cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error)
	ListStacksWithContext(aws.Context, *cloudformation.ListStacksInput, ...request.Option) (*cloudformation.ListStacksOutput, error)
	ListStacksRequest(*cloudformation.ListStacksInput) (*request.Request, *cloudformation.ListStacksOutput)

	ListStacksPages(*cloudformation.ListStacksInput, func(*cloudformation.ListStacksOutput, bool) bool) error
	ListStacksPagesWithContext(aws.Context, *cloudformation.ListStacksInput, func(*cloudformation.ListStacksOutput, bool) bool, ...request.Option) error

	SetStackPolicy(*cloudformation.SetStackPolicyInput) (*cloudformation.SetStackPolicyOutput, error)
	SetStackPolicyWithContext(aws.Context, *cloudformation.SetStackPolicyInput, ...request.Option) (*cloudformation.SetStackPolicyOutput, error)
	SetStackPolicyRequest(*cloudformation.SetStackPolicyInput) (*request.Request, *cloudformation.SetStackPolicyOutput)

	SignalResource(*cloudformation.SignalResourceInput) (*cloudformation.SignalResourceOutput, error)
	SignalResourceWithContext(aws.Context, *cloudformation.SignalResourceInput, ...request.Option) (*cloudformation.SignalResourceOutput, error)
	SignalResourceRequest(*cloudformation.SignalResourceInput) (*request.Request, *cloudformation.SignalResourceOutput)

	StopStackSetOperation(*cloudformation.StopStackSetOperationInput) (*cloudformation.StopStackSetOperationOutput, error)
	StopStackSetOperationWithContext(aws.Context, *cloudformation.StopStackSetOperationInput, ...request.Option) (*cloudformation.StopStackSetOperationOutput, error)
	StopStackSetOperationRequest(*cloudformation.StopStackSetOperationInput) (*request.Request, *cloudformation.StopStackSetOperationOutput)

	UpdateStack(*cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error)
	UpdateStackWithContext(aws.Context, *cloudformation.UpdateStackInput, ...request.Option) (*cloudformation.UpdateStackOutput, error)
	UpdateStackRequest(*cloudformation.UpdateStackInput) (*request.Request, *cloudformation.UpdateStackOutput)

	UpdateStackSet(*cloudformation.UpdateStackSetInput) (*cloudformation.UpdateStackSetOutput, error)
	UpdateStackSetWithContext(aws.Context, *cloudformation.UpdateStackSetInput, ...request.Option) (*cloudformation.UpdateStackSetOutput, error)
	UpdateStackSetRequest(*cloudformation.UpdateStackSetInput) (*request.Request, *cloudformation.UpdateStackSetOutput)

	ValidateTemplate(*cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error)
	ValidateTemplateWithContext(aws.Context, *cloudformation.ValidateTemplateInput, ...request.Option) (*cloudformation.ValidateTemplateOutput, error)
	ValidateTemplateRequest(*cloudformation.ValidateTemplateInput) (*request.Request, *cloudformation.ValidateTemplateOutput)

	WaitUntilChangeSetCreateComplete(*cloudformation.DescribeChangeSetInput) error
	WaitUntilChangeSetCreateCompleteWithContext(aws.Context, *cloudformation.DescribeChangeSetInput, ...request.WaiterOption) error

	WaitUntilStackCreateComplete(*cloudformation.DescribeStacksInput) error
	WaitUntilStackCreateCompleteWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...request.WaiterOption) error

	WaitUntilStackDeleteComplete(*cloudformation.DescribeStacksInput) error
	WaitUntilStackDeleteCompleteWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...request.WaiterOption) error

	WaitUntilStackExists(*cloudformation.DescribeStacksInput) error
	WaitUntilStackExistsWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...request.WaiterOption) error

	WaitUntilStackUpdateComplete(*cloudformation.DescribeStacksInput) error
	WaitUntilStackUpdateCompleteWithContext(aws.Context, *cloudformation.DescribeStacksInput, ...request.WaiterOption) error
}

var _ CloudFormationAPI = (*cloudformation.CloudFormation)(nil)
