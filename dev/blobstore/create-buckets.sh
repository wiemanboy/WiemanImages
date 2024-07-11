#!/bin/sh

isAlive() {
  curl -sf http://127.0.0.1:9000/minio/health/live
}

# Start Minio in the background
minio "$@" --quiet & echo $! > /tmp/minio.pid

# Wait for Minio to be ready
while ! isAlive; do
  sleep 0.01
done

# Authenticate
mc alias set minio http://127.0.0.1:9000 "$MINIO_ROOT_USER" "$MINIO_ROOT_PASSWORD"

# Create buckets
mc mb minio/test-bucket || true
mc anonymous set download minio/test-bucket

# Stop Minio
kill -s INT $(cat /tmp/minio.pid) && rm /tmp/minio.pid

# Wait for Minio to be stopped
while isAlive; do
  sleep 0.01
done

# Start Minio in the foreground
exec minio "$@"