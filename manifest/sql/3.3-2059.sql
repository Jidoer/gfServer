DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS `Device`;
DROP TABLE IF EXISTS `PrintServer`;
DROP TABLE IF EXISTS `MyDevice`;
DROP TABLE IF EXISTS `BanedTokens`;
DROP TABLE IF EXISTS `ABTokens`;
DROP TABLE IF EXISTS `TransferRecord`;

CREATE TABLE `user` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'User ID',
    `passport` VARCHAR(45) NOT NULL UNIQUE COMMENT 'User Passport',
    `password` VARCHAR(255) NOT NULL COMMENT 'User Password (Hashed)',
    `nickname` VARCHAR(45) NOT NULL COMMENT 'User Nickname',
    `avatar` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'User Avatar',
    `email` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'User Email',
    `phone` VARCHAR(45) NOT NULL DEFAULT '' COMMENT 'User Phone',
    `status` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT 'User Status (1: Normal, 2: Blocked)',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_passport` (`passport`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `Device` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'Device ID',
    `token` VARCHAR(255) NOT NULL COMMENT 'Device Token'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `PrintServer` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'Print Server ID',
    `passport` VARCHAR(255) NOT NULL UNIQUE COMMENT 'User Passport',
    `token` VARCHAR(255) NOT NULL UNIQUE COMMENT 'Login Token',
    `name` VARCHAR(45) NOT NULL COMMENT 'Client Name',
    `type` INT NOT NULL COMMENT 'Client Type',
    `location_type` INT NOT NULL COMMENT 'Location Type',
    `location` VARCHAR(45) NOT NULL COMMENT 'Location',
    `usb_product` VARCHAR(255) NOT NULL COMMENT 'Connected Device',
    `money` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT 'Balance',
    `transformed_money` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT 'Withdrawn Money',
    `is_online` TINYINT NOT NULL DEFAULT 0 COMMENT 'Online Status',
    `ban` TINYINT NOT NULL DEFAULT 0 COMMENT 'Banned Status',
    `expiration_time` DATETIME NOT NULL COMMENT 'Expiration Time'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `MyDevice` (
    `user_id` INT UNSIGNED NOT NULL COMMENT 'User ID',
    `prints` INT UNSIGNED NOT NULL COMMENT 'Managed Printer ID',
    PRIMARY KEY (`user_id`, `prints`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`prints`) REFERENCES `PrintServer`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `BanedTokens` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'Ban ID',
    `token` VARCHAR(255) NOT NULL COMMENT 'Blocked Token',
    `baned_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Banned Time',
    `allowed` TINYINT NOT NULL DEFAULT 0 COMMENT 'Allowed Status',
    `dc` TEXT NOT NULL COMMENT 'Remarks'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Blocked Connections';

CREATE TABLE `ABTokens` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'AB Token ID',
    `tokens` VARCHAR(255) NOT NULL COMMENT 'Token Value',
    `allow_type` INT NOT NULL COMMENT 'Allowed Type',
    `dc` TEXT NOT NULL COMMENT 'Remarks'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Manage Token Connections';

CREATE TABLE `TransferRecord` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'Transfer ID',
    `OptionSuperId` INT UNSIGNED NOT NULL COMMENT 'Operator ID',
    `TransferMoney` DECIMAL(10,2) NOT NULL COMMENT 'Transfer Amount',
    `status` TINYINT(3) NOT NULL COMMENT 'Transfer Status',
    `PrinterId` INT UNSIGNED NOT NULL COMMENT 'Target Printer ID',
    FOREIGN KEY (`OptionSuperId`) REFERENCES `user`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`PrinterId`) REFERENCES `PrintServer`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
