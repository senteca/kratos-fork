package x

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertEqualTime(t *testing.T, expected, actual time.Time) {
	assert.EqualValues(t, expected.UTC().Round(time.Second), actual.UTC().Round(time.Second))
}

func RequireEqualTime(t *testing.T, expected, actual time.Time) {
	require.EqualValues(t, expected.UTC().Round(time.Second), actual.UTC().Round(time.Second))
}
