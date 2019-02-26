package config

type OSSConfig struct {
	ENDPOINT          string `yaml:"ENDPOINT"`
	ACCESS_KEY_ID     string `yaml:"ACCESS_KEY_ID"`
	ACCESS_KEY_SECRET string `yaml:"ACCESS_KEY_SECRET"`
	BUCKET_NAME       string `yaml:"BUCKET_NAME"`
}

type DatabaseConfig struct {
	Type string `yaml:"type"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Database string `yaml:"database"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
}

type CacheConfig struct {
	Type string `yaml:"type"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	DB int `yaml:"db"`
	Password string `yaml:"password"`
}

type SettingConfig struct {
	SECRET_KEY string `yaml:"SECRET_KEY"`
	DEBUG string `yaml:"DEBUG"`
	DEFAULT_CHARSET string `yaml:"DEFAULT_CHARSET"`
	DATABASES []DatabaseConfig `yaml:"DATABASES"`
	CACHES []CacheConfig `yaml:"CACHES"`
}
