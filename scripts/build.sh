#!/bin/bash

echo "🐳 Building API Docker image..."
docker build -f docker/api.Dockerfile -t routify-api .

echo "📦 Building UI Docker image..."
docker build -f docker/ui.Dockerfile -t routify-ui .
