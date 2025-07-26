CREATE TABLE `user` (
    `id`        INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    `passport`  VARCHAR(45) NOT NULL UNIQUE COMMENT 'User Passport',
    `password`  VARCHAR(255) NOT NULL COMMENT 'User Password (Hashed)',
    `nickname`  VARCHAR(45) NOT NULL COMMENT 'User Nickname',
    `avatar`    VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'User Avatar',
    `email`     VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'User Email',
    `phone`     VARCHAR(45) NOT NULL DEFAULT '' COMMENT 'User Phone',
    `status`    TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'User Status (1: Normal, 2: Blocked)',

    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_passport` (`passport`)              
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

