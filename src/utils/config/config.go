package config

type Config struct {
	App      App
	Database Database
	Service  Service
}

type App struct {
	Name    string
	Version string
}

type Database struct {
	PSQL SQL
}

type SQL struct {
	Driver   string
	Host     string
	Port     int64
	DB       string
	User     string
	Password string
	SSL      bool
}

type Service struct {
	GRPC struct {
		Port int64
	}
}
