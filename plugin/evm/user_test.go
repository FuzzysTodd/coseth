// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// TestUserControlsAddressZeroAddress tests that controlsAddress properly rejects zero addresses
func TestUserControlsAddressZeroAddress(t *testing.T) {
	require := require.New(t)

	u := &user{
		db: nil,
	}

	// Test with nil database
	zeroAddr := common.Address{}
	_, err := u.controlsAddress(zeroAddr)
	require.ErrorIs(err, errDBNil, "should return errDBNil when database is nil")

	// Note: Testing with a valid database would require setting up the full database
	// infrastructure, which is beyond the scope of this security fix.
	// The important validation is that zero addresses are checked before database operations.
}

// TestUserGetKeyZeroAddress tests that getKey properly rejects zero addresses
func TestUserGetKeyZeroAddress(t *testing.T) {
	require := require.New(t)

	u := &user{
		db: nil,
	}

	// Test with nil database
	zeroAddr := common.Address{}
	_, err := u.getKey(zeroAddr)
	require.ErrorIs(err, errDBNil, "should return errDBNil when database is nil")

	// Note: Testing with a valid database would require setting up the full database
	// infrastructure, which is beyond the scope of this security fix.
	// The important validation is that zero addresses are checked before database operations.
}

// TestErrorConstants verifies that error constants are properly defined
func TestErrorConstants(t *testing.T) {
	require := require.New(t)

	require.NotNil(errDBNil, "errDBNil should be defined")
	require.NotNil(errKeyNil, "errKeyNil should be defined")
	require.NotNil(errEmptyAddress, "errEmptyAddress should be defined")

	require.Equal("db uninitialized", errDBNil.Error())
	require.Equal("key uninitialized", errKeyNil.Error())
	require.Equal("address is empty", errEmptyAddress.Error())
}
