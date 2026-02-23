# Docker Compose Guide

## Goal
Run your Go app together with PostgreSQL and Redis in one command.

## Services you need
1. `postgres` — your database
2. `redis` — your cache
3. `app` — your Go service (built from Dockerfile)

## Key concepts

### Service names = hostnames
Inside the docker network, services talk to each other by service name.
So in your docker-specific config, use:
- `db.host: postgres` (not `localhost`)
- `redis.host: redis` (not `localhost`)

### Build your app
For your own service, point to the Dockerfile:
```yaml
app:
  build: .        # looks for Dockerfile in current directory
  ports:
    - "8080:8080" # host:container
```

### Healthchecks
Postgres and Redis take a moment to be ready. Without healthchecks your app
might start before the DB is up and crash. Add healthchecks to postgres and redis:

**Postgres:**
```yaml
healthcheck:
  test: ["CMD-SHELL", "pg_isready -U youruser"]
  interval: 5s
  timeout: 5s
  retries: 5
```

**Redis (with password):**
```yaml
healthcheck:
  test: ["CMD", "redis-cli", "-a", "yourpassword", "ping"]
  interval: 5s
  timeout: 5s
  retries: 5
```

### depends_on with condition
Make your app wait until postgres and redis are actually healthy:
```yaml
app:
  depends_on:
    postgres:
      condition: service_healthy
    redis:
      condition: service_healthy
```

### Mount your config
Instead of baking secrets into the image, mount a docker-specific config:
```yaml
app:
  volumes:
    - ./src/config/config.docker.yaml:/app/config/config.yaml
```
Create `config.docker.yaml` as a copy of `config.yaml` but with:
- `server.host: 0.0.0.0` (listen on all interfaces inside container)
- `db.host: postgres`
- `redis.host: redis`

### Persistent volumes
Don't lose your data when containers restart:
```yaml
volumes:
  postgres_data:
  redis_data:
```

Then reference them in each service:
```yaml
postgres:
  volumes:
    - postgres_data:/var/lib/postgresql/data
```

## Useful commands
```bash
docker compose up --build       # build and start everything
docker compose up --build -d    # same but in background
docker compose down             # stop and remove containers
docker compose down -v          # also remove volumes (deletes data!)
docker compose logs -f app      # follow logs for your app
docker compose ps               # check status of all services
```

## Things to remember
- Redis password must match in both the redis `command` and your config
- Postgres `POSTGRES_DB` env var must match `db.name` in your config
- The app port inside the container must match what you `EXPOSE` in the Dockerfile
