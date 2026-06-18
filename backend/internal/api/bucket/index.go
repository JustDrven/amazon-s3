package bucket

import (
	"encoding/xml"
	"net/http"

	"justdrven.dev/storage/cmd/configuration"
	"justdrven.dev/storage/internal/api/common"

	bucketManager "justdrven.dev/storage/internal/repository/bucket"

	"justdrven.dev/storage/pkg"
)

func getFullPath(bucket string, key string) string {

	mainFolder := configuration.GetConfig().StorageFolderPath
	pkg.FixFolderPath(&mainFolder)

	return mainFolder + bucket + "/" + key

}

func GetListBucketResultHandler(res http.ResponseWriter, req *http.Request) {
	common.SetXmlType(res)

	bucket := req.PathValue("bucket")
	key := req.PathValue("key")
	path := getFullPath(bucket, key)

	results, err := bucketManager.GetListObjects(bucket, key, path)
	if err != nil {
		xml.NewEncoder(res).Encode(common.APIErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	xml.NewEncoder(res).Encode(results)

}
