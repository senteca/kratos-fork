package identity

import (
	"fmt"
	"sync"
	"time"

	"github.com/ory/jsonschema/v3"

	"github.com/ory/kratos/schema"
)

type SchemaExtensionVerify struct {
	lifespan time.Duration
	l        sync.Mutex
	v        []VerifiableAddress
	i        *Identity
}

func NewSchemaExtensionVerify(i *Identity, lifespan time.Duration) *SchemaExtensionVerify {
	return &SchemaExtensionVerify{i: i, lifespan: lifespan}
}

func (r *SchemaExtensionVerify) Run(ctx jsonschema.ValidationContext, s schema.ExtensionConfig, value interface{}) error {
	r.l.Lock()
	defer r.l.Unlock()

	switch s.Verification.Via {
	case "email":
		if !jsonschema.Formats["email"](value) {
			return ctx.Error("format", "%q is not valid %q", value, "email")
		}

		address, err := NewVerifiableEmailAddress(fmt.Sprintf("%s", value), r.i.ID, r.lifespan)
		if err != nil {
			return err
		}

		if has := r.has(r.i.Addresses, address); has != nil {
			if r.has(r.v, address) == nil {
				r.v = append(r.v, *has)
			}
			return nil
		}

		if has := r.has(r.v, address); has == nil {
			r.v = append(r.v, *address)
		}

		return nil
	case "":
		return nil
	}

	return ctx.Error("", "verification.via has unknown value %q", s.Verification.Via)
}

func (r *SchemaExtensionVerify) has(haystack []VerifiableAddress, needle *VerifiableAddress) *VerifiableAddress {
	for _, has := range haystack {
		if has.Value == needle.Value && has.Via == needle.Via {
			return &has
		}
	}
	return nil
}

func (r *SchemaExtensionVerify) Finish() error {
	r.i.Addresses = r.v
	return nil
}
