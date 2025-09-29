# FlowGrid バックエンドデプロイメントガイド

## 概要

このプロジェクトは、二重管理の問題を解決し、以下の環境で動作する統一されたバックエンドを提供します：

- **通常サーバー**: ポート8080で起動する標準的なGoサーバー
- **Cloudflare Workers**: サーバーレス環境での実行

## アーキテクチャ

### 統一されたコードベース

以前は `main.go` と `main_worker.go` で二重管理されていましたが、現在は以下の構造で統一されています：

```
backend/
├── main.go                 # 通常サーバー用エントリーポイント
├── worker_adapter.go       # Cloudflare Workers用アダプター
├── models/database.go      # 環境対応データベース接続
├── db/database.go          # データベースインターフェース
└── wrangler.toml          # Cloudflare Workers設定
```

### データベース抽象化

- **Databaseインターフェース**: SQLiteとCloudflare D1の両方に対応
- **環境自動検出**: 実行環境に応じて適切なデータベース接続を選択
- **モックサポート**: テスト用のMockDatabase実装

## デプロイ方法

### 通常サーバーの起動

```bash
# ビルド
cd backend
go build -o dist/server main.go

# 実行
./dist/server
```

サーバーは `http://localhost:8080` で起動します。

### Cloudflare Workersへのデプロイ

```bash
# ビルド
cd backend
go build -o dist/worker worker_adapter.go

# デプロイ
wrangler deploy
```

## 環境変数

### 共通環境変数

- `JWT_SECRET`: JWTトークンの署名用シークレットキー
- `ENVIRONMENT`: 実行環境（`production` または `development`）

### 通常サーバー環境変数

- `DB_PATH`: SQLiteデータベースファイルのパス（デフォルト: `./flowgrid.db`）

### Cloudflare Workers環境変数

- `CF_PAGES`: Cloudflare Pages環境かどうか
- `CLOUDFLARE_WORKERS`: Cloudflare Workers環境かどうか

## ビルド設定

### 通常サーバー

```bash
go build -o dist/server main.go
```

### Cloudflare Workers

```bash
go build -o dist/worker worker_adapter.go
```

## メリット

1. **コードの統一**: 同じビジネスロジックを両環境で共有
2. **メンテナンス性向上**: 修正が1箇所で済む
3. **テスト容易性**: モックデータベースによるテストが可能
4. **環境対応**: 自動的に実行環境を検出して適切な設定を適用

## トラブルシューティング

### ビルドエラー

- **依存関係エラー**: `go mod tidy` を実行して依存関係を整理
- **インポートサイクル**: パッケージ構造を確認して循環参照を解消

### 実行時エラー

- **データベース接続エラー**: 環境変数が正しく設定されているか確認
- **CORSエラー**: フロントエンドのオリジンが許可されているか確認

## 今後の拡張

- [ ] Cloudflare D1データベースの完全対応
- [ ] 本番環境用のD1アダプター実装
- [ ] データベースマイグレーションの自動化
- [ ] モニタリングとロギングの強化
