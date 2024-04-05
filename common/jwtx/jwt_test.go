package jwtx

import (
	"testing"
	"time"
)

func TestGetToken(t *testing.T) {
	key := "mySecretKey"
	iat := time.Now().Unix()
	second := int64(1)
	uid := int64(123)

	_, err := GetToken(key, iat, second, uid)
	if err != nil {
		t.Errorf("GetToken returned an error: %v", err)
	}
}

func TestIsExpired(t *testing.T) {
	key := "mySecretKey"
	iat := time.Now().Unix()
	second := int64(1)
	uid := int64(123)

	token, err := GetToken(key, iat, second, uid)
	if err != nil {
		t.Errorf("GetToken returned an error: %v", err)
	}

	expired := IsExpired(token, key)
	if !expired {
		t.Errorf("IsExpired returned true, expected false")
	}

	time.Sleep(time.Second * 2)
	expired = IsExpired(token, key)
	if expired {
		t.Errorf("IsExpired returned false, expected true")
	}
}
