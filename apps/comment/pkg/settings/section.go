package settings

type ServerConfig struct {
	Mode    string `mapstructure:"mode"`
	GinMode string `mapstructure:"gin_mode"`
}

type MySQLConfig struct {
	User         string `mapstructure:"user"`
	Pass         string `mapstructure:"pass"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Name         string `mapstructure:"name"`
	Charset      string `mapstructure:"charset"`
	ParseTime    bool   `mapstructure:"parse_time"`
	Loc          string `mapstructure:"loc"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxLifetime  int    `mapstructure:"max_lifetime"`
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

type ConsulConfig struct {
	Address    string `mapstructure:"address"`
	Scheme     string `mapstructure:"scheme"`
	DataCenter string `mapstructure:"data_center"`
	Token      string `mapstructure:"token"`
}

type OpentelemetryConfig struct {
	Endpoint string `mapstructure:"endpoint"`
}

type ScyllaConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Keyspace string `mapstructure:"keyspace"`
}

type KafkaConfig struct {
	Broker  string `mapstructure:"broker"`
	Port    int    `mapstructure:"port"`
	Topic   string `mapstructure:"topic"`
	GroupID string `mapstructure:"group_id"`
}

type Config struct {
	Server              ServerConfig        `mapstructure:"server"`
	MySQLConfig         MySQLConfig         `mapstructure:"mysql"`
	LogConfig           LogConfig           `mapstructure:"log"`
	RedisConfig         RedisConfig         `mapstructure:"redis"`
	ConsulConfig        ConsulConfig        `mapstructure:"consul"`
	OpentelemetryConfig OpentelemetryConfig `mapstructure:"opentelemetry"`
	ScyllaConfig        ScyllaConfig        `mapstructure:"scylla"`
	KafkaConfig         KafkaConfig         `mapstructure:"kafka"`
}
