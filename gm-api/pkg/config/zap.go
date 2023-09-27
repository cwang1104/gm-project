package config

type zap struct {
	DebugFileName string `yaml:"debugFileName"`
	InfoFileName  string `yaml:"infoFileName"`
	WarnFileName  string `yaml:"warnFileName"`
	MaxSize       int    `yaml:"maxSize"`
	MaxAge        int    `yaml:"maxAge"`
	MaxBackup     int    `yaml:"maxBackup"`
}
