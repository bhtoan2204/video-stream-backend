helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
kubectl create namespace kafka

helm install kafka bitnami/kafka \
  --namespace kafka \
  --set kraft.enabled=true \
  --set replicaCount=3 \
  --set controller.replicaCount=3 \
  --set externalAccess.enabled=false \
  --set listeners.client.protocol=PLAINTEXT \
  --set storageClass=local-path \
  --set persistence.size=8Gi

