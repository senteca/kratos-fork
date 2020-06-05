package password_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/ory/jsonschema/v3"
	"github.com/ory/x/errorsx"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"github.com/ory/x/pointerx"

	"github.com/ory/viper"

	"github.com/ory/kratos/driver/configuration"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/internal"
	"github.com/ory/kratos/internal/testhelpers"
	"github.com/ory/kratos/schema"
	"github.com/ory/kratos/selfservice/flow/login"
	"github.com/ory/kratos/selfservice/form"
	"github.com/ory/kratos/selfservice/strategy/password"
	"github.com/ory/kratos/x"
)

func nlr(exp time.Duration) *login.Request {
	id := x.NewUUID()
	return &login.Request{
		ID:         id,
		IssuedAt:   time.Now().UTC(),
		ExpiresAt:  time.Now().UTC().Add(exp),
		RequestURL: "remove-this-if-test-fails",
		Methods: map[identity.CredentialsType]*login.RequestMethod{
			identity.CredentialsTypePassword: {
				Method: identity.CredentialsTypePassword,
				Config: &login.RequestMethodConfig{
					RequestMethodConfigurator: &form.HTMLForm{
						Method: "POST",
						Action: "/action",
						Fields: form.Fields{
							{
								Name:     "identifier",
								Type:     "text",
								Required: true,
							},
							{
								Name:     "password",
								Type:     "password",
								Required: true,
							},
							{
								Name:     form.CSRFTokenName,
								Type:     "hidden",
								Required: true,
								Value:    x.FakeCSRFToken,
							},
						},
					},
				},
			},
		},
	}
}

