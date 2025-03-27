#!/bin/bash

# install_helm.sh

set -e

# Check if helm is installed
if ! command -v helm &> /dev/null; then
  echo "Helm not found. Installing Helm..."
  curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
else
  echo "Helm is already installed."
fi

# Add HashiCorp Helm repo
helm repo add hashicorp https://helm.releases.hashicorp.com
helm repo update

echo "Helm setup complete."