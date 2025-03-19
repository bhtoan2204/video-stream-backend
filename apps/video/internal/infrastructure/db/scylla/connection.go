package scylla

import (
	"fmt"
	"time"

	"github.com/bhtoan2204/video/global"
	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

func testScript() {
	err := global.ScyllaSession.Query(`
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		name TEXT,
		age INT
	)`).Exec()

	if err != nil {
		global.Logger.Error("Failed to execute test query", zap.Error(err))
		panic(err)
	}
	global.Logger.Info("Test query executed successfully")
}

func InitDB() {
	global.Logger.Info("Connecting to Scylla: ", zap.Any("config", global.Config.ScyllaConfig))

	cluster := gocql.NewCluster(global.Config.ScyllaConfig.Host)
	cluster.Port = global.Config.ScyllaConfig.Port
	cluster.Consistency = gocql.Quorum
	cluster.ConnectTimeout = 10 * time.Second
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: global.Config.ScyllaConfig.Username,
		Password: global.Config.ScyllaConfig.Password,
	}

	tempSession, err := cluster.CreateSession()
	if err != nil {
		global.Logger.Error("Failed to connect to Scylla", zap.Error(err))
		panic(err)
	}
	defer tempSession.Close()

	err = tempSession.Query(fmt.Sprintf(`
	CREATE KEYSPACE IF NOT EXISTS %s 
	WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3};
`, global.Config.ScyllaConfig.Keyspace)).Exec()
	if err != nil {
		global.Logger.Error("Failed to create keyspace", zap.Error(err))
		panic(err)
	}

	cluster.Keyspace = global.Config.ScyllaConfig.Keyspace

	session, err := cluster.CreateSession()
	if err != nil {
		global.Logger.Error("Failed to connect to Scylla", zap.Error(err))
		panic(err)
	}

	global.ScyllaSession = session
	global.Logger.Info("Connected to Scylla successfully")

	testScript()
}

func CloseDB() {
	global.ScyllaSession.Close()
	global.Logger.Info("Scylla connection closed")
}
