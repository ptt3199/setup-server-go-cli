#!/bin/bash

# Build script for Setup Server CLI
# This creates a pre-built binary for distribution

set -e

echo "🏗️  Building Setup Server CLI Tools..."
echo "═══════════════════════════════════════"

# Build full CLI (development + build + deploy)
echo "📦 Building full CLI (setup-server-cli)..."
go build -o setup-server-cli .

if [ $? -eq 0 ]; then
    echo "✅ Full CLI built successfully: setup-server-cli"
else
    echo "❌ Full CLI build failed"
    exit 1
fi

echo ""
echo "🎉 Build completed successfully!" 