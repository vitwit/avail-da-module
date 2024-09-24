package types

import (
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/spf13/cast"
)

// AvailConfig defines the configuration for the in-process Avail relayer.
type AvailConfiguration struct {
	// avail light node url
	LightClientURL string `mapstructure:"light-client-url"`

	// avail chain ID
	// ChainID string `mapstructure:"chain-id"`

	// Overrides built-in app-id used
	AppID int `mapstructure:"app-id"`

	// avail config
	Seed string `json:"seed"`

	// RPC of the cosmos node to fetch the block data
	CosmosNodeRPC string `json:"cosmos-node-rpc"`

	MaxBlobBlocks uint64 `json:"max-blob-blocks"`

	PublishBlobInterval uint64 `json:"publish-blob-interval"`

	VoteInterval uint64 `json:"vote-interval"`
}

const (
	FlagChainID             = "avail.chain-id"
	FlagOverrideAppID       = "avail.override-app-id"
	FlagOverridePubInterval = "avail.override-pub-interval"
	FlagQueryInterval       = "avail.proof-query-interval"
	FlagSeed                = "avail-seed"
	FlagLightClientURL      = "avail.light-client-url"
	FlagCosmosNodeRPC       = "avail.cosmos-node-rpc"
	FlagMaxBlobBlocks       = "avail.max-blob-blocks"
	FlagPublishBlobInterval = "avail.publish-blob-interval"
	FlagVoteInterval        = "avail.vote-interval"
	FlagValidatorKey        = "avail.validator-key"
	FlagCosmosNodeDir       = "avail.cosmos-node-dir"
	FlagKeyringBackendType  = "avail.keyring-backend-type"

	DefaultConfigTemplate = `

	[avail]

	# Avail light client node url for posting data
	light-client-url = "http://127.0.0.1:8000"

	# Overrides the expected  avail app-id, test-only
	override-app-id = "1"

	# Seed for avail
	seed = "bottom drive obey lake curtain smoke basket hold race lonely fit walk//Alice"

	# RPC of cosmos node to get the block data
	cosmos-node-rpc = "http://127.0.0.1:26657"

	# Maximum number of blocks over which blobs can be processed
	max-blob-blocks = 10

	# The frequency at which block data is posted to the Avail Network
	publish-blob-interval = 5

	# It is the period before validators verify whether data is truly included in
	# Avail and confirm it with the network using vote extension
	vote-interval = 5
	`
)

var DefaultAvailConfig = AvailConfiguration{
	// ChainID:             "avail-1",
	Seed:                "bottom drive obey lake curtain smoke basket hold race lonely fit walk//Alice",
	AppID:               1,
	CosmosNodeRPC:       "http://127.0.0.1:26657",
	MaxBlobBlocks:       20,
	PublishBlobInterval: 10,
	VoteInterval:        5,
	LightClientURL:      "http://127.0.0.1:8000",
	// ValidatorKey:        "alice",
	// CosmosNodeDir:          ".simapp",
}

func AvailConfigFromAppOpts(appOpts servertypes.AppOptions) AvailConfiguration {
	return AvailConfiguration{
		// ChainID:             cast.ToString(appOpts.Get(FlagChainID)),
		AppID:               cast.ToInt(appOpts.Get(FlagOverrideAppID)),
		Seed:                cast.ToString(appOpts.Get(FlagSeed)),
		LightClientURL:      cast.ToString(appOpts.Get(FlagLightClientURL)),
		CosmosNodeRPC:       cast.ToString(appOpts.Get(FlagCosmosNodeRPC)),
		MaxBlobBlocks:       cast.ToUint64(appOpts.Get(FlagMaxBlobBlocks)),
		PublishBlobInterval: cast.ToUint64(appOpts.Get(FlagPublishBlobInterval)),
		VoteInterval:        cast.ToUint64(appOpts.Get(FlagVoteInterval)),
		// ValidatorKey:        cast.ToString(appOpts.Get(FlagValidatorKey)),
		// CosmosNodeDir:          cast.ToString(appOpts.Get(FlagCosmosNodeDir)),
	}
}
