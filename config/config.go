package config

type Config struct {
	Name     string
	Mode     string
	Port     int
	Debug    bool
	Timezone string

	Log      LogConfig
	DataBase DataBaseConfig
	JWT      JWTConfig
	Redis    RedisConfig
}
