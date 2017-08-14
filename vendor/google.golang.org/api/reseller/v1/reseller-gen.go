// Package reseller provides access to the Enterprise Apps Reseller API.
//
// See https://developers.google.com/google-apps/reseller/
//
// Usage example:
//
//   import "google.golang.org/api/reseller/v1"
//   ...
//   resellerService, err := reseller.New(oauthHttpClient)
package reseller // import "google.golang.org/api/reseller/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "reseller:v1"
const apiName = "reseller"
const apiVersion = "v1"
const basePath = "https://www.googleapis.com/apps/reseller/v1/"

// OAuth2 scopes used by this API.
const (
	// Manage users on your domain
	AppsOrderScope = "https://www.googleapis.com/auth/apps.order"

	// Manage users on your domain
	AppsOrderReadonlyScope = "https://www.googleapis.com/auth/apps.order.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Customers = NewCustomersService(s)
	s.Resellernotify = NewResellernotifyService(s)
	s.Subscriptions = NewSubscriptionsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Customers *CustomersService

	Resellernotify *ResellernotifyService

	Subscriptions *SubscriptionsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewCustomersService(s *Service) *CustomersService {
	rs := &CustomersService{s: s}
	return rs
}

type CustomersService struct {
	s *Service
}

func NewResellernotifyService(s *Service) *ResellernotifyService {
	rs := &ResellernotifyService{s: s}
	return rs
}

type ResellernotifyService struct {
	s *Service
}

func NewSubscriptionsService(s *Service) *SubscriptionsService {
	rs := &SubscriptionsService{s: s}
	return rs
}

type SubscriptionsService struct {
	s *Service
}

// Address: JSON template for address of a customer.
type Address struct {
	// AddressLine1: A customer's physical address. An address can be
	// composed of one to three lines. The addressline2 and addressLine3 are
	// optional.
	AddressLine1 string `json:"addressLine1,omitempty"`

	// AddressLine2: Line 2 of the address.
	AddressLine2 string `json:"addressLine2,omitempty"`

	// AddressLine3: Line 3 of the address.
	AddressLine3 string `json:"addressLine3,omitempty"`

	// ContactName: The customer contact's name. This is required.
	ContactName string `json:"contactName,omitempty"`

	// CountryCode: For countryCode information, see the ISO 3166 country
	// code elements. Verify that country is approved for resale of Google
	// products. This property is required when creating a new customer.
	CountryCode string `json:"countryCode,omitempty"`

	// Kind: Identifies the resource as a customer address. Value:
	// customers#address
	Kind string `json:"kind,omitempty"`

	// Locality: An example of a locality value is the city of San
	// Francisco.
	Locality string `json:"locality,omitempty"`

	// OrganizationName: The company or company division name. This is
	// required.
	OrganizationName string `json:"organizationName,omitempty"`

	// PostalCode: A postalCode example is a postal zip code such as 94043.
	// This property is required when creating a new customer.
	PostalCode string `json:"postalCode,omitempty"`

	// Region: An example of a region value is CA for the state of
	// California.
	Region string `json:"region,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AddressLine1") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AddressLine1") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Address) MarshalJSON() ([]byte, error) {
	type noMethod Address
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ChangePlanRequest: JSON template for the ChangePlan rpc request.
type ChangePlanRequest struct {
	// DealCode: Google-issued code (100 char max) for discounted pricing on
	// subscription plans. Deal code must be included in changePlan request
	// in order to receive discounted rate. This property is optional. If a
	// deal code has already been added to a subscription, this property may
	// be left empty and the existing discounted rate will still apply (if
	// not empty, only provide the deal code that is already present on the
	// subscription). If a deal code has never been added to a subscription
	// and this property is left blank, regular pricing will apply.
	DealCode string `json:"dealCode,omitempty"`

	// Kind: Identifies the resource as a subscription change plan request.
	// Value: subscriptions#changePlanRequest
	Kind string `json:"kind,omitempty"`

	// PlanName: The planName property is required. This is the name of the
	// subscription's payment plan. For more information about the Google
	// payment plans, see API concepts.
	//
	// Possible values are:
	// - ANNUAL_MONTHLY_PAY - The annual commitment plan with monthly
	// payments
	// - ANNUAL_YEARLY_PAY - The annual commitment plan with yearly payments
	//
	// - FLEXIBLE - The flexible plan
	// - TRIAL - The 30-day free trial plan
	PlanName string `json:"planName,omitempty"`

	// PurchaseOrderId: This is an optional property. This purchase order
	// (PO) information is for resellers to use for their company tracking
	// usage. If a purchaseOrderId value is given it appears in the API
	// responses and shows up in the invoice. The property accepts up to 80
	// plain text characters.
	PurchaseOrderId string `json:"purchaseOrderId,omitempty"`

	// Seats: This is a required property. The seats property is the number
	// of user seat licenses.
	Seats *Seats `json:"seats,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DealCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DealCode") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ChangePlanRequest) MarshalJSON() ([]byte, error) {
	type noMethod ChangePlanRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Customer: JSON template for a customer.
type Customer struct {
	// AlternateEmail: Like the "Customer email" in the reseller tools, this
	// email is the secondary contact used if something happens to the
	// customer's service such as service outage or a security issue. This
	// property is required when creating a new customer and should not use
	// the same domain as customerDomain.
	AlternateEmail string `json:"alternateEmail,omitempty"`

	// CustomerDomain: The customer's primary domain name string.
	// customerDomain is required when creating a new customer. Do not
	// include the www prefix in the domain when adding a customer.
	CustomerDomain string `json:"customerDomain,omitempty"`

	// CustomerDomainVerified: Whether the customer's primary domain has
	// been verified.
	CustomerDomainVerified bool `json:"customerDomainVerified,omitempty"`

	// CustomerId: This property will always be returned in a response as
	// the unique identifier generated by Google. In a request, this
	// property can be either the primary domain or the unique identifier
	// generated by Google.
	CustomerId string `json:"customerId,omitempty"`

	// Kind: Identifies the resource as a customer. Value: reseller#customer
	Kind string `json:"kind,omitempty"`

	// PhoneNumber: Customer contact phone number. This can be continuous
	// numbers, with spaces, etc. But it must be a real phone number and
	// not, for example, "123". See phone  local format conventions.
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// PostalAddress: A customer's address information. Each field has a
	// limit of 255 charcters.
	PostalAddress *Address `json:"postalAddress,omitempty"`

	// ResourceUiUrl: URL to customer's Admin console dashboard. The
	// read-only URL is generated by the API service. This is used if your
	// client application requires the customer to complete a task in the
	// Admin console.
	ResourceUiUrl string `json:"resourceUiUrl,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AlternateEmail") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AlternateEmail") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Customer) MarshalJSON() ([]byte, error) {
	type noMethod Customer
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RenewalSettings: JSON template for a subscription renewal settings.
type RenewalSettings struct {
	// Kind: Identifies the resource as a subscription renewal setting.
	// Value: subscriptions#renewalSettings
	Kind string `json:"kind,omitempty"`

	// RenewalType: Renewal settings for the annual commitment plan. For
	// more detailed information, see renewal options in the administrator
	// help center. When renewing a subscription, the renewalType is a
	// required property.
	RenewalType string `json:"renewalType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Kind") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RenewalSettings) MarshalJSON() ([]byte, error) {
	type noMethod RenewalSettings
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResellernotifyGetwatchdetailsResponse: JSON template for
// resellernotify getwatchdetails response.
type ResellernotifyGetwatchdetailsResponse struct {
	// ServiceAccountEmailAddresses: List of registered service accounts.
	ServiceAccountEmailAddresses []string `json:"serviceAccountEmailAddresses,omitempty"`

	// TopicName: Topic name of the PubSub
	TopicName string `json:"topicName,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "ServiceAccountEmailAddresses") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "ServiceAccountEmailAddresses") to include in API requests with the
	// JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResellernotifyGetwatchdetailsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ResellernotifyGetwatchdetailsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResellernotifyResource: JSON template for resellernotify response.
type ResellernotifyResource struct {
	// TopicName: Topic name of the PubSub
	TopicName string `json:"topicName,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "TopicName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "TopicName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResellernotifyResource) MarshalJSON() ([]byte, error) {
	type noMethod ResellernotifyResource
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Seats: JSON template for subscription seats.
type Seats struct {
	// Kind: Identifies the resource as a subscription change plan request.
	// Value: subscriptions#seats
	Kind string `json:"kind,omitempty"`

	// LicensedNumberOfSeats: Read-only field containing the current number
	// of licensed seats for FLEXIBLE Google-Apps subscriptions and
	// secondary subscriptions such as Google-Vault and Drive-storage.
	LicensedNumberOfSeats int64 `json:"licensedNumberOfSeats,omitempty"`

	// MaximumNumberOfSeats: The maximumNumberOfSeats property is the
	// maximum number of licenses that the customer can purchase. This
	// property applies to plans other than the annual commitment plan. How
	// a user's licenses are managed depends on the subscription's payment
	// plan:
	// - annual commitment plan (with monthly or yearly payments) — For
	// this plan, a reseller is invoiced on the number of user licenses in
	// the numberOfSeats property. The maximumNumberOfSeats property is a
	// read-only property in the API's response.
	// - flexible plan — For this plan, a reseller is invoiced on the
	// actual number of users which is capped by the maximumNumberOfSeats.
	// This is the maximum number of user licenses a customer has for user
	// license provisioning. This quantity can be increased up to the
	// maximum limit defined in the reseller's contract. And the minimum
	// quantity is the current number of users in the customer account.
	// - 30-day free trial plan — A subscription in a 30-day free trial is
	// restricted to maximum 10 seats.
	MaximumNumberOfSeats int64 `json:"maximumNumberOfSeats,omitempty"`

	// NumberOfSeats: The numberOfSeats property holds the customer's number
	// of user licenses. How a user's licenses are managed depends on the
	// subscription's plan:
	// - annual commitment plan (with monthly or yearly pay) — For this
	// plan, a reseller is invoiced on the number of user licenses in the
	// numberOfSeats property. This is the maximum number of user licenses
	// that a reseller's customer can create. The reseller can add more
	// licenses, but once set, the numberOfSeats can not be reduced until
	// renewal. The reseller is invoiced based on the numberOfSeats value
	// regardless of how many of these user licenses are provisioned users.
	//
	// - flexible plan — For this plan, a reseller is invoiced on the
	// actual number of users which is capped by the maximumNumberOfSeats.
	// The numberOfSeats property is not used in the request or response for
	// flexible plan customers.
	// - 30-day free trial plan — The numberOfSeats property is not used
	// in the request or response for an account in a 30-day trial.
	NumberOfSeats int64 `json:"numberOfSeats,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Kind") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Seats) MarshalJSON() ([]byte, error) {
	type noMethod Seats
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Subscription: JSON template for a subscription.
type Subscription struct {
	// BillingMethod: Read-only field that returns the current billing
	// method for a subscription.
	BillingMethod string `json:"billingMethod,omitempty"`

	// CreationTime: The creationTime property is the date when subscription
	// was created. It is in milliseconds using the Epoch format. See an
	// example Epoch converter.
	CreationTime int64 `json:"creationTime,omitempty,string"`

	// CustomerDomain: Primary domain name of the customer
	CustomerDomain string `json:"customerDomain,omitempty"`

	// CustomerId: This property will always be returned in a response as
	// the unique identifier generated by Google. In a request, this
	// property can be either the primary domain or the unique identifier
	// generated by Google.
	CustomerId string `json:"customerId,omitempty"`

	// DealCode: Google-issued code (100 char max) for discounted pricing on
	// subscription plans. Deal code must be included in insert requests in
	// order to receive discounted rate. This property is optional, regular
	// pricing applies if left empty.
	DealCode string `json:"dealCode,omitempty"`

	// Kind: Identifies the resource as a Subscription. Value:
	// reseller#subscription
	Kind string `json:"kind,omitempty"`

	// Plan: The plan property is required. In this version of the API, the
	// G Suite plans are the flexible plan, annual commitment plan, and the
	// 30-day free trial plan. For more information about the API"s payment
	// plans, see the API concepts.
	Plan *SubscriptionPlan `json:"plan,omitempty"`

	// PurchaseOrderId: This is an optional property. This purchase order
	// (PO) information is for resellers to use for their company tracking
	// usage. If a purchaseOrderId value is given it appears in the API
	// responses and shows up in the invoice. The property accepts up to 80
	// plain text characters.
	PurchaseOrderId string `json:"purchaseOrderId,omitempty"`

	// RenewalSettings: Renewal settings for the annual commitment plan. For
	// more detailed information, see renewal options in the administrator
	// help center.
	RenewalSettings *RenewalSettings `json:"renewalSettings,omitempty"`

	// ResourceUiUrl: URL to customer's Subscriptions page in the Admin
	// console. The read-only URL is generated by the API service. This is
	// used if your client application requires the customer to complete a
	// task using the Subscriptions page in the Admin console.
	ResourceUiUrl string `json:"resourceUiUrl,omitempty"`

	// Seats: This is a required property. The number and limit of user seat
	// licenses in the plan.
	Seats *Seats `json:"seats,omitempty"`

	// SkuId: A required property. The skuId is a unique system identifier
	// for a product's SKU assigned to a customer in the subscription. For
	// products and SKUs available in this version of the API, see  Product
	// and SKU IDs.
	SkuId string `json:"skuId,omitempty"`

	// SkuName: Read-only external display name for a product's SKU assigned
	// to a customer in the subscription. SKU names are subject to change at
	// Google's discretion. For products and SKUs available in this version
	// of the API, see  Product and SKU IDs.
	SkuName string `json:"skuName,omitempty"`

	// Status: This is an optional property.
	Status string `json:"status,omitempty"`

	// SubscriptionId: The subscriptionId is the subscription identifier and
	// is unique for each customer. This is a required property. Since a
	// subscriptionId changes when a subscription is updated, we recommend
	// not using this ID as a key for persistent data. Use the
	// subscriptionId as described in retrieve all reseller subscriptions.
	SubscriptionId string `json:"subscriptionId,omitempty"`

	// SuspensionReasons: Read-only field containing an enumerable of all
	// the current suspension reasons for a subscription. It is possible for
	// a subscription to have many concurrent, overlapping suspension
	// reasons. A subscription's STATUS is SUSPENDED until all pending
	// suspensions are removed.
	//
	// Possible options include:
	// - PENDING_TOS_ACCEPTANCE - The customer has not logged in and
	// accepted the G Suite Resold Terms of Services.
	// - RENEWAL_WITH_TYPE_CANCEL - The customer's commitment ended and
	// their service was cancelled at the end of their term.
	// - RESELLER_INITIATED - A manual suspension invoked by a Reseller.
	// - TRIAL_ENDED - The customer's trial expired without a plan selected.
	//
	// - OTHER - The customer is suspended for an internal Google reason
	// (e.g. abuse or otherwise).
	SuspensionReasons []string `json:"suspensionReasons,omitempty"`

	// TransferInfo: Read-only transfer related information for the
	// subscription. For more information, see retrieve transferable
	// subscriptions for a customer.
	TransferInfo *SubscriptionTransferInfo `json:"transferInfo,omitempty"`

	// TrialSettings: The G Suite annual commitment and flexible payment
	// plans can be in a 30-day free trial. For more information, see the
	// API concepts.
	TrialSettings *SubscriptionTrialSettings `json:"trialSettings,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "BillingMethod") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BillingMethod") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Subscription) MarshalJSON() ([]byte, error) {
	type noMethod Subscription
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SubscriptionPlan: The plan property is required. In this version of
// the API, the G Suite plans are the flexible plan, annual commitment
// plan, and the 30-day free trial plan. For more information about the
// API"s payment plans, see the API concepts.
type SubscriptionPlan struct {
	// CommitmentInterval: In this version of the API, annual commitment
	// plan's interval is one year.
	CommitmentInterval *SubscriptionPlanCommitmentInterval `json:"commitmentInterval,omitempty"`

	// IsCommitmentPlan: The isCommitmentPlan property's boolean value
	// identifies the plan as an annual commitment plan:
	// - true — The subscription's plan is an annual commitment plan.
	// - false — The plan is not an annual commitment plan.
	IsCommitmentPlan bool `json:"isCommitmentPlan,omitempty"`

	// PlanName: The planName property is required. This is the name of the
	// subscription's plan. For more information about the Google payment
	// plans, see the API concepts.
	//
	// Possible values are:
	// - ANNUAL_MONTHLY_PAY — The annual commitment plan with monthly
	// payments
	// - ANNUAL_YEARLY_PAY — The annual commitment plan with yearly
	// payments
	// - FLEXIBLE — The flexible plan
	// - TRIAL — The 30-day free trial plan. A subscription in trial will
	// be suspended after the 30th free day if no payment plan is assigned.
	// Calling changePlan will assign a payment plan to a trial but will not
	// activate the plan. A trial will automatically begin its assigned
	// payment plan after its 30th free day or immediately after calling
	// startPaidService.
	PlanName string `json:"planName,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CommitmentInterval")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CommitmentInterval") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SubscriptionPlan) MarshalJSON() ([]byte, error) {
	type noMethod SubscriptionPlan
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SubscriptionPlanCommitmentInterval: In this version of the API,
// annual commitment plan's interval is one year.
type SubscriptionPlanCommitmentInterval struct {
	// EndTime: An annual commitment plan's interval's endTime in
	// milliseconds using the UNIX Epoch format. See an example Epoch
	// converter.
	EndTime int64 `json:"endTime,omitempty,string"`

	// StartTime: An annual commitment plan's interval's startTime in
	// milliseconds using UNIX Epoch format. See an example Epoch converter.
	StartTime int64 `json:"startTime,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "EndTime") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EndTime") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SubscriptionPlanCommitmentInterval) MarshalJSON() ([]byte, error) {
	type noMethod SubscriptionPlanCommitmentInterval
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SubscriptionTransferInfo: Read-only transfer related information for
// the subscription. For more information, see retrieve transferable
// subscriptions for a customer.
type SubscriptionTransferInfo struct {
	// MinimumTransferableSeats: When inserting a subscription, this is the
	// minimum number of seats listed in the transfer order for this
	// product. For example, if the customer has 20 users, the reseller
	// cannot place a transfer order of 15 seats. The minimum is 20 seats.
	MinimumTransferableSeats int64 `json:"minimumTransferableSeats,omitempty"`

	// TransferabilityExpirationTime: The time when transfer token or intent
	// to transfer will expire. The time is in milliseconds using UNIX Epoch
	// format.
	TransferabilityExpirationTime int64 `json:"transferabilityExpirationTime,omitempty,string"`

	// ForceSendFields is a list of field names (e.g.
	// "MinimumTransferableSeats") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MinimumTransferableSeats")
	// to include in API requests with the JSON null value. By default,
	// fields with empty values are omitted from API requests. However, any
	// field with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SubscriptionTransferInfo) MarshalJSON() ([]byte, error) {
	type noMethod SubscriptionTransferInfo
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SubscriptionTrialSettings: The G Suite annual commitment and flexible
// payment plans can be in a 30-day free trial. For more information,
// see the API concepts.
type SubscriptionTrialSettings struct {
	// IsInTrial: Determines if a subscription's plan is in a 30-day free
	// trial or not:
	// - true — The plan is in trial.
	// - false — The plan is not in trial.
	IsInTrial bool `json:"isInTrial,omitempty"`

	// TrialEndTime: Date when the trial ends. The value is in milliseconds
	// using the UNIX Epoch format. See an example Epoch converter.
	TrialEndTime int64 `json:"trialEndTime,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "IsInTrial") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IsInTrial") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *SubscriptionTrialSettings) MarshalJSON() ([]byte, error) {
	type noMethod SubscriptionTrialSettings
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Subscriptions: JSON template for a subscription list.
type Subscriptions struct {
	// Kind: Identifies the resource as a collection of subscriptions.
	// Value: reseller#subscriptions
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Subscriptions: The subscriptions in this page of results.
	Subscriptions []*Subscription `json:"subscriptions,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Kind") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Kind") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Subscriptions) MarshalJSON() ([]byte, error) {
	type noMethod Subscriptions
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "reseller.customers.get":

type CustomersGetCall struct {
	s            *Service
	customerId   string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Get a customer account.
func (r *CustomersService) Get(customerId string) *CustomersGetCall {
	c := &CustomersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.customerId = customerId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersGetCall) Fields(s ...googleapi.Field) *CustomersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CustomersGetCall) IfNoneMatch(entityTag string) *CustomersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersGetCall) Context(ctx context.Context) *CustomersGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "customers/{customerId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"customerId": c.customerId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.customers.get" call.
// Exactly one of *Customer or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Customer.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *CustomersGetCall) Do(opts ...googleapi.CallOption) (*Customer, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Customer{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a customer account.",
	//   "httpMethod": "GET",
	//   "id": "reseller.customers.get",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Either the customer's primary domain name or the customer's unique identifier. If using the domain name, we do not recommend using a customerId as a key for persistent data. If the domain name for a customerId is changed, the Google system automatically updates.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}",
	//   "response": {
	//     "$ref": "Customer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order",
	//     "https://www.googleapis.com/auth/apps.order.readonly"
	//   ]
	// }

}

// method id "reseller.customers.insert":

type CustomersInsertCall struct {
	s          *Service
	customer   *Customer
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Order a new customer's account.
func (r *CustomersService) Insert(customer *Customer) *CustomersInsertCall {
	c := &CustomersInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.customer = customer
	return c
}

// CustomerAuthToken sets the optional parameter "customerAuthToken":
// The customerAuthToken query string is required when creating a resold
// account that transfers a direct customer's subscription or transfers
// another reseller customer's subscription to your reseller management.
// This is a hexadecimal authentication token needed to complete the
// subscription transfer. For more information, see the administrator
// help center.
func (c *CustomersInsertCall) CustomerAuthToken(customerAuthToken string) *CustomersInsertCall {
	c.urlParams_.Set("customerAuthToken", customerAuthToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersInsertCall) Fields(s ...googleapi.Field) *CustomersInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersInsertCall) Context(ctx context.Context) *CustomersInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.customer)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "customers")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.customers.insert" call.
