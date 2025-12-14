package model

type Config struct {
	Version     string `json:"version"`
	Verbose     bool   `json:"verbose"`
	DataDir     string `json:"dataDir"`
	LogsDir     string `json:"logsDir"`
	RegistryUrl string `json:"registryUrl"`
}
