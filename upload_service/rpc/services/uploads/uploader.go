package uploads

import (
	"platform/commons/codes"
	"platform/upload_service/rpc/models"
	"platform/upload_service/rpc/services/uploads/providers"
	"platform/utils"
)

// Uploader represents a upload process abstract object
type Uploader struct {
	content  []byte
	filename string

	uploadModel *models.Upload

	errorCode codes.ErrorCode
}

// NewUploader create a new uploader object
func NewUploader(filename string, content []byte) *Uploader {
	u := new(Uploader)
	u.filename = filename
	u.content = content
	u.uploadModel = &models.Upload{}
	return u
}

// ErrorCode implement the ErrorCoder interface
func (u *Uploader) ErrorCode() codes.ErrorCode {
	return u.errorCode
}

// Do do the uploading job
func (u *Uploader) Do() (err error) {
	defer func() {
		if err != nil {
			utils.GetLog().Error("upload service Uploader Do error: ", err)
		}
	}()

	if err = u.save(); err != nil {
		u.errorCode = codes.ErrorCodeUploadCreate
		return
	}

	if err = u.send(); err != nil {
		u.errorCode = codes.ErrorCodeUploadSend
		return
	}

	return
}

func (u *Uploader) save() (err error) {
	u.uploadModel.Provider = int(models.Qiniu)
	u.uploadModel.FilePath = u.Filename()

	return u.uploadModel.Create()
}

func (u *Uploader) send() (err error) {
	return providers.NewQiniuProvider(u).Do()
}

// Filename implement the Provider interface
func (u *Uploader) Filename() string {
	return u.filename
}

// Content implement the Provider interface
func (u *Uploader) Content() []byte {
	return u.content
}