// Exactly one of *Customer or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Customer.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *CustomersInsertCall) Do(opts ...googleapi.CallOption) (*Customer, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Customer{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Order a new customer's account.",
	//   "httpMethod": "POST",
	//   "id": "reseller.customers.insert",
	//   "parameters": {
	//     "customerAuthToken": {
	//       "description": "The customerAuthToken query string is required when creating a resold account that transfers a direct customer's subscription or transfers another reseller customer's subscription to your reseller management. This is a hexadecimal authentication token needed to complete the subscription transfer. For more information, see the administrator help center.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers",
	//   "request": {
	//     "$ref": "Customer"
	//   },
	//   "response": {
	//     "$ref": "Customer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.customers.patch":

type CustomersPatchCall struct {
	s          *Service
	customerId string
	customer   *Customer
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Patch: Update a customer account's settings. This method supports
// patch semantics.
func (r *CustomersService) Patch(customerId string, customer *Customer) *CustomersPatchCall {
	c := &CustomersPatchCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.customerId = customerId
	c.customer = customer
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersPatchCall) Fields(s ...googleapi.Field) *CustomersPatchCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersPatchCall) Context(ctx context.Context) *CustomersPatchCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersPatchCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersPatchCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.customer)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "customers/{customerId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"customerId": c.customerId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.customers.patch" call.
// Exactly one of *Customer or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Customer.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *CustomersPatchCall) Do(opts ...googleapi.CallOption) (*Customer, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Customer{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a customer account's settings. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "reseller.customers.patch",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Either the customer's primary domain name or the customer's unique identifier. If using the domain name, we do not recommend using a customerId as a key for persistent data. If the domain name for a customerId is changed, the Google system automatically updates.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}",
	//   "request": {
	//     "$ref": "Customer"
	//   },
	//   "response": {
	//     "$ref": "Customer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.customers.update":

