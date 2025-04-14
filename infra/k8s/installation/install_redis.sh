# /bin/bash
# This script installs Redis using Helm and sets up a Redis Cluster with 3 master nodes and 3 replica nodes.
helm install redis bitnami/redis -n redis --set architecture=replication --set cluster.enabled=true --set cluster.nodes=3 --set cluster.slaveCount=3 --set auth.enabled=false