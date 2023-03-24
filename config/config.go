package config

var GlobalConfig = new(Config)

type Config struct {
	Name     string
	Mode     string
	Port     int
	Debug    bool
	Timezone string
	Log      LogConfig
}
