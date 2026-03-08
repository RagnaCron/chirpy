package auth

import (
	"testing"
)

const password = "123456789"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v\n", password)
	}

	match, err := CheckPassword(password, hash)
	if err != nil {
		t.Errorf("Error Checking Hash: %v\n", err)
		return
	}

	if !match {
		t.Error("Password and hash do not match!!!")
	}
}
