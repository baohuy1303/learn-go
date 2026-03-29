# Background Job Queue — 2-Week To-Do List

---

## Week 1 — Foundation
> Go + PostgreSQL + Redis + core queue engine

---

### Day 1 — Go Fundamentals I
**Goal:** Get comfortable with Go syntax coming from JS/TypeScript

#### Setup
- [x] Install Go 1.22+ from go.dev/dl
- [x] Install VS Code extension: `golang.go` (gopls)
- [x] Run `go version` to confirm install
- [x] Read the VS Code Go extension setup guide — enable format on save, lint on save
- [x] Create your first module: `mkdir go-practice && cd go-practice && go mod init practice`

#### Language basics
- [x] Learn basic types: `string`, `int`, `float64`, `bool`, `byte`
- [x] Understand `:=` (short declaration) vs `var` — know when each is used
- [x] Write a function with multiple return values: `func divide(a, b float64) (float64, error)`
- [x] Understand zero values — what is the zero value of int, string, bool, pointer, slice, map
- [ ] Learn named returns and when NOT to use them
- [ ] Understand `const` and `iota` for enums
- [x] Learn `for` loop variants — C-style, range, while-style (Go has no while)
- [ ] Understand `defer` — what order do multiple defers execute?

#### Practice
- [ ] Complete go.dev/tour sections: Basics, Flow control, More types (skip Methods and Interfaces for today)
- [ ] Write a CLI calculator: takes two numbers and an operator (+, -, *, /) as args, prints result
- [ ] Make it handle division by zero correctly using multiple return values

---

### Day 2 — Go Fundamentals II + Interfaces
**Goal:** Understand Go's type system, interfaces, and error handling

#### Structs and methods
- [ ] Define a struct with exported (uppercase) and unexported (lowercase) fields
- [ ] Write methods with value receivers vs pointer receivers — understand the difference
- [ ] Write a constructor function `NewJob(...)` — Go convention for initializing structs
- [ ] Understand struct embedding — Go's version of inheritance
- [ ] Learn struct tags: `json:"name"` and `db:"name"`

#### Interfaces
- [ ] Understand implicit interface satisfaction — no `implements` keyword
- [ ] Write a `JobHandler` interface: `Execute(ctx context.Context, payload []byte) error`
- [ ] Write two structs that satisfy it: `EmailHandler` and `ReportHandler`
- [ ] Understand the empty interface `any` / `interface{}` and when it's used
- [ ] Learn type assertions: `h, ok := handler.(EmailHandler)`
- [ ] Learn type switches

#### Error handling
- [ ] Understand `errors.New` vs `fmt.Errorf`
- [ ] Learn error wrapping: `fmt.Errorf("processing job: %w", err)`
- [ ] Learn `errors.Is` and `errors.As` for unwrapping
- [ ] Write a custom error type: `type JobError struct { JobID string; Err error }`
- [ ] Understand when to return errors vs panic

#### Build target
- [ ] Define your `Job` struct with fields: `ID string`, `Type string`, `Payload []byte`, `Status string`, `Attempts int`, `MaxAttempts int`, `CreatedAt time.Time`, `RunAt time.Time`
- [ ] Define `JobStatus` constants using iota: `StatusPending`, `StatusRunning`, `StatusCompleted`, `StatusFailed`, `StatusDead`
- [ ] Write a `Validate()` method on Job that returns an error if Type is empty or Payload is nil

---

### Day 3 — Go Concurrency + REST API
**Goal:** Understand goroutines and build your first HTTP server

#### Concurrency
- [ ] Understand goroutines: `go func()` — fire and forget
- [ ] Learn unbuffered vs buffered channels — `make(chan Job)` vs `make(chan Job, 100)`
- [ ] Understand channel direction: `chan<- Job` (send only) vs `<-chan Job` (receive only)
- [ ] Learn `select` statement — multiplex across channels, default case
- [ ] Learn `sync.WaitGroup` — wait for N goroutines to finish
- [ ] Learn `sync.Mutex` and `sync.RWMutex` — protect shared state
- [ ] Understand `context.Context` — `WithCancel`, `WithTimeout`, `WithDeadline`
- [ ] Write a worker pool: 5 goroutines all reading from a shared job channel

