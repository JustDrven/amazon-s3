package object

import (
	"encoding/base64"
	"encoding/xml"
	"net/http"
	"strings"

	"justdrven.dev/storage/command/configuration"
	"justdrven.dev/storage/internal/api/common"

	objectManager "justdrven.dev/storage/internal/repository/object"
)

func getFileName(key string) string {
	if !strings.Contains(key, "/") {
		return key
	}

	data := strings.Split(key, "/")
	return data[len(data)-1]
}

func getFullPath() string {
	path := configuration.GetConfig().StorageFolderPath

	if !strings.HasSuffix(path, "/") {
		path += "/"
	}

	return path
}

func GetObjectHandler(res http.ResponseWriter, req *http.Request) {
	common.SetXmlType(res)

	path := getFullPath()
	bucket := req.PathValue("bucket")
	key := req.PathValue("key")
	bucketPath := path + bucket + "/" + key

	data, code, err := objectManager.GetContent(bucketPath)
	encoder := xml.NewEncoder(res)

	if err != nil {
		encoder.Encode(GetObjectFailedResult{
			Code:    code,
			Message: err.Error(),
		})
		return
	}

	fileName := getFileName(key)

	encoded := base64.StdEncoding.EncodeToString(data)

	encoder.Encode(GetObjectResult{
		Prefix: bucket,
		Name:   key,

		FilePath: bucketPath,
		FileName: fileName,

		Size: len(data),

		Content: encoded,
	})

}
