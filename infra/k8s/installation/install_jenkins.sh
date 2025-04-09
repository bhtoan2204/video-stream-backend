helm repo add jenkins https://charts.jenkins.io
helm repo update
helm install jenkins jenkins/jenkins --set controller.installPlugins="{kubernetes-1.29.0,workflow-aggregator-2.7,git-4.10.0,blueocean-1.25.3,configuration-as-code-1.58}" --set controller.serviceType=LoadBalancer --set persistence.enabled=true --set persistence.size=10Gi --set persistence.storageClass=gp2