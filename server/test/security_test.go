package test_test

import (
	"fmt"
	"testing"

	"justdrven.dev/storage/shared/src/security"
)

const (
	PASSWORD = "test-password"
)

func TestHash(t *testing.T) {
	value := security.Hash(PASSWORD)

	if value == "" {
		t.Error("unexpected error with hash password")
	}

	fmt.Printf("hashed password is: %s\n", value)
}

func TestCompare(t *testing.T) {

	hashedPassword := security.Hash(PASSWORD)

	areSame := security.Compare(hashedPassword, PASSWORD)
	if !areSame {
		t.Error("passwords aren't same!")
	}

}
