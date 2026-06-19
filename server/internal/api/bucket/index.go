package bucket

import (
	"encoding/xml"
	"net/http"
	"time"

	"justdrven.dev/storage/internal/api/common"

	bucketManager "justdrven.dev/storage/internal/repository/bucket"
	"justdrven.dev/storage/pkg"
	"justdrven.dev/storage/shared/src/configuration"

	pkgShared "justdrven.dev/storage/shared/src/pkg"
)

var BUCKET_CACHE = pkg.NewCache()

func getFullPath(bucket string, key string) string {

	mainFolder := configuration.GetConfig().StorageFolderPath

	pkgShared.FixFolderPath(&mainFolder)

	return mainFolder + bucket + "/" + key

}

func GetListBucketResultHandler(res http.ResponseWriter, req *http.Request) {
	common.SetXmlType(res)

	bucket := req.PathValue("bucket")
	key := req.PathValue("key")
	path := getFullPath(bucket, key)

	encoder := xml.NewEncoder(res)

	if val, found := BUCKET_CACHE.Get(path); found {
		result := val.(*bucketManager.ListBucketResult)

		encoder.Encode(result)
		return
	}

	results, err := bucketManager.GetListObjects(bucket, key, path)
	if err != nil {
		encoder.Encode(common.APIErrorResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	BUCKET_CACHE.SetWithTTL(path, results, 1, 10*time.Second)

	xml.NewEncoder(res).Encode(results)

}
