package profile_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/nosurf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"github.com/urfave/negroni"

	"github.com/ory/x/httpx"

	"github.com/ory/viper"
	"github.com/ory/x/urlx"

	"github.com/ory/kratos/driver/configuration"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/internal"
	"github.com/ory/kratos/internal/httpclient/client"
	"github.com/ory/kratos/internal/httpclient/client/common"
	"github.com/ory/kratos/internal/httpclient/models"
	"github.com/ory/kratos/selfservice/errorx"
	"github.com/ory/kratos/selfservice/flow/profile"
	"github.com/ory/kratos/selfservice/form"
	"github.com/ory/kratos/session"
	"github.com/ory/kratos/x"
)

func init() {
	internal.RegisterFakes()
}

func fieldsToURLValues(ff models.FormFields) url.Values {
	values := url.Values{}
	for _, f := range ff {
		values.Set(f.Name, fmt.Sprintf("%v", f.Value))
	}
	return values
}

func TestUpdateProfile(t *testing.T) {
	_, reg := internal.NewRegistryDefault(t)
	viper.Set(configuration.ViperKeyDefaultIdentityTraitsSchemaURL, "file://./stub/identity.schema.json")

	ui := func() *httptest.Server {
		router := httprouter.New()
		router.GET("/profile", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			w.WriteHeader(http.StatusNoContent)
		})
		router.GET("/login", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			w.WriteHeader(http.StatusUnauthorized)
		})
		router.GET("/error", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
			w.WriteHeader(http.StatusUnauthorized)
		})
		return httptest.NewServer(router)
	}()
	defer ui.Close()

	errTs := errorx.NewErrorTestServer(t, reg)
	defer errTs.Close()

	viper.Set(configuration.ViperKeyURLsError, errTs.URL)
	viper.Set(configuration.ViperKeyURLsProfile, ui.URL+"/profile")
	viper.Set(configuration.ViperKeyURLsLogin, ui.URL+"/login")
	// set this intermediate because kratos needs some valid url for CRUDE operations
	viper.Set(configuration.ViperKeyURLsSelfPublic, "http://example.com")

	primaryIdentity := &identity.Identity{
		ID: x.NewUUID(),
		Credentials: map[identity.CredentialsType]identity.Credentials{
			"password": {Type: "password", Identifiers: []string{"john@doe.com"}, Config: json.RawMessage(`{"hashed_password":"foo"}`)},
		},
		Traits:         identity.Traits(`{"email":"john@doe.com","stringy":"foobar","booly":false,"numby":2.5}`),
		TraitsSchemaID: configuration.DefaultIdentityTraitsSchemaID,
	}

	publicTS, adminTS := func() (*httptest.Server, *httptest.Server) {
		router := x.NewRouterPublic()
		admin := x.NewRouterAdmin()
		reg.ProfileManagementHandler().RegisterPublicRoutes(router)
		reg.ProfileManagementHandler().RegisterAdminRoutes(admin)
		route, _ := session.MockSessionCreateHandlerWithIdentity(t, reg, primaryIdentity)
		router.GET("/setSession", route)

		other, _ := session.MockSessionCreateHandlerWithIdentity(t, reg, &identity.Identity{ID: x.NewUUID(), Traits: identity.Traits(`{}`)})
		router.GET("/setSession/other-user", other)
		n := negroni.Classic()
		n.UseHandler(router)
		return httptest.NewServer(nosurf.New(n)), httptest.NewServer(admin)
	}()
	defer publicTS.Close()

	viper.Set(configuration.ViperKeyURLsSelfPublic, publicTS.URL)

	primaryUser := func() *http.Client {
		c := session.MockCookieClient(t)
		session.MockHydrateCookieClient(t, c, publicTS.URL+"/setSession")
		return c
	}()

	otherUser := func() *http.Client {
		c := session.MockCookieClient(t)
		session.MockHydrateCookieClient(t, c, publicTS.URL+"/setSession/other-user")
		return c
	}()

	publicClient := client.NewHTTPClientWithConfig(
		nil,
		&client.TransportConfig{Host: urlx.ParseOrPanic(publicTS.URL).Host, BasePath: "/", Schemes: []string{"http"}},
	)

	adminClient := client.NewHTTPClientWithConfig(
		nil,
		&client.TransportConfig{Host: urlx.ParseOrPanic(adminTS.URL).Host, BasePath: "/", Schemes: []string{"http"}},
	)

	makeRequest := func(t *testing.T) *common.GetSelfServiceBrowserProfileManagementRequestOK {
		res, err := primaryUser.Get(publicTS.URL + profile.PublicProfileManagementPath)
		require.NoError(t, err)

		rs, err := publicClient.Common.GetSelfServiceBrowserProfileManagementRequest(
			common.NewGetSelfServiceBrowserProfileManagementRequestParams().WithHTTPClient(primaryUser).
				WithRequest(res.Request.URL.Query().Get("request")),
		)
		require.NoError(t, err)

		return rs
	}

	t.Run("description=call endpoints", func(t *testing.T) {
		pr, ar := x.NewRouterPublic(), x.NewRouterAdmin()
		reg.ProfileManagementHandler().RegisterPublicRoutes(pr)
		reg.ProfileManagementHandler().RegisterAdminRoutes(ar)

		adminTS, publicTS := httptest.NewServer(ar), httptest.NewServer(pr)
		defer adminTS.Close()
		defer publicTS.Close()

		for k, tc := range []*http.Request{
			httpx.MustNewRequest("GET", publicTS.URL+profile.PublicProfileManagementPath, nil, ""),
			httpx.MustNewRequest("GET", publicTS.URL+profile.PublicProfileManagementRequestPath, nil, ""),
			httpx.MustNewRequest("POST", publicTS.URL+profile.PublicProfileManagementUpdatePath, strings.NewReader(url.Values{"foo": {"bar"}}.Encode()), "application/x-www-form-urlencoded"),
			httpx.MustNewRequest("POST", publicTS.URL+profile.PublicProfileManagementUpdatePath, strings.NewReader(`{"foo":"bar"}`), "application/json"),
		} {
			t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
				res, err := http.DefaultClient.Do(tc)
				require.NoError(t, err)
				assert.EqualValues(t, http.StatusUnauthorized, res.StatusCode)
			})
		}

		for name, daemon := range map[string]struct {
			statusCode int
			url        string
		}{
			"public": {statusCode: 401, url: publicTS.URL},
			"admin":  {statusCode: 404, url: adminTS.URL},
		} {
			t.Run("daemon="+name, func(t *testing.T) {
				for k, tc := range []*http.Request{
					httpx.MustNewRequest("GET", daemon.url+profile.PublicProfileManagementRequestPath+"?request=1234", nil, ""),
				} {
					t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
						res, err := http.DefaultClient.Do(tc)
						require.NoError(t, err)
						assert.EqualValues(t, daemon.statusCode, res.StatusCode)
					})
				}
			})
		}
	})

	t.Run("daemon=public/description=fetching a non-existent request should return a 403 error", func(t *testing.T) {
		_, err := publicClient.Common.GetSelfServiceBrowserProfileManagementRequest(
			common.NewGetSelfServiceBrowserProfileManagementRequestParams().WithHTTPClient(otherUser).WithRequest("i-do-not-exist"),
		)
		require.Error(t, err)

		require.IsType(t, &common.GetSelfServiceBrowserProfileManagementRequestForbidden{}, err)
		assert.Equal(t, int64(http.StatusForbidden), err.(*common.GetSelfServiceBrowserProfileManagementRequestForbidden).Payload.Error.Code)
	})

	t.Run("daemon=admin/description=fetching a non-existent request should return a 404 error", func(t *testing.T) {
		_, err := adminClient.Common.GetSelfServiceBrowserProfileManagementRequest(
			common.NewGetSelfServiceBrowserProfileManagementRequestParams().WithHTTPClient(otherUser).WithRequest("i-do-not-exist"),
		)
		require.Error(t, err)

		require.IsType(t, &common.GetSelfServiceBrowserProfileManagementRequestNotFound{}, err)
		assert.Equal(t, int64(http.StatusNotFound), err.(*common.GetSelfServiceBrowserProfileManagementRequestNotFound).Payload.Error.Code)
	})

	t.Run("description=should fail to fetch request if identity changed", func(t *testing.T) {
		res, err := primaryUser.Get(publicTS.URL + profile.PublicProfileManagementPath)
		require.NoError(t, err)

		rid := res.Request.URL.Query().Get("request")
		require.NotEmpty(t, rid)

		_, err = publicClient.Common.GetSelfServiceBrowserProfileManagementRequest(
			common.NewGetSelfServiceBrowserProfileManagementRequestParams().WithHTTPClient(otherUser).WithRequest(rid),
		)
		require.Error(t, err)
		require.IsType(t, &common.GetSelfServiceBrowserProfileManagementRequestForbidden{}, err)
		assert.EqualValues(t, int64(http.StatusForbidden), err.(*common.GetSelfServiceBrowserProfileManagementRequestForbidden).Payload.Error.Code, "should return a 403 error because the identities from the cookies do not match")
	})

	t.Run("description=should fail to post data if CSRF is missing", func(t *testing.T) {
		rs := makeRequest(t)
		f := rs.Payload.Form
		res, err := primaryUser.PostForm(f.Action, url.Values{})
		require.NoError(t, err)
		assert.EqualValues(t, 400, res.StatusCode, "should return a 400 error because CSRF token is not set")
	})

	t.Run("description=should redirect to profile management ui and /profiles/requests?request=... should come back with the right information", func(t *testing.T) {
		res, err := primaryUser.Get(publicTS.URL + profile.PublicProfileManagementPath)
		require.NoError(t, err)

		assert.Equal(t, ui.URL, res.Request.URL.Scheme+"://"+res.Request.URL.Host)
		assert.Equal(t, "/profile", res.Request.URL.Path, "should end up at the profile URL")

		rid := res.Request.URL.Query().Get("request")
		require.NotEmpty(t, rid)

		pr, err := publicClient.Common.GetSelfServiceBrowserProfileManagementRequest(
			common.NewGetSelfServiceBrowserProfileManagementRequestParams().WithHTTPClient(primaryUser).WithRequest(rid),
		)
		require.NoError(t, err, "%s", rid)

		assert.Equal(t, rid, string(pr.Payload.ID))
		assert.NotEmpty(t, pr.Payload.Identity)
		assert.Equal(t, primaryIdentity.ID.String(), string(pr.Payload.Identity.ID))
		assert.JSONEq(t, string(primaryIdentity.Traits), x.MustEncodeJSON(t, pr.Payload.Identity.Traits))
		assert.Equal(t, primaryIdentity.TraitsSchemaID, pr.Payload.Identity.TraitsSchemaID)
		assert.Equal(t, publicTS.URL+profile.PublicProfileManagementPath, pr.Payload.RequestURL)

		found := false
		for i := range pr.Payload.Form.Fields {
			if pr.Payload.Form.Fields[i].Name == form.CSRFTokenName {
				found = true
				require.NotEmpty(t, pr.Payload.Form.Fields[i])
				pr.Payload.Form.Fields = append(pr.Payload.Form.Fields[:i], pr.Payload.Form.Fields[i+1:]...)
				break
			}
		}
		require.True(t, found)

		assert.Equal(t, &models.Form{
			Action: publicTS.URL + profile.PublicProfileManagementUpdatePath + "?request=" + rid,
			Method: "POST",
			Fields: models.FormFields{
				&models.FormField{Name: "traits.booly", Required: false, Type: "checkbox", Value: false},
				&models.FormField{Name: "traits.email", Required: false, Type: "text", Value: "john@doe.com"},
				&models.FormField{Name: "traits.numby", Required: false, Type: "number", Value: json.Number("2.5")},
				&models.FormField{Name: "traits.stringy", Required: false, Type: "text", Value: "foobar"},
			},
		}, pr.Payload.Form)
	})

	submitForm := func(t *testing.T, req *common.GetSelfServiceBrowserProfileManagementRequestOK, values url.Values) (string, *common.GetSelfServiceBrowserProfileManagementRequestOK) {
		res, err := primaryUser.PostForm(req.Payload.Form.Action, values)
		require.NoError(t, err)
		assert.EqualValues(t, http.StatusNoContent, res.StatusCode)

		assert.Equal(t, ui.URL, res.Request.URL.Scheme+"://"+res.Request.URL.Host)
		assert.Equal(t, "/profile", res.Request.URL.Path, "should end up at the profile URL")

		rs, err := publicClient.Common.GetSelfServiceBrowserProfileManagementRequest(
			common.NewGetSelfServiceBrowserProfileManagementRequestParams().WithHTTPClient(primaryUser).
				WithRequest(res.Request.URL.Query().Get("request")),
		)
		require.NoError(t, err)
		body, err := json.Marshal(rs.Payload)
		require.NoError(t, err)
		return string(body), rs
	}

	t.Run("description=should come back with form errors if some profile data is invalid", func(t *testing.T) {
		rs := makeRequest(t)
		values := fieldsToURLValues(rs.Payload.Form.Fields)
		values.Set("traits.should_long_string", "too-short")
		values.Set("traits.stringy", "bazbar") // it should still override new values!
		actual, _ := submitForm(t, rs, values)

		assert.NotEmpty(t, gjson.Get(actual, "form.fields.#(name==csrf_token).value").String(), "%s", actual)
		assert.Equal(t, "too-short", gjson.Get(actual, "form.fields.#(name==traits.should_long_string).value").String(), "%s", actual)
		assert.Equal(t, "bazbar", gjson.Get(actual, "form.fields.#(name==traits.stringy).value").String(), "%s", actual)
		assert.Equal(t, "2.5", gjson.Get(actual, "form.fields.#(name==traits.numby).value").String(), "%s", actual)
		assert.Equal(t, "traits.should_long_string: String length must be greater than or equal to 25", gjson.Get(actual, "form.fields.#(name==traits.should_long_string).errors.0.message").String(), "%s", actual)
	})

	t.Run("description=should come back with form errors if trying to update email", func(t *testing.T) {
		rs := makeRequest(t)
		values := fieldsToURLValues(rs.Payload.Form.Fields)
		values.Set("traits.email", "not-john-doe")
		res, err := primaryUser.PostForm(rs.Payload.Form.Action, values)
		require.NoError(t, err)
		defer res.Body.Close()

		assert.Contains(t, res.Request.URL.String(), errTs.URL)
		assert.EqualValues(t, http.StatusOK, res.StatusCode)

		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)
		assert.Contains(t, gjson.Get(string(body), "0.reason").String(), "A field was modified that updates one or more credentials-related settings", "%s", body)
	})

	t.Run("description=should retry with invalid payloads multiple times before succeeding", func(t *testing.T) {
		t.Run("flow=fail first update", func(t *testing.T) {
			rs := makeRequest(t)
			values := fieldsToURLValues(rs.Payload.Form.Fields)
			values.Set("traits.should_big_number", "1")
			actual, response := submitForm(t, rs, values)
			assert.False(t, response.Payload.UpdateSuccessful, "%s", actual)

			assert.Equal(t, "1", gjson.Get(actual, "form.fields.#(name==traits.should_big_number).value").String(), "%s", actual)
			assert.Equal(t, "traits.should_big_number: Must be greater than or equal to 1200", gjson.Get(actual, "form.fields.#(name==traits.should_big_number).errors.0.message").String(), "%s", actual)

			assert.Equal(t, "foobar", gjson.Get(actual, "form.fields.#(name==traits.stringy).value").String(), "%s", actual) // sanity check if original payload is still here
		})

		t.Run("flow=fail second update", func(t *testing.T) {
			rs := makeRequest(t)
			values := fieldsToURLValues(rs.Payload.Form.Fields)
			values.Del("traits.should_big_number")
			values.Set("traits.should_long_string", "short")
			values.Set("traits.numby", "this-is-not-a-number")
			actual, response := submitForm(t, rs, values)
			assert.False(t, response.Payload.UpdateSuccessful, "%s", actual)

			assert.Empty(t, gjson.Get(actual, "form.fields.#(name==traits.should_big_number).errors.0.message").String(), "%s", actual)
			assert.Empty(t, gjson.Get(actual, "form.fields.#(name==traits.should_big_number).value").String(), "%s", actual)

			assert.Equal(t, "short", gjson.Get(actual, "form.fields.#(name==traits.should_long_string).value").String(), "%s", actual)
			assert.Equal(t, "traits.should_long_string: String length must be greater than or equal to 25", gjson.Get(actual, "form.fields.#(name==traits.should_long_string).errors.0.message").String(), "%s", actual)

			assert.Equal(t, "this-is-not-a-number", gjson.Get(actual, "form.fields.#(name==traits.numby).value").String(), "%s", actual)
			assert.Equal(t, "traits.numby: Invalid type. Expected: number, given: string", gjson.Get(actual, "form.fields.#(name==traits.numby).errors.0.message").String(), "%s", actual)

			assert.Equal(t, "foobar", gjson.Get(actual, "form.fields.#(name==traits.stringy).value").String(), "%s", actual) // sanity check if original payload is still here
		})

		t.Run("flow=succeed with final request", func(t *testing.T) {
			rs := makeRequest(t)
			values := fieldsToURLValues(rs.Payload.Form.Fields)
			values.Set("traits.email", "john@doe.com")
			values.Set("traits.numby", "15")
			values.Set("traits.should_big_number", "9001")
			values.Set("traits.should_long_string", "this is such a long string, amazing stuff!")
			actual, response := submitForm(t, rs, values)
			assert.True(t, response.Payload.UpdateSuccessful, "%s", actual)

			assert.Empty(t, gjson.Get(actual, "form.fields.#(name==traits.numby).errors").Value(), "%s", actual)
			assert.Empty(t, gjson.Get(actual, "form.fields.#(name==traits.should_big_number).errors").Value(), "%s", actual)
			assert.Empty(t, gjson.Get(actual, "form.fields.#(name==traits.should_long_string).errors").Value(), "%s", actual)

			assert.Equal(t, 15.0, gjson.Get(actual, "form.fields.#(name==traits.numby).value").Value(), "%s", actual)
			assert.Equal(t, 9001.0, gjson.Get(actual, "form.fields.#(name==traits.should_big_number).value").Value(), "%s", actual)
			assert.Equal(t, "this is such a long string, amazing stuff!", gjson.Get(actual, "form.fields.#(name==traits.should_long_string).value").Value(), "%s", actual)

			assert.Equal(t, "foobar", gjson.Get(actual, "form.fields.#(name==traits.stringy).value").String(), "%s", actual) // sanity check if original payload is still here
		})

		t.Run("flow=try another update with invalid data", func(t *testing.T) {
			rs := makeRequest(t)
			values := fieldsToURLValues(rs.Payload.Form.Fields)
			values.Set("traits.should_long_string", "short")
			actual, response := submitForm(t, rs, values)
			assert.False(t, response.Payload.UpdateSuccessful, "%s", actual)
		})
	})
}
