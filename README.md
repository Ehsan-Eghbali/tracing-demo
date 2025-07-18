# Tracing Demo Project
This project is a simple Go application designed to demonstrate distributed tracing concepts using Jaeger. It's an educational resource for learning observability patterns in microservices.
---
## 🚀 Running with Docker
### 1. Build Docker Image:
    docker build -t tracing-demo .
### 2. Run Application Container:
    docker run --rm tracing-demo
---
## 🕵️‍♂️ Setting Up Jaeger for Trace Visualization

Run Jaeger in Docker to collect and view tracing data:
```
docker run -d --name jaeger \
-e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
-p 16686:16686 \
-p 14268:14268 \
jaegertracing/all-in-one:latest
```
- Access the Jaeger dashboard at:  
  `http://localhost:16686`
---
**GitHub Repository**:  
https://github.com/Ehsan-Eghbali/tracing-demo.git