type CustomersUpdateCall struct {
	s          *Service
	customerId string
	customer   *Customer
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Update a customer account's settings.
func (r *CustomersService) Update(customerId string, customer *Customer) *CustomersUpdateCall {
	c := &CustomersUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.customerId = customerId
	c.customer = customer
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CustomersUpdateCall) Fields(s ...googleapi.Field) *CustomersUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CustomersUpdateCall) Context(ctx context.Context) *CustomersUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CustomersUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CustomersUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.customer)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "customers/{customerId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"customerId": c.customerId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.customers.update" call.
// Exactly one of *Customer or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Customer.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *CustomersUpdateCall) Do(opts ...googleapi.CallOption) (*Customer, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Customer{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a customer account's settings.",
	//   "httpMethod": "PUT",
	//   "id": "reseller.customers.update",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Either the customer's primary domain name or the customer's unique identifier. If using the domain name, we do not recommend using a customerId as a key for persistent data. If the domain name for a customerId is changed, the Google system automatically updates.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}",
	//   "request": {
	//     "$ref": "Customer"
	//   },
	//   "response": {
	//     "$ref": "Customer"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.resellernotify.getwatchdetails":

type ResellernotifyGetwatchdetailsCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Getwatchdetails: Returns all the details of the watch corresponding
// to the reseller.
func (r *ResellernotifyService) Getwatchdetails() *ResellernotifyGetwatchdetailsCall {
	c := &ResellernotifyGetwatchdetailsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ResellernotifyGetwatchdetailsCall) Fields(s ...googleapi.Field) *ResellernotifyGetwatchdetailsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ResellernotifyGetwatchdetailsCall) IfNoneMatch(entityTag string) *ResellernotifyGetwatchdetailsCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ResellernotifyGetwatchdetailsCall) Context(ctx context.Context) *ResellernotifyGetwatchdetailsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ResellernotifyGetwatchdetailsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ResellernotifyGetwatchdetailsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "resellernotify/getwatchdetails")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.resellernotify.getwatchdetails" call.
