#!/bin/bash

set -e

echo "🏗️  Setup Server CLI Build Script"
echo "═══════════════════════════════════════"

# Check if GoReleaser is available and user wants release build
if command -v goreleaser &> /dev/null && [[ "$1" == "release" ]]; then
    echo "📦 Building release with GoReleaser..."
    
    # Load environment variables from .env file if it exists
    if [[ -f .env ]]; then
        echo "🔑 Loading environment variables from .env..."
        export $(grep -v '^#' .env | xargs)
    fi
    
    # Check if GITHUB_TOKEN is set
    if [[ -z "$GITHUB_TOKEN" ]]; then
        echo "❌ GITHUB_TOKEN not found!"
        echo "💡 Add GITHUB_TOKEN=your_token to your .env file"
        echo "   Or export it: export GITHUB_TOKEN=your_token"
        exit 1
    fi
    
    if [[ "$2" == "snapshot" ]]; then
        echo "🔄 Creating snapshot release (no git tags required)..."
        goreleaser release --snapshot --clean
    else
        echo "🚀 Creating tagged release..."
        goreleaser release --clean
    fi
    echo "✅ Release build completed!"
    echo "📁 Check ./dist/ directory for all binaries"
    exit 0
fi

# Development builds (like before)
echo "🔧 Building development binaries..."
echo ""

echo "📦 Building for Linux x86_64..."
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o setup-server-cli-linux .

echo "📦 Building for macOS ARM64 (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o setup-server-cli-mac .

if [ $? -eq 0 ]; then
    echo ""
    echo "✅ Development builds completed successfully!"
    echo "📁 Available binaries:"
    echo "   • setup-server-cli-linux (Linux x86_64)"
    echo "   • setup-server-cli-mac (macOS ARM64)"
    echo ""
    echo "💡 Tips:"
    echo "   • For release builds: ./build-cli.sh release"
    echo "   • For snapshot builds: ./build-cli.sh release snapshot"
    echo "   • Install GoReleaser: go install github.com/goreleaser/goreleaser@latest"
else
    echo "❌ Build failed"
    exit 1
fi 