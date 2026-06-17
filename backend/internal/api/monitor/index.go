package monitor

import (
	"encoding/json"
	"net/http"
	"time"

	"justdrven.dev/storage/internal/api/common"
)

func MonitorHandler(response http.ResponseWriter, request *http.Request) {
	common.SetJsonType(response)

	now := time.Now().Unix()

	json.NewEncoder(response).Encode(MonitorStatus{
		Status:    "UP",
		Timestamp: time.Duration(now),
	})
}