#### HTTP API
- [ ] Install chi: `go get github.com/go-chi/chi/v5`
- [ ] Install chi middleware: `go get github.com/go-chi/chi/v5/middleware`
- [ ] Write a basic HTTP server on port 8080
- [ ] Add chi middleware: `middleware.Logger`, `middleware.Recoverer`, `middleware.RealIP`
- [ ] Define routes: `POST /api/v1/jobs`, `GET /api/v1/jobs/{id}`, `GET /api/v1/jobs`, `DELETE /api/v1/jobs/{id}`
- [ ] Write request/response structs with JSON tags
- [ ] Handle JSON decode errors — return 400 with descriptive message
- [ ] Add `GET /health` endpoint returning `{"status": "ok", "timestamp": "..."}`

#### Build target
- [ ] In-memory job store: `map[string]*Job` protected by `sync.RWMutex`
- [ ] `POST /api/v1/jobs` stores job in map, returns 201 with job ID
- [ ] `GET /api/v1/jobs/{id}` retrieves job from map, returns 404 if not found
- [ ] Worker goroutine reads from a channel, sleeps 1 second (simulating work), marks job complete

---

### Day 4 — PostgreSQL + pgx
**Goal:** Replace in-memory store with a real database

#### PostgreSQL setup
- [ ] Install PostgreSQL locally (Homebrew: `brew install postgresql@16`)
- [ ] Start PostgreSQL service
- [ ] Create database: `createdb jobqueue`
- [ ] Connect with psql: `psql jobqueue`
- [ ] Learn basic psql commands: `\dt`, `\d table_name`, `\timing`, `\x` for expanded output

#### Schema design
- [ ] Write `001_create_jobs.sql` migration file
- [ ] Jobs table columns: `id UUID PRIMARY KEY DEFAULT gen_random_uuid()`, `type VARCHAR(255) NOT NULL`, `payload JSONB NOT NULL DEFAULT '{}'`, `status VARCHAR(50) NOT NULL DEFAULT 'pending'`, `attempts INT NOT NULL DEFAULT 0`, `max_attempts INT NOT NULL DEFAULT 3`, `last_error TEXT`, `created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()`, `updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()`, `run_at TIMESTAMPTZ NOT NULL DEFAULT NOW()`, `completed_at TIMESTAMPTZ`
- [ ] Add indexes: `(status, run_at)` for worker polling, `(type, status)` for filtering
- [ ] Write a `dead_letters` table for permanently failed jobs: same columns as jobs plus `failed_at TIMESTAMPTZ`
- [ ] Apply migration: `psql jobqueue < migrations/001_create_jobs.sql`

#### pgx connection
- [ ] Install pgx: `go get github.com/jackc/pgx/v5`
- [ ] Write `internal/db/db.go` — creates a pgxpool connection pool
- [ ] Configure pool: `MaxConns: 20`, `MinConns: 2`, connection timeout
- [ ] Write a `Ping(ctx)` method to verify connection on startup
- [ ] Use `os.Getenv("DATABASE_URL")` — no hardcoded connection strings

#### Repository layer
- [ ] Write `internal/repository/job_repo.go`
- [ ] `CreateJob(ctx, job) error` — INSERT with pgx named params
- [ ] `GetJob(ctx, id) (*Job, error)` — SELECT by ID, return `ErrNotFound` if missing
- [ ] `ListJobs(ctx, filter) ([]*Job, error)` — filter by status, type, limit/offset
- [ ] `UpdateJobStatus(ctx, id, status, lastError) error` — UPDATE with `updated_at = NOW()`
- [ ] `ClaimNextJob(ctx, workerID) (*Job, error)` — `SELECT ... FOR UPDATE SKIP LOCKED LIMIT 1` wrapped in a transaction
- [ ] Run `EXPLAIN ANALYZE` on your ClaimNextJob query — understand the query plan

#### Build target
- [ ] Replace in-memory map with PostgreSQL repository
- [ ] All CRUD endpoints read/write from real DB
- [ ] Worker uses `ClaimNextJob` to atomically claim jobs — no two workers process the same job

---

### Day 5 — Core Worker + Retry Logic
**Goal:** Build the actual job processing engine

#### Worker pool
- [ ] Write `internal/worker/pool.go`
- [ ] `Pool` struct: holds concurrency count, job repo, handler registry, stop channel
- [ ] `Start(ctx)` launches N worker goroutines
- [ ] `Stop()` gracefully shuts down — wait for in-flight jobs to finish
- [ ] Each goroutine loops: claim job → execute → update status → repeat
- [ ] Add configurable poll interval (default 1 second) when no jobs are available
- [ ] Use `context.WithTimeout` per job execution — jobs can't run forever

