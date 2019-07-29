package event

import (
	"context"
	"encoding/xml"
	"github.com/videonext/onvif/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type GetServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl GetServiceCapabilitiesResponse"`

	// The capabilities for the event service is returned in the Capabilities element.
	Capabilities Capabilities `xml:"http://www.onvif.org/ver10/events/wsdl Capabilities,omitempty"`
}

type CreatePullPointSubscription struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl CreatePullPointSubscription"`

	// Optional XPATH expression to select specific topics.
	Filter FilterType `xml:"Filter,omitempty"`

	// Initial termination time.
	InitialTerminationTime AbsoluteOrRelativeTimeType `xml:"InitialTerminationTime,omitempty"`

	SubscriptionPolicy struct {
	} `xml:"SubscriptionPolicy,omitempty"`
}

type CreatePullPointSubscriptionResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl CreatePullPointSubscriptionResponse"`

	// Endpoint reference of the subscription to be used for pulling the messages.
	SubscriptionReference EndpointReferenceType `xml:"SubscriptionReference,omitempty"`

	CurrentTime CurrentTime `xml:"CurrentTime,omitempty"`

	TerminationTime TerminationTime `xml:"TerminationTime,omitempty"`
}

type PullMessages struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl PullMessages"`

	// Maximum time to block until this method returns.
	Timeout Duration `xml:"http://www.onvif.org/ver10/schema Timeout,omitempty"`

	// Upper limit for the number of messages to return at once. A server implementation may decide to return less messages.
	MessageLimit int32 `xml:"http://www.onvif.org/ver10/schema MessageLimit,omitempty"`
}

type PullMessagesResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl PullMessagesResponse"`

	// The date and time when the messages have been delivered by the web server to the client.
	CurrentTime string `xml:"http://www.onvif.org/ver10/schema CurrentTime,omitempty"`

	// Date time when the PullPoint will be shut down without further pull requests.
	TerminationTime string `xml:"http://www.onvif.org/ver10/schema TerminationTime,omitempty"`

	NotificationMessage []NotificationMessage `xml:"NotificationMessage,omitempty"`
}

type PullMessagesFaultResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl PullMessagesFaultResponse"`

	// Maximum timeout supported by the device.
	MaxTimeout Duration `xml:"http://www.onvif.org/ver10/schema MaxTimeout,omitempty"`

	// Maximum message limit supported by the device.
	MaxMessageLimit int32 `xml:"http://www.onvif.org/ver10/schema MaxMessageLimit,omitempty"`
}

type Seek struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl Seek"`

	// The date and time to match against stored messages.
	UtcTime string `xml:"http://www.onvif.org/ver10/schema UtcTime,omitempty"`

	// Reverse the pull direction of PullMessages.
	Reverse bool `xml:"http://www.onvif.org/ver10/events/wsdl Reverse,omitempty"`
}

type SeekResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl SeekResponse"`
}

type SetSynchronizationPoint struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl SetSynchronizationPoint"`
}

type SetSynchronizationPointResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl SetSynchronizationPointResponse"`
}

type GetEventProperties struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl GetEventProperties"`
}

type GetEventPropertiesResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/events/wsdl GetEventPropertiesResponse"`

	// List of topic namespaces supported.
	TopicNamespaceLocation AnyURI `xml:"http://www.onvif.org/ver10/schema TopicNamespaceLocation,omitempty"`

	FixedTopicSet FixedTopicSet `xml:"FixedTopicSet,omitempty"`

	TopicSet TopicSet `xml:"TopicSet,omitempty"`

	TopicExpressionDialect []TopicExpressionDialect `xml:"TopicExpressionDialect,omitempty"`

	//
	// Defines the XPath function set supported for message content filtering.
	// The following MessageContentFilterDialects should be returned if a device supports the message content filtering:
	//
	// A device that does not support any MessageContentFilterDialect returns a single empty url.
	//
	MessageContentFilterDialect AnyURI `xml:"http://www.onvif.org/ver10/schema MessageContentFilterDialect,omitempty"`

	//
	// Optional ProducerPropertiesDialects. Refer to  for advanced filtering.
	ProducerPropertiesFilterDialect AnyURI `xml:"http://www.onvif.org/ver10/schema ProducerPropertiesFilterDialect,omitempty"`

	//
	// The Message Content Description Language allows referencing
	// of vendor-specific types. In order to ease the integration of such types into a client application,
	// the GetEventPropertiesResponse shall list all URI locations to schema files whose types are
	// used in the description of notifications, with MessageContentSchemaLocation elements.
	// This list shall at least contain the URI of the ONVIF schema file.
	MessageContentSchemaLocation AnyURI `xml:"http://www.onvif.org/ver10/schema MessageContentSchemaLocation,omitempty"`
}

