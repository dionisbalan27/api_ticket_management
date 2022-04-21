package models

type ServerConfig struct {
	PostgresConfig PostgresConfig
	DbNameLog      string `env:"NAME_POSTGRES_LOG,required"`
}

type PostgresConfig struct {
	Name     string `env:"NAME_POSTGRES,required"`
	Host     string `env:"HOST_POSTGRES,required"`
	Port     string `env:"PORT_POSTGRES,required"`
	User     string `env:"USER_POSTGRES"`
	Password string `env:"PASS_POSTGRES"`
}
