package pcloud

import (
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	authToken = ""
	user := os.Getenv("PCLOUD_USER")
	pass := os.Getenv("PCLOUD_PASS")
	err := LogIn(user, pass)
	if err != nil {
		t.Errorf("LogIn: %v", err)
	}
	want := "123"
	if AuthToken() != want {
		t.Errorf("AuthToken() = %v, want %v", AuthToken(), want)
	}
}
