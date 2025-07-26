DROP TABLE IF EXISTS `User`;
DROP TABLE IF EXISTS `Device`;
-- DROP TABLE IF EXISTS `PrintServer`;
-- DROP TABLE IF EXISTS `UserDevice`;
DROP TABLE IF EXISTS `BannedTokens`;
DROP TABLE IF EXISTS `ABTokens`;
-- DROP TABLE IF EXISTS `TransferRecord`;

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

--!
-- 打印服务器表
-- CREATE TABLE `PrintServer` (
--     `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
--     `passport` VARCHAR(255) NOT NULL UNIQUE COMMENT '唯一登录秘钥',
--     `token` VARCHAR(255) NOT NULL COMMENT '登录秘钥 可以设定为序列号 可以相同',
--     `name` VARCHAR(45) NOT NULL COMMENT '客户端名称',
--     `type` TINYINT NOT NULL COMMENT '客户端类型',
--     `location_type` TINYINT NOT NULL COMMENT '位置类型',
--     `location` VARCHAR(100) NOT NULL COMMENT '地址信息',
--     `usb_product` VARCHAR(255) NOT NULL COMMENT '连接的设备',
--     `balance` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT '余额',
--     `withdrawn_money` DECIMAL(10,2) NOT NULL DEFAULT 0 COMMENT '已提现的金额',
--     `is_online` BOOLEAN NOT NULL DEFAULT FALSE COMMENT '是否在线',
--     `ban` BOOLEAN NOT NULL DEFAULT FALSE COMMENT '是否禁止登录',
--     `expiration_time` DATETIME NOT NULL COMMENT '到期时间',
--     UNIQUE KEY `idx_passport` (`passport`)
-- ) COMMENT '打印服务器表';

-- 用户设备关联表
-- CREATE TABLE `UserDevice` (
--     `user_id` INT NOT NULL COMMENT '用户ID',
--     `device_id` INT NOT NULL COMMENT '设备ID',
--     `bind_time` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '绑定时间',
--     PRIMARY KEY (`user_id`, `device_id`), --同UNIQUE(user_id, device_id)：确保同一用户只能绑定同一设备一次
--     FOREIGN KEY (`user_id`) REFERENCES `User`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
--     FOREIGN KEY (`device_id`) REFERENCES `PrintServer`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
-- ) COMMENT '用户与设备关联表';

--!
-- 用户设备关联表
-- CREATE TABLE `UserDevice` (
--     `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '关联表主键',
--     `user_id` INT UNSIGNED NOT NULL COMMENT '用户ID',
--     `device_id` INT UNSIGNED NOT NULL COMMENT '设备ID',
--     `bind_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '绑定时间',
--     -- `is_admin` BOOLEAN DEFAULT FALSE COMMENT '是否为设备管理员',
--     FOREIGN KEY (`user_id`) REFERENCES `User`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
--     FOREIGN KEY (`device_id`) REFERENCES `PrintServer`(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
--     UNIQUE (`user_id`, `device_id`) -- 确保同一用户和设备的绑定关系唯一
-- ) COMMENT '用户与设备关联表';


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

--!
-- 资金转账记录
-- CREATE TABLE `TransferRecord` (
--     `id` INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
--     `transfer_money` DECIMAL(10,2) NOT NULL COMMENT '转账金额',
--     `status` TINYINT NOT NULL COMMENT '状态',
--     `operate_user` INT UNSIGNED COMMENT '操作用户ID',
--     `printer_id` INT UNSIGNED COMMENT '打印机ID',
--     `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--     FOREIGN KEY (`operate_user`) REFERENCES `User`(`id`) ON DELETE SET NULL ON UPDATE CASCADE,
--     FOREIGN KEY (`printer_id`) REFERENCES `PrintServer`(`id`) ON DELETE SET NULL ON UPDATE CASCADE
-- ) COMMENT '资金转账记录';

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
/**
 * 建议：路由 path 路径与文件夹名称相同，找文件可浏览器地址找，方便定位文件位置
 *
 * 路由meta对象参数说明
 * meta: {
 *      title:          菜单栏及 tagsView 栏、菜单搜索名称（国际化）
 *      isLink：        是否超链接菜单，开启外链条件，`1、isLink: 链接地址不为空 2、isIframe:false`
 *      isHide：        是否隐藏此路由
 *      isKeepAlive：   是否缓存组件状态
 *      isAffix：       是否固定在 tagsView 栏上
 *      isIframe：      是否内嵌窗口，开启条件，`1、isIframe:true 2、isLink：链接地址不为空`
 *      roles：         当前路由权限标识，取角色管理。控制路由显示、隐藏。超级管理员：admin 普通角色：common
 *      icon：          菜单、tagsView 图标，阿里：加 `iconfont xxx`，fontawesome：加 `fa xxx`
 * }
 */
CREATE TABLE `SystemMenus` (
    `id`            INT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '主键 ID',
    `parent_id`     INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级路由 ID，0 表示顶级路由',
    `path`          VARCHAR(255) NOT NULL COMMENT '路由路径',
    `name`          VARCHAR(255) NOT NULL COMMENT '路由名称',
    `component`     VARCHAR(255) NOT NULL COMMENT '组件路径',
    `redirect`      VARCHAR(255) DEFAULT NULL COMMENT '重定向路径', --可无

    --meta
    `title`         VARCHAR(255) NOT NULL COMMENT '标题',
    `is_link`       VARCHAR(255) DEFAULT NULL COMMENT '外链地址',
    `is_hide`       BOOLEAN DEFAULT FALSE COMMENT '是否隐藏',
    `is_keep_alive` BOOLEAN DEFAULT FALSE COMMENT '是否缓存',
    `is_affix`      BOOLEAN DEFAULT FALSE COMMENT '是否固定标签',
    `is_iframe`     BOOLEAN DEFAULT FALSE COMMENT '是否嵌套 iframe',
    `roles`         TEXT COMMENT '允许访问的角色，JSON 存储',
    `icon`          VARCHAR(255) NOT NULL DEFAULT 'https://....' COMMENT '图标', ---need-fix

    `created_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

    FOREIGN KEY (`parent_id`) REFERENCES `SystemMenus`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='路由表';


INSERT INTO `SystemMenus` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `is_link`, `is_hide`, `is_keep_alive`, `is_affix`, `is_iframe`, `roles`, `icon`, `created_at`, `updated_at`) VALUES
(1, 0, '/dashboard', 'Dashboard', 'dashboard/index', NULL, '仪表盘', NULL, FALSE, TRUE, TRUE, FALSE, '["admin", "user"]', 'https://example.com/icon1.png', NOW(), NOW()),

(2, 0, '/system', 'System', 'system/index', NULL, '系统管理', NULL, FALSE, TRUE, FALSE, FALSE, '["admin"]', 'https://example.com/icon2.png', NOW(), NOW()),

(3, 2, '/system/user', 'User', 'system/user', NULL, '用户管理', NULL, FALSE, TRUE, FALSE, FALSE, '["admin"]', 'https://example.com/icon3.png', NOW(), NOW()),

(4, 2, '/system/role', 'Role', 'system/role', NULL, '角色管理', NULL, FALSE, TRUE, FALSE, FALSE, '["admin"]', 'https://example.com/icon4.png', NOW(), NOW()),

(5, 0, '/reports', 'Reports', 'reports/index', NULL, '报表管理', NULL, FALSE, TRUE, FALSE, FALSE, '["admin", "user"]', 'https://example.com/icon5.png', NOW(), NOW());




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
