helm repo add jenkins https://charts.jenkins.io
helm repo update
helm install jenkins jenkins/jenkins \
  --set controller.serviceType=ClusterIP \
  --set persistence.enabled=true \
  --set persistence.size=10Gi \
  --namespace jenkins --create-namespace
