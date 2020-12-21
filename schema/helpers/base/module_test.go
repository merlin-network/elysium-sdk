package base

import (
	"encoding/json"
	clientContext "github.com/cosmos/cosmos-sdk/client/context"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkModule "github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/gorilla/mux"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	baseTestUtilities "github.com/persistenceOne/persistenceSDK/utilities/test/schema/helpers/base"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"math/rand"
	"testing"
)

var auxiliariesPrototype = func() helpers.Auxiliaries {
	return auxiliaries{[]helpers.Auxiliary{NewAuxiliary("testAuxiliary", baseTestUtilities.TestAuxiliaryKeeperPrototype)}}
}
var genesisPrototype = func() helpers.Genesis {
	return NewGenesis(baseTestUtilities.KeyPrototype, baseTestUtilities.MappablePrototype,
		[]helpers.Mappable{baseTestUtilities.NewMappable("test", "testValue")},
		[]types.Parameter{base.NewParameter(base.NewID("testParameter"), base.NewStringData("testData"), func(interface{}) error { return nil })})
}
var mapperPrototype = func() helpers.Mapper {
	return NewMapper(baseTestUtilities.KeyPrototype, baseTestUtilities.MappablePrototype)
}
var parametersPrototype = func() helpers.Parameters {
	return NewParameters(base.NewParameter(base.NewID("testParameter"), base.NewStringData("testData"), func(interface{}) error { return nil }))
}
var queriesPrototype = func() helpers.Queries {
	return queries{[]helpers.Query{NewQuery("testQuery", "q", "testQuery", "test", baseTestUtilities.TestQueryRequestPrototype,
		baseTestUtilities.TestQueryResponsePrototype, baseTestUtilities.TestQueryKeeperPrototype)}}
}
var simulatorPrototype = func() helpers.Simulator { return nil }
var transactionsPrototype = func() helpers.Transactions {
	return transactions{[]helpers.Transaction{NewTransaction("testMsg", "", "", baseTestUtilities.TestTransactionRequestPrototype, baseTestUtilities.TestMessagePrototype,
		baseTestUtilities.TestTransactionKeeperPrototype)}}
}

func TestModule(t *testing.T) {
	context, storeKey, transientStoreKey := baseTestUtilities.SetupTest(t)
	codec := baseTestUtilities.MakeCodec()
	subspace := params.NewSubspace(codec, storeKey, transientStoreKey, "test") //.WithKeyTable(parametersPrototype().GetKeyTable())
	//subspace.SetParamSet(context, parametersPrototype())
	Module := NewModule("test", auxiliariesPrototype, genesisPrototype,
		mapperPrototype, parametersPrototype, queriesPrototype, simulatorPrototype, transactionsPrototype).Initialize(storeKey, subspace).(module)

	// AppModuleBasic
	require.Equal(t, "test", Module.Name())

	// RegisterCodec
	Module.RegisterCodec(codec)

	//TODO DefaultGenesis
	//Module.DefaultGenesis()

	//TODO ValidateGenesis
	// Module.ValidateGenesis(json.RawMessage{})

	// RegisterRESTRoutes
	cliContext := clientContext.NewCLIContext().WithCodec(codec).WithChainID("test")
	router := mux.NewRouter()
	Module.RegisterRESTRoutes(cliContext, router)

	// GetTxCmd
	require.Equal(t, "test", Module.GetTxCmd(codec).Name())
	require.Equal(t, "test", Module.GetQueryCmd(codec).Name())

	//AppModule
	Module.RegisterInvariants(nil) //No return
	require.Equal(t, "test", Module.Route())

	response, Error := Module.NewHandler()(context, baseTestUtilities.NewTestMessage(sdkTypes.AccAddress("addr"), "id"))
	require.Nil(t, Error)
	require.NotNil(t, response)

	require.Equal(t, "test", Module.QuerierRoute())

	encodedRequest, Error := Module.queries.Get("testQuery").(query).requestPrototype().Encode()
	require.Nil(t, Error)

	queryResponse, Error := Module.NewQuerierHandler()(context, []string{"testQuery"}, abciTypes.RequestQuery{Data: encodedRequest})
	require.Nil(t, Error)
	require.NotNil(t, queryResponse)

	Module.BeginBlock(context, abciTypes.RequestBeginBlock{})
	Module.EndBlock(context, abciTypes.RequestEndBlock{})
	//TODO	Module.InitGenesis(context, json.RawMessage{})
	// Module.ExportGenesis(context)

	// AppModuleSimulation
	Module.GenerateGenesisState(&sdkModule.SimulationState{})
	Module.ProposalContents(sdkModule.SimulationState{})
	Module.RandomizedParams(&rand.Rand{})
	Module.RegisterStoreDecoder(sdkTypes.StoreDecoderRegistry{})
	Module.WeightedOperations(sdkModule.SimulationState{})

	//types.Module
	require.Equal(t, "testAuxiliary", Module.GetAuxiliary("testAuxiliary").GetName())
	_, Error = Module.DecodeModuleTransactionRequest("testMsg", json.RawMessage(`{"BaseReq":{"from":"addr"},"ID":"id"}`))
	require.Nil(t, Error)
}
