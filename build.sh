#!/bin/bash
set -e

echo "Building Go application..."
cd "$(dirname "$0")"
go build -tags netgo -ldflags '-s -w' -o app ./cmd/api

echo "Build complete!"
ls -lah app
