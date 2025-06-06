# Task Scheduler (gRPC Job Queue)

## 🚀 Overview
A production-style, observable gRPC-based job queue with basic token-based access control. Built for reliability, observability, and developer-friendliness.

---

## 🛠️ Stack Summary
- **Go**: Main application and worker logic
- **gRPC**: API for submitting and tracking jobs
- **Redis**: Job queue (streams), job status, and job logs
- **Prometheus**: Metrics for observability
- **Docker Compose**: One-command local stack

---

## 🧱 Code Structure
```
.
├── cmd/                    # Application entry points
│   └── main.go            # Main application entry point
├── internal/              # Private application code
│   ├── metrics/          # Prometheus metrics definitions
│   ├── server/           # gRPC server implementation
│   └── worker/           # Job processing worker
├── proto/                # Protocol Buffer definitions
│   └── task.proto        # Service and message definitions
├── docker-compose.yml    # Docker services orchestration
├── Dockerfile           # Application container definition
├── prometheus.yml       # Prometheus configuration
├── go.mod              # Go module definition
└── README.md           # Project documentation
```

---

## 📜 Protobuf Definitions
Located in `/proto/task.proto`.

- `SubmitJob(JobRequest) returns (JobResponse)`
- `GetJobStatus(JobStatusRequest) returns (JobStatusResponse)`
- `GetJobLogs(JobStatusRequest) returns (JobLogsResponse)`

[📄 View the Protobuf definitions →](./proto/task.proto)

---

## ⚙️ How It Works
1. **Submit a Job**: Clients send jobs via gRPC (`SubmitJob`), including a payload and job ID. Authentication is enforced via a token in metadata.
2. **Job Queueing**: Jobs are enqueued in a Redis stream.
3. **Worker**: A Go worker consumes jobs, processes them, handles retries, and updates job status and logs in Redis.
4. **Job Status & Logs**: Clients can query job status (`GetJobStatus`) and fetch job logs (`GetJobLogs`).
5. **Observability**: Prometheus scrapes `/metrics` for job counts, failures, and latency histograms.

---

## 📈 Metrics Overview
Exposed at `/metrics`:
- `jobs_processed_total` — count of all processed jobs
- `jobs_failed_total` — count of failed jobs
- `job_duration_seconds` — histogram of processing latency
- `jobs_processed_by_type_total{type="email"}` — per-job-type counts
- `jobs_failed_by_type_total{type="email"}` — per-job-type failures

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

## 🔁 Retry Strategy
- Jobs are retried up to 3 times on failure
- Exponential backoff strategy (`2^retryCount` seconds)
- Final failures are tracked with `status=failed` in Redis

---

## 📦 Features
- Job queueing, retries, and status tracking
- Per-job logs stored in Redis
- Prometheus metrics for processed/failed jobs and latency
- Token-based authentication
- One-command local stack with Docker Compose

---

## 🧪 Testing
- Unit tests for core retry and processing logic (coming soon)
- Recommended: use [`redis-mock`](https://github.com/go-redis/redismock) or [`testcontainers-go`](https://github.com/testcontainers/testcontainers-go) for integration tests
- Even a small `worker_test.go` with a fake processor would impress

---

> This project was built to demonstrate distributed job queuing, retry logic, and observability using real-world tools. It's ideal as a foundation for task scheduling systems, async pipelines, or devtools infra.