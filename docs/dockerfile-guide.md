# Dockerfile Guide

## Goal
A production-ready multi-stage Dockerfile for a Go app.

## Why multi-stage?
Two `FROM` blocks in one file:
1. **Builder stage** — has Go toolchain, compiles the binary
2. **Final stage** — tiny Alpine image, only contains the compiled binary

This keeps the final image small (no Go compiler, no source code).

## Stages

### Stage 1: Builder
- Base image: `golang:1.24-alpine`
- Set a working directory
- Copy `go.mod` and `go.sum` first, then run `go mod download`
  - This is a Docker cache trick: dependency layer only rebuilds when go.mod changes
- Copy the rest of your source code
- Build the binary:
  ```
  CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o bedrok .
  ```
  - `CGO_ENABLED=0` — static binary, no C dependencies
  - `-trimpath` — removes local file paths from the binary
  - `-ldflags="-s -w"` — strips debug info, reduces binary size

### Stage 2: Final image
- Base image: `alpine:3.21`
- Install only what you need at runtime:
  ```
  apk --no-cache add ca-certificates tzdata
  ```
  - `ca-certificates` — needed for HTTPS outbound calls
  - `tzdata` — needed if you use time zones
- Copy the binary from the builder stage:
  ```
  COPY --from=builder /build/bedrok .
  ```
- Expose your port (matches `config.yaml` server port)
- Set the entrypoint: `CMD ["./bedrok"]`

## Things to remember
- Your app reads `config/config.yaml` relative to its working directory
  - Either `COPY` the config into the image, or mount it via docker-compose volume
  - For production: mounting is better so you don't bake secrets into the image
- `WORKDIR` in the final stage should match where you copy the binary to
- Keep the final image as minimal as possible
