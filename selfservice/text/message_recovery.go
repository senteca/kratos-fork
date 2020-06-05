package text

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/tidwall/sjson"

	"github.com/ory/herodot"
)

func NewRecoverySuccessful(privilegedSessionExpiresAt time.Time) *Message {
	hasLeft := time.Until(privilegedSessionExpiresAt)
	context, _ := sjson.Set("{}", "expires_at", privilegedSessionExpiresAt)
	return &Message{
		ID:      InfoSelfServiceRecoverySuccessful,
		Type:    Info,
		Text:    fmt.Sprintf("You successfully recovered your accent. Please change your password or set up an alternative login method (e.g. social sign in) within the next %.2f minutes.", hasLeft.Minutes()),
		Context: []byte(context),
	}
}

func NewRecoveryEmailSent() *Message {
	return &Message{
		ID:      InfoSelfServiceRecoveryEmailSent,
		Type:    Info,
		Text:    "An email containing a recovery link has been sent to the email address you provided.",
		Context: []byte("{}"),
	}
}

func NewErrorValidationRecoveryMissingRecoveryToken() error {
	return errors.WithStack(herodot.
		ErrBadRequest.
		WithDetail("error_id", ErrorValidationRecoveryMissingRecoveryToken).
		WithReason("A recovery request was made but no recovery token was included in the request, please retry the flow."))
}

func NewErrorValidationRecoveryRecoveryTokenInvalidOrAlreadyUsed() *Message {
	return &Message{
		ID:      ErrorValidationRecoveryRecoveryTokenInvalidOrAlreadyUsed,
		Text:    "The recovery token is invalid or has already been used. Please retry the flow.",
		Type:    Error,
		Context: nil,
	}
}

func NewErrorValidationRetrySuccess() *Message {
	return &Message{
		ID:      ErrorValidationRecoveryRetrySuccess,
		Text:    "The request was already completed successfully and can not be retried.",
		Type:    Error,
		Context: nil,
	}
}

func NewErrorValidationUnexpectedState() *Message {
	return &Message{
		ID:      ErrorValidationRecoveryStateFailure,
		Text:    "The recovery flow reached a failure state and must be retried.",
		Type:    Error,
		Context: nil,
	}
}
