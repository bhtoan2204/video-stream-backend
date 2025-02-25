#!/bin/bash

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
    \"database.server.name\": \"dbserver1\",
    \"database.whitelist\": \"user,refresh_tokens\",
    \"database.history.kafka.bootstrap.servers\": \"kafka:29092\",
    \"database.history.kafka.topic\": \"dbhistory.user\",
    \"schema.history.internal.kafka.topic\": \"schemahistory.user\",
    \"schema.history.internal.kafka.bootstrap.servers\": \"kafka:29092\",
    \"database.parse_time\": \"true\",
    \"database.charset\": \"utf8mb4\",
    \"database.serverTimezone\": \"UTC\",
    \"topic.prefix\": \"dbserver1\",
    \"max_idle_conns\": \"10\",
    \"max_open_conns\": \"100\",
    \"max_lifetime\": \"3600\"
  }
}"


# mysql:
#   user: root
#   pass: root
#   host: mysql
#   port: 3306
#   name: user
#   parse_time: true
#   charset: utf8mb4
#   loc: UTC
#   max_idle_conns: 10
#   max_open_conns: 100
#   max_lifetime: 3600