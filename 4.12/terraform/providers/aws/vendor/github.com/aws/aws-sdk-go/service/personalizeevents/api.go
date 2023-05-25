// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalizeevents

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

const opPutEvents = "PutEvents"

// PutEventsRequest generates a "aws/request.Request" representing the
// client's request for the PutEvents operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See PutEvents for more information on using the PutEvents
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the PutEventsRequest method.
//	req, resp := client.PutEventsRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/personalize-events-2018-03-22/PutEvents
func (c *PersonalizeEvents) PutEventsRequest(input *PutEventsInput) (req *request.Request, output *PutEventsOutput) {
	op := &request.Operation{
		Name:       opPutEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/events",
	}

	if input == nil {
		input = &PutEventsInput{}
	}

	output = &PutEventsOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(restjson.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// PutEvents API operation for Amazon Personalize Events.
//
// Records user interaction event data. For more information see Recording Events
// (https://docs.aws.amazon.com/personalize/latest/dg/recording-events.html).
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Personalize Events's
// API operation PutEvents for usage and error information.
//
// Returned Error Types:
//   - InvalidInputException
//     Provide a valid value for the field or parameter.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/personalize-events-2018-03-22/PutEvents
func (c *PersonalizeEvents) PutEvents(input *PutEventsInput) (*PutEventsOutput, error) {
	req, out := c.PutEventsRequest(input)
	return out, req.Send()
}

// PutEventsWithContext is the same as PutEvents with the addition of
// the ability to pass a context and additional request options.
//
// See PutEvents for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PersonalizeEvents) PutEventsWithContext(ctx aws.Context, input *PutEventsInput, opts ...request.Option) (*PutEventsOutput, error) {
	req, out := c.PutEventsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opPutItems = "PutItems"

// PutItemsRequest generates a "aws/request.Request" representing the
// client's request for the PutItems operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See PutItems for more information on using the PutItems
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the PutItemsRequest method.
//	req, resp := client.PutItemsRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/personalize-events-2018-03-22/PutItems
func (c *PersonalizeEvents) PutItemsRequest(input *PutItemsInput) (req *request.Request, output *PutItemsOutput) {
	op := &request.Operation{
		Name:       opPutItems,
		HTTPMethod: "POST",
		HTTPPath:   "/items",
	}

	if input == nil {
		input = &PutItemsInput{}
	}

	output = &PutItemsOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(restjson.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// PutItems API operation for Amazon Personalize Events.
//
// Adds one or more items to an Items dataset. For more information see Importing
// Items Incrementally (https://docs.aws.amazon.com/personalize/latest/dg/importing-items.html).
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Personalize Events's
// API operation PutItems for usage and error information.
//
// Returned Error Types:
//
//   - InvalidInputException
//     Provide a valid value for the field or parameter.
//
//   - ResourceNotFoundException
//     Could not find the specified resource.
//
//   - ResourceInUseException
//     The specified resource is in use.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/personalize-events-2018-03-22/PutItems
func (c *PersonalizeEvents) PutItems(input *PutItemsInput) (*PutItemsOutput, error) {
	req, out := c.PutItemsRequest(input)
	return out, req.Send()
}

// PutItemsWithContext is the same as PutItems with the addition of
// the ability to pass a context and additional request options.
//
// See PutItems for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PersonalizeEvents) PutItemsWithContext(ctx aws.Context, input *PutItemsInput, opts ...request.Option) (*PutItemsOutput, error) {
	req, out := c.PutItemsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opPutUsers = "PutUsers"

// PutUsersRequest generates a "aws/request.Request" representing the
// client's request for the PutUsers operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See PutUsers for more information on using the PutUsers
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the PutUsersRequest method.
//	req, resp := client.PutUsersRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/personalize-events-2018-03-22/PutUsers
func (c *PersonalizeEvents) PutUsersRequest(input *PutUsersInput) (req *request.Request, output *PutUsersOutput) {
	op := &request.Operation{
		Name:       opPutUsers,
		HTTPMethod: "POST",
		HTTPPath:   "/users",
	}

	if input == nil {
		input = &PutUsersInput{}
	}

	output = &PutUsersOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(restjson.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// PutUsers API operation for Amazon Personalize Events.
//
// Adds one or more users to a Users dataset. For more information see Importing
// Users Incrementally (https://docs.aws.amazon.com/personalize/latest/dg/importing-users.html).
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Personalize Events's
// API operation PutUsers for usage and error information.
//
// Returned Error Types:
//
//   - InvalidInputException
//     Provide a valid value for the field or parameter.
//
//   - ResourceNotFoundException
//     Could not find the specified resource.
//
//   - ResourceInUseException
//     The specified resource is in use.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/personalize-events-2018-03-22/PutUsers
func (c *PersonalizeEvents) PutUsers(input *PutUsersInput) (*PutUsersOutput, error) {
	req, out := c.PutUsersRequest(input)
	return out, req.Send()
}

// PutUsersWithContext is the same as PutUsers with the addition of
// the ability to pass a context and additional request options.
//
// See PutUsers for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PersonalizeEvents) PutUsersWithContext(ctx aws.Context, input *PutUsersInput, opts ...request.Option) (*PutUsersOutput, error) {
	req, out := c.PutUsersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// Represents user interaction event information sent using the PutEvents API.
type Event struct {
	_ struct{} `type:"structure"`

	// An ID associated with the event. If an event ID is not provided, Amazon Personalize
	// generates a unique ID for the event. An event ID is not used as an input
	// to the model. Amazon Personalize uses the event ID to distinquish unique
	// events. Any subsequent events after the first with the same event ID are
	// not used in model training.
	EventId *string `locationName:"eventId" min:"1" type:"string"`

	// The type of event, such as click or download. This property corresponds to
	// the EVENT_TYPE field of your Interactions schema and depends on the types
	// of events you are tracking.
	//
	// EventType is a required field
	EventType *string `locationName:"eventType" min:"1" type:"string" required:"true"`

	// The event value that corresponds to the EVENT_VALUE field of the Interactions
	// schema.
	EventValue *float64 `locationName:"eventValue" type:"float"`

	// A list of item IDs that represents the sequence of items you have shown the
	// user. For example, ["itemId1", "itemId2", "itemId3"].
	Impression []*string `locationName:"impression" min:"1" type:"list"`

	// The item ID key that corresponds to the ITEM_ID field of the Interactions
	// schema.
	ItemId *string `locationName:"itemId" min:"1" type:"string"`

	// A string map of event-specific data that you might choose to record. For
	// example, if a user rates a movie on your site, other than movie ID (itemId)
	// and rating (eventValue) , you might also send the number of movie ratings
	// made by the user.
	//
	// Each item in the map consists of a key-value pair. For example,
	//
	// {"numberOfRatings": "12"}
	//
	// The keys use camel case names that match the fields in the Interactions schema.
	// In the above example, the numberOfRatings would match the 'NUMBER_OF_RATINGS'
	// field defined in the Interactions schema.
	Properties aws.JSONValue `locationName:"properties" type:"jsonvalue"`

	// The ID of the recommendation.
	RecommendationId *string `locationName:"recommendationId" min:"1" type:"string"`

	// The timestamp (in Unix time) on the client side when the event occurred.
	//
	// SentAt is a required field
	SentAt *time.Time `locationName:"sentAt" type:"timestamp" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Event) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Event) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Event) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Event"}
	if s.EventId != nil && len(*s.EventId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("EventId", 1))
	}
	if s.EventType == nil {
		invalidParams.Add(request.NewErrParamRequired("EventType"))
	}
	if s.EventType != nil && len(*s.EventType) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("EventType", 1))
	}
	if s.Impression != nil && len(s.Impression) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Impression", 1))
	}
	if s.ItemId != nil && len(*s.ItemId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ItemId", 1))
	}
	if s.RecommendationId != nil && len(*s.RecommendationId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("RecommendationId", 1))
	}
	if s.SentAt == nil {
		invalidParams.Add(request.NewErrParamRequired("SentAt"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEventId sets the EventId field's value.
func (s *Event) SetEventId(v string) *Event {
	s.EventId = &v
	return s
}

// SetEventType sets the EventType field's value.
func (s *Event) SetEventType(v string) *Event {
	s.EventType = &v
	return s
}

// SetEventValue sets the EventValue field's value.
func (s *Event) SetEventValue(v float64) *Event {
	s.EventValue = &v
	return s
}

// SetImpression sets the Impression field's value.
func (s *Event) SetImpression(v []*string) *Event {
	s.Impression = v
	return s
}

// SetItemId sets the ItemId field's value.
func (s *Event) SetItemId(v string) *Event {
	s.ItemId = &v
	return s
}

// SetProperties sets the Properties field's value.
func (s *Event) SetProperties(v aws.JSONValue) *Event {
	s.Properties = v
	return s
}

// SetRecommendationId sets the RecommendationId field's value.
func (s *Event) SetRecommendationId(v string) *Event {
	s.RecommendationId = &v
	return s
}

// SetSentAt sets the SentAt field's value.
func (s *Event) SetSentAt(v time.Time) *Event {
	s.SentAt = &v
	return s
}

// Provide a valid value for the field or parameter.
type InvalidInputException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidInputException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s InvalidInputException) GoString() string {
	return s.String()
}

func newErrorInvalidInputException(v protocol.ResponseMetadata) error {
	return &InvalidInputException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InvalidInputException) Code() string {
	return "InvalidInputException"
}

// Message returns the exception's message.
func (s *InvalidInputException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InvalidInputException) OrigErr() error {
	return nil
}

func (s *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InvalidInputException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InvalidInputException) RequestID() string {
	return s.RespMetadata.RequestID
}

// Represents item metadata added to an Items dataset using the PutItems API.
// For more information see Importing Items Incrementally (https://docs.aws.amazon.com/personalize/latest/dg/importing-items.html).
type Item struct {
	_ struct{} `type:"structure"`

	// The ID associated with the item.
	//
	// ItemId is a required field
	ItemId *string `locationName:"itemId" min:"1" type:"string" required:"true"`

	// A string map of item-specific metadata. Each element in the map consists
	// of a key-value pair. For example, {"numberOfRatings": "12"}.
	//
	// The keys use camel case names that match the fields in the schema for the
	// Items dataset. In the previous example, the numberOfRatings matches the 'NUMBER_OF_RATINGS'
	// field defined in the Items schema. For categorical string data, to include
	// multiple categories for a single item, separate each category with a pipe
	// separator (|). For example, \"Horror|Action\".
	Properties aws.JSONValue `locationName:"properties" type:"jsonvalue"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Item) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s Item) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Item) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Item"}
	if s.ItemId == nil {
		invalidParams.Add(request.NewErrParamRequired("ItemId"))
	}
	if s.ItemId != nil && len(*s.ItemId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ItemId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetItemId sets the ItemId field's value.
func (s *Item) SetItemId(v string) *Item {
	s.ItemId = &v
	return s
}

// SetProperties sets the Properties field's value.
func (s *Item) SetProperties(v aws.JSONValue) *Item {
	s.Properties = v
	return s
}

type PutEventsInput struct {
	_ struct{} `type:"structure"`

	// A list of event data from the session.
	//
	// EventList is a required field
	EventList []*Event `locationName:"eventList" min:"1" type:"list" required:"true"`

	// The session ID associated with the user's visit. Your application generates
	// the sessionId when a user first visits your website or uses your application.
	// Amazon Personalize uses the sessionId to associate events with the user before
	// they log in. For more information, see Recording Events (https://docs.aws.amazon.com/personalize/latest/dg/recording-events.html).
	//
	// SessionId is a required field
	SessionId *string `locationName:"sessionId" min:"1" type:"string" required:"true"`

	// The tracking ID for the event. The ID is generated by a call to the CreateEventTracker
	// (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateEventTracker.html)
	// API.
	//
	// TrackingId is a required field
	TrackingId *string `locationName:"trackingId" min:"1" type:"string" required:"true"`

	// The user associated with the event.
	UserId *string `locationName:"userId" min:"1" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutEventsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutEventsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutEventsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutEventsInput"}
	if s.EventList == nil {
		invalidParams.Add(request.NewErrParamRequired("EventList"))
	}
	if s.EventList != nil && len(s.EventList) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("EventList", 1))
	}
	if s.SessionId == nil {
		invalidParams.Add(request.NewErrParamRequired("SessionId"))
	}
	if s.SessionId != nil && len(*s.SessionId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("SessionId", 1))
	}
	if s.TrackingId == nil {
		invalidParams.Add(request.NewErrParamRequired("TrackingId"))
	}
	if s.TrackingId != nil && len(*s.TrackingId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("TrackingId", 1))
	}
	if s.UserId != nil && len(*s.UserId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("UserId", 1))
	}
	if s.EventList != nil {
		for i, v := range s.EventList {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "EventList", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEventList sets the EventList field's value.
func (s *PutEventsInput) SetEventList(v []*Event) *PutEventsInput {
	s.EventList = v
	return s
}

// SetSessionId sets the SessionId field's value.
func (s *PutEventsInput) SetSessionId(v string) *PutEventsInput {
	s.SessionId = &v
	return s
}

// SetTrackingId sets the TrackingId field's value.
func (s *PutEventsInput) SetTrackingId(v string) *PutEventsInput {
	s.TrackingId = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *PutEventsInput) SetUserId(v string) *PutEventsInput {
	s.UserId = &v
	return s
}

type PutEventsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutEventsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutEventsOutput) GoString() string {
	return s.String()
}

type PutItemsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the Items dataset you are adding the item
	// or items to.
	//
	// DatasetArn is a required field
	DatasetArn *string `locationName:"datasetArn" type:"string" required:"true"`

	// A list of item data.
	//
	// Items is a required field
	Items []*Item `locationName:"items" min:"1" type:"list" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutItemsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutItemsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutItemsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutItemsInput"}
	if s.DatasetArn == nil {
		invalidParams.Add(request.NewErrParamRequired("DatasetArn"))
	}
	if s.Items == nil {
		invalidParams.Add(request.NewErrParamRequired("Items"))
	}
	if s.Items != nil && len(s.Items) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Items", 1))
	}
	if s.Items != nil {
		for i, v := range s.Items {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Items", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDatasetArn sets the DatasetArn field's value.
func (s *PutItemsInput) SetDatasetArn(v string) *PutItemsInput {
	s.DatasetArn = &v
	return s
}

// SetItems sets the Items field's value.
func (s *PutItemsInput) SetItems(v []*Item) *PutItemsInput {
	s.Items = v
	return s
}

type PutItemsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutItemsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutItemsOutput) GoString() string {
	return s.String()
}

type PutUsersInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the Users dataset you are adding the user
	// or users to.
	//
	// DatasetArn is a required field
	DatasetArn *string `locationName:"datasetArn" type:"string" required:"true"`

	// A list of user data.
	//
	// Users is a required field
	Users []*User `locationName:"users" min:"1" type:"list" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutUsersInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutUsersInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutUsersInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutUsersInput"}
	if s.DatasetArn == nil {
		invalidParams.Add(request.NewErrParamRequired("DatasetArn"))
	}
	if s.Users == nil {
		invalidParams.Add(request.NewErrParamRequired("Users"))
	}
	if s.Users != nil && len(s.Users) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Users", 1))
	}
	if s.Users != nil {
		for i, v := range s.Users {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Users", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDatasetArn sets the DatasetArn field's value.
func (s *PutUsersInput) SetDatasetArn(v string) *PutUsersInput {
	s.DatasetArn = &v
	return s
}

// SetUsers sets the Users field's value.
func (s *PutUsersInput) SetUsers(v []*User) *PutUsersInput {
	s.Users = v
	return s
}

type PutUsersOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutUsersOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s PutUsersOutput) GoString() string {
	return s.String()
}

// The specified resource is in use.
type ResourceInUseException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ResourceInUseException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ResourceInUseException) GoString() string {
	return s.String()
}

func newErrorResourceInUseException(v protocol.ResponseMetadata) error {
	return &ResourceInUseException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ResourceInUseException) Code() string {
	return "ResourceInUseException"
}

// Message returns the exception's message.
func (s *ResourceInUseException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ResourceInUseException) OrigErr() error {
	return nil
}

func (s *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ResourceInUseException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ResourceInUseException) RequestID() string {
	return s.RespMetadata.RequestID
}

// Could not find the specified resource.
type ResourceNotFoundException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ResourceNotFoundException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s ResourceNotFoundException) GoString() string {
	return s.String()
}

func newErrorResourceNotFoundException(v protocol.ResponseMetadata) error {
	return &ResourceNotFoundException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ResourceNotFoundException) Code() string {
	return "ResourceNotFoundException"
}

// Message returns the exception's message.
func (s *ResourceNotFoundException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ResourceNotFoundException) OrigErr() error {
	return nil
}

func (s *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ResourceNotFoundException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ResourceNotFoundException) RequestID() string {
	return s.RespMetadata.RequestID
}

// Represents user metadata added to a Users dataset using the PutUsers API.
// For more information see Importing Users Incrementally (https://docs.aws.amazon.com/personalize/latest/dg/importing-users.html).
type User struct {
	_ struct{} `type:"structure"`

	// A string map of user-specific metadata. Each element in the map consists
	// of a key-value pair. For example, {"numberOfVideosWatched": "45"}.
	//
	// The keys use camel case names that match the fields in the schema for the
	// Users dataset. In the previous example, the numberOfVideosWatched matches
	// the 'NUMBER_OF_VIDEOS_WATCHED' field defined in the Users schema. For categorical
	// string data, to include multiple categories for a single user, separate each
	// category with a pipe separator (|). For example, \"Member|Frequent shopper\".
	Properties aws.JSONValue `locationName:"properties" type:"jsonvalue"`

	// The ID associated with the user.
	//
	// UserId is a required field
	UserId *string `locationName:"userId" min:"1" type:"string" required:"true"`
}

// String returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s User) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation.
//
// API parameter values that are decorated as "sensitive" in the API will not
// be included in the string output. The member name will be present, but the
// value will be replaced with "sensitive".
func (s User) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *User) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "User"}
	if s.UserId == nil {
		invalidParams.Add(request.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("UserId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetProperties sets the Properties field's value.
func (s *User) SetProperties(v aws.JSONValue) *User {
	s.Properties = v
	return s
}

// SetUserId sets the UserId field's value.
func (s *User) SetUserId(v string) *User {
	s.UserId = &v
	return s
}
