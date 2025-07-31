#!/bin/bash

echo "🔧 Starting backend..."
go run ./cmd/api &

echo "🌐 Starting frontend..."
cd ui
npm install
npm run dev
