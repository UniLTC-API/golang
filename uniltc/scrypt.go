package uniltc

import (
	"encoding/base64"
	"fmt"

	"code.google.com/p/go.crypto/scrypt"
)

const (
	n                = 1 << 14
	r                = 1 << 3
	p                = 1 << 0
	keyLen           = 1 << 7
	Op_newbid        = "newbid"
	Op_newask        = "newask"
	Op_cancelbid     = "cancelbid"
	Op_cancelask     = "cancelask"
	Op_currentorders = "currentorders"
	Op_historytrades = "historytrades"
	Op_userinfo      = "userinfo"
	Op_withdraw      = "withdraw"
)

func EncryptedRes(sec string, op string, args ...interface{}) (string, error) {
	data := fmt.Sprintf("%v", op)
	for _, v := range args {
		data += fmt.Sprintf("%v", v)
	}
	b, err := scrypt.Key([]byte(data), []byte(sec), n, r, p, keyLen)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
