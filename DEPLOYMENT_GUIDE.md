# FlowGrid デプロイメントガイド

## 概要

FlowGridは以下の2つのコンポーネントで構成されています：

1. **フロントエンド**: Svelteアプリケーション（Cloudflare Pagesでホスト）
2. **バックエンド**: Go APIサーバー（Cloudflare Workersでホスト）

## デプロイ方法

### フロントエンド（Cloudflare Pages）

フロントエンドはCloudflare Pagesでホストされます：

```bash
# フロントエンドディレクトリに移動
cd front

# 依存関係のインストール
npm install

# ビルド
npm run build

# Cloudflare Pagesにデプロイ
# （Cloudflare Pagesの設定で自動デプロイが設定されている場合）
```

**Cloudflare Pages設定**:
- ビルドコマンド: `npm run build`
- ビルド出力ディレクトリ: `build`

### バックエンド（Cloudflare Workers）

バックエンドAPIはCloudflare Workersでホストされます：

```bash
# バックエンドディレクトリに移動
cd backend

# 依存関係の整理
go mod tidy

# 方法1: 自動デプロイ（推奨）
wrangler deploy

# 方法2: デプロイスクリプトを使用
./deploy.sh
```

**Cloudflare Workers設定**:
- メインファイル: `dist/worker`
- ビルドコマンド: `go build -o dist/worker worker_adapter.go`
- **重要**: デプロイコマンドは `wrangler deploy` のみ（`dist/main` を指定しない）

**デプロイエラーの解決**:
- エラー: `The expected output file at "dist/main" was not found`
- 解決策: `wrangler deploy dist/main` ではなく `wrangler deploy` を使用

## 環境変数

### バックエンド環境変数

`backend/wrangler.toml` で設定：

```toml
[vars]
JWT_SECRET = "your-jwt-secret-key-here"
```

### フロントエンド環境変数

Cloudflare Pagesのダッシュボードで設定：

- `VITE_API_BASE_URL`: バックエンドAPIのベースURL

## ビルドエラーの解決

### Cloudflare PagesでのGoビルドエラー

**問題**: Cloudflare PagesがGoコードをビルドしようとする

**解決策**: 
- Cloudflare Pagesはフロントエンドのみをビルド
- バックエンドはCloudflare Workersで別途デプロイ
- フロントエンドの `package.json` に正しいビルドコマンドを設定

### 依存関係エラー

```bash
# バックエンド
cd backend && go mod tidy

# フロントエンド  
cd front && npm install
```

## 開発環境

### ローカル開発

```bash
# バックエンド起動
cd backend && go run main.go

# フロントエンド起動（別ターミナル）
cd front && npm run dev
```

### 本番環境

1. フロントエンドをCloudflare Pagesにデプロイ
2. バックエンドをCloudflare Workersにデプロイ
3. 環境変数を適切に設定

## トラブルシューティング

### CORSエラー

フロントエンドとバックエンドのドメインが異なる場合、CORS設定を確認：

- バックエンドの `worker_adapter.go` でCORSヘッダーを設定
- フロントエンドのAPIリクエストURLを確認

### データベース接続エラー

- Cloudflare D1データベースの設定を確認
- ローカル開発時はSQLiteを使用

## 参考リンク

- [Cloudflare Pages Documentation](https://developers.cloudflare.com/pages/)
- [Cloudflare Workers Documentation](https://developers.cloudflare.com/workers/)
- [Cloudflare D1 Documentation](https://developers.cloudflare.com/d1/)
