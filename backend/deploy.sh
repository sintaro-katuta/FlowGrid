#!/bin/bash

# FlowGrid Backend Deployment Script
# This script ensures correct deployment to Cloudflare Workers

echo "🚀 Starting FlowGrid Backend Deployment..."

# Check if we're in the backend directory
if [ ! -f "wrangler.toml" ]; then
    echo "❌ Error: Please run this script from the backend directory"
    echo "Usage: cd backend && ./deploy.sh"
    exit 1
fi

# Clean up any previous builds
echo "🧹 Cleaning previous builds..."
rm -rf dist/
mkdir -p dist/

# Build the worker using Docker
echo "🔨 Building worker using Docker..."
docker run --rm -v $(pwd):/app -w /app golang:1.24.6 go build -o dist/worker .

# Verify the build
if [ ! -f "dist/worker" ]; then
    echo "❌ Build failed: dist/worker not found"
    exit 1
fi

echo "✅ Build successful: dist/worker created"

# Deploy to Cloudflare Workers
echo "☁️  Deploying to Cloudflare Workers..."

# Check if wrangler is available, if not try to install it
if ! command -v wrangler &> /dev/null; then
    echo "⚠️  Wrangler not found, trying to install..."
    if command -v npm &> /dev/null; then
        npm install -g wrangler
    elif command -v yarn &> /dev/null; then
        yarn global add wrangler
    else
        echo "❌ Cannot install wrangler: npm or yarn not available"
        echo "Please install wrangler manually: npm install -g wrangler"
        exit 1
    fi
fi

# Deploy using wrangler with environment specification
echo "🚀 Deploying to production environment..."
wrangler deploy --env production

echo "🎉 Deployment completed!"
