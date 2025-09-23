-- テストデータ挿入スクリプト
-- ステータステーブルにデータを挿入
INSERT IGNORE INTO statuses (name) VALUES 
('todo'),
('in progress'),
('done');

-- ロールテーブルにデータを挿入
INSERT IGNORE INTO roles (name) VALUES 
('admin'),
('user'),
('manager');

-- プロジェクトテーブルにデータを挿入
INSERT IGNORE INTO projects (name) VALUES 
('FlowGrid Development'),
('Marketing Campaign'),
('Customer Support');

-- ユーザーテーブルにデータを挿入
INSERT IGNORE INTO users (name, email, role_id) VALUES 
('山田太郎', 'taro.yamada@example.com', 1),
('鈴木花子', 'hanako.suzuki@example.com', 2),
('佐藤健', 'ken.sato@example.com', 3);

-- スプリントテーブルにデータを挿入
INSERT IGNORE INTO sprint (name, project_id, start_date, end_date) VALUES 
('Sprint 1', 1, '2024-01-01', '2024-01-14'),
('Sprint 2', 1, '2024-01-15', '2024-01-28'),
('Marketing Q1', 2, '2024-01-01', '2024-03-31');

-- タスクテーブルにテストデータを挿入
INSERT IGNORE INTO tasks (title, description, user_id, project_id, sprint_id, status_id) VALUES 
('認証機能の実装', 'JWTベースの認証システムを実装する', 1, 1, 1, 1),
('データベース設計', 'タスク管理システムのデータベース設計', 2, 1, 1, 2),
('フロントエンド開発', 'Reactコンポーネントの作成', 3, 1, 1, 3),
('APIエンドポイント作成', 'RESTful APIの実装', 1, 1, 2, 1),
('テストケース作成', '単体テストと統合テストの作成', 2, 1, 2, 2),
('デプロイ設定', '本番環境へのデプロイ設定', 3, 1, 2, 3),
('広告キャンペーン計画', 'Q1のマーケティング計画立案', 1, 2, 3, 1),
('SNS投稿作成', 'ソーシャルメディア向けコンテンツ作成', 2, 2, 3, 2),
('成果レポート作成', 'キャンペーンの成果分析レポート', 3, 2, 3, 3);
