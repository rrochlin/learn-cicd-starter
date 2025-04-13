package auth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"net/http"
	"testing"
)

func TestPasswordCreate(t *testing.T) {
	t.Error("intentional fail")
	secret, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Errorf("Could not generate key secret %v", err)
		return
	}

	headers := http.Header{}
	headers.Set(
		"Authorization",
		fmt.Sprintf("ApiKey %v", secret.X.String()),
	)
	key, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("Could not extract API key: %v", err)
		return
	}
	if key != secret.X.String() {
		t.Errorf("key: %v\nsecret:%v\nValues do not match", key, secret)
		return
	}
}
