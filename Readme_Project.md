#Data Ingestion Service

#Overview
This Go service ingests data from a public API, adds metadata, and stores the transformed data to AWS S3. It supports testing, Docker deployment, and cloud-native storage.

---

#Features
- Fetches data from `https://jsonplaceholder.typicode.com/posts`
- Transforms records by adding:
  - `ingested_at` (UTC timestamp)
  - `source` (e.g., "placeholder_api")
- 
Stores results:
  - Locally to JSON file (for testing)
  - To AWS S3 (production---- Bucket name & crdential)
- 
Includes:
  - Unit and integration tests (Script)
  - Docker containerization (YML File)
  - Modular codebase (`fetcher.go`, `transformer.go`, `storage.go`)

---

#How to Run

 Run locally (without Docker)
```bash
go run main.go
```

Run tests
```bash
go test ./...
```

 Run with Docker
```bash
docker build -t go-ingestor .
docker run go-ingestor
```

---

#File Structure

```
.
├── main.go
├── fetcher.go
├── transformer.go
├── storage.go
├── transformer_test.go
├── ingestion_test.go
├── go.mod
├── go.sum
├── Dockerfile
└── docker-compose.yml
```

---

AWS S3 Setup
environment is configured to use AWS credentials:

```bash
aws configure
```

Update `bucketName` and `objectKey` inside `storage.go`.

---

#Docker Compose

Use this to manage multi-service testing (e.g., local S3 via MinIO).

```bash
docker-compose up 
```

---

 Trade-offs and Notes
- Modular code, testable
- Clean transformation logic
- Ready for CI/CD container pipeline
- Add retry logic for HTTP and S3
- Add logging and metrics
- Add support for batching and message queues

---

##Future Improvements
- Retry and backoff mechanisms
- Observability (Prometheus + Grafana)
- Parallel ingestion pipelines


## Improvements if I Had More Time

Add S3 Writer
Right now, I’m saving the ingested data in a local JSON file for simplicity.
If I had more time, I would connect the app to Amazon S3 so that the data could be uploaded to the cloud.
That would make it easier to access from anywhere and would be more scalable.

Add REST API to Fetch Ingested Data
I could build a small API using something like net/http or Gin that lets users hit a URL like localhost:8080/data to view the ingested records.
This would make it easier to view or integrate the data in other systems.

## Trade-offs I Made
Used File Storage for Easier Testing
I chose to write the output to a local file (ingested_data.json) instead of using a database or S3.
This kept things simple during development and made it easy to test locally without setting up cloud permissions.

## Hardest Parts
Making the API Call Resilient
Handling API timeouts or failures gracefully was tricky.
I had to think about how to set a timeout, retry if needed, and handle any non-200 errors properly.

Designing in a Modular Way
Breaking the project into reusable pieces (like fetch, transform, store) took effort.
But doing it this way makes the code cleaner and easier to maintain or extend later.