// Exactly one of *ResellernotifyGetwatchdetailsResponse or error will
// be non-nil. Any non-2xx status code is an error. Response headers are
// in either
// *ResellernotifyGetwatchdetailsResponse.ServerResponse.Header or (if a
// response was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ResellernotifyGetwatchdetailsCall) Do(opts ...googleapi.CallOption) (*ResellernotifyGetwatchdetailsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ResellernotifyGetwatchdetailsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns all the details of the watch corresponding to the reseller.",
	//   "httpMethod": "GET",
	//   "id": "reseller.resellernotify.getwatchdetails",
	//   "path": "resellernotify/getwatchdetails",
	//   "response": {
	//     "$ref": "ResellernotifyGetwatchdetailsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order",
	//     "https://www.googleapis.com/auth/apps.order.readonly"
	//   ]
	// }

}

// method id "reseller.resellernotify.register":

type ResellernotifyRegisterCall struct {
	s          *Service
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Register: Registers a Reseller for receiving notifications.
func (r *ResellernotifyService) Register() *ResellernotifyRegisterCall {
	c := &ResellernotifyRegisterCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// ServiceAccountEmailAddress sets the optional parameter
// "serviceAccountEmailAddress": The service account which will own the
// created Cloud-PubSub topic.
func (c *ResellernotifyRegisterCall) ServiceAccountEmailAddress(serviceAccountEmailAddress string) *ResellernotifyRegisterCall {
	c.urlParams_.Set("serviceAccountEmailAddress", serviceAccountEmailAddress)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ResellernotifyRegisterCall) Fields(s ...googleapi.Field) *ResellernotifyRegisterCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ResellernotifyRegisterCall) Context(ctx context.Context) *ResellernotifyRegisterCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ResellernotifyRegisterCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ResellernotifyRegisterCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "resellernotify/register")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.resellernotify.register" call.
