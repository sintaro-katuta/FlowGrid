# FlowGrid Backend - Cloudflare Workers デプロイガイド

## 概要

このガイドでは、FlowGridバックエンドをCloudflare Workers + D1 Databaseにデプロイする手順を説明します。

## 前提条件

- Cloudflareアカウント
- Wrangler CLIのインストール
- Node.js 18以上

## セットアップ手順

### 1. Wrangler CLIのインストール

```bash
npm install -g wrangler
```

### 2. Cloudflareへのログイン

```bash
wrangler login
```

### 3. D1データベースの作成

開発環境用データベースを作成:

```bash
wrangler d1 create flowgrid-db-dev
```

本番環境用データベースを作成:

```bash
wrangler d1 create flowgrid-db
```

### 4. データベースIDの設定

作成されたデータベースIDを `wrangler.toml` に設定:

```toml
# 開発環境
[[d1_databases]]
binding = "DB"
database_name = "flowgrid-db-dev"
database_id = "取得した開発環境のデータベースID"

# 本番環境
[env.production]
[[d1_databases]]
binding = "DB"
database_name = "flowgrid-db"
database_id = "取得した本番環境のデータベースID"
```

### 5. スキーマのデプロイ

SQLiteスキーマをD1データベースに適用:

```bash
wrangler d1 execute flowgrid-db-dev --file migrations/init.sqlite.sql
wrangler d1 execute flowgrid-db --file migrations/init.sqlite.sql
```

### 6. 環境変数の設定

Cloudflareダッシュボードで環境変数を設定:

**開発環境:**
- `JWT_SECRET`: 開発用JWTシークレットキー

**本番環境:**
- `JWT_SECRET`: 本番用JWTシークレットキー

### 7. アプリケーションのデプロイ

開発環境へのデプロイ:

```bash
wrangler deploy
```

本番環境へのデプロイ:

```bash
wrangler deploy --env production
```

## ビルドとテスト

### ローカル開発

```bash
# 依存関係のインストール
go mod tidy

# ローカルビルド
go build -o dist/main main_worker.go

# ローカルテスト
wrangler dev
```

### 環境変数の確認

```bash
wrangler secret list
```

## APIエンドポイント

デプロイ後、以下のエンドポイントが利用可能になります:

- `https://flowgrid-backend.your-subdomain.workers.dev/auth/*` - 認証API
- `https://flowgrid-backend.your-subdomain.workers.dev/projects/*` - プロジェクトAPI
- `https://flowgrid-backend.your-subdomain.workers.dev/tasks/*` - タスクAPI
- `https://flowgrid-backend.your-subdomain.workers.dev/health` - ヘルスチェック

## トラブルシューティング

### データベース接続エラー

```bash
# データベースの状態確認
wrangler d1 info flowgrid-db-dev

# クエリのテスト
wrangler d1 execute flowgrid-db-dev --command "SELECT name FROM sqlite_master WHERE type='table';"
```

### デプロイエラー

```bash
# ログの確認
wrangler tail

# 環境変数の再設定
wrangler secret put JWT_SECRET
```

## カスタムドメインの設定（オプション）

Cloudflareダッシュボードでカスタムドメインを設定できます。

## コスト管理

Cloudflare Workers + D1は無料枠が充実していますが、使用量に応じて課金される場合があります。定期的に使用量を確認してください。

## 参考リンク

- [Cloudflare Workers Documentation](https://developers.cloudflare.com/workers/)
- [D1 Database Documentation](https://developers.cloudflare.com/d1/)
- [Wrangler CLI Documentation](https://developers.cloudflare.com/workers/wrangler/)
