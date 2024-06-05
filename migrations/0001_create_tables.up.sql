CREATE TABLE `todos` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT comment 'ID',
    `task` VARCHAR(128) NOT NULL comment 'タスク',
    `status` VARCHAR(20) NOT NULL comment 'タスクステータス',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `recipes` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT comment 'ID',
    `name` VARCHAR(128) NOT NULL comment 'レシピ名',
    `type` VARCHAR(20) NOT NULL comment 'レシピタイプ',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP comment '作成日時',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新日時',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `users` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT comment 'ID',
    `name` VARCHAR(128) NOT NULL comment 'ユーザー名',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP comment '作成日時',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新日時',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `grades` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT comment 'ID',
    `user_id` BIGINT(20) NOT NULL comment 'ユーザーID',
    `score` INT NOT NULL comment 'グレード',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP comment '作成日時',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新日時',
    PRIMARY KEY(`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;