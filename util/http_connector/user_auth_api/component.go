package user_auth_api

import (
	"time"
	"github.com/nexsoft-git/nexcommon/model"
)

// Component Schemas
// Generated from components.schemas.UserAuthRequest
type UserAuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	CountryCode string `json:"country_code"`
	ResourceId string `json:"resource_id"`
	Scope string `json:"scope"`
	Locale string `json:"locale"`
	AdditionalInformation []string `json:"additional_information"`
	EmailMessage string `json:"email_message"`
	EmailLinkMessage string `json:"email_link_message"`
	PhoneMessage string `json:"phone_message"`
}

// Generated from components.schemas.ErrorPayload
type ErrorPayload struct {
	Message string `json:"message"`
	Code string `json:"code"`
}

// Generated from components.schemas.NexsoftResponse
type NexsoftResponse struct {
	Nexsoft Nexsoft `json:"nexsoft"`
}

// Generated from components.schemas.Nexsoft
type Nexsoft struct {
	Header Header `json:"header"`
	Payload Payload `json:"payload"`
}

// Generated from components.schemas.Header
type Header struct {
	RequestId string `json:"request_id"`
	Version string `json:"version"`
	Timestamp time.Time `json:"timestamp"`
}

// Generated from components.schemas.Payload
type Payload struct {
	Status Status `json:"status"`
	Data Data `json:"data"`
	Other interface{} `json:"other"`
}

// Generated from components.schemas.Status
type Status struct {
	Success bool `json:"success"`
	Code string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

// Generated from components.schemas.Data
type Data struct {
	Meta interface{} `json:"meta"`
	Content Content `json:"content"`
}

// Generated from components.schemas.Content
type Content struct {
	UserId int `json:"user_id"`
	ClientId string `json:"client_id"`
	NotifyStatus NotifyStatus `json:"notify_status"`
}

// Generated from components.schemas.NotifyStatus
type NotifyStatus struct {
	EmailStatus EmailStatus `json:"email_status"`
	PhoneStatus PhoneStatus `json:"phone_status"`
}

// Generated from components.schemas.EmailStatus
type EmailStatus struct {
	EmailNotify bool `json:"email_notify"`
	EmailNotifyStatus bool `json:"email_notify_status"`
	EmailNotifyMessage string `json:"email_notify_message"`
}

// Generated from components.schemas.PhoneStatus
type PhoneStatus struct {
	PhoneNotify bool `json:"phone_notify"`
	PhoneNotifyStatus bool `json:"phone_notify_status"`
	PhoneNotifyMessage string `json:"phone_notify_message"`
}

// Generated from components.schemas.ErrorResponse
type ErrorResponse struct {
	Status string `json:"status"`
	Payload ErrorPayload `json:"payload"`
}

// Generated from components.schemas.UserUploadMultipart
type UserUploadMultipart struct {
	Data UserAuthRequest `json:"data"`
	Photo model.FileURLPath `json:"photo"`
}
