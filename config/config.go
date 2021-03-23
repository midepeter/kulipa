package config

type MainConfig struct {
	Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	//Database url is an optional parameter that will contain a full connection option string

	Host string `yaml:"host" env:"DATABASE_HOST" env-description:"database host"`

	//The post we are going to be hosting the database
	Port string `yaml:"port" env:"DATABASE_PORT" env-description:"database port"`

	//Database name
	Database string `yaml:"database" env:"DATABASE_NAME" env-description:"database name"`

	//Database Username 
	Username string `yaml:"username" env:"DATABASE_USERNAME" env-description:"database username"`

	//Database password
	Password string `yaml:"password" env:"DATABASE_PASSWORD" env-description:"database password"`

	//Database SSL status 
	SSL string `yaml:"ssl" env:"DATABASE_SSL" env-description:"SSL status"`
}
