package text

import (
	"fmt"
	"time"
)

const (
	InfoSelfServiceLogin ID = 1010000 + iota
)

const (
	ErrorValidationLogin            ID = 4010000 + iota // 4010000
	ErrorValidationLoginFlowExpired                     // 4010001
)

func NewErrorValidationLoginFlowExpired(ago time.Duration) *Message {
	return &Message{
		ID:   ErrorValidationLoginFlowExpired,
		Text: fmt.Sprintf("The login flow expired %.2f minutes ago, please try again.", ago.Minutes()),
		Type: Error,
		Context: context(map[string]interface{}{
			"expired_at": time.Now().Add(ago),
		}),
	}
}
