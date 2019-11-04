package selfservice_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/nosurf"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"github.com/ory/viper"

	"github.com/ory/hive/driver"
	"github.com/ory/hive/driver/configuration"
	"github.com/ory/hive/identity"
	"github.com/ory/hive/internal"
	. "github.com/ory/hive/selfservice"
	"github.com/ory/hive/selfservice/oidc"
	"github.com/ory/hive/selfservice/password"
	"github.com/ory/hive/session"
	"github.com/ory/hive/x"
)

func newErrTs(t *testing.T, reg driver.Registry) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		e, err := reg.ErrorManager().Read(r.URL.Query().Get("error"))
		require.NoError(t, err)
		reg.Writer().Write(w, r, e)
	}))
}

func TestLogoutHandler(t *testing.T) {
	_, reg := internal.NewMemoryRegistry(t)
	handler := reg.StrategyHandler()

	router := x.NewRouterPublic()
	handler.RegisterPublicRoutes(router)
	reg.WithCSRFHandler(x.NewCSRFHandler(router, reg.Writer(), logrus.New(), "/", "", false))
	ts := httptest.NewServer(reg.CSRFHandler())
	defer ts.Close()

	var sess session.Session
	sess.SID = uuid.New().String()
	sess.Identity = new(identity.Identity)
	require.NoError(t, reg.SessionManager().Create(context.Background(), &sess))

	router.GET("/set", session.MockSetSession(t, reg))

	router.GET("/csrf", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		_, _ = w.Write([]byte(nosurf.Token(r)))
	})

	redirTS := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer redirTS.Close()

	viper.Set(configuration.ViperKeyDefaultIdentityTraitsSchemaURL, "file://./stub/registration.schema.json")
	viper.Set(configuration.ViperKeySelfServiceLogoutRedirectURL, redirTS.URL)
	viper.Set(configuration.ViperKeyURLsSelfPublic, ts.URL)

	client := session.MockCookieClient(t)

	t.Run("case=set initial session", func(t *testing.T) {
		session.MockHydrateCookieClient(t, client, ts.URL+"/set")
	})

	var token string
	t.Run("case=get csrf token", func(t *testing.T) {
		res, err := ts.Client().Get(ts.URL + "/csrf")
		require.NoError(t, err)
		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)
		require.NoError(t, res.Body.Close())
		token = string(body)
		require.NotEmpty(t, token)
	})

	t.Run("case=log out", func(t *testing.T) {
		res, err := client.Get(ts.URL + BrowserLogoutPath)
		require.NoError(t, err)

		var found bool
		for _, c := range res.Cookies() {
			if c.Name == session.DefaultSessionCookieName {
				found = true
			}
		}
		require.False(t, found)
	})

	t.Run("case=csrf token should be reset", func(t *testing.T) {
		res, err := ts.Client().Get(ts.URL + "/csrf")
		require.NoError(t, err)
		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)
		require.NoError(t, res.Body.Close())
		require.NotEmpty(t, body)
		assert.NotEqual(t, token, string(body))
	})

}

func TestEnsureSessionRedirect(t *testing.T) {
	_, reg := internal.NewMemoryRegistry(t)
	handler := reg.StrategyHandler()

	router := x.NewRouterPublic()
	handler.RegisterPublicRoutes(router)
	ts := httptest.NewServer(router)
	defer ts.Close()

	redirTS := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("already authenticated"))
	}))
	defer redirTS.Close()

	viper.Set(configuration.ViperKeyURLsDefaultReturnTo, redirTS.URL)
	viper.Set(configuration.ViperKeyURLsSelfPublic, ts.URL)
	viper.Set(configuration.ViperKeyDefaultIdentityTraitsSchemaURL, "file://./stub/registration.schema.json")

	for k, tc := range [][]string{
		{"GET", BrowserLoginPath},
		{"GET", BrowserRegistrationPath},

		{"POST", password.LoginPath},
		{"POST", password.RegistrationPath},

		// it is ok that these contain the parameters as arw strings as we are only interested in checking if the middleware is working
		{"POST", oidc.AuthPath},
		{"GET", oidc.AuthPath},
		{"GET", oidc.CallbackPath},
	} {
		t.Run(fmt.Sprintf("case=%d/method=%s/path=%s", k, tc[0], tc[1]), func(t *testing.T) {
			body, _ := session.MockMakeAuthenticatedRequest(t, reg, router.Router, x.NewTestHTTPRequest(t, tc[0], ts.URL+tc[1], nil))
			assert.EqualValues(t, "already authenticated", string(body))
		})
	}
}

