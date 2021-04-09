package testhelpers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/ory/kratos/x"
	"github.com/ory/x/pointerx"

	"github.com/ory/kratos-client-go"
)

func NewSDKClient(ts *httptest.Server) *kratos.APIClient {
	return NewSDKClientFromURL(ts.URL)
}

func NewSDKCustomClient(ts *httptest.Server, client *http.Client) *kratos.APIClient {
	conf := kratos.NewConfiguration()
	conf.Servers = kratos.ServerConfigurations{{URL: ts.URL}}
	conf.HTTPClient = client
	return kratos.NewAPIClient(conf)
}

func NewSDKClientFromURL(u string) *kratos.APIClient {
	conf := kratos.NewConfiguration()
	conf.Servers = kratos.ServerConfigurations{{URL: u}}
	return kratos.NewAPIClient(conf)
}

func SDKFormFieldsToURLValues(ff []kratos.UiNode) url.Values {
	values := url.Values{}
	for _, f := range ff {
		attr := f.Attributes.UiNodeInputAttributes
		if attr == nil {
			continue
		}

		val := attr.Value.GetActualInstance()
		if val == nil {
			continue
		}

		switch v := val.(type) {
		case *bool:
			values.Set(attr.Name, fmt.Sprintf("%v", *v))
		case *string:
			values.Set(attr.Name, fmt.Sprintf("%v", *v))
		case *float32:
			values.Set(attr.Name, fmt.Sprintf("%v", *v))
		case *float64:
			values.Set(attr.Name, fmt.Sprintf("%v", *v))
		}
	}
	return values
}

func NewFakeCSRFNode() *kratos.UiNode {
	return &kratos.UiNode{
		Group: "default",
		Type:  "input",
		Attributes: kratos.UiNodeInputAttributesAsUiNodeAttributes(&kratos.UiNodeInputAttributes{
			Name:     "csrf_token",
			Required: pointerx.Bool(true),
			Type:     "hidden",
			Value: &kratos.UiNodeInputAttributesValue{
				String: pointerx.String(x.FakeCSRFToken),
			},
		}),
	}
}

func NewSDKEmailNode(group string) *kratos.UiNode {
	return &kratos.UiNode{
		Type:  "input",
		Group: group,
		Attributes: kratos.UiNodeInputAttributesAsUiNodeAttributes(&kratos.UiNodeInputAttributes{
			Name:     "email",
			Type:     "email",
			Required: pointerx.Bool(true),
			Value:    &kratos.UiNodeInputAttributesValue{String: pointerx.String("email")},
		}),
	}
}

func NewSDKOIDCNode(name, provider string) *kratos.UiNode {
	return &kratos.UiNode{
		Attributes: kratos.UiNodeInputAttributesAsUiNodeAttributes(&kratos.UiNodeInputAttributes{
			Name:  name,
			Type:  "submit",
			Value: &kratos.UiNodeInputAttributesValue{String: pointerx.String(provider)},
		}),
	}
}
