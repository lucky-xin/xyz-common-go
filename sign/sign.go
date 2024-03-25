package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"time"
)

func SignWithTimestamp(secret, content string) (string, string) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
	toSign := []byte(timestamp + "\n" + secret + "\n" + content)
	sign := Sign([]byte(secret), toSign)
	return timestamp, sign
}

func Sign(secret, toSign []byte) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write(toSign)
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return sign
}
