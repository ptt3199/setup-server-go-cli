#!/bin/bash

set -e

echo "🏗️  Building Setup Server CLI Tools..."
echo "═══════════════════════════════════════"

echo "📦 Building full CLI for Linux x86_64 (setup-server-cli)..."
GOOS=linux GOARCH=amd64 go build -o setup-server-cli-linux .

echo "📦 Building full CLI for Mac Silicon (setup-server-cli)..."
GOOS=darwin GOARCH=arm64 go build -o setup-server-cli-mac .

if [ $? -eq 0 ]; then
    echo "✅ Full CLI built successfully: setup-server-cli"
else
    echo "❌ Full CLI build failed"
    exit 1
fi

echo ""
echo "🎉 Build completed successfully!" 