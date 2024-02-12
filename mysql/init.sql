-- ユーザーを作成する
CREATE USER 'user'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'user'@'%';
FLUSH PRIVILEGES;

-- testデータベースを作成する
CREATE DATABASE IF NOT EXISTS test DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;;

-- testデータベースを使用する
USE test;

-- userテーブルを作成する
CREATE TABLE IF NOT EXISTS user (
    id VARCHAR(64) PRIMARY KEY,
    name VARCHAR(100),
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
