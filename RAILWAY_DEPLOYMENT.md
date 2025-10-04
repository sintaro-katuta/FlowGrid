# Railway デプロイガイド

このガイドでは、FlowGridアプリケーションをRailwayでデプロイする手順を説明します。

## プロジェクト構成

- **バックエンド**: `backend/` ディレクトリ (Go + Gin)
- **フロントエンド**: `front/` ディレクトリ (Svelte + Vite)

## Railwayでのデプロイ手順

### 1. Railwayアカウントの準備

1. [Railway](https://railway.app/) にアクセスしてアカウントを作成またはログイン
2. GitHubアカウントと連携

### 2. バックエンドのデプロイ

1. Railwayダッシュボードで「New Project」をクリック
2. 「Deploy from GitHub repo」を選択
3. FlowGridリポジトリを選択
4. 以下の設定を選択：
   - **Root Directory**: `backend`
   - **Branch**: `main` (またはデプロイしたいブランチ)

5. 環境変数の設定：
   - `DATABASE_URL`: データベース接続URL (RailwayのPostgreSQLアドオンを使用する場合、自動的に設定されます)
   - その他必要な環境変数

### 3. フロントエンドのデプロイ

1. 別のプロジェクトとして「New Project」をクリック
2. 「Deploy from GitHub repo」を選択
3. FlowGridリポジトリを選択
4. 以下の設定を選択：
   - **Root Directory**: `front`
   - **Branch**: `main` (またはデプロイしたいブランチ)

5. 環境変数の設定：
   - `PUBLIC_API_URL`: バックエンドのURL (例: `https://your-backend-service.up.railway.app`)

### 4. ドメイン設定

各サービスにカスタムドメインを設定できます：

1. プロジェクト設定で「Domains」を選択
2. カスタムドメインを追加

## 環境変数

### バックエンド環境変数

```env
DATABASE_URL=postgresql://username:password@host:port/database
PORT=8080
# その他必要な環境変数
```

### フロントエンド環境変数

```env
PUBLIC_API_URL=https://your-backend-service.up.railway.app
```

## ビルド設定

### バックエンド
- **Builder**: NIXPACKS
- **Start Command**: `./main`

### フロントエンド
- **Builder**: NIXPACKS
- **Build Command**: `npm run build`
- **Start Command**: `npm start`

## 注意事項

1. **データベース**: RailwayのPostgreSQLアドオンを使用することを推奨します
2. **CORS設定**: バックエンドでフロントエンドのドメインを許可するようにCORS設定を確認してください
3. **環境変数**: 本番環境用の環境変数を適切に設定してください
4. **ビルドキャッシュ**: Railwayは自動的にビルドキャッシュを管理します

## トラブルシューティング

### ビルド失敗時
- ログを確認して依存関係の問題を特定
- Node.js/Pnpmのバージョン互換性を確認

### 起動失敗時
- ポート設定が正しいか確認
- 環境変数が正しく設定されているか確認

## 参考リンク

- [Railway Documentation](https://docs.railway.app/)
- [Svelte Adapter Node](https://kit.svelte.dev/docs/adapter-node)
- [Go on Railway](https://docs.railway.app/languages/go)
