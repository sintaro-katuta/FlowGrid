#!/bin/bash

# FlowGrid Backend Cloudflareデプロイスクリプト

set -e  # エラーが発生したらスクリプトを終了

echo "🚀 FlowGrid Backend Cloudflareデプロイを開始します..."

# 環境変数の確認
if [ -z "$CLOUDFLARE_ACCOUNT_ID" ]; then
    echo "❌ CLOUDFLARE_ACCOUNT_IDが設定されていません"
    echo "CloudflareダッシュボードからアカウントIDを取得して設定してください"
    exit 1
fi

# 依存関係のインストール
echo "📦 Go依存関係をインストール中..."
go mod tidy

# ビルド
echo "🔨 アプリケーションをビルド中..."
go build -o dist/main main_worker.go

# 環境の確認
if [ "$1" = "production" ]; then
    ENV="production"
    DB_NAME="flowgrid-db"
    echo "🎯 本番環境にデプロイします"
else
    ENV="development"
    DB_NAME="flowgrid-db-dev"
    echo "🔧 開発環境にデプロイします"
fi

# データベースの存在確認
echo "🗄️  D1データベースを確認中..."
if ! wrangler d1 list | grep -q "$DB_NAME"; then
    echo "📊 データベース $DB_NAME を作成します"
    wrangler d1 create "$DB_NAME"
    
    # スキーマの適用
    echo "📋 データベーススキーマを適用中..."
    wrangler d1 execute "$DB_NAME" --file migrations/init.sqlite.sql
else
    echo "✅ データベース $DB_NAME は既に存在します"
fi

# 環境変数の設定確認
echo "🔐 環境変数を確認中..."
if ! wrangler secret list | grep -q "JWT_SECRET"; then
    echo "⚠️  JWT_SECRETが設定されていません"
    echo "以下のコマンドで設定してください:"
    echo "wrangler secret put JWT_SECRET"
    echo "またはCloudflareダッシュボードから設定してください"
fi

# デプロイ
echo "🚀 Cloudflare Workersにデプロイ中..."
if [ "$ENV" = "production" ]; then
    wrangler deploy --env production
else
    wrangler deploy
fi

echo "✅ デプロイが完了しました！"
echo ""
echo "🌐 APIエンドポイント:"
if [ "$ENV" = "production" ]; then
    echo "   https://flowgrid-backend.your-subdomain.workers.dev"
else
    echo "   https://flowgrid-backend.your-subdomain.workers.dev"
fi
echo ""
echo "🔍 ヘルスチェック:"
echo "   curl https://flowgrid-backend.your-subdomain.workers.dev/health"
echo ""
echo "📊 ログの確認:"
echo "   wrangler tail"
echo ""
echo "💡 トラブルシューティング:"
echo "   - 環境変数: wrangler secret list"
echo "   - データベース: wrangler d1 info $DB_NAME"
echo "   - デプロイ状態: wrangler deployments list"
