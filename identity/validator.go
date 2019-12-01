package identity

import (
	"github.com/ory/gojsonschema"
	"github.com/ory/x/errorsx"
	"github.com/ory/x/stringsx"

	"github.com/ory/kratos/driver/configuration"
	"github.com/ory/kratos/schema"
)

type Validator struct {
	c configuration.Provider
	v *schema.Validator
}

type ValidationProvider interface {
	IdentityValidator() *Validator
}

func NewValidator(c configuration.Provider) *Validator {
	return &Validator{
		c: c,
		v: schema.NewValidator(),
	}
}

type ValidationExtender interface {
	WithIdentity(*Identity) ValidationExtender
	schema.ValidationExtender
}

func (v *Validator) Validate(i *Identity) error {
	es := []schema.ValidationExtender{
		NewValidationExtensionIdentifier().WithIdentity(i),
	}

	err := v.v.Validate(
		stringsx.Coalesce(
			i.TraitsSchemaURL,
			v.c.DefaultIdentityTraitsSchemaURL().String(),
		),
		gojsonschema.NewBytesLoader(i.Traits),
		es...,
	)

	switch errs := errorsx.Cause(err).(type) {
	case schema.ResultErrors:
		for k, err := range errs {
			errs[k].SetContext(schema.ContextSetRoot(schema.ContextRemoveRootStub(err.Context()), "traits"))
		}
		return errs
	}
	return err
}
