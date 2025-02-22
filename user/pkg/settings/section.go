package settings

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
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

type SecurityConfig struct {
	JWTAccessSecret      string `mapstructure:"jwt_access_secret"`
	JWTRefreshSecret     string `mapstructure:"jwt_refresh_secret"`
	JWTAccessExpiration  int    `mapstructure:"jwt_access_expiration"`
	JWTRefreshExpiration int    `mapstructure:"jwt_refresh_expiration"`
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

// type S3Config struct {
// 	AccessKeyID     string `mapstructure:"access_key_id"`
// 	SecretAccessKey string `mapstructure:"secret_access_key"`
// 	Region          string `mapstructure:"region"`
// 	Bucket          string `mapstructure:"bucket"`
// }

type ConsulConfig struct {
	Address    string `mapstructure:"address"`
	Scheme     string `mapstructure:"scheme"`
	DataCenter string `mapstructure:"data_center"`
	Token      string `mapstructure:"token"`
}

type KafkaConfig struct {
	Broker  string `mapstructure:"broker"`
	Port    int    `mapstructure:"port"`
	Topic   string `mapstructure:"topic"`
	GroupID string `mapstructure:"group_id"`
}

type ElasticSearchConfig struct {
	Address  string `mapstructure:"address"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Config struct {
	Server              ServerConfig        `mapstructure:"server"`
	MySQLConfig         MySQLConfig         `mapstructure:"mysql"`
	LogConfig           LogConfig           `mapstructure:"log"`
	SecurityConfig      SecurityConfig      `mapstructure:"security"`
	RedisConfig         RedisConfig         `mapstructure:"redis"`
	ConsulConfig        ConsulConfig        `mapstructure:"consul"`
	KafkaConfig         KafkaConfig         `mapstructure:"kafka"`
	ElasticSearchConfig ElasticSearchConfig `mapstructure:"elasticsearch"`
	// S3Config       S3Config       `mapstructure:"s3"`
}
