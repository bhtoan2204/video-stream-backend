#!/bin/bash

until curl -s http://localhost:8083/ | grep -q "connectors"; do
  sleep 5
done

curl -X POST -H "Content-Type: application/json" --data '{
  "name": "mysql-connector",
  "config": {
    "connector.class": "io.debezium.connector.mysql.MySqlConnector",
    "tasks.max": "1",
    "database.hostname": "mysql",
    "database.port": "3306",
    "database.user": "root",
    "database.password": "root",
    "database.server.id": "184054",
    "database.server.name": "mysql_server",
    "database.include.list": "appdb",
    "table.include.list": "appdb.user",
    "database.history.kafka.bootstrap.servers": "kafka:29092",
    "database.history.kafka.topic": "dbhistory.appdb"
  }
}' http://localhost:8083/connectors

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