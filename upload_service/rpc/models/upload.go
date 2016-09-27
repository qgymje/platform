package models

import "time"

// UploadProvider upload service provider
type UploadProvider int

const (
	// Qiniu used in China mainland
	Qiniu UploadProvider = iota + 1
)

// Upload uploaded file record
type Upload struct {
	FilePath       string    `bson:"file_path"`
	FileRemotePath string    `bson:"file_remote_path"`
	Provider       int       `bson:"provider"`
	CreatedAt      time.Time `bson:"created_At"`
}

// Create create a file upload record
func (u *Upload) Create() error {
	session := GetMongo()
	defer session.Close()

	u.CreatedAt = time.Now()
	return session.DB(DBName).C(ColNameUpload).Insert(&u)
}
