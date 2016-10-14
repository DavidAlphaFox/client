// Auto-generated by avdl-compiler v1.3.9 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/local.avdl

package chat1

import (
	"errors"
	gregor1 "github.com/keybase/client/go/protocol/gregor1"
	keybase1 "github.com/keybase/client/go/protocol/keybase1"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type MessageText struct {
	Body string `codec:"body" json:"body"`
}

type MessageConversationMetadata struct {
	ConversationTitle string `codec:"conversationTitle" json:"conversationTitle"`
}

type MessageEdit struct {
	MessageID MessageID `codec:"messageID" json:"messageID"`
	Body      string    `codec:"body" json:"body"`
}

type MessageDelete struct {
	MessageID MessageID `codec:"messageID" json:"messageID"`
}

type MessageHeadline struct {
	Headline string `codec:"headline" json:"headline"`
}

type Asset struct {
	Filename string `codec:"filename" json:"filename"`
	Region   string `codec:"region" json:"region"`
	Endpoint string `codec:"endpoint" json:"endpoint"`
	Bucket   string `codec:"bucket" json:"bucket"`
	Path     string `codec:"path" json:"path"`
	Size     int    `codec:"size" json:"size"`
	MimeType string `codec:"mimeType" json:"mimeType"`
	EncHash  Hash   `codec:"encHash" json:"encHash"`
	Key      []byte `codec:"key" json:"key"`
}

type MessageAttachment struct {
	Object   Asset  `codec:"object" json:"object"`
	Preview  *Asset `codec:"preview,omitempty" json:"preview,omitempty"`
	Metadata []byte `codec:"metadata" json:"metadata"`
}

type MessageBody struct {
	MessageType__ MessageType                  `codec:"messageType" json:"messageType"`
	Text__        *MessageText                 `codec:"text,omitempty" json:"text,omitempty"`
	Attachment__  *MessageAttachment           `codec:"attachment,omitempty" json:"attachment,omitempty"`
	Edit__        *MessageEdit                 `codec:"edit,omitempty" json:"edit,omitempty"`
	Delete__      *MessageDelete               `codec:"delete,omitempty" json:"delete,omitempty"`
	Metadata__    *MessageConversationMetadata `codec:"metadata,omitempty" json:"metadata,omitempty"`
	Headline__    *MessageHeadline             `codec:"headline,omitempty" json:"headline,omitempty"`
}

func (o *MessageBody) MessageType() (ret MessageType, err error) {
	switch o.MessageType__ {
	case MessageType_TEXT:
		if o.Text__ == nil {
			err = errors.New("unexpected nil value for Text__")
			return ret, err
		}
	case MessageType_ATTACHMENT:
		if o.Attachment__ == nil {
			err = errors.New("unexpected nil value for Attachment__")
			return ret, err
		}
	case MessageType_EDIT:
		if o.Edit__ == nil {
			err = errors.New("unexpected nil value for Edit__")
			return ret, err
		}
	case MessageType_DELETE:
		if o.Delete__ == nil {
			err = errors.New("unexpected nil value for Delete__")
			return ret, err
		}
	case MessageType_METADATA:
		if o.Metadata__ == nil {
			err = errors.New("unexpected nil value for Metadata__")
			return ret, err
		}
	case MessageType_HEADLINE:
		if o.Headline__ == nil {
			err = errors.New("unexpected nil value for Headline__")
			return ret, err
		}
	}
	return o.MessageType__, nil
}

func (o MessageBody) Text() MessageText {
	if o.MessageType__ != MessageType_TEXT {
		panic("wrong case accessed")
	}
	if o.Text__ == nil {
		return MessageText{}
	}
	return *o.Text__
}

func (o MessageBody) Attachment() MessageAttachment {
	if o.MessageType__ != MessageType_ATTACHMENT {
		panic("wrong case accessed")
	}
	if o.Attachment__ == nil {
		return MessageAttachment{}
	}
	return *o.Attachment__
}

func (o MessageBody) Edit() MessageEdit {
	if o.MessageType__ != MessageType_EDIT {
		panic("wrong case accessed")
	}
	if o.Edit__ == nil {
		return MessageEdit{}
	}
	return *o.Edit__
}

func (o MessageBody) Delete() MessageDelete {
	if o.MessageType__ != MessageType_DELETE {
		panic("wrong case accessed")
	}
	if o.Delete__ == nil {
		return MessageDelete{}
	}
	return *o.Delete__
}

func (o MessageBody) Metadata() MessageConversationMetadata {
	if o.MessageType__ != MessageType_METADATA {
		panic("wrong case accessed")
	}
	if o.Metadata__ == nil {
		return MessageConversationMetadata{}
	}
	return *o.Metadata__
}

func (o MessageBody) Headline() MessageHeadline {
	if o.MessageType__ != MessageType_HEADLINE {
		panic("wrong case accessed")
	}
	if o.Headline__ == nil {
		return MessageHeadline{}
	}
	return *o.Headline__
}

func NewMessageBodyWithText(v MessageText) MessageBody {
	return MessageBody{
		MessageType__: MessageType_TEXT,
		Text__:        &v,
	}
}

func NewMessageBodyWithAttachment(v MessageAttachment) MessageBody {
	return MessageBody{
		MessageType__: MessageType_ATTACHMENT,
		Attachment__:  &v,
	}
}

func NewMessageBodyWithEdit(v MessageEdit) MessageBody {
	return MessageBody{
		MessageType__: MessageType_EDIT,
		Edit__:        &v,
	}
}

func NewMessageBodyWithDelete(v MessageDelete) MessageBody {
	return MessageBody{
		MessageType__: MessageType_DELETE,
		Delete__:      &v,
	}
}

func NewMessageBodyWithMetadata(v MessageConversationMetadata) MessageBody {
	return MessageBody{
		MessageType__: MessageType_METADATA,
		Metadata__:    &v,
	}
}

func NewMessageBodyWithHeadline(v MessageHeadline) MessageBody {
	return MessageBody{
		MessageType__: MessageType_HEADLINE,
		Headline__:    &v,
	}
}

type HeaderPlaintextVersion int

const (
	HeaderPlaintextVersion_V1 HeaderPlaintextVersion = 1
)

var HeaderPlaintextVersionMap = map[string]HeaderPlaintextVersion{
	"V1": 1,
}

var HeaderPlaintextVersionRevMap = map[HeaderPlaintextVersion]string{
	1: "V1",
}

type HeaderPlaintextV1 struct {
	Conv            ConversationIDTriple     `codec:"conv" json:"conv"`
	TlfName         string                   `codec:"tlfName" json:"tlfName"`
	TlfPublic       bool                     `codec:"tlfPublic" json:"tlfPublic"`
	MessageType     MessageType              `codec:"messageType" json:"messageType"`
	Prev            []MessagePreviousPointer `codec:"prev" json:"prev"`
	Sender          gregor1.UID              `codec:"sender" json:"sender"`
	SenderDevice    gregor1.DeviceID         `codec:"senderDevice" json:"senderDevice"`
	BodyHash        Hash                     `codec:"bodyHash" json:"bodyHash"`
	HeaderSignature *SignatureInfo           `codec:"headerSignature,omitempty" json:"headerSignature,omitempty"`
}

type HeaderPlaintext struct {
	Version__ HeaderPlaintextVersion `codec:"version" json:"version"`
	V1__      *HeaderPlaintextV1     `codec:"v1,omitempty" json:"v1,omitempty"`
}

func (o *HeaderPlaintext) Version() (ret HeaderPlaintextVersion, err error) {
	switch o.Version__ {
	case HeaderPlaintextVersion_V1:
		if o.V1__ == nil {
			err = errors.New("unexpected nil value for V1__")
			return ret, err
		}
	}
	return o.Version__, nil
}

func (o HeaderPlaintext) V1() HeaderPlaintextV1 {
	if o.Version__ != HeaderPlaintextVersion_V1 {
		panic("wrong case accessed")
	}
	if o.V1__ == nil {
		return HeaderPlaintextV1{}
	}
	return *o.V1__
}

func NewHeaderPlaintextWithV1(v HeaderPlaintextV1) HeaderPlaintext {
	return HeaderPlaintext{
		Version__: HeaderPlaintextVersion_V1,
		V1__:      &v,
	}
}

type BodyPlaintextVersion int

const (
	BodyPlaintextVersion_V1 BodyPlaintextVersion = 1
)

var BodyPlaintextVersionMap = map[string]BodyPlaintextVersion{
	"V1": 1,
}

var BodyPlaintextVersionRevMap = map[BodyPlaintextVersion]string{
	1: "V1",
}

type BodyPlaintextV1 struct {
	MessageBody MessageBody `codec:"messageBody" json:"messageBody"`
}

type BodyPlaintext struct {
	Version__ BodyPlaintextVersion `codec:"version" json:"version"`
	V1__      *BodyPlaintextV1     `codec:"v1,omitempty" json:"v1,omitempty"`
}

func (o *BodyPlaintext) Version() (ret BodyPlaintextVersion, err error) {
	switch o.Version__ {
	case BodyPlaintextVersion_V1:
		if o.V1__ == nil {
			err = errors.New("unexpected nil value for V1__")
			return ret, err
		}
	}
	return o.Version__, nil
}

func (o BodyPlaintext) V1() BodyPlaintextV1 {
	if o.Version__ != BodyPlaintextVersion_V1 {
		panic("wrong case accessed")
	}
	if o.V1__ == nil {
		return BodyPlaintextV1{}
	}
	return *o.V1__
}

func NewBodyPlaintextWithV1(v BodyPlaintextV1) BodyPlaintext {
	return BodyPlaintext{
		Version__: BodyPlaintextVersion_V1,
		V1__:      &v,
	}
}

type MessagePlaintext struct {
	ClientHeader MessageClientHeader `codec:"clientHeader" json:"clientHeader"`
	MessageBody  MessageBody         `codec:"messageBody" json:"messageBody"`
}

type MessageUnboxedState int

const (
	MessageUnboxedState_VALID MessageUnboxedState = 1
	MessageUnboxedState_ERROR MessageUnboxedState = 2
)

var MessageUnboxedStateMap = map[string]MessageUnboxedState{
	"VALID": 1,
	"ERROR": 2,
}

var MessageUnboxedStateRevMap = map[MessageUnboxedState]string{
	1: "VALID",
	2: "ERROR",
}

type MessageUnboxedValid struct {
	ClientHeader     MessageClientHeader `codec:"clientHeader" json:"clientHeader"`
	ServerHeader     MessageServerHeader `codec:"serverHeader" json:"serverHeader"`
	MessageBody      MessageBody         `codec:"messageBody" json:"messageBody"`
	SenderUsername   string              `codec:"senderUsername" json:"senderUsername"`
	SenderDeviceName string              `codec:"senderDeviceName" json:"senderDeviceName"`
	HeaderHash       Hash                `codec:"headerHash" json:"headerHash"`
}

type MessageUnboxedError struct {
	ErrMsg      string      `codec:"errMsg" json:"errMsg"`
	MessageID   MessageID   `codec:"messageID" json:"messageID"`
	MessageType MessageType `codec:"messageType" json:"messageType"`
}

type MessageUnboxed struct {
	State__ MessageUnboxedState  `codec:"state" json:"state"`
	Valid__ *MessageUnboxedValid `codec:"valid,omitempty" json:"valid,omitempty"`
	Error__ *MessageUnboxedError `codec:"error,omitempty" json:"error,omitempty"`
}

func (o *MessageUnboxed) State() (ret MessageUnboxedState, err error) {
	switch o.State__ {
	case MessageUnboxedState_VALID:
		if o.Valid__ == nil {
			err = errors.New("unexpected nil value for Valid__")
			return ret, err
		}
	case MessageUnboxedState_ERROR:
		if o.Error__ == nil {
			err = errors.New("unexpected nil value for Error__")
			return ret, err
		}
	}
	return o.State__, nil
}

func (o MessageUnboxed) Valid() MessageUnboxedValid {
	if o.State__ != MessageUnboxedState_VALID {
		panic("wrong case accessed")
	}
	if o.Valid__ == nil {
		return MessageUnboxedValid{}
	}
	return *o.Valid__
}

func (o MessageUnboxed) Error() MessageUnboxedError {
	if o.State__ != MessageUnboxedState_ERROR {
		panic("wrong case accessed")
	}
	if o.Error__ == nil {
		return MessageUnboxedError{}
	}
	return *o.Error__
}

func NewMessageUnboxedWithValid(v MessageUnboxedValid) MessageUnboxed {
	return MessageUnboxed{
		State__: MessageUnboxedState_VALID,
		Valid__: &v,
	}
}

func NewMessageUnboxedWithError(v MessageUnboxedError) MessageUnboxed {
	return MessageUnboxed{
		State__: MessageUnboxedState_ERROR,
		Error__: &v,
	}
}

type ThreadView struct {
	Messages   []MessageUnboxed `codec:"messages" json:"messages"`
	Pagination *Pagination      `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type UnreadFirstNumLimit struct {
	NumRead int `codec:"NumRead" json:"NumRead"`
	AtLeast int `codec:"AtLeast" json:"AtLeast"`
	AtMost  int `codec:"AtMost" json:"AtMost"`
}

type ConversationInfoLocal struct {
	Id         ConversationID       `codec:"id" json:"id"`
	Triple     ConversationIDTriple `codec:"triple" json:"triple"`
	TlfName    string               `codec:"tlfName" json:"tlfName"`
	TopicName  string               `codec:"topicName" json:"topicName"`
	Visibility TLFVisibility        `codec:"visibility" json:"visibility"`
}

type ConversationLocal struct {
	Error       *string                `codec:"error,omitempty" json:"error,omitempty"`
	Info        ConversationInfoLocal  `codec:"info" json:"info"`
	ReaderInfo  ConversationReaderInfo `codec:"readerInfo" json:"readerInfo"`
	MaxMessages []MessageUnboxed       `codec:"maxMessages" json:"maxMessages"`
}

type GetThreadQuery struct {
	MarkAsRead   bool          `codec:"markAsRead" json:"markAsRead"`
	MessageTypes []MessageType `codec:"messageTypes" json:"messageTypes"`
	Before       *gregor1.Time `codec:"before,omitempty" json:"before,omitempty"`
	After        *gregor1.Time `codec:"after,omitempty" json:"after,omitempty"`
}

type GetThreadLocalRes struct {
	Thread     ThreadView  `codec:"thread" json:"thread"`
	RateLimits []RateLimit `codec:"rateLimits" json:"rateLimits"`
}

type GetInboxLocalRes struct {
	ConversationsUnverified []Conversation `codec:"conversationsUnverified" json:"conversationsUnverified"`
	Pagination              *Pagination    `codec:"pagination,omitempty" json:"pagination,omitempty"`
	RateLimits              []RateLimit    `codec:"rateLimits" json:"rateLimits"`
}

type GetInboxLocalQuery struct {
	TlfName           *string         `codec:"tlfName,omitempty" json:"tlfName,omitempty"`
	TopicName         *string         `codec:"topicName,omitempty" json:"topicName,omitempty"`
	ConvID            *ConversationID `codec:"convID,omitempty" json:"convID,omitempty"`
	TopicType         *TopicType      `codec:"topicType,omitempty" json:"topicType,omitempty"`
	TlfVisibility     *TLFVisibility  `codec:"tlfVisibility,omitempty" json:"tlfVisibility,omitempty"`
	Before            *gregor1.Time   `codec:"before,omitempty" json:"before,omitempty"`
	After             *gregor1.Time   `codec:"after,omitempty" json:"after,omitempty"`
	OneChatTypePerTLF *bool           `codec:"oneChatTypePerTLF,omitempty" json:"oneChatTypePerTLF,omitempty"`
	UnreadOnly        bool            `codec:"unreadOnly" json:"unreadOnly"`
	ReadOnly          bool            `codec:"readOnly" json:"readOnly"`
}

type GetInboxAndUnboxLocalRes struct {
	Conversations []ConversationLocal `codec:"conversations" json:"conversations"`
	Pagination    *Pagination         `codec:"pagination,omitempty" json:"pagination,omitempty"`
	RateLimits    []RateLimit         `codec:"rateLimits" json:"rateLimits"`
}

type PostLocalRes struct {
	RateLimits []RateLimit `codec:"rateLimits" json:"rateLimits"`
}

type NewConversationLocalRes struct {
	Conv       ConversationLocal `codec:"conv" json:"conv"`
	RateLimits []RateLimit       `codec:"rateLimits" json:"rateLimits"`
}

type GetInboxSummaryForCLILocalQuery struct {
	TopicType           TopicType           `codec:"topicType" json:"topicType"`
	After               string              `codec:"after" json:"after"`
	Before              string              `codec:"before" json:"before"`
	Visibility          TLFVisibility       `codec:"visibility" json:"visibility"`
	UnreadFirst         bool                `codec:"unreadFirst" json:"unreadFirst"`
	UnreadFirstLimit    UnreadFirstNumLimit `codec:"unreadFirstLimit" json:"unreadFirstLimit"`
	ActivitySortedLimit int                 `codec:"activitySortedLimit" json:"activitySortedLimit"`
}

type GetInboxSummaryForCLILocalRes struct {
	Conversations []ConversationLocal `codec:"conversations" json:"conversations"`
	RateLimits    []RateLimit         `codec:"rateLimits" json:"rateLimits"`
}

type GetConversationForCLILocalQuery struct {
	MarkAsRead     bool                `codec:"markAsRead" json:"markAsRead"`
	MessageTypes   []MessageType       `codec:"MessageTypes" json:"MessageTypes"`
	Since          *string             `codec:"Since,omitempty" json:"Since,omitempty"`
	Limit          UnreadFirstNumLimit `codec:"limit" json:"limit"`
	ConversationId ConversationID      `codec:"conversationId" json:"conversationId"`
}

type GetConversationForCLILocalRes struct {
	Conversation ConversationLocal `codec:"conversation" json:"conversation"`
	Messages     []MessageUnboxed  `codec:"messages" json:"messages"`
	RateLimits   []RateLimit       `codec:"rateLimits" json:"rateLimits"`
}

type GetMessagesLocalRes struct {
	Messages   []MessageUnboxed `codec:"messages" json:"messages"`
	RateLimits []RateLimit      `codec:"rateLimits" json:"rateLimits"`
}

type DownloadAttachmentLocalRes struct {
	RateLimits []RateLimit `codec:"rateLimits" json:"rateLimits"`
}

type GetThreadLocalArg struct {
	ConversationID ConversationID  `codec:"conversationID" json:"conversationID"`
	Query          *GetThreadQuery `codec:"query,omitempty" json:"query,omitempty"`
	Pagination     *Pagination     `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type GetInboxLocalArg struct {
	Query      *GetInboxLocalQuery `codec:"query,omitempty" json:"query,omitempty"`
	Pagination *Pagination         `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type GetInboxAndUnboxLocalArg struct {
	Query      *GetInboxLocalQuery `codec:"query,omitempty" json:"query,omitempty"`
	Pagination *Pagination         `codec:"pagination,omitempty" json:"pagination,omitempty"`
}

type PostLocalArg struct {
	ConversationID ConversationID   `codec:"conversationID" json:"conversationID"`
	Msg            MessagePlaintext `codec:"msg" json:"msg"`
}

type PostAttachmentLocalArg struct {
	SessionID      int                 `codec:"sessionID" json:"sessionID"`
	ConversationID ConversationID      `codec:"conversationID" json:"conversationID"`
	ClientHeader   MessageClientHeader `codec:"clientHeader" json:"clientHeader"`
	Source         keybase1.Stream     `codec:"source" json:"source"`
	Filename       string              `codec:"filename" json:"filename"`
	Size           int                 `codec:"size" json:"size"`
}

type NewConversationLocalArg struct {
	TlfName       string        `codec:"tlfName" json:"tlfName"`
	TopicType     TopicType     `codec:"topicType" json:"topicType"`
	TlfVisibility TLFVisibility `codec:"tlfVisibility" json:"tlfVisibility"`
	TopicName     *string       `codec:"topicName,omitempty" json:"topicName,omitempty"`
}

type GetInboxSummaryForCLILocalArg struct {
	Query GetInboxSummaryForCLILocalQuery `codec:"query" json:"query"`
}

type GetConversationForCLILocalArg struct {
	Query GetConversationForCLILocalQuery `codec:"query" json:"query"`
}

type GetMessagesLocalArg struct {
	ConversationID ConversationID `codec:"conversationID" json:"conversationID"`
	MessageIDs     []MessageID    `codec:"messageIDs" json:"messageIDs"`
}

type DownloadAttachmentLocalArg struct {
	SessionID      int             `codec:"sessionID" json:"sessionID"`
	ConversationID ConversationID  `codec:"conversationID" json:"conversationID"`
	MessageID      MessageID       `codec:"messageID" json:"messageID"`
	Sink           keybase1.Stream `codec:"sink" json:"sink"`
}

type LocalInterface interface {
	GetThreadLocal(context.Context, GetThreadLocalArg) (GetThreadLocalRes, error)
	GetInboxLocal(context.Context, GetInboxLocalArg) (GetInboxLocalRes, error)
	GetInboxAndUnboxLocal(context.Context, GetInboxAndUnboxLocalArg) (GetInboxAndUnboxLocalRes, error)
	PostLocal(context.Context, PostLocalArg) (PostLocalRes, error)
	PostAttachmentLocal(context.Context, PostAttachmentLocalArg) (PostLocalRes, error)
	NewConversationLocal(context.Context, NewConversationLocalArg) (NewConversationLocalRes, error)
	GetInboxSummaryForCLILocal(context.Context, GetInboxSummaryForCLILocalQuery) (GetInboxSummaryForCLILocalRes, error)
	GetConversationForCLILocal(context.Context, GetConversationForCLILocalQuery) (GetConversationForCLILocalRes, error)
	GetMessagesLocal(context.Context, GetMessagesLocalArg) (GetMessagesLocalRes, error)
	DownloadAttachmentLocal(context.Context, DownloadAttachmentLocalArg) (DownloadAttachmentLocalRes, error)
}

func LocalProtocol(i LocalInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "chat.1.local",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getThreadLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetThreadLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetThreadLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetThreadLocalArg)(nil), args)
						return
					}
					ret, err = i.GetThreadLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getInboxLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetInboxLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetInboxLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetInboxLocalArg)(nil), args)
						return
					}
					ret, err = i.GetInboxLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getInboxAndUnboxLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetInboxAndUnboxLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetInboxAndUnboxLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetInboxAndUnboxLocalArg)(nil), args)
						return
					}
					ret, err = i.GetInboxAndUnboxLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"postLocal": {
				MakeArg: func() interface{} {
					ret := make([]PostLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PostLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]PostLocalArg)(nil), args)
						return
					}
					ret, err = i.PostLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"postAttachmentLocal": {
				MakeArg: func() interface{} {
					ret := make([]PostAttachmentLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]PostAttachmentLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]PostAttachmentLocalArg)(nil), args)
						return
					}
					ret, err = i.PostAttachmentLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"newConversationLocal": {
				MakeArg: func() interface{} {
					ret := make([]NewConversationLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]NewConversationLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]NewConversationLocalArg)(nil), args)
						return
					}
					ret, err = i.NewConversationLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getInboxSummaryForCLILocal": {
				MakeArg: func() interface{} {
					ret := make([]GetInboxSummaryForCLILocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetInboxSummaryForCLILocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetInboxSummaryForCLILocalArg)(nil), args)
						return
					}
					ret, err = i.GetInboxSummaryForCLILocal(ctx, (*typedArgs)[0].Query)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"getConversationForCLILocal": {
				MakeArg: func() interface{} {
					ret := make([]GetConversationForCLILocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetConversationForCLILocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetConversationForCLILocalArg)(nil), args)
						return
					}
					ret, err = i.GetConversationForCLILocal(ctx, (*typedArgs)[0].Query)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"GetMessagesLocal": {
				MakeArg: func() interface{} {
					ret := make([]GetMessagesLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetMessagesLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetMessagesLocalArg)(nil), args)
						return
					}
					ret, err = i.GetMessagesLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"DownloadAttachmentLocal": {
				MakeArg: func() interface{} {
					ret := make([]DownloadAttachmentLocalArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]DownloadAttachmentLocalArg)
					if !ok {
						err = rpc.NewTypeError((*[]DownloadAttachmentLocalArg)(nil), args)
						return
					}
					ret, err = i.DownloadAttachmentLocal(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type LocalClient struct {
	Cli rpc.GenericClient
}

func (c LocalClient) GetThreadLocal(ctx context.Context, __arg GetThreadLocalArg) (res GetThreadLocalRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.local.getThreadLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) GetInboxLocal(ctx context.Context, __arg GetInboxLocalArg) (res GetInboxLocalRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.local.getInboxLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) GetInboxAndUnboxLocal(ctx context.Context, __arg GetInboxAndUnboxLocalArg) (res GetInboxAndUnboxLocalRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.local.getInboxAndUnboxLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) PostLocal(ctx context.Context, __arg PostLocalArg) (res PostLocalRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.local.postLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) PostAttachmentLocal(ctx context.Context, __arg PostAttachmentLocalArg) (res PostLocalRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.local.postAttachmentLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) NewConversationLocal(ctx context.Context, __arg NewConversationLocalArg) (res NewConversationLocalRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.local.newConversationLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) GetInboxSummaryForCLILocal(ctx context.Context, query GetInboxSummaryForCLILocalQuery) (res GetInboxSummaryForCLILocalRes, err error) {
	__arg := GetInboxSummaryForCLILocalArg{Query: query}
	err = c.Cli.Call(ctx, "chat.1.local.getInboxSummaryForCLILocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) GetConversationForCLILocal(ctx context.Context, query GetConversationForCLILocalQuery) (res GetConversationForCLILocalRes, err error) {
	__arg := GetConversationForCLILocalArg{Query: query}
	err = c.Cli.Call(ctx, "chat.1.local.getConversationForCLILocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) GetMessagesLocal(ctx context.Context, __arg GetMessagesLocalArg) (res GetMessagesLocalRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.local.GetMessagesLocal", []interface{}{__arg}, &res)
	return
}

func (c LocalClient) DownloadAttachmentLocal(ctx context.Context, __arg DownloadAttachmentLocalArg) (res DownloadAttachmentLocalRes, err error) {
	err = c.Cli.Call(ctx, "chat.1.local.DownloadAttachmentLocal", []interface{}{__arg}, &res)
	return
}
