#!/bin/bash

# install_vault.sh

set -e

# Namespace for Vault
NAMESPACE="vault"

# Create namespace if not exists
kubectl get namespace "$NAMESPACE" &> /dev/null || kubectl create namespace "$NAMESPACE"

# Install Vault in dev mode
helm install vault hashicorp/vault \
  --namespace "$NAMESPACE" \
  --set "server.dev.enabled=true"

echo "Vault has been installed in dev mode in namespace '$NAMESPACE'."

# Optional: Wait for pod to be ready
echo "Waiting for Vault pod to be ready..."
kubectl wait --for=condition=ready pod -l app.kubernetes.io/name=vault -n "$NAMESPACE" --timeout=120s

# Optional: Port forward command
echo "Run the following to port-forward Vault:"
echo "kubectl port-forward svc/vault 8200:8200 -n $NAMESPACE"
