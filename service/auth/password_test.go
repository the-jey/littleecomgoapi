package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "" {
		t.Errorf("expected hash to be not empty")
	}

	if hash == "password" {
		t.Errorf("expected hash to be different from password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if !ComparePasswords(hash, []byte("password")) {
		t.Errorf("expected password to match the hash")
	}

	if ComparePasswords(hash, []byte("notpassword")) {
		t.Error("expected password to not match the hash")
	}
}