#### Handler registry
- [ ] Write `internal/worker/registry.go`
- [ ] `Registry` struct: `map[string]JobHandler`
- [ ] `Register(jobType string, handler JobHandler)` 
- [ ] `Execute(ctx, job) error` — looks up handler by job type, returns error if not found
- [ ] Add panic recovery: `defer func() { if r := recover(); r != nil { ... } }()` — a panicking handler must not crash the worker

#### Retry logic
- [ ] On job failure: increment `attempts` counter in DB
- [ ] Implement exponential backoff: `nextRunAt = now + (baseDelay * 2^attempts) + jitter`
- [ ] Base delay: 30 seconds. Max delay cap: 1 hour
- [ ] Jitter: add random 0–10 seconds to prevent thundering herd
- [ ] If `attempts >= max_attempts`: move job to `dead_letters` table, set status to `dead`
- [ ] Write `MoveToDeadLetter(ctx, job, lastError) error` in repository
- [ ] Log every retry attempt with job ID, attempt number, next run time, error message

#### Build target
- [ ] Enqueue a job via API that always fails → watch it retry 3 times → appear in dead_letters
- [ ] Enqueue 10 jobs with 3 concurrent workers → all complete without duplicates
- [ ] Kill the API mid-processing → restart → in-progress jobs reset to pending on startup

---

### Day 6 — Redis + Idempotency
**Goal:** Add Redis for idempotency, deduplication, and rate limiting

#### Redis setup
- [ ] Install Redis locally (Homebrew: `brew install redis`)
- [ ] Start Redis: `redis-server`
- [ ] Install go-redis: `go get github.com/redis/go-redis/v9`
- [ ] Write `internal/cache/redis.go` — client setup, Ping on startup
- [ ] Learn Redis commands in redis-cli: `SET`, `GET`, `SETNX`, `EXPIRE`, `TTL`, `DEL`, `INCR`, `ZADD`

#### Idempotency keys
- [ ] `POST /api/v1/jobs` now accepts `Idempotency-Key` header (UUID from client)
- [ ] Write `idempotency/middleware.go`
- [ ] On request: check Redis for key `idempotency:{key}`
- [ ] If exists: return cached response immediately (HTTP 200, same body as original)
- [ ] If not exists: process request, store response in Redis with 24-hour TTL
- [ ] Return `Idempotency-Replayed: true` header on cache hit
- [ ] Test: send same POST twice with same key → second call returns identical response, job not duplicated in DB

#### Job deduplication
- [ ] Add optional `DeduplicationKey` field to job creation request
- [ ] If provided: check Redis for `dedup:{key}` before inserting
- [ ] If key exists: return existing job ID, don't insert duplicate
- [ ] Set `dedup:{key}` with TTL equal to job's scheduled run time + 1 hour

#### Rate limiting
- [ ] Write `ratelimit/middleware.go` using sliding window algorithm
- [ ] Key pattern: `ratelimit:{api_key}:{window_start}`
- [ ] `INCR` counter, `EXPIRE` on first request in window
- [ ] Allow 1000 requests/minute per API key — return `HTTP 429` with `Retry-After` header when exceeded
- [ ] Add rate limit headers to every response: `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset`

#### Build target
- [ ] Send identical `POST /jobs` with same `Idempotency-Key` 5 times → 1 job in DB, all 5 return same response
- [ ] Send 1001 requests in a minute → 1001st gets 429
- [ ] Enqueue same job type twice with same `DeduplicationKey` → only 1 job created

---

### Day 7 — Auth + API Keys
**Goal:** Secure every endpoint with real authentication

#### User model
- [ ] Write `002_create_users.sql` migration
- [ ] Users table: `id`, `email`, `password_hash`, `role` (admin/member), `created_at`
- [ ] API keys table: `id`, `user_id FK`, `key_hash` (bcrypt), `key_prefix` (first 8 chars for display), `name`, `last_used_at`, `expires_at`, `created_at`
- [ ] Install bcrypt: `go get golang.org/x/crypto/bcrypt`

#### JWT auth
- [ ] Install jwt library: `go get github.com/golang-jwt/jwt/v5`
- [ ] `POST /auth/register` — hash password with bcrypt cost 12, store user
- [ ] `POST /auth/login` — verify password, return access token (15min TTL) + refresh token (7 day)
- [ ] Access token in response body, refresh token in `httpOnly` `Secure` cookie
- [ ] `POST /auth/refresh` — validate refresh token from cookie, issue new access token
- [ ] `POST /auth/logout` — add refresh token to Redis blocklist with TTL matching remaining token lifetime: `SET blocklist:{jti} 1 EX {remaining_seconds}`
- [ ] JWT claims: `sub` (user ID), `jti` (unique token ID for blocklist), `role`, `exp`

