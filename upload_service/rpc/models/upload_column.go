package models

type _UploadColumn struct {
	CreatedAt string
	FilePath  string
	Provider  string
}

// UploadColumns upload columns name
var UploadColumns _UploadColumn

func init() {
	UploadColumns.CreatedAt = "created_at"
	UploadColumns.FilePath = "file_path"
	UploadColumns.Provider = "provider"

}
