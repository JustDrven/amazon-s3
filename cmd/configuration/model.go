package configuration

type ConfigData struct {
	Port              int    `json:"port"`
	StorageFolderPath string `json:"storageFolderPath"`
	EndUserFilePath   string `json:"enduserFilePath"`
}
