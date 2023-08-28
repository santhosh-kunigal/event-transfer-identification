package utils

import (
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"math/big"
)

const (
	TransferEvent = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
	ERC20         = "ERC20"
	ERC721        = "ERC721"
)

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func HexToIntNoError(end string) *big.Int {
	if end == "0x" {
		return big.NewInt(0)
	}
	toInt, err := HexToInt(end)
	if err != nil {
		logrus.Error(err)
	}
	return toInt
}

func HexToInt(data string) (*big.Int, error) {
	if has0xPrefix(data) {
		data = data[2:]
	}
	var n big.Int
	_, ok := n.SetString(data, 16)
	if ok {
		return &n, nil
	}

	return nil, errors.New("Failed to decode")
}

func has0xPrefix(s string) bool {
	return len(s) >= 2 && s[0] == '0' && (s[1]|32) == 'x'
}

func GetHexString(wholeHex string) string {
	if len(wholeHex) < 64 {
		return wholeHex
	}
	wholeHex = wholeHex[2:]
	wholeHex = wholeHex[:64]
	wholeHex = wholeHex[len(wholeHex)-40:]
	wholeHex = "0x" + wholeHex
	return wholeHex
}
