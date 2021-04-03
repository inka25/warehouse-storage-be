package config

const (
	EnvProduction = "production"
	EnvUat        = "uat"
	EnvTest       = "test"
)

type Config struct {
	AppName string `ini:"appname"`
	Host    string `ini:"host"`
	Port    string `ini:"port"`
	Env     string `ini:"env"`
	//Loadtest    bool   `ini:"loadtest"` // not needed, maybe later
	HTTPTimeout int    `ini:"http_timeout"`
	JWTSecret   string `ini:"jwt_secret"`

	MySQL
	Logger
}

type MySQL struct {
	Host            string `ini:"host"`
	Port            int    `ini:"port"`
	User            string `ini:"user"`
	Pass            string `ini:"password"`
	DB              string `ini:"db"`
	MaxIdle         int    `ini:"max_idle"`
	MaxOpen         int    `ini:"max_open"`
	MaxLifetime     int64  `ini:"max_lifetime"`
	PaginationLimit int64  `ini:"pagination_limit"`
}

type Logger struct {
	Location string `ini:"location"`
}
