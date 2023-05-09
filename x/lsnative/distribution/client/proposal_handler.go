package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/merlin-network/elysium-sdk/v2/x/lsnative/distribution/client/cli"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal)
)
