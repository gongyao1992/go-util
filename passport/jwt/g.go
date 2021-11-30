package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/url"
	"strings"
	"time"
)

// Payload ...
type Payload struct {
	Exp  int64  `json:"exp"`
	ID   string `json:"id"`
	Tp   string `json:"tp"`
	Sign string
}

// Decode ...
func Decode(token, header, key string) (payload Payload, err error) {
	// format: payload + "." + signature
	tokenSafe, _ := url.QueryUnescape(token)
	// sign[0]: payload
	// sign[1]: signature
	sign := strings.Split(tokenSafe, ".")
	if len(sign) != 2 {
		return payload, errors.New("AUTH_INFO_ERR")
	}
	payloadJSON, err := decodeBase64(sign[0])
	if err != nil {
		return
	}
	json.Unmarshal(payloadJSON, &payload)
	messageMAC, err := decodeBase64(sign[1])
	if err != nil {
		return
	}
	signStr := header + "." + sign[0]
	signKey := key + payload.ID
	if !CheckMAC([]byte(signStr), messageMAC, []byte(signKey)) {
		return payload, errors.New("AUTH_INFO_ERR")
	}
	now := time.Now().Unix()
	if payload.Exp < now {
		return payload, errors.New("AUTH_INFO_EXPIRED")
	}
	return
}

// CheckMAC https://golang.org/pkg/crypto/hmac/
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

// check phpweb: Base64UrlSafeEncoder
func encodeBase64(str string) (byt []byte) {
	str = base64.StdEncoding.EncodeToString([]byte(str))
	str = strings.Replace(str, "+", "-", -1)
	str = strings.Replace(str, "/", "_", -1)
	str = strings.Replace(str, "=", "", -1)
	byt = []byte(str)
	return
}

// check phpweb: Base64UrlSafeEncoder
func decodeBase64(str string) (byt []byte, err error) {
	str = strings.Replace(str, "-", "+", -1)
	str = strings.Replace(str, "_", "/", -1)
	byt, err = base64.RawStdEncoding.DecodeString(str)
	if err != nil {
		return byt, errors.New("AUTH_INFO_ERR")
	}
	return
}
