package profile_test

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/x/pointerx"

	"github.com/ory/viper"
	"github.com/ory/x/urlx"

	"github.com/ory/kratos/driver/configuration"
	"github.com/ory/kratos/identity"
	"github.com/ory/kratos/internal"
	"github.com/ory/kratos/internal/httpclient/client/common"
	"github.com/ory/kratos/internal/testhelpers"
	"github.com/ory/kratos/selfservice/flow/login"
	"github.com/ory/kratos/selfservice/flow/profile"
	"github.com/ory/kratos/session"
	"github.com/ory/kratos/x"
)

func init() {
	internal.RegisterFakes()
}

func TestHandler(t *testing.T) {
	_, reg := internal.NewRegistryDefault(t)
	viper.Set(configuration.ViperKeyDefaultIdentityTraitsSchemaURL, "file://./stub/identity.schema.json")

	_ = testhelpers.NewProfileUITestServer(t)
	_ = testhelpers.NewErrorTestServer(t, reg)

	viper.Set(configuration.ViperKeySelfServicePrivilegedAuthenticationAfter, "1ns")
	primaryIdentity := &identity.Identity{ID: x.NewUUID(), Traits: identity.Traits(`{}`)}
	publicTS, adminTS := testhelpers.NewProfileAPIServer(t, reg, []identity.Identity{
		*primaryIdentity, {ID: x.NewUUID(), Traits: identity.Traits(`{}`)}})

	primaryUser, otherUser := testhelpers.NewSessionClient(t, publicTS.URL+"/sessions/set/0"),
		testhelpers.NewSessionClient(t, publicTS.URL+"/sessions/set/1")
	publicClient, adminClient := testhelpers.NewSDKClient(publicTS), testhelpers.NewSDKClient(adminTS)
	newExpiredRequest := func() *profile.Request {
		return profile.NewRequest(-time.Minute,
			&http.Request{URL: urlx.ParseOrPanic(publicTS.URL + login.BrowserLoginPath)},
			&session.Session{Identity: primaryIdentity})
	}

	t.Run("daemon=admin", func(t *testing.T) {
		t.Run("description=fetching a non-existent request should return a 404 error", func(t *testing.T) {
			_, err := adminClient.Common.GetSelfServiceBrowserProfileManagementRequest(
				common.NewGetSelfServiceBrowserProfileManagementRequestParams().WithHTTPClient(otherUser).WithRequest("i-do-not-exist"),
			)
			require.Error(t, err)

			require.IsType(t, &common.GetSelfServiceBrowserProfileManagementRequestNotFound{}, err)
			assert.Equal(t, int64(http.StatusNotFound), err.(*common.GetSelfServiceBrowserProfileManagementRequestNotFound).Payload.Error.Code)
		})

		t.Run("description=fetching an expired request returns 410", func(t *testing.T) {
			pr := newExpiredRequest()
			require.NoError(t, reg.ProfileRequestPersister().CreateProfileRequest(context.Background(), pr))

			_, err := adminClient.Common.GetSelfServiceBrowserProfileManagementRequest(
				common.NewGetSelfServiceBrowserProfileManagementRequestParams().WithHTTPClient(primaryUser).WithRequest(pr.ID.String()),
			)
			require.Error(t, err)

			require.IsType(t, &common.GetSelfServiceBrowserProfileManagementRequestGone{}, err, "%+v", err)
			assert.Equal(t, int64(http.StatusGone), err.(*common.GetSelfServiceBrowserProfileManagementRequestGone).Payload.Error.Code)
		})
	})

	t.Run("daemon=public", func(t *testing.T) {
		t.Run("description=fetching a non-existent request should return a 403 error", func(t *testing.T) {
			_, err := publicClient.Common.GetSelfServiceBrowserProfileManagementRequest(
				common.NewGetSelfServiceBrowserProfileManagementRequestParams().WithHTTPClient(otherUser).WithRequest("i-do-not-exist"),
			)
			require.Error(t, err)

			require.IsType(t, &common.GetSelfServiceBrowserProfileManagementRequestForbidden{}, err)
			assert.Equal(t, int64(http.StatusForbidden), err.(*common.GetSelfServiceBrowserProfileManagementRequestForbidden).Payload.Error.Code)
		})

		t.Run("description=fetching an expired request returns 410", func(t *testing.T) {
			pr := newExpiredRequest()
			require.NoError(t, reg.ProfileRequestPersister().CreateProfileRequest(context.Background(), pr))

			_, err := publicClient.Common.GetSelfServiceBrowserProfileManagementRequest(
				common.NewGetSelfServiceBrowserProfileManagementRequestParams().WithHTTPClient(primaryUser).WithRequest(pr.ID.String()),
			)
			require.Error(t, err)

			require.IsType(t, &common.GetSelfServiceBrowserProfileManagementRequestGone{}, err)
			assert.Equal(t, int64(http.StatusGone), err.(*common.GetSelfServiceBrowserProfileManagementRequestGone).Payload.Error.Code)
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
			f := testhelpers.GetProfileManagementRequestMethodConfig(t, primaryUser, publicTS, profile.StrategyTraitsID)
			res, err := primaryUser.PostForm(pointerx.StringR(f.Action), url.Values{})
			require.NoError(t, err)
			assert.EqualValues(t, 400, res.StatusCode, "should return a 400 error because CSRF token is not set")
		})
	})
}
