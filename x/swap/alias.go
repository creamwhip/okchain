package swap

import (
	"github.com/okex/okchain/x/swap/keeper"
	"github.com/okex/okchain/x/swap/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	//QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper          = keeper.NewKeeper
	NewQuerier         = keeper.NewQuerier
	RegisterCodec      = types.RegisterCodec
	NewMsgAddLiquidity = types.NewMsgAddLiquidity

	// variable aliases
	ModuleCdc = types.ModuleCdc
)

type (
	Keeper = keeper.Keeper
	Params = types.Params

	SwapTokenPair = types.SwapTokenPair
)