func TestLoginNew(t *testing.T) {
	_, reg := internal.NewFastRegistryWithMocks(t)

	ts, _ := testhelpers.NewKratosServer(t, reg)

	errTs := testhelpers.NewErrorTestServer(t, reg)
	uiTs := testhelpers.NewLoginUIRequestEchoServer(t, reg)
	newReturnTs(t, reg)

	// Overwrite these two:
	viper.Set(configuration.ViperKeyURLsError, errTs.URL+"/error-ts")
	viper.Set(configuration.ViperKeyURLsLogin, uiTs.URL+"/login-ts")

	viper.Set(configuration.ViperKeyDefaultIdentityTraitsSchemaURL, "file://./stub/login.schema.json")
	viper.Set(configuration.ViperKeySecretsSession, []string{"not-a-secure-session-key"})

	mr := func(t *testing.T, payload string, requestID string, c *http.Client) (*http.Response, []byte) {
		res, err := c.Post(ts.URL+password.LoginPath+"?request="+requestID, "application/x-www-form-urlencoded", strings.NewReader(payload))
		require.NoError(t, err)
		defer res.Body.Close()
		require.EqualValues(t, http.StatusOK, res.StatusCode, "Request: %+v\n\t\tResponse: %s", res.Request, res)
		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)
		return res, body
	}

	makeRequest := func(t *testing.T, payload string, jar *cookiejar.Jar, force bool) (*http.Response, []byte) {
		c := &http.Client{Jar: jar}
		if jar == nil {
			c.Jar, _ = cookiejar.New(&cookiejar.Options{})
		}

		u := ts.URL + login.BrowserLoginPath
		if force {
			u = u + "?prompt=login"
		}

		res, err := c.Get(u)
		require.NoError(t, err)
		require.EqualValues(t, http.StatusOK, res.StatusCode, "Request: %+v\n\t\tResponse: %s", res.Request, res)
		assert.NotEmpty(t, res.Request.URL.Query().Get("request"))

		return mr(t, payload, res.Request.URL.Query().Get("request"), c)
	}

	fakeRequest := func(t *testing.T, lr *login.Request, payload string, forceRequestID *string, jar *cookiejar.Jar) (*http.Response, []byte) {
		lr.RequestURL = ts.URL
		require.NoError(t, reg.LoginRequestPersister().CreateLoginRequest(context.TODO(), lr))

		requestID := lr.ID.String()
		if forceRequestID != nil {
			requestID = *forceRequestID
		}

		c := &http.Client{Jar: jar}
		if jar == nil {
			c.Jar, _ = cookiejar.New(&cookiejar.Options{})
		}

		return mr(t, payload, requestID, c)
	}

	ensureFieldsExist := func(t *testing.T, body []byte) {
		checkFormContent(t, body, "identifier",
			"password",
			"csrf_token")
	}

	createIdentity := func(identifier, password string) {
		p, _ := reg.Hasher().Generate([]byte(password))
		require.NoError(t, reg.PrivilegedIdentityPool().CreateIdentity(context.Background(), &identity.Identity{
			ID:     x.NewUUID(),
			Traits: identity.Traits(fmt.Sprintf(`{"subject":"%s"}`, identifier)),
			Credentials: map[identity.CredentialsType]identity.Credentials{
				identity.CredentialsTypePassword: {
					Type:        identity.CredentialsTypePassword,
					Identifiers: []string{identifier},
					Config:      json.RawMessage(`{"hashed_password":"` + string(p) + `"}`),
				},
			},
		}))
	}

	t.Run("should show the error ui because the request is malformed", func(t *testing.T) {
		lr := nlr(0)
		res, body := fakeRequest(t, lr, "14=)=!(%)$/ZP()GHIÖ", nil, nil)

		require.Contains(t, res.Request.URL.Path, "login-ts", "%+v", res.Request)
		assert.Equal(t, lr.ID.String(), gjson.GetBytes(body, "id").String(), "%s", body)
		assert.Equal(t, "/action", gjson.GetBytes(body, "methods.password.config.action").String(), "%s", body)
		assert.Contains(t, gjson.GetBytes(body, "methods.password.config.errors.0.message").String(), `invalid URL escape`)
	})

	t.Run("should show the error ui because the request id missing", func(t *testing.T) {
		lr := nlr(time.Minute)
		res, body := fakeRequest(t, lr, url.Values{}.Encode(), pointerx.String(""), nil)

		require.Contains(t, res.Request.URL.Path, "error-ts")
		assert.Equal(t, int64(http.StatusBadRequest), gjson.GetBytes(body, "0.code").Int(), "%s", body)
		assert.Equal(t, "Bad Request", gjson.GetBytes(body, "0.status").String(), "%s", body)
		assert.Contains(t, gjson.GetBytes(body, "0.reason").String(), "request query parameter is missing or invalid", "%s", body)
	})

	t.Run("should return an error because the request does not exist", func(t *testing.T) {
		lr := nlr(0)
		res, body := fakeRequest(t, lr, url.Values{
			"identifier": {"identifier"},
			"password":   {"password"},
		}.Encode(), pointerx.String(x.NewUUID().String()), nil)

		require.Contains(t, res.Request.URL.Path, "error-ts")
		assert.Equal(t, int64(http.StatusNotFound), gjson.GetBytes(body, "0.code").Int(), "%s", body)
		assert.Equal(t, "Not Found", gjson.GetBytes(body, "0.status").String(), "%s", body)
		assert.Contains(t, gjson.GetBytes(body, "0.message").String(), "Unable to locate the resource", "%s", body)
	})

	t.Run("should redirect to login init because the request is expired", func(t *testing.T) {
		lr := nlr(-time.Hour)
		res, body := fakeRequest(t, lr, url.Values{
			"identifier": {"identifier"},
			"password":   {"password"},
		}.Encode(), nil, nil)

		require.Contains(t, res.Request.URL.Path, "login-ts")
		assert.NotEqual(t, lr.ID, gjson.GetBytes(body, "id"))
		assert.Contains(t, gjson.GetBytes(body, "methods.oidc.config.errors.0").String(), "expired", "%s", body)
		assert.Contains(t, gjson.GetBytes(body, "methods.password.config.errors.0").String(), "expired", "%s", body)
	})

	t.Run("should return an error because the credentials are invalid (user does not exist)", func(t *testing.T) {
		lr := nlr(time.Hour)
		res, body := fakeRequest(t, lr, url.Values{
			"identifier": {"identifier"},
			"password":   {"password"},
		}.Encode(), nil, nil)

		require.Contains(t, res.Request.URL.Path, "login-ts")
		assert.Equal(t, lr.ID.String(), gjson.GetBytes(body, "id").String(), "%s", body)
		assert.Equal(t, "/action", gjson.GetBytes(body, "methods.password.config.action").String())
		assert.Equal(t, `the provided credentials are invalid, check for spelling mistakes in your password or username, email address, or phone number`, gjson.GetBytes(body, "methods.password.config.errors.0.message").String())
	})

	t.Run("should return an error because no identifier is set", func(t *testing.T) {
		lr := nlr(time.Hour)
		res, body := fakeRequest(t, lr, url.Values{
			"password": {"password"},
		}.Encode(), nil, nil)

		require.Contains(t, res.Request.URL.Path, "login-ts")
		// Let's ensure that the payload is being propagated properly.
		assert.Equal(t, lr.ID.String(), gjson.GetBytes(body, "id").String())
		assert.Equal(t, "/action", gjson.GetBytes(body, "methods.password.config.action").String())
		ensureFieldsExist(t, body)
		assert.Equal(t, "missing properties: identifier", gjson.GetBytes(body, "methods.password.config.fields.#(name==identifier).errors.0.message").String(), "%s", body)

		// The password value should not be returned!
		assert.Empty(t, gjson.GetBytes(body, "methods.password.config.fields.#(name==password).value").String())
	})

	t.Run("should return an error because no password is set", func(t *testing.T) {
		lr := nlr(time.Hour)
		res, body := fakeRequest(t, lr, url.Values{
			"identifier": {"identifier"},
		}.Encode(), nil, nil)

		require.Contains(t, res.Request.URL.Path, "login-ts")
		// Let's ensure that the payload is being propagated properly.
		assert.Equal(t, lr.ID.String(), gjson.GetBytes(body, "id").String())
		assert.Equal(t, "/action", gjson.GetBytes(body, "methods.password.config.action").String())
		ensureFieldsExist(t, body)
		assert.Equal(t, "missing properties: password", gjson.GetBytes(body, "methods.password.config.fields.#(name==password).errors.0.message").String(), "%s", body)

		assert.Equal(t, x.FakeCSRFToken, gjson.GetBytes(body, "methods.password.config.fields.#(name==csrf_token).value").String())
		assert.Equal(t, "identifier", gjson.GetBytes(body, "methods.password.config.fields.#(name==identifier).value").String(), "%s", body)

		// This must not include the password!
		assert.Empty(t, gjson.GetBytes(body, "methods.password.config.fields.#(name==password).value").String())
	})

	t.Run("should return an error because the credentials are invalid (password not correct)", func(t *testing.T) {
		identifier, pwd := "login-identifier-6", "password"
		createIdentity(identifier, pwd)

		lr := nlr(time.Hour)
		res, body := fakeRequest(t, lr, url.Values{
			"identifier": {identifier},
			"password":   {"not-password"},
		}.Encode(), nil, nil)

		require.Contains(t, res.Request.URL.Path, "login-ts")

		assert.Equal(t, lr.ID.String(), gjson.GetBytes(body, "id").String())
		assert.Equal(t, "/action", gjson.GetBytes(body, "methods.password.config.action").String())
		ensureFieldsExist(t, body)
		assert.Equal(t,
			errorsx.Cause(schema.NewInvalidCredentialsError()).(*jsonschema.ValidationError).Message,
			gjson.GetBytes(body, "methods.password.config.errors.0.message").String(),
			"%s", body,
		)

		// This must not include the password!
		assert.Empty(t, gjson.GetBytes(body, "methods.password.config.fields.#(name==password).value").String())
	})

	t.Run("should pass because with fake request", func(t *testing.T) {
		identifier, pwd := "login-identifier-7", "password"
		createIdentity(identifier, pwd)

		lr := nlr(time.Hour)
		res, body := fakeRequest(t, lr, url.Values{
			"identifier": {identifier},
			"password":   {pwd},
		}.Encode(), nil, nil)

		require.Contains(t, res.Request.URL.Path, "return-ts", "%s", res.Request.URL.String())
		assert.Equal(t, identifier, gjson.GetBytes(body, "identity.traits.subject").String(), "%s", body)
	})

	t.Run("should pass with real request", func(t *testing.T) {
		identifier, pwd := "login-identifier-8", "password"
		createIdentity(identifier, pwd)

		jar, _ := cookiejar.New(nil)
		res, body := makeRequest(t, url.Values{
			"identifier": {identifier},
			"password":   {pwd},
		}.Encode(), jar, true)

		require.Contains(t, res.Request.URL.Path, "return-ts", "%s", res.Request.URL.String())
		assert.Equal(t, identifier, gjson.GetBytes(body, "identity.traits.subject").String(), "%s", body)

		t.Run("retry with different prompts", func(t *testing.T) {
			c := &http.Client{Jar: jar}

			t.Run("redirect to returnTS if prompt is missing", func(t *testing.T) {
				res, err := c.Get(ts.URL + login.BrowserLoginPath)
				require.NoError(t, err)
				require.EqualValues(t, http.StatusOK, res.StatusCode)
			})

			t.Run("show UI and hint at username", func(t *testing.T) {
				res, err := c.Get(ts.URL + login.BrowserLoginPath + "?prompt=login")
				require.NoError(t, err)
				require.EqualValues(t, http.StatusOK, res.StatusCode)

				rid := res.Request.URL.Query().Get("request")
				assert.NotEmpty(t, rid, "%s", res.Request.URL)

				res, err = c.Get(ts.URL + login.BrowserLoginRequestsPath + "?request=" + rid)
				require.NoError(t, err)
				require.EqualValues(t, http.StatusOK, res.StatusCode)

				body, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				assert.True(t, gjson.GetBytes(body, "forced").Bool())
				assert.Equal(t, identifier, gjson.GetBytes(body, "methods.password.config.fields.#(name==identifier).value").String(), "%s", body)
				assert.Empty(t, gjson.GetBytes(body, "methods.password.config.fields.#(name==password).value").String(), "%s", body)
			})
		})
	})

	t.Run("should return an error because not passing validation and reset previous errors and values", func(t *testing.T) {
		lr := &login.Request{
			ID:        x.NewUUID(),
			ExpiresAt: time.Now().Add(time.Minute),
			Methods: map[identity.CredentialsType]*login.RequestMethod{
				identity.CredentialsTypePassword: {
					Method: identity.CredentialsTypePassword,
					Config: &login.RequestMethodConfig{
						RequestMethodConfigurator: &password.RequestMethod{
							HTMLForm: &form.HTMLForm{
								Method: "POST",
								Action: "/action",
								Errors: []form.Error{{Message: "some error"}},
								Fields: form.Fields{
									{
										Value:  "baz",
										Name:   "identifier",
										Errors: []form.Error{{Message: "err"}},
									},
									{
										Value:  "bar",
										Name:   "password",
										Errors: []form.Error{{Message: "err"}},
									},
								},
							},
						},
					},
				},
			},
		}

		res, body := fakeRequest(t, lr, url.Values{
			"identifier": {"registration-identifier-9"},
			// "password": {uuid.New().String()},
		}.Encode(), nil, nil)

		require.Contains(t, res.Request.URL.Path, "login-ts")
		assert.Equal(t, lr.ID.String(), gjson.GetBytes(body, "id").String())
		assert.Equal(t, "/action", gjson.GetBytes(body, "methods.password.config.action").String())
		ensureFieldsExist(t, body)

		assert.Empty(t, gjson.GetBytes(body, "methods.password.config.fields.#(name==identity).value"))
		assert.Empty(t, gjson.GetBytes(body, "methods.password.config.fields.#(name==identity).error"))
		assert.Empty(t, gjson.GetBytes(body, "methods.password.config.error"))
		assert.Contains(t, gjson.GetBytes(body, "methods.password.config.fields.#(name==password).errors.0").String(), "missing properties: password", "%s", body)
	})

	t.Run("should be a new session with forced flag", func(t *testing.T) {
		identifier, pwd := "login-identifier-reauth", "password"
		createIdentity(identifier, pwd)

		jar, err := cookiejar.New(&cookiejar.Options{})
		require.NoError(t, err)
		_, body1 := fakeRequest(t, nlr(time.Hour), url.Values{
			"identifier": {identifier},
			"password":   {pwd},
		}.Encode(), nil, jar)

		lr2 := nlr(time.Hour)
		lr2.Forced = true
		res, body2 := fakeRequest(t, lr2, url.Values{
			"identifier": {identifier},
			"password":   {pwd},
		}.Encode(), nil, jar)

		require.Contains(t, res.Request.URL.Path, "return-ts", "%s", res.Request.URL.String())
		assert.Equal(t, identifier, gjson.GetBytes(body2, "identity.traits.subject").String(), "%s", body2)
		assert.NotEqual(t, gjson.GetBytes(body1, "sid").String(), gjson.GetBytes(body2, "sid").String(), "%s\n\n%s\n", body1, body2)
	})

	t.Run("should be the same session without forced flag", func(t *testing.T) {
		identifier, pwd := "login-identifier-no-reauth", "password"
		createIdentity(identifier, pwd)

		jar, err := cookiejar.New(&cookiejar.Options{})
		require.NoError(t, err)
		_, body1 := fakeRequest(t, nlr(time.Hour), url.Values{
			"identifier": {identifier},
			"password":   {pwd},
		}.Encode(), nil, jar)

		lr2 := nlr(time.Hour)
		res, body2 := fakeRequest(t, lr2, url.Values{
			"identifier": {identifier},
			"password":   {pwd},
		}.Encode(), nil, jar)

		require.Contains(t, res.Request.URL.Path, "return-ts", "%s", res.Request.URL.String())
		assert.Equal(t, identifier, gjson.GetBytes(body2, "identity.traits.subject").String(), "%s", body2)
		assert.Equal(t, gjson.GetBytes(body1, "sid").String(), gjson.GetBytes(body2, "sid").String(), "%s\n\n%s\n", body1, body2)
	})
}
