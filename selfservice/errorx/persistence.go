package errorx

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/herodot"

	"github.com/ory/kratos/x"
)

type (
	Persister interface {
		// Add adds an error to the manager and returns a unique identifier or an error if insertion fails.
		Add(ctx context.Context, errs ...error) (uuid.UUID, error)

		// Read returns an error by its unique identifier and marks the error as read. If an error occurs during retrieval
		// the second return parameter is an error.
		Read(ctx context.Context, id uuid.UUID) ([]json.RawMessage, error)

		// Clear clears read containers that are older than a certain amount of time. If force is set to true, unread
		// errors will be cleared as well.
		Clear(ctx context.Context, olderThan time.Duration, force bool) error
	}

	PersistenceProvider interface {
		SelfServiceErrorPersister() Persister
	}
)

func TestPersister(p Persister) func(t *testing.T) {
	toJSON := func(t *testing.T, in interface{}) string {
		out, err := json.Marshal(in)
		require.NoError(t, err)
		return string(out)
	}

	return func(t *testing.T) {
		t.Run("case=not found", func(t *testing.T) {
			_, err := p.Read(context.Background(), x.NewUUID())
			require.Error(t, err)
		})

		t.Run("case=en- and decode properly", func(t *testing.T) {
			expected := herodot.ErrNotFound.WithReason("foobar")
			actualID, err := p.Add(context.Background(), expected)
			require.NoError(t, err)

			actual, err := p.Read(context.Background(), actualID)
			require.NoError(t, err)

			assert.JSONEq(t, toJSON(t, []error{expected}), toJSON(t, actual))
		})

		t.Run("case=clear", func(t *testing.T) {
			expected := herodot.ErrNotFound.WithReason("foobar")
			actualID, err := p.Add(context.Background(), expected)
			require.NoError(t, err)

			_, err = p.Read(context.Background(), actualID)
			require.NoError(t, err)

			time.Sleep(time.Millisecond * 100)

			require.NoError(t, p.Clear(context.Background(), time.Millisecond, false))
			_, err = p.Read(context.Background(), actualID)
			require.Error(t, err)
		})
	}
}
