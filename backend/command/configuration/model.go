package configuration

type ConfigData struct {
	JwtSecret         string `json:"jwtSecret"`
	Port              int    `json:"port"`
	StorageFolderPath string `json:"storageFolderPath"`
}