// Exactly one of *ResellernotifyResource or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ResellernotifyResource.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ResellernotifyRegisterCall) Do(opts ...googleapi.CallOption) (*ResellernotifyResource, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ResellernotifyResource{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Registers a Reseller for receiving notifications.",
	//   "httpMethod": "POST",
	//   "id": "reseller.resellernotify.register",
	//   "parameters": {
	//     "serviceAccountEmailAddress": {
	//       "description": "The service account which will own the created Cloud-PubSub topic.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "resellernotify/register",
	//   "response": {
	//     "$ref": "ResellernotifyResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.resellernotify.unregister":

type ResellernotifyUnregisterCall struct {
	s          *Service
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Unregister: Unregisters a Reseller for receiving notifications.
func (r *ResellernotifyService) Unregister() *ResellernotifyUnregisterCall {
	c := &ResellernotifyUnregisterCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// ServiceAccountEmailAddress sets the optional parameter
// "serviceAccountEmailAddress": The service account which owns the
// Cloud-PubSub topic.
func (c *ResellernotifyUnregisterCall) ServiceAccountEmailAddress(serviceAccountEmailAddress string) *ResellernotifyUnregisterCall {
	c.urlParams_.Set("serviceAccountEmailAddress", serviceAccountEmailAddress)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ResellernotifyUnregisterCall) Fields(s ...googleapi.Field) *ResellernotifyUnregisterCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ResellernotifyUnregisterCall) Context(ctx context.Context) *ResellernotifyUnregisterCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ResellernotifyUnregisterCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ResellernotifyUnregisterCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "resellernotify/unregister")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.resellernotify.unregister" call.
// Exactly one of *ResellernotifyResource or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ResellernotifyResource.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ResellernotifyUnregisterCall) Do(opts ...googleapi.CallOption) (*ResellernotifyResource, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ResellernotifyResource{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Unregisters a Reseller for receiving notifications.",
	//   "httpMethod": "POST",
	//   "id": "reseller.resellernotify.unregister",
	//   "parameters": {
	//     "serviceAccountEmailAddress": {
	//       "description": "The service account which owns the Cloud-PubSub topic.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "resellernotify/unregister",
	//   "response": {
	//     "$ref": "ResellernotifyResource"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.activate":

type SubscriptionsActivateCall struct {
	s              *Service
	customerId     string
	subscriptionId string
	urlParams_     gensupport.URLParams
	ctx_           context.Context
	header_        http.Header
}

// Activate: Activates a subscription previously suspended by the
// reseller
func (r *SubscriptionsService) Activate(customerId string, subscriptionId string) *SubscriptionsActivateCall {
	c := &SubscriptionsActivateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.customerId = customerId
	c.subscriptionId = subscriptionId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsActivateCall) Fields(s ...googleapi.Field) *SubscriptionsActivateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SubscriptionsActivateCall) Context(ctx context.Context) *SubscriptionsActivateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SubscriptionsActivateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SubscriptionsActivateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "customers/{customerId}/subscriptions/{subscriptionId}/activate")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.subscriptions.activate" call.
// Exactly one of *Subscription or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Subscription.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SubscriptionsActivateCall) Do(opts ...googleapi.CallOption) (*Subscription, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Subscription{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Activates a subscription previously suspended by the reseller",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.activate",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Either the customer's primary domain name or the customer's unique identifier. If using the domain name, we do not recommend using a customerId as a key for persistent data. If the domain name for a customerId is changed, the Google system automatically updates.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "This is a required property. The subscriptionId is the subscription identifier and is unique for each customer. Since a subscriptionId changes when a subscription is updated, we recommend to not use this ID as a key for persistent data. And the subscriptionId can be found using the retrieve all reseller subscriptions method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}/activate",
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.changePlan":

type SubscriptionsChangePlanCall struct {
	s                 *Service
	customerId        string
	subscriptionId    string
	changeplanrequest *ChangePlanRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// ChangePlan: Update a subscription plan. Use this method to update a
// plan for a 30-day trial or a flexible plan subscription to an annual
// commitment plan with monthly or yearly payments.
func (r *SubscriptionsService) ChangePlan(customerId string, subscriptionId string, changeplanrequest *ChangePlanRequest) *SubscriptionsChangePlanCall {
	c := &SubscriptionsChangePlanCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.customerId = customerId
	c.subscriptionId = subscriptionId
	c.changeplanrequest = changeplanrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsChangePlanCall) Fields(s ...googleapi.Field) *SubscriptionsChangePlanCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SubscriptionsChangePlanCall) Context(ctx context.Context) *SubscriptionsChangePlanCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SubscriptionsChangePlanCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SubscriptionsChangePlanCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.changeplanrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "customers/{customerId}/subscriptions/{subscriptionId}/changePlan")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.subscriptions.changePlan" call.
