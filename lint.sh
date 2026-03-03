#!/usr/bin/env bash
set -euo pipefail

GOLANGCI_VERSION="v2.1.6"

# Check if golangci-lint is installed
if ! command -v golangci-lint &>/dev/null; then
  echo "golangci-lint not found, installing ${GOLANGCI_VERSION}..."
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
    | sh -s -- -b "$(go env GOPATH)/bin" "${GOLANGCI_VERSION}"
fi

echo "Running golangci-lint..."
cd "$(dirname "$0")/src"
golangci-lint run ./...