type Capabilities struct {

	// Indicates that the WS Subscription policy is supported.

	WSSubscriptionPolicySupport bool `xml:"http://www.onvif.org/ver10/events/wsdl WSSubscriptionPolicySupport,attr,omitempty"`

	// Indicates that the WS Pull Point is supported.

	WSPullPointSupport bool `xml:"http://www.onvif.org/ver10/events/wsdl WSPullPointSupport,attr,omitempty"`

	// Indicates that the WS Pausable Subscription Manager Interface is supported.

	WSPausableSubscriptionManagerInterfaceSupport bool `xml:"http://www.onvif.org/ver10/events/wsdl WSPausableSubscriptionManagerInterfaceSupport,attr,omitempty"`

	// Maximum number of supported notification producers as defined by WS-BaseNotification.

	MaxNotificationProducers int32 `xml:"http://www.onvif.org/ver10/schema MaxNotificationProducers,attr,omitempty"`

	// Maximum supported number of notification pull points.

	MaxPullPoints int32 `xml:"http://www.onvif.org/ver10/schema MaxPullPoints,attr,omitempty"`

	// Indication if the device supports persistent notification storage.

	PersistentNotificationStorage bool `xml:"http://www.onvif.org/ver10/events/wsdl PersistentNotificationStorage,attr,omitempty"`
}

type RelationshipTypeOpenEnum string

type RelationshipType AnyURI

const (
	RelationshipTypeHttpwwww3org200508addressingreply RelationshipType = "http://www.w3.org/2005/08/addressing/reply"
)

type FaultCodesOpenEnumType string

type FaultCodesType QName

const (
	FaultCodesTypeTnsInvalidAddressingHeader FaultCodesType = "tns:InvalidAddressingHeader"

	FaultCodesTypeTnsInvalidAddress FaultCodesType = "tns:InvalidAddress"

	FaultCodesTypeTnsInvalidEPR FaultCodesType = "tns:InvalidEPR"

	FaultCodesTypeTnsInvalidCardinality FaultCodesType = "tns:InvalidCardinality"

	FaultCodesTypeTnsMissingAddressInEPR FaultCodesType = "tns:MissingAddressInEPR"

	FaultCodesTypeTnsDuplicateMessageID FaultCodesType = "tns:DuplicateMessageID"

	FaultCodesTypeTnsActionMismatch FaultCodesType = "tns:ActionMismatch"

	FaultCodesTypeTnsMessageAddressingHeaderRequired FaultCodesType = "tns:MessageAddressingHeaderRequired"

	FaultCodesTypeTnsDestinationUnreachable FaultCodesType = "tns:DestinationUnreachable"

	FaultCodesTypeTnsActionNotSupported FaultCodesType = "tns:ActionNotSupported"

	FaultCodesTypeTnsEndpointUnavailable FaultCodesType = "tns:EndpointUnavailable"
)

type EndpointReference EndpointReferenceType

type Metadata MetadataType

type MessageID AttributedURIType

type RelatesTo RelatesToType

type ReplyTo EndpointReferenceType

type From EndpointReferenceType

type FaultTo EndpointReferenceType

type To AttributedURIType

type Action AttributedURIType

type RetryAfter AttributedUnsignedLongType

type ProblemHeaderQName AttributedQNameType

type ProblemHeader AttributedAnyType

