package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"justdrven.dev/storage/cmd/configuration"
	"justdrven.dev/storage/internal/api/bucket"
	"justdrven.dev/storage/internal/api/common"
	"justdrven.dev/storage/internal/api/object"
	"justdrven.dev/storage/internal/repository/router"
	"justdrven.dev/storage/pkg"
)

func Main(file *os.File) {

	log.Info("Starting engine..")

	config, err := configuration.LoadConfig(file)
	if err != nil {
		panic(err)
	}

	Start(*config)
}

func isPortAvaliable(port int) error {
	return pkg.IsPortAvaliable(port)
}

func NotFoundHandler(res http.ResponseWriter, req *http.Request) {
	common.SetJsonType(res)

	encoder := json.NewEncoder(res)
	encoder.Encode(common.APIErrorResponse{
		Code:    404,
		Message: "NOT_FOUND",
	})
}

func registerCommonEndpoints() {

	router.Endpoint{
		Path:    "/",
		Handler: NotFoundHandler,
	}.Register()

}

func registerUserEndpoints() {

}

func registerBucketEndpoints() {

	router.Endpoint{
		Method:  router.GET,
		Path:    "/bucket/{bucket}/{key...}",
		Handler: bucket.GetListBucketResultHandler,
	}.Register()

}

func registerObjectEndpoints() {

	router.Endpoint{
		Method:  router.GET,
		Path:    "/object/view/{bucket}/{key...}",
		Handler: object.GetObjectHandler,
	}.Register()

}

func registerRouters() {
	log.Info("Enabling endpoints..")

	registerCommonEndpoints()
	registerUserEndpoints()

	registerBucketEndpoints()
	registerObjectEndpoints()

	log.Info("Endpoints are complete!")
}

func Start(config configuration.ConfigData) {
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
