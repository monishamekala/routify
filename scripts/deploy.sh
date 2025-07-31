#!/bin/bash

echo "🚀 Deploying Routify to Kubernetes..."

kubectl apply -f deployments/api-deployment.yaml
kubectl apply -f deployments/api-service.yaml

# Add other components if needed, e.g. Redis/Postgres
