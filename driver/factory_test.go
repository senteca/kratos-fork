package driver_test

import (
	"context"
	"testing"

	"github.com/ory/x/configx"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	driver "github.com/ory/kratos/driver"
	"github.com/ory/kratos/driver/config"
)

func TestDriverNew(t *testing.T) {
	r := driver.New(
		context.Background(),
		configx.WithValue(config.ViperKeyDSN, config.DefaultSQLiteMemoryDSN),
		configx.SkipValidation())

	assert.EqualValues(t, config.DefaultSQLiteMemoryDSN, r.Config(context.Background()).DSN())
	require.NoError(t, r.Persister().Ping())
}
