# GitHub Actions CI Guide

## Overview

GitHub Actions CI runs automated tasks (like linting, tests) on your repository events (push, PR).
Workflow files live in `.github/workflows/` and use YAML format.

---

## File Location

```
.github/
  workflows/
    lint.yml      ← your CI workflow file
```

---

## Required Fields

### `name`
The display name of the workflow shown in the GitHub Actions tab.
```yaml
name: Lint
```

---

### `on` — Trigger Events
Defines when the workflow runs.

```yaml
on:
  push:
    branches: [master]       # runs on push to master
  pull_request:
    branches: [master]       # runs on PRs targeting master
```

Other common triggers:
```yaml
on:
  schedule:
    - cron: '0 0 * * *'     # runs every day at midnight (UTC)
  workflow_dispatch:          # allows manual trigger from GitHub UI
```

---

### `jobs`
A workflow must have at least one job. Each job runs independently on a runner.

```yaml
jobs:
  job-name:                  # arbitrary name, used as identifier
    runs-on: ubuntu-latest   # the OS/runner to use
    steps:
      - ...
```

**Common `runs-on` values:**
- `ubuntu-latest`
- `windows-latest`
- `macos-latest`

---

### `steps`
A list of sequential actions inside a job. Each step either runs a shell command or uses a prebuilt action.

#### Run a shell command:
```yaml
steps:
  - name: Print message
    run: echo "Hello from CI"
```

#### Use a prebuilt action:
```yaml
steps:
  - name: Checkout code
    uses: actions/checkout@v4
```

The `uses` field references actions from GitHub Marketplace in format `owner/repo@version`.

---

## Common Actions

| Action | Purpose |
|--------|---------|
| `actions/checkout@v4` | Clones your repo into the runner |
| `actions/setup-go@v5` | Installs a specific Go version |
| `golangci/golangci-lint-action@v6` | Runs golangci-lint |

---

## Full Example — Go Linter CI

```yaml
name: Lint

on:
  push:
    branches: [master]
  pull_request:
    branches: [master]

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version-file: src/go.mod   # reads Go version from go.mod
          cache: false

      - name: Run golangci-lint
        uses: golangci/golangci-lint-action@v6
        with:
          version: v2.1.6
          working-directory: src         # path to your Go module root
```

### Key `with` parameters:
- `go-version-file` — path to `go.mod` so Go version is read automatically
- `cache: false` — disables Go module cache (golangci-lint-action manages its own)
- `version` — exact golangci-lint version to use (match your `lint.sh`)
- `working-directory` — directory where `golangci-lint run` is executed

---

## How to Activate

1. Create the file at `.github/workflows/lint.yml`
2. Commit and push to `master`
3. Go to your repo on GitHub → **Actions** tab to see the run

The workflow will now run automatically on every push to `master` and on every pull request targeting `master`.
