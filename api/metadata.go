// Package api is the structure of the JSON returned by the pCloud API.
package api

// Metadata is returned when listing a folder
type Metadata struct {
	Path           string     `json:"path"`
	Name           string     `json:"name"`
	Created        string     `json:"created"`
	Modified       string     `json:"modified"`
	IsFolder       bool       `json:"isfolder"`
	FileID         int64      `json:"fileid"`
	Hash           int64      `json:"hash"`
	ID             string     `json:"id"`
	Size           int64      `json:"size"`
	ParentFolderID int64      `json:"parentfolderid"`
	ContentType    string     `json:"contenttype"`
	Contents       []Metadata `json:"contents"`
}