type ProblemIRI AttributedURIType

type ProblemAction ProblemActionType

type EndpointReferenceType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing EndpointReference"`

	Address AttributedURIType `xml:"Address,omitempty"`

	ReferenceParameters ReferenceParametersType `xml:"ReferenceParameters,omitempty"`

	Metadata Metadata `xml:"Metadata,omitempty"`
}

type ReferenceParametersType struct {
}

type MetadataType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing Metadata"`
}

type RelatesToType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing RelatesTo"`

	Value AnyURI

	RelationshipType RelationshipTypeOpenEnum `xml:"RelationshipType,attr,omitempty"`
}

type AttributedURIType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing MessageID"`

	Value AnyURI
}

type AttributedUnsignedLongType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing RetryAfter"`

	Value uint64
}

type AttributedQNameType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing ProblemHeaderQName"`

	Value QName
}

type AttributedAnyType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing ProblemHeader"`
}

type ProblemActionType struct {
	XMLName xml.Name `xml:"http://www.w3.org/2005/08/addressing ProblemAction"`

	Action Action `xml:"Action,omitempty"`

	SoapAction AnyURI `xml:"http://www.onvif.org/ver10/schema SoapAction,omitempty"`
}

type FullTopicExpression string

type ConcreteTopicExpression string

type SimpleTopicExpression QName

type TopicNamespace TopicNamespaceType

type TopicSet TopicSetType

type Documentation struct {
}

type ExtensibleDocumented struct {
	Documentation Documentation `xml:"documentation,omitempty"`
}

type QueryExpressionType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/t-1 ProducerProperties"`

	Dialect AnyURI `xml:"http://www.onvif.org/ver10/schema Dialect,attr,omitempty"`
}

type TopicNamespaceType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/t-1 TopicNamespace"`

	*ExtensibleDocumented

	Topic []struct {
		*TopicType

		Parent ConcreteTopicExpression `xml:"parent,attr,omitempty"`
	} `xml:"Topic,omitempty"`

	Name NCName `xml:"name,attr,omitempty"`

	TargetNamespace AnyURI `xml:"targetNamespace,attr,omitempty"`

	Final bool `xml:"final,attr,omitempty"`
}

type TopicType struct {
	*ExtensibleDocumented

	MessagePattern QueryExpressionType `xml:"MessagePattern,omitempty"`

	Topic []TopicType `xml:"Topic,omitempty"`

	Name NCName `xml:"name,attr,omitempty"`

	MessageTypes string `xml:"messageTypes,attr,omitempty"`

	Final bool `xml:"final,attr,omitempty"`
}

type TopicSetType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/t-1 TopicSet"`

	*ExtensibleDocumented
}

type BaseFault BaseFaultType

type BaseFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsrf/bf-2 BaseFault"`

	Timestamp string `xml:"http://www.onvif.org/ver10/schema Timestamp,omitempty"`

	Originator EndpointReferenceType `xml:"Originator,omitempty"`

	ErrorCode struct {
		Dialect AnyURI `xml:"dialect,attr,omitempty"`
	} `xml:"ErrorCode,omitempty"`

	Description []struct {
		Value string

		string `xml:",attr,omitempty"`
	} `xml:"Description,omitempty"`

	FaultCause struct {
	} `xml:"FaultCause,omitempty"`
}

type AbsoluteOrRelativeTimeType string

type TopicExpression TopicExpressionType

type FixedTopicSet bool

type TopicExpressionDialect AnyURI

type NotificationProducerRP struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 NotificationProducerRP"`

	TopicExpression []TopicExpression `xml:"TopicExpression,omitempty"`

	FixedTopicSet FixedTopicSet `xml:"FixedTopicSet,omitempty"`

	TopicExpressionDialect []TopicExpressionDialect `xml:"TopicExpressionDialect,omitempty"`

	TopicSet TopicSet `xml:"TopicSet,omitempty"`
}

type ConsumerReference EndpointReferenceType

type Filter FilterType

type SubscriptionPolicy SubscriptionPolicyType

type CreationTime time.Time

