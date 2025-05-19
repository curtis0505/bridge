package conf

type Config []RepositoryConfig

type RepositoryConfig struct {
	Type       string `mapstructure:"type"`
	RedisDB    string `mapstructure:"redisDB"`
	DataSource string `mapstructure:"datasource"`
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	TLS        bool   `mapstructure:"tls"`
}
