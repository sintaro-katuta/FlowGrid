# タスクの優先度と分割

## フロント

- パッケージ選定（React/Vue/Next.js など）
- アーキテクチャ設計（状態管理、Atomic Design 採用可否など）
- UI仕様定義（Figmaなど使用可）
- UI実装
  - トップページ
  - ログイン/サインアップ
  - ダッシュボード（チーム一覧、詳細、編集など）
- フロント側バリデーション実装
- 状態管理（Redux/Zustand/Recoilなど）
- API連携（フロント→バック連携）
- エラーハンドリング対応
- 単体テスト（Jest/Testing Library）
- E2Eテスト（Playwright/Cypress）←必要であれば

## バックエンド
 - 使用技術選定（Node.js/Express/FastAPIなど）
- パッケージ設計（機能ごとのモジュール分離）
- アーキテクチャ設計（MVC/Clean Architectureなど）
- DB設計
  - ER図作成
  - テーブル定義書作成
  - マイグレーションツール導入（Prisma/Migration/Fluentなど）
- API設計
  - OpenAPI (Swagger) 定義
  - 認証（JWT/OAuthなど）
  - CRUD実装（各モデルに対して）
  - バリデーション
  - エラーハンドリング
  - 単体テスト
  - 統合テスト（必要であれば）
- ロギング・モニタリング機能（必要に応じて）

## インフラ

- パッケージ選定（AWS/GCP/Azure など）
- アーキテクチャ設計（サーバレス、マイクロサービスなど）
- 構成図作成（draw.io/diagrams.netなど）
- デプロイ環境構築
  - フロント：S3 + CloudFront / Vercel / Netlify など
  - バック：EC2, ECS, Lambda, Railway, Renderなど
  - DB：RDS、Supabase、PlanetScaleなど
- CI/CD パイプライン設計・構築
  - GitHub Actions / CircleCI / GitLab CI など
- IaC（Terraform/CDKなど）
- SSL証明書・HTTPS対応
- 環境変数管理（Secrets Manager/.envなど）
- 単体テスト（インフラ用のテストスクリプト、疎通確認）
- モニタリング & アラート（CloudWatch, Sentry, Datadog など）