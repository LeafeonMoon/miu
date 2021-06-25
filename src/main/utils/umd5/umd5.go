package umd5

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(val string) string {
	m := md5.New()
	m.Write([]byte(val))
	return hex.EncodeToString(m.Sum(nil))
}
