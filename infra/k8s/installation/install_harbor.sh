# /bin/bash

helm repo add harbor https://helm.goharbor.io
helm repo update

kubectl create namespace harbor


helm install harbor harbor/harbor \
  -n harbor \
  --set harborAdminPassword=sasuketamin \
  --set persistence.persistentVolumeClaim.registry.size=10Gi \
  --set persistence.persistentVolumeClaim.database.size=5Gi \
  --set persistence.persistentVolumeClaim.jobservice.size=1Gi \
  --set persistence.persistentVolumeClaim.redis.size=1Gi \
  --set persistence.persistentVolumeClaim.trivy.size=5Gi \
  --set trivy.enabled=true