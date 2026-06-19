package object

import (
	"encoding/base64"
	"encoding/xml"
	"net/http"
	"strings"
	"time"

	"justdrven.dev/storage/internal/api/common"
	"justdrven.dev/storage/pkg"
	"justdrven.dev/storage/shared/src/configuration"

	objectManager "justdrven.dev/storage/internal/repository/object"
)

var OBJECT_CACHE = pkg.NewCache()

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
	encoder := xml.NewEncoder(res)

	if val, found := OBJECT_CACHE.Get(bucketPath); found {
		result := val.(GetObjectResult)
		encoder.Encode(result)
		return
	}

	data, code, err := objectManager.GetContent(bucketPath)
	if err != nil {
		encoder.Encode(common.APIErrorResponse{
			Code:    code,
			Message: err.Error(),
		})
		return
	}

	fileName := getFileName(key)

	encoded := base64.StdEncoding.EncodeToString(data)

	result := GetObjectResult{
		Prefix: bucket,
		Name:   key,

		FilePath: bucketPath,
		FileName: fileName,

		Size: len(data),

		Content: encoded,
	}

	OBJECT_CACHE.SetWithTTL(bucketPath, result, 1, 5*time.Second)

	encoder.Encode(result)

}
