package md5

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Check(encrypted, content string, salt ...string) bool {
	return strings.EqualFold(Encode(content, salt...), encrypted)
}

func Encode(data string, salt ...string) string {
	h := md5.New()
	h.Write([]byte(data + strings.Join(salt, "")))
	return hex.EncodeToString(h.Sum(nil))
}