// Exactly one of *Subscription or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Subscription.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SubscriptionsChangePlanCall) Do(opts ...googleapi.CallOption) (*Subscription, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Subscription{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a subscription plan. Use this method to update a plan for a 30-day trial or a flexible plan subscription to an annual commitment plan with monthly or yearly payments.",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.changePlan",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Either the customer's primary domain name or the customer's unique identifier. If using the domain name, we do not recommend using a customerId as a key for persistent data. If the domain name for a customerId is changed, the Google system automatically updates.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "This is a required property. The subscriptionId is the subscription identifier and is unique for each customer. Since a subscriptionId changes when a subscription is updated, we recommend to not use this ID as a key for persistent data. And the subscriptionId can be found using the retrieve all reseller subscriptions method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}/changePlan",
	//   "request": {
	//     "$ref": "ChangePlanRequest"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.changeRenewalSettings":

type SubscriptionsChangeRenewalSettingsCall struct {
	s               *Service
	customerId      string
	subscriptionId  string
	renewalsettings *RenewalSettings
	urlParams_      gensupport.URLParams
	ctx_            context.Context
	header_         http.Header
}

// ChangeRenewalSettings: Update a user license's renewal settings. This
// is applicable for accounts with annual commitment plans only.
func (r *SubscriptionsService) ChangeRenewalSettings(customerId string, subscriptionId string, renewalsettings *RenewalSettings) *SubscriptionsChangeRenewalSettingsCall {
	c := &SubscriptionsChangeRenewalSettingsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.customerId = customerId
	c.subscriptionId = subscriptionId
	c.renewalsettings = renewalsettings
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsChangeRenewalSettingsCall) Fields(s ...googleapi.Field) *SubscriptionsChangeRenewalSettingsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SubscriptionsChangeRenewalSettingsCall) Context(ctx context.Context) *SubscriptionsChangeRenewalSettingsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SubscriptionsChangeRenewalSettingsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SubscriptionsChangeRenewalSettingsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.renewalsettings)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "customers/{customerId}/subscriptions/{subscriptionId}/changeRenewalSettings")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.subscriptions.changeRenewalSettings" call.
// Exactly one of *Subscription or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Subscription.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SubscriptionsChangeRenewalSettingsCall) Do(opts ...googleapi.CallOption) (*Subscription, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Subscription{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a user license's renewal settings. This is applicable for accounts with annual commitment plans only.",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.changeRenewalSettings",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Either the customer's primary domain name or the customer's unique identifier. If using the domain name, we do not recommend using a customerId as a key for persistent data. If the domain name for a customerId is changed, the Google system automatically updates.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "This is a required property. The subscriptionId is the subscription identifier and is unique for each customer. Since a subscriptionId changes when a subscription is updated, we recommend to not use this ID as a key for persistent data. And the subscriptionId can be found using the retrieve all reseller subscriptions method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}/changeRenewalSettings",
	//   "request": {
	//     "$ref": "RenewalSettings"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.changeSeats":

type SubscriptionsChangeSeatsCall struct {
	s              *Service
	customerId     string
	subscriptionId string
	seats          *Seats
	urlParams_     gensupport.URLParams
	ctx_           context.Context
	header_        http.Header
}

// ChangeSeats: Update a subscription's user license settings.
func (r *SubscriptionsService) ChangeSeats(customerId string, subscriptionId string, seats *Seats) *SubscriptionsChangeSeatsCall {
	c := &SubscriptionsChangeSeatsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.customerId = customerId
	c.subscriptionId = subscriptionId
	c.seats = seats
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsChangeSeatsCall) Fields(s ...googleapi.Field) *SubscriptionsChangeSeatsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SubscriptionsChangeSeatsCall) Context(ctx context.Context) *SubscriptionsChangeSeatsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SubscriptionsChangeSeatsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SubscriptionsChangeSeatsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.seats)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "customers/{customerId}/subscriptions/{subscriptionId}/changeSeats")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.subscriptions.changeSeats" call.
// Exactly one of *Subscription or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Subscription.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SubscriptionsChangeSeatsCall) Do(opts ...googleapi.CallOption) (*Subscription, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Subscription{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update a subscription's user license settings.",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.changeSeats",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Either the customer's primary domain name or the customer's unique identifier. If using the domain name, we do not recommend using a customerId as a key for persistent data. If the domain name for a customerId is changed, the Google system automatically updates.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "This is a required property. The subscriptionId is the subscription identifier and is unique for each customer. Since a subscriptionId changes when a subscription is updated, we recommend to not use this ID as a key for persistent data. And the subscriptionId can be found using the retrieve all reseller subscriptions method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}/changeSeats",
	//   "request": {
	//     "$ref": "Seats"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.delete":

type SubscriptionsDeleteCall struct {
	s              *Service
	customerId     string
	subscriptionId string
	urlParams_     gensupport.URLParams
	ctx_           context.Context
	header_        http.Header
}

// Delete: Cancel, suspend or transfer a subscription to direct.
func (r *SubscriptionsService) Delete(customerId string, subscriptionId string, deletionType string) *SubscriptionsDeleteCall {
	c := &SubscriptionsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.customerId = customerId
	c.subscriptionId = subscriptionId
	c.urlParams_.Set("deletionType", deletionType)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsDeleteCall) Fields(s ...googleapi.Field) *SubscriptionsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SubscriptionsDeleteCall) Context(ctx context.Context) *SubscriptionsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SubscriptionsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SubscriptionsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "customers/{customerId}/subscriptions/{subscriptionId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.subscriptions.delete" call.
