package container

import (
	"testing"

	"github.com/Manta-Network/manta-fp/testutil"

	"github.com/stretchr/testify/require"
)

// ImageConfig contains all images and their respective tags
// needed for running e2e tests.
type ImageConfig struct {
	BabylonRepository string
	BabylonVersion    string
}

//nolint:deadcode
const (
	dockerBabylondRepository = "babylonlabs/babylond"
)

// NewImageConfig returns ImageConfig needed for running e2e test.
func NewImageConfig(t *testing.T) ImageConfig {
	babylondVersion, err := testutil.GetBabylonVersion()
	require.NoError(t, err)
	return ImageConfig{
		BabylonRepository: dockerBabylondRepository,
		BabylonVersion:    babylondVersion,
	}
}
