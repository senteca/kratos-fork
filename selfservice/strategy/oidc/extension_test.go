package oidc_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/gojsonschema"

	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/schema"
	"github.com/ory/kratos/selfservice/strategy/oidc"
)

func TestValidationExtension(t *testing.T) {
	ts := httptest.NewServer(http.FileServer(http.Dir("stub")))
	defer ts.Close()

	sv := schema.NewValidator()
	i := identity.NewIdentity(ts.URL + "/registration.schema.json")

	ve := oidc.NewValidationExtension()
	ve.WithIdentity(i)
	require.NoError(t, sv.Validate(
		ts.URL+"/extension.schema.json",
		gojsonschema.NewReferenceLoader("file://stub/extension.data.json"),
		ve,
	))

	assert.JSONEq(t, `{"email": "someone@email.org","names": ["peter","pan"]}`, string(i.Traits))
	assert.JSONEq(t, `{"email": "someone@email.org","names": ["peter","pan"]}`, string(ve.Values()))
}
