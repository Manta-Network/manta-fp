package store_test

import (
	"math/rand"
	"os"
	"testing"

	"github.com/Manta-Network/manta-fp/eotsmanager/config"
	"github.com/Manta-Network/manta-fp/eotsmanager/store"
	"github.com/Manta-Network/manta-fp/testutil"

	"github.com/babylonlabs-io/babylon/testutil/datagen"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/stretchr/testify/require"
)

// FuzzEOTSStore tests save and show EOTS key names properly
func FuzzEOTSStore(f *testing.F) {
	testutil.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		t.Parallel()
		r := rand.New(rand.NewSource(seed))

		homePath := t.TempDir()
		cfg := config.DefaultDBConfigWithHomePath(homePath)

		dbBackend, err := cfg.GetDBBackend()
		require.NoError(t, err)

		vs, err := store.NewEOTSStore(dbBackend)
		require.NoError(t, err)

		defer func() {
			dbBackend.Close()
			err := os.RemoveAll(homePath)
			require.NoError(t, err)
		}()

		expectedKeyName := testutil.GenRandomHexStr(r, 10)
		_, btcPk, err := datagen.GenRandomBTCKeyPair(r)
		require.NoError(t, err)

		// add key name for the first time
		err = vs.AddEOTSKeyName(
			btcPk,
			expectedKeyName,
		)
		require.NoError(t, err)

		// add duplicate key name
		err = vs.AddEOTSKeyName(
			btcPk,
			expectedKeyName,
		)
		require.ErrorIs(t, err, store.ErrDuplicateEOTSKeyName)

		keyNameFromDB, err := vs.GetEOTSKeyName(schnorr.SerializePubKey(btcPk))
		require.NoError(t, err)
		require.Equal(t, expectedKeyName, keyNameFromDB)

		_, randomBtcPk, err := datagen.GenRandomBTCKeyPair(r)
		require.NoError(t, err)
		_, err = vs.GetEOTSKeyName(schnorr.SerializePubKey(randomBtcPk))
		require.ErrorIs(t, err, store.ErrEOTSKeyNameNotFound)
	})
}

// FuzzSignStore tests save sign records
func FuzzSignStore(f *testing.F) {
	testutil.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		t.Parallel()
		r := rand.New(rand.NewSource(seed))

		homePath := t.TempDir()
		cfg := config.DefaultDBConfigWithHomePath(homePath)

		dbBackend, err := cfg.GetDBBackend()
		require.NoError(t, err)

		vs, err := store.NewEOTSStore(dbBackend)
		require.NoError(t, err)

		defer func() {
			dbBackend.Close()
			err := os.RemoveAll(homePath)
			require.NoError(t, err)
		}()

		expectedHeight := r.Uint64()
		pk := testutil.GenRandomByteArray(r, 32)
		msg := testutil.GenRandomByteArray(r, 32)
		eotsSig := testutil.GenRandomByteArray(r, 32)

		chainID := []byte("test-chain")
		// save for the first time
		err = vs.SaveSignRecord(
			expectedHeight,
			chainID,
			msg,
			pk,
			eotsSig,
		)
		require.NoError(t, err)

		// try to save the record at the same height
		err = vs.SaveSignRecord(
			expectedHeight,
			chainID,
			msg,
			pk,
			eotsSig,
		)
		require.ErrorIs(t, err, store.ErrDuplicateSignRecord)

		signRecordFromDB, found, err := vs.GetSignRecord(pk, chainID, expectedHeight)
		require.True(t, found)
		require.NoError(t, err)
		require.Equal(t, msg, signRecordFromDB.Msg)
		require.Equal(t, eotsSig, signRecordFromDB.Signature)

		rndHeight := r.Uint64()
		_, found, err = vs.GetSignRecord(pk, chainID, rndHeight)
		require.NoError(t, err)
		require.False(t, found)
	})
}
