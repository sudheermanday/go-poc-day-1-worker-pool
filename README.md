# Go Worker Pool System

## Features
- Concurrent task processing using goroutines
- Worker pool pattern
- Retry mechanism with backoff
- Config-driven system (JSON + CLI args)
- Graceful shutdown using context

## Tech
- Go concurrency (goroutines, channels)
- Context cancellation
- JSON config

## Learnings
- Avoid goroutine leaks
- Design retry strategies
- Build scalable worker systems