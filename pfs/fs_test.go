package pfs

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"testing/fstest"

	"github.com/as27/pcloud/api"
)

func loadJSON(t *testing.T, filename string, v interface{}) error {
	data, err := os.Open(filename)
	if err != nil {
		t.Errorf("os.Open: %v", err)
	}
	defer data.Close()
	err = json.NewDecoder(data).Decode(v)
	if err != nil {
		t.Errorf("json.Decode: %v", err)
	}
	return nil
}

func TestFSStruct(t *testing.T) {
	type respStr struct {
		Result       int `json:"result"`
		api.Metadata `json:"metadata"`
	}
	md := &respStr{}
	err := loadJSON(t, "testfiles/listfolder.json", md)
	if err != nil {
		t.Errorf("loadJSON: %v", err)
	}
	f := &FS{
		metadata: md.Metadata,
	}
	err = fstest.TestFS(f, "file1.txt", "file2.wav")
	if err != nil {
		t.Errorf("fstest.TestFS: %v", err)
	}
}

func TestFileStruct1(t *testing.T) {
	type respStr struct {
		Result       int `json:"result"`
		api.Metadata `json:"metadata"`
	}
	md := &respStr{}
	err := loadJSON(t, "testfiles/listfolder.json", md)
	if err != nil {
		t.Errorf("loadJSON: %v", err)
	}
	f := &FS{
		metadata: md.Metadata,
	}
	f1, _ := f.Open("file1.txt")
	fmt.Printf("%#v", f1)
	fmt.Println(f1.Stat())
	//fmt.Println(f1.Info())
	t.Error("hej")
}