#### API key auth
- [ ] `POST /api/v1/keys` — generate 32-byte random key, prefix with `jq_live_`, bcrypt hash before storing, return plaintext key ONCE (never retrievable again)
- [ ] `GET /api/v1/keys` — list keys by prefix + name (never expose hash)
- [ ] `DELETE /api/v1/keys/{id}` — revoke key
- [ ] Auth middleware: check `Authorization: Bearer {token}` — try JWT first, then API key lookup

#### RBAC
- [ ] Middleware: `RequireRole(roles ...string)` — reads role from JWT claims
- [ ] Admin-only routes: `GET /admin/jobs` (all users' jobs), `GET /admin/metrics`
- [ ] User routes: only see their own jobs (filter by `user_id` in all queries)
- [ ] Add `user_id UUID FK` to jobs table: `003_add_user_id_to_jobs.sql`

#### Build target
- [ ] Register → login → get JWT → create job with Bearer token → get job → logout → confirm token rejected
- [ ] Create API key → use it to enqueue a job → revoke key → confirm it no longer works

---

## Week 2 — Infrastructure
> WebSockets + Docker + Testing + CI/CD + Kubernetes + Ship

---

### Day 8 — WebSockets + Real-time Dashboard
**Goal:** Build the live job feed that makes your demo compelling

#### Backend WebSocket server
- [ ] Install gorilla/websocket: `go get github.com/gorilla/websocket`
- [ ] Write `internal/ws/hub.go` — central Hub managing all client connections
- [ ] Hub fields: `clients map[*Client]bool`, `broadcast chan []byte`, `register chan *Client`, `unregister chan *Client`
- [ ] `Hub.Run()` goroutine: handles register/unregister/broadcast in a single goroutine (no mutex needed — single owner pattern)
- [ ] Write `internal/ws/client.go` — one per connected browser
- [ ] Client read pump: reads messages from browser (commands), handles close
- [ ] Client write pump: pulls from send channel, writes to WebSocket with write deadline
- [ ] Implement ping/pong heartbeat — client must respond to server pings within 10s or connection is closed
- [ ] WebSocket upgrade endpoint: `GET /ws` — validate JWT before upgrading
- [ ] Scope broadcasts by user ID — users only see their own job events

#### Event emission
- [ ] Write `internal/events/emitter.go` — thin wrapper around Hub.Broadcast
- [ ] Define event types: `job.created`, `job.started`, `job.completed`, `job.failed`, `job.retrying`, `job.dead`
- [ ] Event payload: `{ "event": "job.completed", "data": { ...job fields } }`
- [ ] Hook emitter into worker pool — emit after every status transition
- [ ] Hook emitter into API — emit `job.created` after successful POST

#### Frontend
- [ ] Create Next.js app: `npx create-next-app@latest frontend`
- [ ] Write `useJobStream` hook — connects to WebSocket, maintains job list in state, handles reconnect with exponential backoff
- [ ] Job card component: type badge, status badge (color-coded), attempt counter, duration, error message (expandable)
- [ ] Job feed: sorted by `created_at DESC`, max 100 items, older items fade out
- [ ] Status badge colors: pending=gray, running=blue (pulse animation), completed=green, failed=red, dead=red+skull
- [ ] Enqueue button panel with demo job types: slow job (5s), fast job (100ms), failing job, batch (enqueue 20 at once)
- [ ] Concurrency slider: adjust worker pool size 1–20, see throughput change in real time
- [ ] Stats bar: jobs/sec throughput, success rate %, avg duration

#### Build target
- [ ] Open dashboard, click "Enqueue 20 jobs", watch all 20 process live with status transitions
- [ ] Click "Failing job" — watch it fail, show retry countdown, retry, eventually go dead
- [ ] Drag concurrency slider from 2 to 10 — visibly faster throughput

---

### Day 9 — Docker + Multi-service Setup
**Goal:** Run the entire stack with one command

#### Dockerfile
- [ ] Write multi-stage `Dockerfile`
- [ ] Stage 1 (builder): `FROM golang:1.22-alpine`, copy go.mod + go.sum, `go mod download`, copy source, `CGO_ENABLED=0 go build -ldflags="-s -w" -o /app ./cmd/server`
- [ ] Stage 2 (final): `FROM scratch`, copy binary + CA certs (`ca-certificates`), copy migrations directory
- [ ] Set `USER 1001` — don't run as root
- [ ] `EXPOSE 8080`
- [ ] `ENTRYPOINT ["/app"]`
- [ ] Confirm final image is under 20MB: `docker image ls`
- [ ] Write separate `Dockerfile.worker` for the worker binary

#### docker-compose.yml
- [ ] Service: `postgres` — `postgres:16-alpine`, named volume for data, healthcheck with `pg_isready`
- [ ] Service: `redis` — `redis:7-alpine`, named volume, healthcheck with `redis-cli ping`
- [ ] Service: `api` — built from Dockerfile, `depends_on: {postgres: {condition: service_healthy}, redis: {condition: service_healthy}}`, env vars from `.env`
- [ ] Service: `worker` — built from Dockerfile.worker, same dependencies as api
- [ ] Service: `prometheus` — `prom/prometheus:latest`, mount `prometheus.yml` config
- [ ] Service: `grafana` — `grafana/grafana:latest`, mount dashboard JSON, provision datasource
- [ ] Named volumes: `postgres_data`, `redis_data`, `grafana_data`
- [ ] All services on same Docker network

#### Configuration
- [ ] Write `.env.example` with all required variables and descriptions
- [ ] Write `.env` for local dev (in .gitignore)
- [ ] Variables: `DATABASE_URL`, `REDIS_URL`, `JWT_SECRET`, `JWT_REFRESH_SECRET`, `PORT`, `WORKER_CONCURRENCY`, `LOG_LEVEL`, `ENVIRONMENT`
- [ ] API reads all config from env — no hardcoded values anywhere

#### Build target
- [ ] `docker compose up --build` starts entire stack
- [ ] API healthy at `localhost:8080/health`
- [ ] Prometheus scraping at `localhost:9090`
- [ ] Grafana accessible at `localhost:3000`
- [ ] `docker compose down -v` tears down everything including volumes

---

### Day 10 — Testing
**Goal:** Write tests that actually catch bugs and run in CI

#### Test setup
- [ ] Install testcontainers: `go get github.com/testcontainers/testcontainers-go`
- [ ] Write `internal/testutil/db.go` — spins up a real PostgreSQL container for tests, runs migrations, returns connection pool, tears down after test
- [ ] Write `internal/testutil/redis.go` — same for Redis
- [ ] Write `TestMain` in repository package — start containers once, share across all tests in package

#### Unit tests
- [ ] Test backoff calculation: `TestExponentialBackoff` — assert correct delays for attempts 1–5, assert max cap respected, assert jitter within range
- [ ] Test job validation: `TestJob_Validate` — empty type, nil payload, negative max attempts
- [ ] Test idempotency key logic: `TestIdempotencyKey` — first call processes, second returns cached, expired key re-processes
- [ ] Test rate limiter: `TestSlidingWindowRateLimiter` — under limit succeeds, at limit succeeds, over limit returns error
- [ ] Test JWT: `TestGenerateAndValidateToken` — valid token, expired token, tampered token, revoked token
- [ ] Test RBAC middleware: `TestRequireRole` — correct role passes, wrong role returns 403, no token returns 401

#### Integration tests
- [ ] `TestJobLifecycle_Success` — enqueue job via API → worker processes → assert status=completed in DB, WebSocket event received
- [ ] `TestJobLifecycle_RetryAndDead` — enqueue always-failing job → assert 3 retries with correct delays → assert status=dead in dead_letters
- [ ] `TestJobLifecycle_Idempotency` — enqueue same job twice with same idempotency key → assert 1 job in DB
- [ ] `TestWorkerPool_NoDuplicates` — 100 jobs, 10 concurrent workers → assert each job processed exactly once
- [ ] `TestWorkerPool_GracefulShutdown` — start processing, call Stop(), assert in-flight jobs complete before shutdown
- [ ] `TestAuth_FullFlow` — register, login, use token, refresh, logout, confirm revoked

#### Code quality
- [ ] Set up `.golangci.yml` config — enable `errcheck`, `staticcheck`, `gosec`, `govet`, `gofmt`
- [ ] Run `golangci-lint run ./...` — fix all warnings
- [ ] Run `go test ./... -race` — fix any data races
- [ ] Aim for >80% coverage on repository and worker packages: `go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out`

#### Build target
- [ ] `go test ./... -race` passes with zero failures
- [ ] `golangci-lint run ./...` passes with zero warnings
- [ ] Coverage report shows >80% on core packages

---

### Day 11 — Observability: Logging + Metrics
**Goal:** Make the system debuggable in production

#### Structured logging
- [ ] Install zerolog: `go get github.com/rs/zerolog`
- [ ] Write `internal/logger/logger.go` — singleton zerolog instance, configured from `LOG_LEVEL` env var
- [ ] Every log line must include: `service`, `environment`, `version` (from build-time ldflags)
- [ ] Job-related logs must include: `job_id`, `job_type`, `worker_id`, `attempt`
- [ ] Request logs must include: `method`, `path`, `status`, `duration_ms`, `request_id`
- [ ] Error logs must include: `error`, stack trace for unexpected errors
- [ ] Use log levels correctly: `DEBUG` for verbose dev info, `INFO` for state transitions, `WARN` for retries, `ERROR` for failures
- [ ] Replace all `fmt.Println` and `log.Println` with zerolog

#### Prometheus metrics
- [ ] Install client_golang: `go get github.com/prometheus/client_golang`
- [ ] Write `internal/metrics/metrics.go` — define all metrics as package-level vars
- [ ] `jobs_enqueued_total` — Counter, labels: `job_type`
- [ ] `jobs_processed_total` — Counter, labels: `job_type`, `status` (completed/failed/dead)
- [ ] `job_processing_duration_seconds` — Histogram, labels: `job_type`, buckets: 0.1s, 0.5s, 1s, 5s, 10s, 30s
- [ ] `job_queue_depth` — Gauge, labels: `status` (pending/running/failed) — update every 30s
- [ ] `worker_active_count` — Gauge — increment on job start, decrement on finish
- [ ] `http_requests_total` — Counter, labels: `method`, `path`, `status_code`
- [ ] `http_request_duration_seconds` — Histogram, labels: `method`, `path`
- [ ] `redis_operations_total` — Counter, labels: `operation`, `result` (hit/miss)
- [ ] Mount `/metrics` endpoint with `promhttp.Handler()`

#### Prometheus + Grafana config
- [ ] Write `prometheus.yml` — scrape `api:8080/metrics` every 15s
- [ ] Create Grafana dashboard JSON with panels:
  - [ ] Jobs/sec throughput (rate over 1m)
  - [ ] Job success vs failure rate (stacked bar)
  - [ ] Queue depth over time (line chart)
  - [ ] p50 / p95 / p99 job duration (using histogram_quantile)
  - [ ] Active worker count
  - [ ] HTTP request rate by endpoint
  - [ ] Redis cache hit ratio
- [ ] Set up Grafana alert: fire if queue depth > 1000 for 5 minutes

#### Build target
- [ ] Enqueue 100 jobs → open Grafana → see throughput spike, queue depth rise then fall, duration distribution
- [ ] Every log line is valid JSON with all required fields
- [ ] `curl localhost:8080/metrics` shows all custom metrics with correct labels

---

### Day 12 — CI/CD with GitHub Actions
**Goal:** Automate testing, building, and deploying

#### Repository setup
- [ ] Push project to GitHub (private repo)
- [ ] Create `.github/workflows/` directory
- [ ] Add secrets to GitHub repo: `KUBECONFIG`, `GHCR_TOKEN`, `DATABASE_URL_TEST`, `REDIS_URL_TEST`
- [ ] Set up branch protection on `main`: require PR, require status checks to pass, no force push

#### Workflow 1: PR checks (runs on every PR)
- [ ] Trigger: `on: pull_request`
- [ ] Job 1 — lint: checkout, setup Go 1.22, cache Go modules, run `golangci-lint run ./...`
- [ ] Job 2 — test: checkout, setup Go, start PostgreSQL + Redis via `services:` block in Actions, run `go test ./... -race -count=1`
- [ ] Job 3 — build: confirm `go build ./...` succeeds
- [ ] All three jobs must pass for PR to be mergeable

#### Workflow 2: Build and push (runs on merge to main)
- [ ] Trigger: `on: push: branches: [main]`
- [ ] Login to GHCR: `docker/login-action`
- [ ] Build and push API image: `docker/build-push-action`, tag with `latest` and git SHA
- [ ] Build and push worker image: same
- [ ] Output image digest for deployment step

#### Workflow 3: Deploy (runs after successful build)
- [ ] Write `kubeconfig` secret to `~/.kube/config`
- [ ] Run `helm upgrade --install jobqueue ./helm/jobqueue --set image.tag=${{ github.sha }} --namespace prod`
- [ ] Wait for rollout: `kubectl rollout status deployment/jobqueue-api -n prod`
- [ ] Smoke test: `curl https://yourdomain.com/health` — fail workflow if non-200

#### golangci-lint config
- [ ] Write `.golangci.yml`
- [ ] Enable linters: `errcheck`, `staticcheck`, `gosec`, `govet`, `gofmt`, `goimports`, `unused`, `misspell`
- [ ] Set `run.timeout: 5m`
- [ ] Exclude test files from `gosec`

#### Build target
- [ ] Open a PR → see all 3 check jobs run in GitHub Actions tab
- [ ] Merge PR → image built, pushed to GHCR, deployed to K8s automatically
- [ ] Intentionally break a test → confirm PR is blocked

---

### Day 13 — Kubernetes: Production Deployment
**Goal:** Run on a real cluster with TLS, scaling, and persistence

#### Cluster setup
- [ ] Provision Hetzner VPS (CX21 — 2vCPU, 4GB RAM, $6/month)
- [ ] SSH in, install k3s: `curl -sfL https://get.k3s.io | sh -`
- [ ] Copy kubeconfig to local machine: `scp root@your-ip:/etc/rancher/k3s/k3s.yaml ~/.kube/config`
- [ ] Confirm cluster: `kubectl get nodes` — should show Ready
- [ ] Install Helm locally: `brew install helm`

#### Cluster tooling
- [ ] Install cert-manager: `helm install cert-manager jetstack/cert-manager --set installCRDs=true -n cert-manager --create-namespace`
- [ ] Create `ClusterIssuer` for Let's Encrypt (staging first, then prod)
- [ ] Install nginx ingress: `helm install ingress-nginx ingress-nginx/ingress-nginx -n ingress-nginx --create-namespace`
- [ ] Point your domain DNS A record to the VPS IP
- [ ] Confirm nginx ingress gets an external IP: `kubectl get svc -n ingress-nginx`

#### Kubernetes manifests
- [ ] Create namespaces: `prod`, `staging`, `monitoring`
- [ ] Install sealed-secrets: `helm install sealed-secrets sealed-secrets/sealed-secrets -n kube-system`
- [ ] Seal your secrets: `kubectl create secret generic jobqueue-secrets --dry-run=client -o yaml | kubeseal > sealed-secrets.yaml`
- [ ] Write `Deployment` for API: 2 replicas, image pull policy Always, resource requests (100m CPU, 128Mi RAM), limits (500m CPU, 512Mi RAM)
- [ ] Liveness probe: `GET /health` every 10s, fail after 3 attempts
- [ ] Readiness probe: `GET /ready` (checks DB + Redis connectivity) — pod won't receive traffic until this passes
- [ ] Write `Deployment` for worker: 3 replicas
- [ ] Write `Service` for API: ClusterIP, port 8080
- [ ] Write `Ingress`: host `api.yourdomain.com`, TLS with cert-manager annotation, routes to API service
- [ ] Write `HorizontalPodAutoscaler` for worker: min 2, max 10 replicas, scale on CPU > 60%
- [ ] Deploy PostgreSQL as `StatefulSet` with `PersistentVolumeClaim` (10GB)
- [ ] Deploy Redis as `StatefulSet` with `PersistentVolumeClaim` (2GB)

#### Helm chart
- [ ] Write `helm/jobqueue/Chart.yaml`
- [ ] Write `helm/jobqueue/values.yaml` — image repo/tag, replica counts, resource limits, ingress host, env vars
- [ ] Write `helm/jobqueue/values.prod.yaml` — production overrides
- [ ] Write `helm/jobqueue/values.staging.yaml` — staging overrides (lower replicas, staging domain)
- [ ] Template all manifests: `deployment.yaml`, `worker-deployment.yaml`, `service.yaml`, `ingress.yaml`, `hpa.yaml`
- [ ] `helm lint ./helm/jobqueue` — zero errors
- [ ] Deploy: `helm upgrade --install jobqueue ./helm/jobqueue -f values.prod.yaml -n prod`

#### Deploy monitoring
- [ ] Deploy Prometheus via `kube-prometheus-stack` Helm chart to `monitoring` namespace
- [ ] Write `ServiceMonitor` that scrapes your API `/metrics` endpoint
- [ ] Deploy Grafana, import your dashboard JSON
- [ ] Confirm metrics appearing in Grafana from live cluster traffic

#### Build target
- [ ] `https://api.yourdomain.com/health` returns 200 with valid TLS cert
- [ ] `kubectl get pods -n prod` — all pods Running
- [ ] Enqueue 50 jobs via the live URL — watch them process in Grafana dashboard
- [ ] `kubectl scale deployment jobqueue-worker --replicas=0 -n prod` — watch HPA bring workers back up

---

### Day 14 — Polish, Networking Deep Dive + Ship
**Goal:** Understand what you built, launch it, get first users

#### TCP/IP and networking — learn by doing
- [ ] Watch your WebSocket upgrade: `curl -v --include --no-buffer -H "Connection: Upgrade" -H "Upgrade: websocket" https://api.yourdomain.com/ws` — read the HTTP 101 response headers
- [ ] Use `tcpdump` on the VPS to watch job API requests: `tcpdump -i eth0 port 8080 -A`
- [ ] Deliberately drop Redis connection mid-request — observe TCP RST in logs, confirm your reconnect logic handles it
- [ ] Test DNS TTL: change your A record, observe how long old IP is cached
- [ ] Load test with `wrk`: `wrk -t4 -c50 -d30s https://api.yourdomain.com/health` — watch K8s HPA respond
- [ ] Inspect HTTP/2 headers on your API: `curl --http2 -I https://api.yourdomain.com/health`

#### Distributed systems — document your decisions
- [ ] Write `ARCHITECTURE.md` explaining: why `SELECT FOR UPDATE SKIP LOCKED` (optimistic concurrency, prevents double processing), why idempotency keys exist (at-least-once delivery guarantee), what happens to in-flight jobs if a worker pod crashes (they reset to pending after a timeout — explain how), why exponential backoff with jitter (thundering herd problem)
- [ ] Add a sequence diagram in the README showing job lifecycle: enqueue → claim → execute → complete/retry

#### Go SDK
- [ ] Create `sdk/` directory
- [ ] Write `client.go`: `NewClient(apiKey string) *Client`, `Enqueue(ctx, JobRequest) (*Job, error)`, `GetJob(ctx, id) (*Job, error)`
- [ ] Write `handler.go`: `RegisterHandler(jobType string, fn HandlerFunc)`, `StartWorker(ctx) error`
- [ ] Write `sdk/README.md` with quick start example
- [ ] `go mod tidy`, confirm SDK has zero external dependencies except standard library
- [ ] Tag release: `git tag v0.1.0 && git push origin v0.1.0`
- [ ] Publish: `GOPROXY=proxy.golang.org go list -m github.com/yourusername/jobqueue-sdk@v0.1.0`

#### Landing page
- [ ] Hero: "The Go-native job queue with a real dashboard. Free up to 100k jobs/month."
- [ ] Live demo playground embedded on homepage (the WebSocket dashboard from Day 8)
- [ ] One-line install: `go get github.com/yourusername/jobqueue-sdk`
- [ ] Three-line quickstart code snippet
- [ ] Comparison table vs Inngest, Trigger.dev, asynq — highlight Go-native, hosted, generous free tier
- [ ] Pricing section: Free (100k jobs/month), Pro ($19/month, 2M jobs), Team ($79/month, unlimited)

#### Analytics
- [ ] Install PostHog (free tier): `posthog.com`
- [ ] Track: `job_enqueued`, `sdk_install`, `dashboard_opened`, `demo_playground_used`
- [ ] Set up Sentry for error tracking in production API

#### Distribution
- [ ] Post in Gopher Slack `#show-and-tell`: include live URL, dashboard screenshot, one-line install
- [ ] Post in r/golang: title "I built a hosted Go-native job queue — feedback welcome"
- [ ] Post on Hacker News Show HN
- [ ] Post on X/Twitter with a Loom video showing the live dashboard under load
- [ ] DM 5 developers you know who have Go projects — ask them to try it

#### Build target
- [ ] Live URL working with TLS
- [ ] SDK published to pkg.go.dev
- [ ] First community post live
- [ ] PostHog dashboard showing first events

---

## Priority coverage summary

| # | Skill | Days covered |
|---|---|---|
| 1 | Go + PostgreSQL + REST API + architecture | 1–5 |
| 2 | TCP/IP + networking | 14 |
| 3 | Distributed systems theory | 5, 6, 14 |
| 4 | Security + auth | 7 |
| 5 | Redis + job queues + retries + idempotency | 5, 6 |
| 6 | Docker + containerization | 9 |
| 7 | WebSockets + real-time | 8 |
| 8 | Testing + code quality | 10 |
| 9 | CI/CD + GitHub Actions | 12 |
| 10 | Observability + Prometheus + Grafana | 11 |
| 11 | Kubernetes | 13 |
| 12 | Helm | 13 |
| 13–18 | gRPC, Kafka, OTel, AWS, Terraform, distributed patterns | Project 2 |
