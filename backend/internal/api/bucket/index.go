package bucket

import (
	"encoding/xml"
	"net/http"

	"justdrven.dev/storage/internal/api/common"
)

func GetListBucketResultHandler(res http.ResponseWriter, req *http.Request) {
	common.SetXmlType(res)

	bucket := req.PathValue("bucket")
	key := req.PathValue("key")

	xml.NewEncoder(res).Encode(ListBucketResult{
		Name:   key,
		Prefix: bucket,

		KeyCount: 0,
		Contents: make([]ObjectContent, 0),
	})

}
