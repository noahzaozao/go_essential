package essential

type OSSConfig struct {
	ENDPOINT          string `json:"ENDPOINT"`
	ACCESS_KEY_ID     string `json:"ACCESS_KEY_ID"`
	ACCESS_KEY_SECRET string `json:"ACCESS_KEY_SECRET"`
	BUCKET_NAME       string `json:"BUCKET_NAME"`
}

type DatabaseConfig struct {
	Type string `json:"type"`
	Host string `json:"host"`
	Port string `json:"port"`
	Database string `json:"database"`
	User string `json:"user"`
	Password string `json:"password"`
}

type CacheConfig struct {
	Type string `json:"type"`
	Host string `json:"host"`
	Port string `json:"port"`
	Password string `json:"password"`
}

type SettingConfig struct {
	SECRET_KEY string `json:"SECRET_KEY"`
	DEBUG string `json:"DEBUG"`
	DEFAULT_CHARSET string `json:"DEFAULT_CHARSET"`
	DATABASES []DatabaseConfig `json:"DATABASES"`
	CACHES []CacheConfig `json:"CACHES"`
}
