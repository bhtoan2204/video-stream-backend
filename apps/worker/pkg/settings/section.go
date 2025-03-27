package settings

type ServerConfig struct {
	Mode    string `mapstructure:"mode"`
	GinMode string `mapstructure:"gin_mode"`
}

type LogConfig struct {
	LogLevel   string `mapstructure:"log_level"`
	FilePath   string `mapstructure:"file_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type OpentelemetryConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}

type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type S3Config struct {
	Bucket string `mapstructure:"bucket"`
	Region string `mapstructure:"region"`
}

type Config struct {
	Server              ServerConfig        `mapstructure:"server"`
	RedisConfig         RedisConfig         `mapstructure:"redis"`
	SMTPConfig          SMTPConfig          `mapstructure:"smtp"`
	OpentelemetryConfig OpentelemetryConfig `mapstructure:"opentelemetry"`
	LogConfig           LogConfig           `mapstructure:"log"`
	S3Config            S3Config            `mapstructure:"s3"`
}