func TestLoginHandler(t *testing.T) {
	_, reg := internal.NewMemoryRegistry(t)
	handler := reg.StrategyHandler()

	router := x.NewRouterPublic()
	handler.RegisterPublicRoutes(router)
	ts := httptest.NewServer(nosurf.New(router))
	defer ts.Close()

	redirTS := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer redirTS.Close()

	loginTS := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get(ts.URL + BrowserLoginRequestsPath + "?request=" + r.URL.Query().Get("request"))
		require.NoError(t, err)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)
		_, _ = w.Write(body)
	}))
	defer loginTS.Close()

	errTS := newErrTs(t, reg)
	defer errTS.Close()

	viper.Set(configuration.ViperKeyURLsLogin, loginTS.URL)
	viper.Set(configuration.ViperKeyURLsSelfPublic, ts.URL)
	viper.Set(configuration.ViperKeyURLsError, errTS.URL)

	for k := range []struct {
	}{
		{},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			res, err := ts.Client().Get(ts.URL + BrowserLoginPath)
			require.NoError(t, err)

			defer res.Body.Close()
			require.Equal(t, http.StatusOK, res.StatusCode)

			body, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)

			assert.Equal(t, "password", gjson.GetBytes(body, "methods.password.method").String(), "%s", body)
			assert.NotEmpty(t, gjson.GetBytes(body, "methods.password.config.fields.csrf_token.value").String(), "%s", body)
			assert.NotEmpty(t, gjson.GetBytes(body, "id").String(), "%s", body)
			assert.Contains(t, gjson.GetBytes(body, "methods.password.config.action").String(), gjson.GetBytes(body, "id").String(), "%s", body)
			assert.Contains(t, gjson.GetBytes(body, "methods.password.config.action").String(), ts.URL, "%s", body)
		})
	}
}

func TestRegistrationHandler(t *testing.T) {
	_, reg := internal.NewMemoryRegistry(t)
	handler := reg.StrategyHandler()

	router := x.NewRouterPublic()
	handler.RegisterPublicRoutes(router)
	ts := httptest.NewServer(nosurf.New(router))
	defer ts.Close()

	redirTS := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer redirTS.Close()

	loginTS := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get(ts.URL + BrowserRegistrationRequestsPath + "?request=" + r.URL.Query().Get("request"))
		require.NoError(t, err)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)
		_, _ = w.Write(body)
	}))
	defer loginTS.Close()

	errTS := newErrTs(t, reg)
	defer errTS.Close()

	viper.Set(configuration.ViperKeyURLsRegistration, loginTS.URL)
	viper.Set(configuration.ViperKeyURLsSelfPublic, ts.URL)
	viper.Set(configuration.ViperKeyURLsError, errTS.URL)

	for k := range []struct {
	}{
		{},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			res, err := ts.Client().Get(ts.URL + BrowserRegistrationPath)
			require.NoError(t, err)
			defer res.Body.Close()
			require.Equal(t, http.StatusOK, res.StatusCode)

			body, err := ioutil.ReadAll(res.Body)
			require.NoError(t, err)

			assert.Equal(t, "password", gjson.GetBytes(body, "methods.password.method").String(), "%s", body)
			assert.NotEmpty(t, gjson.GetBytes(body, "methods.password.config.fields.csrf_token.value").String(), "%s", body)
			assert.NotEmpty(t, gjson.GetBytes(body, "id").String(), "%s", body)
			assert.Contains(t, gjson.GetBytes(body, "methods.password.config.action").String(), gjson.GetBytes(body, "id").String(), "%s", body)
			assert.Contains(t, gjson.GetBytes(body, "methods.password.config.action").String(), ts.URL, "%s", body)
		})
	}
}
