# Task Scheduler (gRPC Job Queue)

## 🚀 Overview
A production-grade, observable, and secure gRPC-based job queue system with Redis, Prometheus metrics, and Docker Compose orchestration.

---

## 🛠️ Stack Summary
- **Go**: Main application and worker logic
- **gRPC**: API for submitting and tracking jobs
- **Redis**: Job queue (streams), job status, and job logs
- **Prometheus**: Metrics for observability
- **Docker Compose**: One-command local stack

---

## ⚙️ How It Works
1. **Submit a Job**: Clients send jobs via gRPC (`SubmitJob`), including a payload and job ID. Authentication is enforced via a token in metadata.
2. **Job Queueing**: Jobs are enqueued in a Redis stream.
3. **Worker**: A Go worker consumes jobs, processes them, handles retries, and updates job status and logs in Redis.
4. **Job Status & Logs**: Clients can query job status (`GetJobStatus`) and fetch job logs (`GetJobLogs`).
5. **Observability**: Prometheus scrapes `/metrics` for job counts, failures, and latency histograms.

---

## 🗺️ Architecture Diagram

```
+-------------------+         +-------------------+         +-------------------+
|    gRPC Client    | <-----> |      App/API      | <-----> |      Redis        |
| (grpcurl/Postman) |         |  (Go + gRPC)      |         | (Streams, Hashes) |
+-------------------+         +-------------------+         +-------------------+
         |                            |                             |
         |                            v                             |
         |                  +-------------------+                   |
         |                  |     Worker        |-------------------+
         |                  |  (Go, Prometheus) |                   |
         |                  +-------------------+                   |
         |                            |                             |
         |                            v                             |
         |                  +-------------------+                   |
         |                  |   Prometheus      |                   |
         |                  +-------------------+                   |
```

---

## 📝 gRPC Sample Request (grpcurl)

**Submit a Job:**
```sh
grpcurl -plaintext \
  -H "authorization: my-secret-token" \
  -d '{
    "job_id": "test-1",
    "payload": "{\"type\":\"email\",\"data\":{\"to\":\"test@example.com\",\"subject\":\"Test Email\",\"body\":\"This is a test email\"},\"timeout\":30}"
  }' \
  localhost:50053 task.TaskService/SubmitJob
```

**Get Job Status:**
```sh
grpcurl -plaintext \
  -H "authorization: my-secret-token" \
  -d '{"job_id": "test-1"}' \
  localhost:50053 task.TaskService/GetJobStatus
```

**Get Job Logs:**
```sh
grpcurl -plaintext \
  -H "authorization: my-secret-token" \
  -d '{"job_id": "test-1"}' \
  localhost:50053 task.TaskService/GetJobLogs
```

---

## 🐳 Quick Start (Docker Compose)
```sh
docker-compose up --build
```
- App: gRPC on `localhost:50053`, metrics on `localhost:2113/metrics`
- Redis: `localhost:6379`
- Prometheus: `localhost:9090`

---

## 🔒 Authentication
- All gRPC endpoints require an `authorization` token in metadata.
- Default token: `my-secret-token` (set `AUTH_TOKEN` env var to override)

---

## 📦 Features
- Job queueing, retries, and status tracking
- Per-job logs stored in Redis
- Prometheus metrics for processed/failed jobs and latency
- Token-based authentication
- One-command local stack with Docker Compose

---

## 🏆 Stretch Goals
- Dead-letter queue for failed jobs
- Job priority levels
- Kubernetes deployment
- Web frontend for job submission/tracking

---

## 📣 Show it off!
- Publish this repo to GitHub
- Share your architecture diagram and Prometheus dashboard on LinkedIn or a blog
- This is FAANG-level backend engineering—congrats!