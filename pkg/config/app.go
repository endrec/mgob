package config

type AppConfig struct {
	LogLevel    string `json:"log_level"`
	JSONLog     bool   `json:"json_log"`
	Port        int    `json:"port"`
	ConfigPath  string `json:"config_path"`
	StoragePath string `json:"storage_path"`
	TmpPath     string `json:"tmp_path"`
	DataPath    string `json:"data_path"`
	Version     string `json:"version"`
	UseAwsCli   bool   `json:"use_aws_cli"`
}
