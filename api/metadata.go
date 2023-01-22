package api

// Metadata is returned when listfolder method is used.
// Documentation of the API:
// https://docs.pcloud.com/methods/folder/listfolder.html
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
