package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"justdrven.dev/storage/command/configuration"
	"justdrven.dev/storage/internal/api/common"
	apiMonitor "justdrven.dev/storage/internal/api/monitor"
	"justdrven.dev/storage/internal/repository/router"
	"justdrven.dev/storage/pkg"
)

func Main(file *os.File) {
	log.Info("Starting engine..")

	config, err := configuration.LoadConfig(file)
	if err != nil {
		panic(err)
	}

	StartAPI(*config)
}

func isPortAvaliable(port int) error {
	return pkg.IsPortAvaliable(port)
}

func NotFoundHandler(res http.ResponseWriter, req *http.Request) {
	common.SetJsonType(res)

	encoder := json.NewEncoder(res)
	encoder.Encode(NotFoundResponse{
		Code:    404,
		Message: "NOT_FOUND",
	})
}

func registerRouters() {
	log.Info("Enabling endpoints..")

	router.Endpoint{
		Method:  router.ALL,
		Path:    "/",
		Handler: NotFoundHandler,
	}.Register()

	router.Endpoint{
		Method:  router.GET,
		Path:    "/monitor/health",
		Handler: apiMonitor.MonitorHandler,
	}.Register()

}

func StartAPI(config configuration.ConfigData) {
	port := config.Port

	err := isPortAvaliable(config.Port)
	if err != nil {
		panic(err)
	}

	registerRouters()

	address := fmt.Sprintf(":%d", port)
	log.Info("The service is listening!", "port", port)

	http.ListenAndServe(address, nil)
}
