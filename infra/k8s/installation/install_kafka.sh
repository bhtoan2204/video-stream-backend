helm install kafka bitnami/kafka -n kafka -f kafka-value.yml
kubectl run kafka-client --restart='Never' --image docker.io/bitnami/kafka:latest -n kafka --command -- sleep infinity
