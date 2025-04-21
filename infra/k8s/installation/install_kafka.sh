helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
kubectl create namespace kafka

helm upgrade --install kafka bitnami/kafka \
  --namespace kafka \
  --create-namespace \
  --set persistence.size=8Gi,logPersistence.size=8Gi,volumePermissions.enabled=true,persistence.enabled=true,logPersistence.enabled=true,serviceAccount.create=true,rbac.create=true \
  --version 23.0.7 \
  -f kafka-values.yaml




