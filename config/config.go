package config

type Config struct {
	App        App
	Server     Server
	PostgreSql PostgreSql
	Redis      Redis
}

type App struct {
	Name               string
	Env                string
	Ssl                bool
	AssetsCacheVersion int8
	IpRateLimiter      bool
}

type Server struct {
	Host          string
	Port          string
	SessionExpire int
}

type PostgreSql struct {
	Username string
	Password string
	Host     string
	Port     int
	DbName   string
}

type Redis struct {
	DbName   int
	Host     string
	Port     int
	Password string
}
