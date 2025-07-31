# Docker Files

- `api.Dockerfile`: Builds and runs the Go gRPC backend
- `ui.Dockerfile`: Builds and serves the React frontend with nginx

To build:
```bash
docker build -f docker/api.Dockerfile -t routify-api .
docker build -f docker/ui.Dockerfile -t routify-ui .
