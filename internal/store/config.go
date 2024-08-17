package store

type DBConnectionConfig struct {
	DatabaseDriver  string
	DatabaseName    string
	Host            string
	Password        string
	QueryBufferSize uint8
	User            string
}

func NewDBConnectionConfig() DBConnectionConfig {
	return DBConnectionConfig{
		DatabaseDriver:  "postgres",
		DatabaseName:    "app",
		Host:            "localhost",
		Password:        "password",
		QueryBufferSize: 10,
		User:            "admin",
	}
}

func (config DBConnectionConfig) WithDatabaseDriver(driver string) DBConnectionConfig {
	config.DatabaseDriver = driver
	return config
}
func (config DBConnectionConfig) WithDatabaseName(name string) DBConnectionConfig {
	config.DatabaseName = name
	return config
}
func (config DBConnectionConfig) WithHost(host string) DBConnectionConfig {
	config.Host = host
	return config
}
func (config DBConnectionConfig) WithPassword(password string) DBConnectionConfig {
	config.Password = password
	return config
}
func (config DBConnectionConfig) WithQueryBufferSize(size uint8) DBConnectionConfig {
	config.QueryBufferSize = size
	return config
}
func (config DBConnectionConfig) WithUser(user string) DBConnectionConfig {
	config.User = user
	return config
}
