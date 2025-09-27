# FlowGrid

![FlowGrid Logo](https://github.com/user-attachments/assets/7a188c2b-9d31-43da-b881-e512192f4d70)

FlowGridは、シンプルで直感的なタスク管理アプリケーションです。プロジェクトベースのタスク管理、優先度設定、進捗トラッキング機能を提供します。

## 特徴

- 📋 **プロジェクトベースのタスク管理**
- 🎯 **優先度設定（高/中/低）**
- 📊 **進捗トラッキング**
- 👥 **ユーザー認証システム**
- 🎨 **カラフルなプロジェクトカテゴリ**
- 📱 **レスポンシブデザイン**

## 技術スタック

### フロントエンド
- **Svelte** - モダンなフロントエンドフレームワーク
- **Vite** - 高速なビルドツール
- **Tailwind CSS** - ユーティリティファーストCSS

### バックエンド
- **Go** - 高速で信頼性の高いバックエンド
- **Gin** - Webフレームワーク
- **MySQL** - データベース

### インフラストラクチャ
- **Docker Compose** - ローカル開発環境
- **Cloudflare Pages** - フロントエンドホスティング（推奨）
- **Cloudflare Workers** - サーバーレスAPI（推奨）
- **Cloudflare D1** - エッジデータベース（推奨）

## クイックスタート

### 前提条件

- Node.js 18+
- Go 1.19+
- Docker & Docker Compose
- Git

### ローカル開発環境のセットアップ

1. **リポジトリのクローン**
```bash
git clone https://github.com/sintaro-katuta/FlowGrid.git
cd FlowGrid
```

2. **バックエンドのセットアップ**
```bash
cd backend
go mod download
```

3. **フロントエンドのセットアップ**
```bash
cd ../front
npm install
```

4. **Docker Composeでの起動**
```bash
# プロジェクトルートで実行
docker-compose up -d
```

5. **アプリケーションへのアクセス**
- フロントエンド: http://localhost:3000
- バックエンドAPI: http://localhost:8080

## デプロイオプション

### 推奨: Cloudflare（月額0円〜）

**メリット**:
- 無料枠が充実（実質0円で運用可能）
- グローバルCDNで高速
- 設定がシンプル

**デプロイ手順**:

1. **Cloudflareアカウントの作成**
   - [Cloudflare](https://dash.cloudflare.com/sign-up)で無料アカウント作成

2. **フロントエンドのデプロイ（Cloudflare Pages）**
   - Cloudflareダッシュボードで「Pages」を選択
   - GitHubリポジトリを接続
   - ビルド設定:
     - ビルドコマンド: `cd front && npm run build`
     - ビルド出力ディレクトリ: `front/dist`
   - 環境変数設定（必要に応じて）

3. **バックエンドのデプロイ（Cloudflare Workers + D1）**
   - WorkersでサーバーレスAPIをデプロイ
   - D1データベースを作成
   - 環境変数を設定

### 代替オプション: 従来のホスティング

**Vercel + PlanetScale**:
- フロントエンド: Vercel（無料枠あり）
- バックエンド: Vercel Serverless Functions
- データベース: PlanetScale（MySQL互換、無料枠あり）

**Netlify + Supabase**:
- フロントエンド: Netlify（無料枠あり）
- バックエンド: Netlify Functions
- データベース: Supabase（PostgreSQL、無料枠あり）

## プロジェクト構造

```
FlowGrid/
├── backend/          # Goバックエンド
│   ├── api/          # APIハンドラー
│   ├── models/       # データモデル
│   └── migrations/   # データベースマイグレーション
├── front/            # Svelteフロントエンド
│   ├── src/
│   │   ├── routes/   # ページルート
│   │   ├── lib/      # ユーティリティ
│   │   └── components/# UIコンポーネント
└── docs/             # ドキュメント
```

## 開発ガイド

### バックエンド開発
```bash
cd backend
go run main.go
```

### フロントエンド開発
```bash
cd front
npm run dev
```

### データベースマイグレーション
```bash
# ローカル環境
docker-compose exec mysql mysql -u user -p flowgrid < backend/migrations/init.sql
```

## コスト見積もり

### Cloudflare構成（推奨）
- **月額コスト**: 0円〜数百円
- 無料枠内で十分な機能を提供

### 従来のAWS構成
- **月額コスト**: 3,700円〜7,300円
- RDS、Lambda、S3、CloudFrontを使用

## ライセンス

MIT License

## 貢献

バグ報告や機能リクエストはGitHub Issuesまでお願いします。

## サポート

問題が発生した場合は、以下の順序で調査してください：
1. Docker Composeのログを確認
2. ブラウザの開発者ツールでエラーを確認
3. GitHub Issuesで問題を報告
