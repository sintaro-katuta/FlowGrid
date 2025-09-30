#!/bin/bash

# Cloudflare Workersデプロイスクリプト
# シンプルなHello World APIをデプロイ

set -e

echo "🚀 Cloudflare Workersデプロイを開始します..."

# ビルドディレクトリの作成
mkdir -p dist

# Goのビルド（Dockerを使用）
echo "📦 Go WASM workerをビルド中..."
docker run --rm -v $(pwd):/app -w /app golang:1.24.6 go build -o dist/worker.wasm workers_main.go

# ファイルの存在確認
if [ ! -f "dist/worker.wasm" ]; then
    echo "❌ ビルドに失敗しました: dist/worker.wasm が見つかりません"
    exit 1
fi

echo "✅ ビルド完了: dist/worker.wasm"

# 環境の選択
ENV=${1:-"development"}
echo "🌍 環境: $ENV"

# デプロイ実行
echo "🚀 Cloudflare Workersにデプロイ中..."
wrangler deploy --env $ENV

echo "🎉 デプロイ完了!"
echo "📡 API URL: https://flowgrid.sintaro-katuta.workers.dev"
echo "🔍 テストコマンド: curl https://flowgrid.sintaro-katuta.workers.dev"
