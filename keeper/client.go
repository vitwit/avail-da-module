package keeper

import (
	cometrpc "github.com/cometbft/cometbft/rpc/client/http"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authTx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/go-bip39"

	// "github.com/tendermint/starport/starport/pkg/xfilepath"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

const (
	defaultGasAdjustment = 1.0
	defaultGasLimit      = 300000
)

// var availdHomePath = xfilepath.JoinFromHome(xfilepath.Path("availsdk"))

func NewClientCtx(kr keyring.Keyring, c *cometrpc.HTTP, chainID string, cdc codec.BinaryCodec) client.Context {
	encodingConfig := MakeEncodingConfig()
	// authtypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	// cryptocodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	// sdk.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	// staking.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	// cryptocodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	// chainID := ctx.ChainID()

	// fmt.Println("address heree......", address)
	// sdk.AccAddressFromBech32()
	// fromAddress, err := sdk.AccAddressFromBech32(kr.FromAddress)
	// fmt.Println("here errorr...", err)
	// if err != nil {
	// 	// return err
	// }
	// fmt.Println("from addresss.........", fromAddress)
	// Assuming you have access to the keyring and broadcast mode
	// broadcastMode := "block"
	broadcastMode := flags.BroadcastSync

	homepath := "/home/vitwit/.availsdk/keyring-test"

	return client.Context{}.
		WithCodec(cdc.(codec.Codec)).
		WithChainID(chainID).
		// WithFromAddress(kr.fromAddress).
		WithFromName("testkey").
		WithKeyringDir(homepath).
		WithBroadcastMode(broadcastMode).
		WithTxConfig(authTx.NewTxConfig(cdc.(codec.Codec), authTx.DefaultSignModes)).
		WithKeyring(kr).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithClient(c).WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithSkipConfirmation(true)
}

// NewFactory creates a new Factory.
func NewFactory(clientCtx client.Context) tx.Factory {
	return tx.Factory{}.
		WithChainID(clientCtx.ChainID).
		WithKeybase(clientCtx.Keyring).
		WithGas(defaultGasLimit).
		WithGasAdjustment(defaultGasAdjustment).
		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithTxConfig(clientCtx.TxConfig)
}

// MakeEncodingConfig creates an EncodingConfig for an amino based test configuration.
func MakeEncodingConfig() EncodingConfig {
	aminoCodec := codec.NewLegacyAmino()
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)
	txCfg := authTx.NewTxConfig(codec, authTx.DefaultSignModes)

	encCfg := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          txCfg,
		Amino:             aminoCodec,
	}

	std.RegisterLegacyAminoCodec(encCfg.Amino)
	std.RegisterInterfaces(encCfg.InterfaceRegistry)
	// mb.RegisterLegacyAminoCodec(encCfg.Amino)
	// mb.RegisterInterfaces(encCfg.InterfaceRegistry)

	return encCfg
}

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

// ImportMnemonic is to import existing account mnemonic in keyring
func ImportMnemonic(keyName, mnemonic, hdPath string, c client.Context) (*keyring.Record, error) {
	info, err := AccountCreate(keyName, mnemonic, hdPath, c) // return account also
	// fmt.Println("here the accc details.......", keyName, mnemonic, hdPath)
	if err != nil {
		return nil, err
	}

	return info, nil
}

// AccountCreate creates an account by name and mnemonic (optional) in the keyring.
func AccountCreate(accountName, mnemonic, hdPath string, c client.Context) (*keyring.Record, error) {
	if mnemonic == "" {
		entropySeed, err := bip39.NewEntropy(256)
		if err != nil {
			return nil, err
		}
		mnemonic, err = bip39.NewMnemonic(entropySeed)
		// fmt.Println("mnemoniccccc here.....", mnemonic)
		if err != nil {
			return nil, err
		}
	}

	algos, _ := c.Keyring.SupportedAlgorithms()
	algo, err := keyring.NewSigningAlgoFromString(string(hd.Secp256k1Type), algos)
	if err != nil {
		return nil, err
	}

	path := hd.CreateHDPath(118, 0, 0).String()
	// fmt.Println("pathhh......", path)

	// record, str, err := c.Keyring.NewMnemonic("test_key1", keyring.English, path, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	// fmt.Println("recorddddd.......", err, str, record)

	// k, _, err = kb.NewMnemonic("test", English, types.FullFundraiserPath, DefaultBIP39Passphrase, hd.Secp256k1)
	info, err := c.Keyring.NewAccount(accountName, mnemonic, keyring.DefaultBIP39Passphrase, path, algo)
	// fmt.Println("after creationnnn.........", info, err)
	if err != nil {
		return nil, err
	}
	// pk, err := info.GetPubKey()
	// if err != nil {
	// 	return err
	// }
	// addr := sdk.AccAddress(pk.Address())
	// fmt.Println("address hereee...", addr)
	// account := c.ToAccount(info)
	// account.Mnemonic = mnemonic
	return info, nil
	// return nil
}
