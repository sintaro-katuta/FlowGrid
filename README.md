# FlowGrid

![FlowGrid Logo](https://github.com/user-attachments/assets/7a188c2b-9d31-43da-b881-e512192f4d70)

FlowGridは、シンプルで直感的なタスク管理アプリケーションです。プロジェクトベースのタスク管理、優先度設定、進捗トラッキング機能を提供します。

## 技術スタック

### フロントエンド
![Svelte](https://img.shields.io/badge/Svelte-5.13.0-blue)
![Vite](https://img.shields.io/badge/Vite-5.13.0-blue)
![Tailwind CSS](https://img.shields.io/badge/Tailwind%20CSS-5.13.0-blue)

### バックエンド
![Go](https://img.shields.io/badge/Go-1.24-blue)
![MySQL](https://img.shields.io/badge/MySQL-8.0-blue)

### インフラストラクチャ
![Docker Compose](https://img.shields.io/badge/Docker%20Compose-1.24-blue)

## クイックスタート

### 前提条件

- Node.js 20+
- Go 1.24+
- Docker & Docker Compose
- Git
- pnpm

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
pnpm install
```

4. **Docker Composeでの起動**
```bash
# プロジェクトルートで実行
docker-compose up -d
```

5. **アプリケーションへのアクセス**
- フロントエンド: http://localhost:3000
- バックエンドAPI: http://localhost:8080

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
pnpm run dev
```
