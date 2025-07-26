DROP TABLE IF EXISTS `User`;
DROP TABLE IF EXISTS `Device`;
DROP TABLE IF EXISTS `PrintServer`;
DROP TABLE IF EXISTS `UserDevice`;
DROP TABLE IF EXISTS `BannedTokens`;
DROP TABLE IF EXISTS `ABTokens`;
DROP TABLE IF EXISTS `TransferRecord`;

-- 用户表
CREATE TABLE `User` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `passport` VARCHAR(45) NOT NULL UNIQUE COMMENT '账号uid',
    `password` VARCHAR(255) NOT NULL COMMENT '密码',
    `nickname` VARCHAR(45) NOT NULL COMMENT '用户昵称',
    `role` TINYINT NOT NULL DEFAULT 0 COMMENT '角色(0-普通用户 10管理员)',
    `avatar` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '头像url',
    `email` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '邮箱',
    `phone` VARCHAR(45) NOT NULL DEFAULT '' COMMENT '手机号',
    `status` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '账户状态 (1: Normal, 2: Blocked)',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_passport` (`passport`)
) COMMENT '用户信息表';

-- -- 设备表
-- CREATE TABLE `Device` (
--     `id` INT PRIMARY KEY AUTO_INCREMENT,
--     `token` VARCHAR(255) NOT NULL UNIQUE
-- ) COMMENT '设备表';

-- 打印服务器表
CREATE TABLE `PrintServer` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `passport` VARCHAR(255) NOT NULL UNIQUE COMMENT '唯一登录秘钥',
    `token` VARCHAR(255) NOT NULL COMMENT '登录秘钥 可以设定为序列号 可以相同',
    `name` VARCHAR(45) NOT NULL COMMENT '客户端名称',
    `type` TINYINT NOT NULL COMMENT '客户端类型',
    `location_type` TINYINT NOT NULL COMMENT '位置类型',
    `location` VARCHAR(100) NOT NULL COMMENT '地址信息',
    `usb_product` VARCHAR(255) NOT NULL COMMENT '连接的设备',
    `balance` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT '余额',
    `withdrawn_money` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT '已提现的金额',
    `is_online` BOOLEAN NOT NULL DEFAULT FALSE COMMENT '是否在线',
    `ban` BOOLEAN NOT NULL DEFAULT FALSE COMMENT '是否禁止登录',
    `expiration_time` DATETIME NOT NULL COMMENT '到期时间',
    UNIQUE KEY `idx_passport` (`passport`)
) COMMENT '打印服务器表';

-- 用户设备关联表
-- CREATE TABLE `UserDevice` (
--     `user_id` INT NOT NULL COMMENT '用户ID',
--     `device_id` INT NOT NULL COMMENT '设备ID',
--     `bind_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '绑定时间',
--     PRIMARY KEY (`user_id`, `device_id`), --同UNIQUE(user_id, device_id)：确保同一用户只能绑定同一设备一次
--     FOREIGN KEY (`user_id`) REFERENCES `User`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
--     FOREIGN KEY (`device_id`) REFERENCES `PrintServer`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
-- ) COMMENT '用户与设备关联表';

-- 用户设备关联表
CREATE TABLE `UserDevice` (
    `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '关联表主键',
    `user_id` INT UNSIGNED NOT NULL COMMENT '用户ID',
    `device_id` INT UNSIGNED NOT NULL COMMENT '设备ID',
    `bind_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '绑定时间',
    -- `is_admin` BOOLEAN DEFAULT FALSE COMMENT '是否为设备管理员',
    FOREIGN KEY (`user_id`) REFERENCES `User`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (`device_id`) REFERENCES `PrintServer`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (`user_id`, `device_id`) -- 确保同一用户和设备的绑定关系唯一
) COMMENT '用户与设备关联表';


-- 被禁止的Token
CREATE TABLE `BannedTokens` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `token` VARCHAR(255) NOT NULL UNIQUE COMMENT '被阻止登录的token',
    `banned_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '禁用时间',
    `allowed` BOOLEAN NOT NULL DEFAULT FALSE COMMENT '是否已允许',
    `remark` TEXT COMMENT '备注'
) COMMENT '被禁止的Token';

-- 允许的Token管理
CREATE TABLE `ABTokens` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `token` VARCHAR(255) NOT NULL UNIQUE,
    `allow_type` TINYINT NOT NULL COMMENT '允许类型',
    `remark` TEXT COMMENT '备注'
) COMMENT '管理Token连接';

-- 资金转账记录
CREATE TABLE `TransferRecord` (
    `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `transfer_money` DECIMAL(10,2) NOT NULL COMMENT '转账金额',
    `status` TINYINT NOT NULL COMMENT '状态',
    `operate_user` INT UNSIGNED COMMENT '操作用户ID',
    `printer_id` INT UNSIGNED COMMENT '打印机ID',
    `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    FOREIGN KEY (`operate_user`) REFERENCES `User`(`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (`printer_id`) REFERENCES `PrintServer`(`id`) ON DELETE SET NULL ON UPDATE CASCADE
) COMMENT '资金转账记录';


-- -- 资金转账记录
-- CREATE TABLE `TransferRecord` (
--     `id` INT PRIMARY KEY AUTO_INCREMENT,
--     `transfer_money` DECIMAL(10,2) NOT NULL COMMENT '转账金额',
--     `status` TINYINT NOT NULL COMMENT '状态',
--     `operate_user` INT NOT NULL COMMENT '操作用户ID',
--     `printer_id` INT NOT NULL COMMENT '打印机ID',
--     `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--     FOREIGN KEY (`operate_user`) REFERENCES `User`(`id`) ON DELETE SET NULL ON UPDATE CASCADE,
--     FOREIGN KEY (`printer_id`) REFERENCES `PrintServer`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
-- ) COMMENT '资金转账记录';

-- 索引优化 NOT USE
-- CREATE INDEX idx_user_passport ON `User` (`passport`);
-- CREATE INDEX idx_user_phone ON `User` (`phone`);
-- CREATE INDEX idx_device_token ON `Device` (`token`);
-- CREATE INDEX idx_printserver_token ON `PrintServer` (`token`);