type SubscriptionManagerRP struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 SubscriptionManagerRP"`

	ConsumerReference ConsumerReference `xml:"ConsumerReference,omitempty"`

	Filter Filter `xml:"Filter,omitempty"`

	SubscriptionPolicy SubscriptionPolicy `xml:"SubscriptionPolicy,omitempty"`

	CreationTime CreationTime `xml:"CreationTime,omitempty"`
}

type SubscriptionReference EndpointReferenceType

type Topic TopicExpressionType

type ProducerReference EndpointReferenceType

type NotificationMessage NotificationMessageHolderType

type Notify struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Notify"`

	NotificationMessage []NotificationMessage `xml:"NotificationMessage,omitempty"`
}

type CurrentTime time.Time

type TerminationTime time.Time

type ProducerProperties QueryExpressionType

type MessageContent QueryExpressionType

type UseRaw struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UseRaw"`
}

type Subscribe struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Subscribe"`

	ConsumerReference EndpointReferenceType `xml:"ConsumerReference,omitempty"`

	Filter FilterType `xml:"Filter,omitempty"`

	InitialTerminationTime AbsoluteOrRelativeTimeType `xml:"InitialTerminationTime,omitempty"`

	SubscriptionPolicy struct {
	} `xml:"SubscriptionPolicy,omitempty"`
}

type SubscribeResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 SubscribeResponse"`

	SubscriptionReference EndpointReferenceType `xml:"SubscriptionReference,omitempty"`

	CurrentTime CurrentTime `xml:"CurrentTime,omitempty"`

	TerminationTime TerminationTime `xml:"TerminationTime,omitempty"`
}

type GetCurrentMessage struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 GetCurrentMessage"`

	Topic TopicExpressionType `xml:"Topic,omitempty"`
}

type GetCurrentMessageResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 GetCurrentMessageResponse"`
}

type SubscribeCreationFailedFault SubscribeCreationFailedFaultType

type InvalidFilterFault InvalidFilterFaultType

type TopicExpressionDialectUnknownFault TopicExpressionDialectUnknownFaultType

type InvalidTopicExpressionFault InvalidTopicExpressionFaultType

type TopicNotSupportedFault TopicNotSupportedFaultType

type MultipleTopicsSpecifiedFault MultipleTopicsSpecifiedFaultType

type InvalidProducerPropertiesExpressionFault InvalidProducerPropertiesExpressionFaultType

type InvalidMessageContentExpressionFault InvalidMessageContentExpressionFaultType

type UnrecognizedPolicyRequestFault UnrecognizedPolicyRequestFaultType

type UnsupportedPolicyRequestFault UnsupportedPolicyRequestFaultType

type NotifyMessageNotSupportedFault NotifyMessageNotSupportedFaultType

type UnacceptableInitialTerminationTimeFault UnacceptableInitialTerminationTimeFaultType

type NoCurrentMessageOnTopicFault NoCurrentMessageOnTopicFaultType

type GetMessages struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 GetMessages"`

	MaximumNumber NonNegativeInteger `xml:"http://www.onvif.org/ver10/schema MaximumNumber,omitempty"`
}

type GetMessagesResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 GetMessagesResponse"`

	NotificationMessage []NotificationMessage `xml:"NotificationMessage,omitempty"`
}

type DestroyPullPoint struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 DestroyPullPoint"`
}

type DestroyPullPointResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 DestroyPullPointResponse"`
}

type UnableToGetMessagesFault UnableToGetMessagesFaultType

type UnableToDestroyPullPointFault UnableToDestroyPullPointFaultType

type CreatePullPoint struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 CreatePullPoint"`
}

type CreatePullPointResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 CreatePullPointResponse"`

	PullPoint EndpointReferenceType `xml:"PullPoint,omitempty"`
}

type UnableToCreatePullPointFault UnableToCreatePullPointFaultType

type Renew struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Renew"`

	TerminationTime AbsoluteOrRelativeTimeType `xml:"TerminationTime,omitempty"`
}

type RenewResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 RenewResponse"`

	TerminationTime TerminationTime `xml:"TerminationTime,omitempty"`

	CurrentTime CurrentTime `xml:"CurrentTime,omitempty"`
}

type UnacceptableTerminationTimeFault UnacceptableTerminationTimeFaultType

type Unsubscribe struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Unsubscribe"`
}

type UnsubscribeResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnsubscribeResponse"`
}

type UnableToDestroySubscriptionFault UnableToDestroySubscriptionFaultType

type PauseSubscription struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 PauseSubscription"`
}

type PauseSubscriptionResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 PauseSubscriptionResponse"`
}

type ResumeSubscription struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 ResumeSubscription"`
}

type ResumeSubscriptionResponse struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 ResumeSubscriptionResponse"`
}

type PauseFailedFault PauseFailedFaultType

type ResumeFailedFault ResumeFailedFaultType

// Removed QueryExpressionType

type TopicExpressionType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 TopicExpression"`

	Dialect AnyURI `xml:"http://www.onvif.org/ver10/schema Dialect,attr,omitempty"`
}

type FilterType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 Filter"`
}

type SubscriptionPolicyType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 SubscriptionPolicy"`
}

type NotificationMessageHolderType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 NotificationMessage"`

	SubscriptionReference SubscriptionReference `xml:"SubscriptionReference,omitempty"`

	Topic Topic `xml:"Topic,omitempty"`

	ProducerReference ProducerReference `xml:"ProducerReference,omitempty"`

	Message struct {
	} `xml:"Message,omitempty"`
}

type SubscribeCreationFailedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 SubscribeCreationFailedFault"`

	*BaseFaultType
}

type InvalidFilterFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 InvalidFilterFault"`

	*BaseFaultType

	UnknownFilter QName `xml:"http://www.onvif.org/ver10/schema UnknownFilter,omitempty"`
}

type TopicExpressionDialectUnknownFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 TopicExpressionDialectUnknownFault"`

	*BaseFaultType
}

type InvalidTopicExpressionFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 InvalidTopicExpressionFault"`

	*BaseFaultType
}

type TopicNotSupportedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 TopicNotSupportedFault"`

	*BaseFaultType
}

type MultipleTopicsSpecifiedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 MultipleTopicsSpecifiedFault"`

	*BaseFaultType
}

type InvalidProducerPropertiesExpressionFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 InvalidProducerPropertiesExpressionFault"`

	*BaseFaultType
}

type InvalidMessageContentExpressionFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 InvalidMessageContentExpressionFault"`

	*BaseFaultType
}

type UnrecognizedPolicyRequestFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnrecognizedPolicyRequestFault"`

	*BaseFaultType

	UnrecognizedPolicy QName `xml:"http://www.onvif.org/ver10/schema UnrecognizedPolicy,omitempty"`
}

type UnsupportedPolicyRequestFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnsupportedPolicyRequestFault"`

	*BaseFaultType

	UnsupportedPolicy QName `xml:"http://www.onvif.org/ver10/schema UnsupportedPolicy,omitempty"`
}

type NotifyMessageNotSupportedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 NotifyMessageNotSupportedFault"`

	*BaseFaultType
}

type UnacceptableInitialTerminationTimeFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnacceptableInitialTerminationTimeFault"`

	*BaseFaultType

	MinimumTime string `xml:"http://www.onvif.org/ver10/schema MinimumTime,omitempty"`

	MaximumTime string `xml:"http://www.onvif.org/ver10/schema MaximumTime,omitempty"`
}

type NoCurrentMessageOnTopicFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 NoCurrentMessageOnTopicFault"`

	*BaseFaultType
}

type UnableToGetMessagesFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnableToGetMessagesFault"`

	*BaseFaultType
}

type UnableToDestroyPullPointFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnableToDestroyPullPointFault"`

	*BaseFaultType
}

type UnableToCreatePullPointFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnableToCreatePullPointFault"`

	*BaseFaultType
}

type UnacceptableTerminationTimeFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnacceptableTerminationTimeFault"`

	*BaseFaultType

	MinimumTime string `xml:"http://www.onvif.org/ver10/schema MinimumTime,omitempty"`

	MaximumTime string `xml:"http://www.onvif.org/ver10/schema MaximumTime,omitempty"`
}

type UnableToDestroySubscriptionFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 UnableToDestroySubscriptionFault"`

	*BaseFaultType
}

type PauseFailedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 PauseFailedFault"`

	*BaseFaultType
}

type ResumeFailedFaultType struct {
	XMLName xml.Name `xml:"http://docs.oasis-open.org/wsn/b-2 ResumeFailedFault"`

	*BaseFaultType
}

type EventPortType interface {

	/* Returns the capabilities of the event service. The result is returned in a typed answer. */
	GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	// Error can be either of the following types:
	//
	//   - ResourceUnknownFault
	//   - InvalidFilterFault
	//   - TopicExpressionDialectUnknownFault
	//   - InvalidTopicExpressionFault
	//   - TopicNotSupportedFault
	//   - InvalidProducerPropertiesExpressionFault
	//   - InvalidMessageContentExpressionFault
	//   - UnacceptableInitialTerminationTimeFault
	//   - UnrecognizedPolicyRequestFault
	//   - UnsupportedPolicyRequestFault
	//   - NotifyMessageNotSupportedFault
	//   - SubscribeCreationFailedFault
	/* This method returns a PullPointSubscription that can be polled using PullMessages.
	This message contains the same elements as the SubscriptionRequest of the WS-BaseNotification without the ConsumerReference.
	If no Filter is specified the pullpoint notifies all occurring events to the client.
	This method is mandatory. */
	CreatePullPointSubscription(request *CreatePullPointSubscription) (*CreatePullPointSubscriptionResponse, error)

	CreatePullPointSubscriptionContext(ctx context.Context, request *CreatePullPointSubscription) (*CreatePullPointSubscriptionResponse, error)

	/* The WS-BaseNotification specification defines a set of OPTIONAL WS-ResouceProperties.
	This specification does not require the implementation of the WS-ResourceProperty interface.
	Instead, the subsequent direct interface shall be implemented by an ONVIF compliant device
	in order to provide information about the FilterDialects, Schema files and topics supported by
	the device. */
	GetEventProperties(request *GetEventProperties) (*GetEventPropertiesResponse, error)

	GetEventPropertiesContext(ctx context.Context, request *GetEventProperties) (*GetEventPropertiesResponse, error)
}

type eventPortType struct {
	client *soap.Client
	xaddr  string
}

func NewEventPortType(client *soap.Client, xaddr string) EventPortType {
	return &eventPortType{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *eventPortType) GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	response := new(GetServiceCapabilitiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/events/wsdl/GetServiceCapabilities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *eventPortType) GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	return service.GetServiceCapabilitiesContext(
		context.Background(),
		request,
	)
}

