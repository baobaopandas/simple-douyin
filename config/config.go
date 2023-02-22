package config

type JWTConfig struct {
	JWTSecret  string `mapstructure:"jwt_secret"`
	ExpireTime int    `mapstructure:"expire_time"`
}

type DbConfig struct {
	DbDriver string `mapstructure:"db_driver"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructrue:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"db_name"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	port string `mapstructure:"port"`
}

type VideoConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type System struct {
	JWTConfig   *JWTConfig   `mapstructure:"jwt"`
	DbConfig    *DbConfig    `mapstructure:"db"`
	RedisConfig *RedisConfig `mapstructure:"redis"`
	VideoConfig *VideoConfig `mapstructure:"video"`
}
