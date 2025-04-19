# /bin/bash
# This script installs Redis using Helm and sets up a Redis Cluster with 3 master nodes and 3 replica nodes.
helm install redis-single \
  oci://registry-1.docker.io/bitnamicharts/redis \
  --namespace redis \
  --set architecture=standalone \
  --set auth.enabled=false \
  --set persistence.size=8Gi

# optional
helm install redisinsight redisinsight/redisinsight \
  --namespace redis \
  --set service.type=ClusterIP            \
  --set ingress.enabled=false             \
  --set resources.requests.memory=256Mi   \
  --set persistence.enabled=true
