/*
 Copyright [2019] - [2021], ELYSIUM TECHNOLOGIES PTE. LTD. and the elysiumCore contributors
 SPDX-License-Identifier: Apache-2.0
*/

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/merlin-network/elysium-sdk/v2/x/halving/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(context context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(context)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}
