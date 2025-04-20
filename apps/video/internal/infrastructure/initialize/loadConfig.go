package initialize

import (
	"fmt"
	"os"
	"strings"

	"github.com/bhtoan2204/video/global"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitLoadConfig() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "local"
	}

	if env != "production" {
		if err := godotenv.Load(
			fmt.Sprintf("config/.env.%s", env),
		); err != nil {
			panic(fmt.Errorf("error loading .env files: %w", err))
		}
	}

	v := viper.New()
	v.SetEnvPrefix("")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	bindEnv(v)

	v.AutomaticEnv()

	config := &global.Config
	if err := v.Unmarshal(config); err != nil {
		panic(fmt.Errorf("unable to decode configuration: %w", err))
	}
}

func bindEnv(v *viper.Viper) {
	// Server mappings
	v.BindEnv("server.mode", "SERVER_MODE")
	v.BindEnv("server.gin_mode", "GIN_MODE")

	// MySQL mappings
	v.BindEnv("mysql.user", "MYSQL_USER")
	v.BindEnv("mysql.pass", "MYSQL_PASS")
	v.BindEnv("mysql.host", "MYSQL_HOST")
	v.BindEnv("mysql.port", "MYSQL_PORT")
	v.BindEnv("mysql.name", "MYSQL_NAME")
	v.BindEnv("mysql.charset", "MYSQL_CHARSET")
	v.BindEnv("mysql.parse_time", "MYSQL_PARSE_TIME")
	v.BindEnv("mysql.loc", "MYSQL_LOC")
	v.BindEnv("mysql.max_idle_conns", "MYSQL_MAX_IDLE_CONNS")
	v.BindEnv("mysql.max_open_conns", "MYSQL_MAX_OPEN_CONNS")
	v.BindEnv("mysql.max_lifetime", "MYSQL_MAX_LIFETIME")

	// Security mappings
	v.BindEnv("security.jwt_access_secret", "SECURITY_JWT_ACCESS_SECRET")
	v.BindEnv("security.jwt_refresh_secret", "SECURITY_JWT_REFRESH_SECRET")
	v.BindEnv("security.jwt_access_expiration", "SECURITY_JWT_ACCESS_EXPIRATION")
	v.BindEnv("security.jwt_refresh_expiration", "SECURITY_JWT_REFRESH_EXPIRATION")

	// Log mappings
	v.BindEnv("log.log_level", "LOG_LOG_LEVEL")
	v.BindEnv("log.file_path", "LOG_FILE_PATH")
	v.BindEnv("log.max_size", "LOG_MAX_SIZE")
	v.BindEnv("log.max_backups", "LOG_MAX_BACKUPS")
	v.BindEnv("log.max_age", "LOG_MAX_AGE")
	v.BindEnv("log.compress", "LOG_COMPRESS")

	// Redis mappings
	v.BindEnv("redis.host", "REDIS_HOST")
	v.BindEnv("redis.port", "REDIS_PORT")
	v.BindEnv("redis.password", "REDIS_PASSWORD")
	v.BindEnv("redis.database", "REDIS_DATABASE")

	// S3 mappings
	v.BindEnv("s3.region", "S3_REGION")
	v.BindEnv("s3.bucket", "S3_BUCKET")

	// Consul mappings
	v.BindEnv("consul.address", "CONSUL_ADDRESS")
	v.BindEnv("consul.scheme", "CONSUL_SCHEME")
	v.BindEnv("consul.data_center", "CONSUL_DATA_CENTER")
	v.BindEnv("consul.token", "CONSUL_TOKEN")

	// Opentelemetry mappings
	v.BindEnv("opentelemetry.endpoint", "OPENTELEMETRY_ENDPOINT")

	// Scylla mappings
	v.BindEnv("scylla.host", "SCYLLA_HOST")
	v.BindEnv("scylla.port", "SCYLLA_PORT")
	v.BindEnv("scylla.username", "SCYLLA_USERNAME")
	v.BindEnv("scylla.password", "SCYLLA_PASSWORD")
	v.BindEnv("scylla.keyspace", "SCYLLA_KEYSPACE")
}
