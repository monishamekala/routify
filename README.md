# Routify – Real-Time Transit Optimizer and Trip Planner

**Personal Project**  
Tech Stack: Go, gRPC, Redis, PostgreSQL, Docker, Kubernetes, Mapbox, React, Prometheus

## 🔍 Overview

Routify is built to solve the challenge of planning efficient public transit trips across multiple transportation modes in real time. It features:

- 🚦 Real-time routing engine with Dijkstra/A* search
- 🌐 Mapbox-powered frontend for trip visualization
- ⚡ gRPC backend supporting 1000+ concurrent requests/sec
- 🧠 Live traffic simulation and Redis caching
- 📊 Full observability with Prometheus, Grafana, and Jaeger
- 🐳 Microservice deployment via Docker and Kubernetes

### ⚙️ Features

- **Real-Time Routing**: Computes optimal routes using Dijkstra/A* with traffic-aware weights.
- **High Throughput**: Handles up to 1,000 concurrent queries/sec with <150ms response via Redis + PostgreSQL.
- **Mapbox Integration**: Visualizes routes and simulations live in the frontend.
- **Modern Stack**: Built in Go with gRPC microservices, deployed on Kubernetes.
- **Observability**: Full metrics and alerting with Prometheus + Grafana.

## 🧱 Architecture

- **Backend**: Go + gRPC microservices
- **Frontend**: React + Mapbox
- **Database**: PostgreSQL for static data, Redis for caching
- **Orchestration**: Docker + Kubernetes
- **Monitoring**: Prometheus + Grafana dashboards

## 📦 Getting Started

```bash
# Clone the repo
git clone https://github.com/monishamekala/routify.git
cd routify-transit-optimizer

# Start services
docker-compose up --build
