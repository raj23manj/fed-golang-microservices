package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(input string) string {
	hash := md5.New()
	defer hash.Reset()
	// Convert string to bytes
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