func (c *SubscriptionsDeleteCall) Do(opts ...googleapi.CallOption) error {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if err != nil {
		return err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return err
	}
	return nil
	// {
	//   "description": "Cancel, suspend or transfer a subscription to direct.",
	//   "httpMethod": "DELETE",
	//   "id": "reseller.subscriptions.delete",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId",
	//     "deletionType"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Either the customer's primary domain name or the customer's unique identifier. If using the domain name, we do not recommend using a customerId as a key for persistent data. If the domain name for a customerId is changed, the Google system automatically updates.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "deletionType": {
	//       "description": "The deletionType query string enables the cancellation, downgrade, or suspension of a subscription.",
	//       "enum": [
	//         "cancel",
	//         "downgrade",
	//         "suspend",
	//         "transfer_to_direct"
	//       ],
	//       "enumDescriptions": [
	//         "Cancels the subscription immediately. This does not apply to a G Suite subscription.",
	//         "Downgrades a G Suite subscription to a Google Apps Free edition subscription only if the customer was initially subscribed to a Google Apps Free edition (also known as the Standard edition). Once downgraded, the customer no longer has access to the previous G Suite subscription and is no longer managed by the reseller.\n\nA G Suite subscription's downgrade cannot be invoked if an active or suspended Google Drive or Google Vault subscription is present. The Google Drive or Google Vault subscription must be cancelled before the G Suite subscription's downgrade is invoked.\n\nThe downgrade deletionType does not apply to other products or G Suite SKUs.",
	//         "(DEPRECATED) The G Suite account is suspended for four days and then cancelled. Once suspended, an administrator has access to the suspended account, but the account users can not access their services. A suspension can be lifted, using the reseller tools.\n\nA G Suite subscription's suspension can not be invoked if an active or suspended Google Drive or Google Vault subscription is present. The Google Drive or Google Vault subscription must be cancelled before the G Suite subscription's suspension is invoked.",
	//         "Transfers a subscription directly to Google.  The customer is immediately transferred to a direct billing relationship with Google and is given a short amount of time with no service interruption. The customer can then choose to set up billing directly with Google by using a credit card, or they can transfer to another reseller."
	//       ],
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "This is a required property. The subscriptionId is the subscription identifier and is unique for each customer. Since a subscriptionId changes when a subscription is updated, we recommend to not use this ID as a key for persistent data. And the subscriptionId can be found using the retrieve all reseller subscriptions method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}",
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.get":

type SubscriptionsGetCall struct {
	s              *Service
	customerId     string
	subscriptionId string
	urlParams_     gensupport.URLParams
	ifNoneMatch_   string
	ctx_           context.Context
	header_        http.Header
}

// Get: Get a specific subscription.
func (r *SubscriptionsService) Get(customerId string, subscriptionId string) *SubscriptionsGetCall {
	c := &SubscriptionsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.customerId = customerId
	c.subscriptionId = subscriptionId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsGetCall) Fields(s ...googleapi.Field) *SubscriptionsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SubscriptionsGetCall) IfNoneMatch(entityTag string) *SubscriptionsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SubscriptionsGetCall) Context(ctx context.Context) *SubscriptionsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SubscriptionsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SubscriptionsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "customers/{customerId}/subscriptions/{subscriptionId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.subscriptions.get" call.
// Exactly one of *Subscription or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Subscription.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SubscriptionsGetCall) Do(opts ...googleapi.CallOption) (*Subscription, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Subscription{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Get a specific subscription.",
	//   "httpMethod": "GET",
	//   "id": "reseller.subscriptions.get",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Either the customer's primary domain name or the customer's unique identifier. If using the domain name, we do not recommend using a customerId as a key for persistent data. If the domain name for a customerId is changed, the Google system automatically updates.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "This is a required property. The subscriptionId is the subscription identifier and is unique for each customer. Since a subscriptionId changes when a subscription is updated, we recommend to not use this ID as a key for persistent data. And the subscriptionId can be found using the retrieve all reseller subscriptions method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}",
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order",
	//     "https://www.googleapis.com/auth/apps.order.readonly"
	//   ]
	// }

}

// method id "reseller.subscriptions.insert":

type SubscriptionsInsertCall struct {
	s            *Service
	customerId   string
	subscription *Subscription
	urlParams_   gensupport.URLParams
	ctx_         context.Context
	header_      http.Header
}

// Insert: Create or transfer a subscription.
func (r *SubscriptionsService) Insert(customerId string, subscription *Subscription) *SubscriptionsInsertCall {
	c := &SubscriptionsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.customerId = customerId
	c.subscription = subscription
	return c
}

// CustomerAuthToken sets the optional parameter "customerAuthToken":
// The customerAuthToken query string is required when creating a resold
// account that transfers a direct customer's subscription or transfers
// another reseller customer's subscription to your reseller management.
// This is a hexadecimal authentication token needed to complete the
// subscription transfer. For more information, see the administrator
// help center.
func (c *SubscriptionsInsertCall) CustomerAuthToken(customerAuthToken string) *SubscriptionsInsertCall {
	c.urlParams_.Set("customerAuthToken", customerAuthToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsInsertCall) Fields(s ...googleapi.Field) *SubscriptionsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SubscriptionsInsertCall) Context(ctx context.Context) *SubscriptionsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SubscriptionsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SubscriptionsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.subscription)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "customers/{customerId}/subscriptions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"customerId": c.customerId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.subscriptions.insert" call.
// Exactly one of *Subscription or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Subscription.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SubscriptionsInsertCall) Do(opts ...googleapi.CallOption) (*Subscription, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Subscription{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Create or transfer a subscription.",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.insert",
	//   "parameterOrder": [
	//     "customerId"
	//   ],
	//   "parameters": {
	//     "customerAuthToken": {
	//       "description": "The customerAuthToken query string is required when creating a resold account that transfers a direct customer's subscription or transfers another reseller customer's subscription to your reseller management. This is a hexadecimal authentication token needed to complete the subscription transfer. For more information, see the administrator help center.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customerId": {
	//       "description": "Either the customer's primary domain name or the customer's unique identifier. If using the domain name, we do not recommend using a customerId as a key for persistent data. If the domain name for a customerId is changed, the Google system automatically updates.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions",
	//   "request": {
	//     "$ref": "Subscription"
	//   },
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.list":

type SubscriptionsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: List of subscriptions managed by the reseller. The list can be
// all subscriptions, all of a customer's subscriptions, or all of a
// customer's transferable subscriptions.
func (r *SubscriptionsService) List() *SubscriptionsListCall {
	c := &SubscriptionsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// CustomerAuthToken sets the optional parameter "customerAuthToken":
// The customerAuthToken query string is required when creating a resold
// account that transfers a direct customer's subscription or transfers
// another reseller customer's subscription to your reseller management.
// This is a hexadecimal authentication token needed to complete the
// subscription transfer. For more information, see the administrator
// help center.
func (c *SubscriptionsListCall) CustomerAuthToken(customerAuthToken string) *SubscriptionsListCall {
	c.urlParams_.Set("customerAuthToken", customerAuthToken)
	return c
}

// CustomerId sets the optional parameter "customerId": Either the
// customer's primary domain name or the customer's unique identifier.
// If using the domain name, we do not recommend using a customerId as a
// key for persistent data. If the domain name for a customerId is
// changed, the Google system automatically updates.
func (c *SubscriptionsListCall) CustomerId(customerId string) *SubscriptionsListCall {
	c.urlParams_.Set("customerId", customerId)
	return c
}

// CustomerNamePrefix sets the optional parameter "customerNamePrefix":
// When retrieving all of your subscriptions and filtering for specific
// customers, you can enter a prefix for a customer name. Using an
// example customer group that includes exam.com, example20.com and
// example.com:
// - exa -- Returns all customer names that start with 'exa' which could
// include exam.com, example20.com, and example.com. A name prefix is
// similar to using a regular expression's asterisk, exa*.
// - example -- Returns example20.com and example.com.
func (c *SubscriptionsListCall) CustomerNamePrefix(customerNamePrefix string) *SubscriptionsListCall {
	c.urlParams_.Set("customerNamePrefix", customerNamePrefix)
	return c
}

// MaxResults sets the optional parameter "maxResults": When retrieving
// a large list, the maxResults is the maximum number of results per
// page. The nextPageToken value takes you to the next page. The default
// is 20.
func (c *SubscriptionsListCall) MaxResults(maxResults int64) *SubscriptionsListCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to specify
// next page in the list
func (c *SubscriptionsListCall) PageToken(pageToken string) *SubscriptionsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsListCall) Fields(s ...googleapi.Field) *SubscriptionsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *SubscriptionsListCall) IfNoneMatch(entityTag string) *SubscriptionsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SubscriptionsListCall) Context(ctx context.Context) *SubscriptionsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SubscriptionsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SubscriptionsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "subscriptions")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.subscriptions.list" call.
// Exactly one of *Subscriptions or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Subscriptions.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *SubscriptionsListCall) Do(opts ...googleapi.CallOption) (*Subscriptions, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Subscriptions{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List of subscriptions managed by the reseller. The list can be all subscriptions, all of a customer's subscriptions, or all of a customer's transferable subscriptions.",
	//   "httpMethod": "GET",
	//   "id": "reseller.subscriptions.list",
	//   "parameters": {
	//     "customerAuthToken": {
	//       "description": "The customerAuthToken query string is required when creating a resold account that transfers a direct customer's subscription or transfers another reseller customer's subscription to your reseller management. This is a hexadecimal authentication token needed to complete the subscription transfer. For more information, see the administrator help center.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customerId": {
	//       "description": "Either the customer's primary domain name or the customer's unique identifier. If using the domain name, we do not recommend using a customerId as a key for persistent data. If the domain name for a customerId is changed, the Google system automatically updates.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "customerNamePrefix": {
	//       "description": "When retrieving all of your subscriptions and filtering for specific customers, you can enter a prefix for a customer name. Using an example customer group that includes exam.com, example20.com and example.com:  \n- exa -- Returns all customer names that start with 'exa' which could include exam.com, example20.com, and example.com. A name prefix is similar to using a regular expression's asterisk, exa*. \n- example -- Returns example20.com and example.com.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "When retrieving a large list, the maxResults is the maximum number of results per page. The nextPageToken value takes you to the next page. The default is 20.",
	//       "format": "uint32",
	//       "location": "query",
	//       "maximum": "100",
	//       "minimum": "1",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to specify next page in the list",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "subscriptions",
	//   "response": {
	//     "$ref": "Subscriptions"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order",
	//     "https://www.googleapis.com/auth/apps.order.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *SubscriptionsListCall) Pages(ctx context.Context, f func(*Subscriptions) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "reseller.subscriptions.startPaidService":

type SubscriptionsStartPaidServiceCall struct {
	s              *Service
	customerId     string
	subscriptionId string
	urlParams_     gensupport.URLParams
	ctx_           context.Context
	header_        http.Header
}

// StartPaidService: Immediately move a 30-day free trial subscription
// to a paid service subscription.
func (r *SubscriptionsService) StartPaidService(customerId string, subscriptionId string) *SubscriptionsStartPaidServiceCall {
	c := &SubscriptionsStartPaidServiceCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.customerId = customerId
	c.subscriptionId = subscriptionId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsStartPaidServiceCall) Fields(s ...googleapi.Field) *SubscriptionsStartPaidServiceCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SubscriptionsStartPaidServiceCall) Context(ctx context.Context) *SubscriptionsStartPaidServiceCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SubscriptionsStartPaidServiceCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SubscriptionsStartPaidServiceCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "customers/{customerId}/subscriptions/{subscriptionId}/startPaidService")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.subscriptions.startPaidService" call.
// Exactly one of *Subscription or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Subscription.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SubscriptionsStartPaidServiceCall) Do(opts ...googleapi.CallOption) (*Subscription, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Subscription{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Immediately move a 30-day free trial subscription to a paid service subscription.",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.startPaidService",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Either the customer's primary domain name or the customer's unique identifier. If using the domain name, we do not recommend using a customerId as a key for persistent data. If the domain name for a customerId is changed, the Google system automatically updates.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "This is a required property. The subscriptionId is the subscription identifier and is unique for each customer. Since a subscriptionId changes when a subscription is updated, we recommend to not use this ID as a key for persistent data. And the subscriptionId can be found using the retrieve all reseller subscriptions method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}/startPaidService",
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}

// method id "reseller.subscriptions.suspend":

type SubscriptionsSuspendCall struct {
	s              *Service
	customerId     string
	subscriptionId string
	urlParams_     gensupport.URLParams
	ctx_           context.Context
	header_        http.Header
}

// Suspend: Suspends an active subscription.
func (r *SubscriptionsService) Suspend(customerId string, subscriptionId string) *SubscriptionsSuspendCall {
	c := &SubscriptionsSuspendCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.customerId = customerId
	c.subscriptionId = subscriptionId
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SubscriptionsSuspendCall) Fields(s ...googleapi.Field) *SubscriptionsSuspendCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *SubscriptionsSuspendCall) Context(ctx context.Context) *SubscriptionsSuspendCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *SubscriptionsSuspendCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *SubscriptionsSuspendCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "customers/{customerId}/subscriptions/{subscriptionId}/suspend")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"customerId":     c.customerId,
		"subscriptionId": c.subscriptionId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "reseller.subscriptions.suspend" call.
// Exactly one of *Subscription or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *Subscription.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *SubscriptionsSuspendCall) Do(opts ...googleapi.CallOption) (*Subscription, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Subscription{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Suspends an active subscription.",
	//   "httpMethod": "POST",
	//   "id": "reseller.subscriptions.suspend",
	//   "parameterOrder": [
	//     "customerId",
	//     "subscriptionId"
	//   ],
	//   "parameters": {
	//     "customerId": {
	//       "description": "Either the customer's primary domain name or the customer's unique identifier. If using the domain name, we do not recommend using a customerId as a key for persistent data. If the domain name for a customerId is changed, the Google system automatically updates.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "subscriptionId": {
	//       "description": "This is a required property. The subscriptionId is the subscription identifier and is unique for each customer. Since a subscriptionId changes when a subscription is updated, we recommend to not use this ID as a key for persistent data. And the subscriptionId can be found using the retrieve all reseller subscriptions method.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "customers/{customerId}/subscriptions/{subscriptionId}/suspend",
	//   "response": {
	//     "$ref": "Subscription"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/apps.order"
	//   ]
	// }

}
