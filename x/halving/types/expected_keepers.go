/*
 Copyright [2019] - [2021], ELYSIUM TECHNOLOGIES PTE. LTD. and the elysiumCore contributors
 SPDX-License-Identifier: Apache-2.0
*/

package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	mintTypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

type MintKeeper interface {
	GetParams(ctx sdk.Context) (params mintTypes.Params)
	SetParams(ctx sdk.Context, params mintTypes.Params)
}
