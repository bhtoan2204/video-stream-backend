#!/bin/bash
# USER DATABASE
curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/ -d "{
  \"name\": \"mysql-user-connector\",
  \"config\": {
    \"connector.class\": \"io.debezium.connector.mysql.MySqlConnector\",
    \"tasks.max\": \"1\",
    \"database.hostname\": \"mysql\",
    \"database.port\": \"3306\",
    \"database.user\": \"root\",
    \"database.password\": \"root\",
    \"database.server.id\": \"184054\",
    \"database.server.name\": \"user_database\",
    \"database.whitelist\": \"users,refresh_tokens,activity_logs,permissions,role_permissions,roles,user_roles,user_settings\",
    \"database.history.kafka.bootstrap.servers\": \"kafka:29092\",
    \"database.history.kafka.topic\": \"dbhistory.user\",
    \"schema.history.internal.kafka.topic\": \"schemahistory.user\",
    \"schema.history.internal.kafka.bootstrap.servers\": \"kafka:29092\",
    \"database.parse_time\": \"true\",
    \"database.charset\": \"utf8mb4\",
    \"database.serverTimezone\": \"UTC\",
    \"topic.prefix\": \"user_database\",
    \"max_idle_conns\": \"10\",
    \"max_open_conns\": \"100\",
    \"max_lifetime\": \"3600\"
  }
}"

# VIDEO DATABASE
curl --location 'localhost:8083/connectors/' \
--header 'Accept: application/json' \
--header 'Content-Type: application/json' \
--data '{
  "name": "mysql-video-connector",
  "config": {
    "connector.class": "io.debezium.connector.mysql.MySqlConnector",
    "tasks.max": "1",
    "database.hostname": "mysql",
    "database.port": "3306",
    "database.user": "root",
    "database.password": "root",
    "database.server.id": "184054",
    "database.server.name": "video_database",
    "database.whitelist": "videos,video_metadata,user_video_metadata,playlists,playlist_videos",
    "database.history.kafka.bootstrap.servers": "kafka:29092",
    "database.history.kafka.topic": "dbhistory.video",
    "schema.history.internal.kafka.topic": "schemahistory.video",
    "schema.history.internal.kafka.bootstrap.servers": "kafka:29092",
    "database.parse_time": "true",
    "database.charset": "utf8mb4",
    "database.serverTimezone": "UTC",
    "topic.prefix": "video_database",
    "max_idle_conns": "10",
    "max_open_conns": "100",
    "max_lifetime": "3600"
  }
}'