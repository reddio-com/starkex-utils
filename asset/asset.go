package asset

import (
	"errors"
	"strings"
)

func GetSelector(typ string) (string, error) {
	switch typ {
	case "ETH":
		return "0x8322fff2", nil
	case "ERC20":
		return "0xf47261b0", nil
	case "ERC721":
		return "0x02571792", nil
	case "MINTABLEERC721":
		return "0xb8b86672", nil
	case "MINTABLEERC20":
		return "0x68646e2d", nil
	}
	return "", errors.New("unknown asset type")
}

func GetAssetInfo(typ, address string) (string, error) {
	selector, err := GetSelector(typ)
	if err != nil {
		return "", err
	}
	if typ == "ETH" {
		return selector, nil
	}

	asset_info := selector + strings.Repeat("0", 24) + address[2:]
	return asset_info, nil
}
