package global

var AppConfig Application

type Application struct {
	App      App      `mapstructure:"app"`
	Server   Server   `mapstructure:"server"`
	Security Security `mapstructure:"security"`
	Swagger  Swagger  `mapstructure:"swagger"`
	Redis    Redis    `mapstructure:"redis"`
	Database Database `mapstructure:"database"`
}

type App struct {
	Title       string `mapstructure:"title"`
	Name        string `mapstructure:"name"`
	Description string `mapstructure:"description"`
	Version     string `mapstructure:"version"`
}

type Server struct {
	Port   int64  `mapstructure:"port"`
	Logger Logger `mapstructure:"logger"`
}

type Logger struct {
	Enable bool `mapstructure:"enable"`
}

type Security struct {
	AES string `mapstructure:"aes"`
}

type Swagger struct {
	Port int64  `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Base string `mapstructure:"base"`
}

type Redis struct {
	Environment  string `mapstructure:"env"`
	Network      string `mapstructure:"network"`
	Url          string `mapstructure:"url"`
	database     string `mapstructure:"database"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Pool         int64  `mapstructure:"pool"`
	MinIdleConns int64  `mapstructure:"minIdleConns"`
	MaxRetries   int64  `mapstructure:"maxRetries"`
	Tls          bool   `mapstructure:"tls"`
	Sleep        int64  `mapstructure:"sleep"`
	ServiceName  string `mapstructure:"serviceName"`
}

type Database struct {
	KeepALive   int64      `mapstructure:"keepALive"`
	LogLevel    string     `mapstructure:"logLevel"`
	MaxIdleConn int64      `mapstructure:"maxIdleConn"`
	MinOpenConn int64      `mapstructure:"minOpenConn"`
	MaxLifetime int64      `mapstructure:"maxLifetime"`
	Connection  Connection `mapstructure:"connection"`
}

type Connection struct {
	Host     string `mapstructure:"host"`
	Port     int64  `mapstructure:"port"`
	Dialect  string `mapstructure:"dialect"`
	Schema   string `mapstructure:"schema"`
	Username string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	SSLMode  string `mapstructure:"sslMode"`
}
