package testslashing

import (
	"github.com/merlin-network/elysium-sdk/v2/x/lsnative/slashing/types"
)

// TestParams construct default slashing params for tests.
// Have to change these parameters for tests
// lest the tests take forever
func TestParams() types.Params {
	params := types.DefaultParams()
	params.SignedBlocksWindow = 1000
	params.DowntimeJailDuration = 60 * 60

	return params
}
