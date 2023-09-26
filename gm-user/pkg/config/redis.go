package config

type redisConfig struct {
	Address     []string `yaml:"address"`
	MaxIdle     int      `yaml:"max_idle"`
	MaxActive   int      `yaml:"max_active"`
	IdleTimeout int      `yaml:"idle_timeout"`
	Username    string   `yaml:"username"`
	Password    string   `yaml:"password"`
	Cluster     bool     `yaml:"cluster"`
}
