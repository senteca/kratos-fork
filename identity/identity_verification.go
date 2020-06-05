package identity

import (
	"time"

	"github.com/gofrs/uuid"

	"github.com/ory/kratos/otp"
)

const (
	VerifiableAddressTypeEmail VerifiableAddressType = "email"

	VerifiableAddressStatusPending   VerifiableAddressStatus = "pending"
	VerifiableAddressStatusCompleted VerifiableAddressStatus = "completed"
)

type (
	// VerifiableAddressType must not exceed 16 characters as that is the limitation in the SQL Schema.
	VerifiableAddressType string

	// VerifiableAddressStatus must not exceed 16 characters as that is the limitation in the SQL Schema.
	VerifiableAddressStatus string

	// swagger:model verifiableIdentityAddress
	VerifiableAddress struct {
		// required: true
		ID uuid.UUID `json:"id" db:"id" faker:"-"`

		// required: true
		Value string `json:"value" db:"value"`

		// required: true
		Verified bool `json:"verified" db:"verified"`

		// required: true
		Via VerifiableAddressType `json:"via" db:"via"`

		VerifiedAt *time.Time `json:"verified_at" faker:"-" db:"verified_at"`

		// required: true
		ExpiresAt time.Time `json:"expires_at" faker:"time_type" db:"expires_at"`

		// IdentityID is a helper struct field for gobuffalo.pop.
		IdentityID uuid.UUID `json:"-" faker:"-" db:"identity_id"`
		// CreatedAt is a helper struct field for gobuffalo.pop.
		CreatedAt time.Time `json:"-" faker:"-" db:"created_at"`
		// UpdatedAt is a helper struct field for gobuffalo.pop.
		UpdatedAt time.Time `json:"-" faker:"-" db:"updated_at"`
		// Code is the verification code, never to be shared as JSON
		Code   string                  `json:"-" db:"code"`
		Status VerifiableAddressStatus `json:"-" db:"status"`
	}
)

func (v VerifiableAddressType) HTMLFormInputType() string {
	switch v {
	case VerifiableAddressTypeEmail:
		return "email"
	}
	return ""
}

func (a VerifiableAddress) TableName() string {
	return "identity_verifiable_addresses"
}

func NewVerifiableEmailAddress(
	value string,
	identity uuid.UUID,
	expiresIn time.Duration,
) (*VerifiableAddress, error) {
	code, err := otp.New()
	if err != nil {
		return nil, err
	}

	return &VerifiableAddress{
		Code:       code,
		Value:      value,
		Verified:   false,
		Status:     VerifiableAddressStatusPending,
		Via:        VerifiableAddressTypeEmail,
		ExpiresAt:  time.Now().Add(expiresIn).UTC(),
		IdentityID: identity,
	}, nil
}