func (service *eventPortType) CreatePullPointSubscriptionContext(ctx context.Context, request *CreatePullPointSubscription) (*CreatePullPointSubscriptionResponse, error) {
	response := new(CreatePullPointSubscriptionResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/events/wsdl/CreatePullPointSubscription", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *eventPortType) CreatePullPointSubscription(request *CreatePullPointSubscription) (*CreatePullPointSubscriptionResponse, error) {
	return service.CreatePullPointSubscriptionContext(
		context.Background(),
		request,
	)
}

func (service *eventPortType) GetEventPropertiesContext(ctx context.Context, request *GetEventProperties) (*GetEventPropertiesResponse, error) {
	response := new(GetEventPropertiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/events/wsdl/GetEventProperties", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *eventPortType) GetEventProperties(request *GetEventProperties) (*GetEventPropertiesResponse, error) {
	return service.GetEventPropertiesContext(
		context.Background(),
		request,
	)
}

type PullPointSubscription interface {

	// Error can be either of the following types:
	//
	//   - PullMessagesFaultResponse
	/*
		This method pulls one or more messages from a PullPoint.
		The device shall provide the following PullMessages command for all SubscriptionManager
		endpoints returned by the CreatePullPointSubscription command. This method shall not wait until
		the requested number of messages is available but return as soon as at least one message is available.
		The command shall at least support a Timeout of one minute. In case a device supports retrieval of less messages
		than requested it shall return these without generating a fault. */
	PullMessages(request *PullMessages) (*PullMessagesResponse, error)

	PullMessagesContext(ctx context.Context, request *PullMessages) (*PullMessagesResponse, error)

	/*
		This method readjusts the pull pointer into the past.
		A device supporting persistent notification storage shall provide the
		following Seek command for all SubscriptionManager endpoints returned by
		the CreatePullPointSubscription command. The optional Reverse argument can
		be used to reverse the pull direction of the PullMessages command.
		The UtcTime argument will be matched against the UtcTime attribute on a
		NotificationMessage.
	*/
	Seek(request *Seek) (*SeekResponse, error)

	SeekContext(ctx context.Context, request *Seek) (*SeekResponse, error)

	/* Properties inform a client about property creation, changes and
	deletion in a uniform way. When a client wants to synchronize its properties with the
	properties of the device, it can request a synchronization point which repeats the current
	status of all properties to which a client has subscribed. The PropertyOperation of all
	produced notifications is set to “Initialized”. The Synchronization Point is
	requested directly from the SubscriptionManager which was returned in either the
	SubscriptionResponse or in the CreatePullPointSubscriptionResponse. The property update is
	transmitted via the notification transportation of the notification interface. This method is mandatory.
	*/
	SetSynchronizationPoint(request *SetSynchronizationPoint) (*SetSynchronizationPointResponse, error)

	SetSynchronizationPointContext(ctx context.Context, request *SetSynchronizationPoint) (*SetSynchronizationPointResponse, error)

	// Error can be either of the following types:
	//
	//   - ResourceUnknownFault
	//   - UnableToDestroySubscriptionFault
	/* The device shall provide the following Unsubscribe command for all SubscriptionManager endpoints returned by the CreatePullPointSubscription command.
	This command shall terminate the lifetime of a pull point.
	*/
	Unsubscribe() error

	UnsubscribeContext(ctx context.Context) error
}

type pullPointSubscription struct {
	client *soap.Client
	xaddr  string
}

func NewPullPointSubscription(client *soap.Client, xaddr string) PullPointSubscription {
	return &pullPointSubscription{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *pullPointSubscription) PullMessagesContext(ctx context.Context, request *PullMessages) (*PullMessagesResponse, error) {
	response := new(PullMessagesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/events/wsdl/PullMessages", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pullPointSubscription) PullMessages(request *PullMessages) (*PullMessagesResponse, error) {
	return service.PullMessagesContext(
		context.Background(),
		request,
	)
}

func (service *pullPointSubscription) SeekContext(ctx context.Context, request *Seek) (*SeekResponse, error) {
	response := new(SeekResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/events/wsdl/Seek", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pullPointSubscription) Seek(request *Seek) (*SeekResponse, error) {
	return service.SeekContext(
		context.Background(),
		request,
	)
}

func (service *pullPointSubscription) SetSynchronizationPointContext(ctx context.Context, request *SetSynchronizationPoint) (*SetSynchronizationPointResponse, error) {
	response := new(SetSynchronizationPointResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/events/wsdl/SetSynchronizationPoint", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *pullPointSubscription) SetSynchronizationPoint(request *SetSynchronizationPoint) (*SetSynchronizationPointResponse, error) {
	return service.SetSynchronizationPointContext(
		context.Background(),
		request,
	)
}

func (service *pullPointSubscription) UnsubscribeContext(ctx context.Context) error {

	err := service.client.CallContext(ctx, service.xaddr, "''", nil, struct{}{})
	if err != nil {
		return err
	}

	return nil
}

func (service *pullPointSubscription) Unsubscribe() error {
	return service.UnsubscribeContext(
		context.Background(),
	)
}

type AnyURI string
type Duration string
type QName string
type NCName string
type NonNegativeInteger int64
type PositiveInteger int64
type NonPositiveInteger int64
type AnySimpleType string
type String string
