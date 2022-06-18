package asset

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/ngaut/log"
	"github.com/reddio-com/starkex-utils/types"
)

func GetAssetInfo(tokenType types.TokenType, address string) string {
	var assetInfo string
	switch tokenType {
	case types.ETH:
		assetInfo = "0x8322fff2"
	case types.ERC20:
		assetInfo = "0xf47261b0" + rightJust(address[2:], 64, "0")
	case types.ERC721:
		assetInfo = "0x02571792" + rightJust(address[2:], 64, "0")
	case types.ERC721M:
		assetInfo = "0xb8b86672" + rightJust(address[2:], 64, "0")
	case types.ERC20M:
		assetInfo = "0x68646e2d" + rightJust(address[2:], 64, "0")
	}
	return assetInfo
}

func GetAssetType(tokenType types.TokenType, address string, tokenQuantum *big.Int) *big.Int {
	assetInfo := GetAssetInfo(tokenType, address)
	quantum := tokenQuantum.Uint64()

	hash := solsha3.SoliditySHA3(
		solsha3.Bytes32(assetInfo),
		solsha3.Uint256(fmt.Sprint(quantum)),
	)

	if tokenType == types.ETH {
		hash = solsha3.SoliditySHA3(
			solsha3.Bytes4(assetInfo),
			solsha3.Uint256(fmt.Sprint(quantum)),
		)
	}

	assetType := new(big.Int)
	assetType.SetString(hex.EncodeToString(hash), 16)
	arg := new(big.Int)
	arg.SetString("03FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
	return assetType.And(assetType, arg)
}

// TODO: support minting_blob
func GetAssetID(tokenType types.TokenType, address string, tokenQuantum *big.Int, tokenId *big.Int) *big.Int {
	assetId := GetAssetType(tokenType, address, tokenQuantum)
	switch tokenType {
	case types.ERC721:
		hash := solsha3.SoliditySHA3(
			solsha3.String("NFT:"),
			solsha3.Uint256(assetId.String()),
			solsha3.Uint256(tokenId.String()),
		)

		assetId.SetString(hex.EncodeToString(hash), 16)
		arg := new(big.Int)
		arg.SetString("03FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
		assetId = assetId.And(assetId, arg)
	case types.ERC20M, types.ERC721M:
		blobHash := new(big.Int)
		blobHash.SetString(hex.EncodeToString(solsha3.SoliditySHA3(
			[]string{
				"string",
			},
			[]interface{}{
				tokenId.String(), // TODO: support minting_blob
			},
		)), 16)
		assetId.SetString(hex.EncodeToString(solsha3.SoliditySHA3(
			[]string{
				"string",
				"uint256",
				"uint256",
			},
			[]interface{}{
				"MINTABLE:",
				assetId,
				blobHash,
			},
		)), 16)
		arg := new(big.Int)
		arg.SetString("0000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
		arg1 := new(big.Int)
		arg1.SetString("400000000000000000000000000000000000000000000000000000000000000", 16)
		assetId = assetId.And(assetId, arg)
		assetId = assetId.Or(assetId, arg1)
	}
	return assetId
}

func GetAssetIDByAssetType(tokenType types.TokenType, assetType *big.Int, tokenId *big.Int) *big.Int {
	assetId := assetType
	switch tokenType {
	case types.ERC721:
		hash := solsha3.SoliditySHA3(
			solsha3.String("NFT:"),
			solsha3.Uint256(assetId.String()),
			solsha3.Uint256(tokenId.String()),
		)

		assetId.SetString(hex.EncodeToString(hash), 16)
		arg := new(big.Int)
		arg.SetString("03FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
		assetId = assetId.And(assetId, arg)
	case types.ERC20M, types.ERC721M:
		blobHash := new(big.Int)
		blobHash.SetString(hex.EncodeToString(solsha3.SoliditySHA3(
			[]string{
				"string",
			},
			[]interface{}{
				tokenId.String(), // TODO: support minting_blob
			},
		)), 16)
		assetId.SetString(hex.EncodeToString(solsha3.SoliditySHA3(
			[]string{
				"string",
				"uint256",
				"uint256",
			},
			[]interface{}{
				"MINTABLE:",
				assetId,
				blobHash,
			},
		)), 16)
		arg := new(big.Int)
		arg.SetString("0000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
		arg1 := new(big.Int)
		arg1.SetString("400000000000000000000000000000000000000000000000000000000000000", 16)
		assetId = assetId.And(assetId, arg)
		assetId = assetId.Or(assetId, arg1)
	}
	return assetId
}

func rightJust(s string, n int, fill string) string {
	log.Debugf("rightJust: s: %s, n: %d, fill: %s", s, n, fill)
	return strings.Repeat(fill, n-len(s)) + s
}
