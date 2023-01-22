package pcloud

import (
	"fmt"
	"os"
	"testing"

	"github.com/as27/pcloud/api"
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

func TestFolderList(t *testing.T) {
	authToken = os.Getenv("PCLOUD_TOKEN")

	type respStr struct {
		Result       int `json:"result"`
		api.Metadata `json:"metadata"`
	}
	md := &respStr{}
	err := ApiRequest(md, "listfolder",
		Param{"path", "/Ida"},
		Param{"recursive", "1"},
	)
	if err != nil {
		t.Errorf("apiRequest: %v", err)
	}
	fmt.Printf("%#v", md)
	t.Errorf("ab")
}
