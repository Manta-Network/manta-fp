package store_test

import (
	"math/rand"
	"testing"

	"github.com/Manta-Network/manta-fp/bbn-fp/proto"
	"github.com/Manta-Network/manta-fp/testutil"

	"github.com/stretchr/testify/require"
)

func TestShouldStart(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name           string
		currFpStatus   proto.FinalityProviderStatus
		expShouldStart bool
	}{
		{
			"Slashed: Should NOT start",
			proto.FinalityProviderStatus_SLASHED,
			false,
		},
		{
			"Inactive: Should start",
			proto.FinalityProviderStatus_INACTIVE,
			true,
		},
		{
			"Registered: Should start",
			proto.FinalityProviderStatus_REGISTERED,
			true,
		},
		{
			"Active: Should start",
			proto.FinalityProviderStatus_ACTIVE,
			true,
		},
	}

	r := rand.New(rand.NewSource(10))
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fp := testutil.GenRandomFinalityProvider(r, t)
			fp.Status = tc.currFpStatus

			shouldStart := fp.ShouldStart()
			require.Equal(t, tc.expShouldStart, shouldStart)
		})
	}
}
