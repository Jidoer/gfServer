DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS `Device`;
DROP TABLE IF EXISTS `PrintServer`;
DROP TABLE IF EXISTS `MyDevice`;
DROP TABLE IF EXISTS `BanedTokens`;
DROP TABLE IF EXISTS `ABTokens`;
DROP TABLE IF EXISTS `TransferRecord`;


CREATE TABLE `user` (
`id` INT PRIMARY KEY NOT NULL,
`passport` VARCHAR(45) NOT NULL UNIQUE COMMENT '账号uid',
`password` VARCHAR(255) NOT NULL COMMENT '密码',
`nickname` VARCHAR(45) NOT NULL COMMENT '用户昵称',
`role` INT NOT NULL COMMENT '角色',
`avatar` VARCHAR(255) NOT NULL COMMENT '头像url',
`email` VARCHAR(255) NOT NULL,
`phone` VARCHAR(45) NOT NULL,
`status` INT NOT NULL COMMENT '账户状态',
`create_at` DATETIME NOT NULL,
`update_at` DATETIME NOT NULL);
/**
@table: user
@columnsDescription:  id() passport(账号uid) password(密码) nickname(用户昵称) role(角色) avatar(头像url) email() phone() status(账户状态) create_at() update_at()
*/

CREATE TABLE `Device` (
`id` INT PRIMARY KEY NOT NULL,
`token` VARCHAR(255) NOT NULL);
/**
@table: Device
@columnsDescription:  id() token()
*/

CREATE TABLE `PrintServer` (
`id` INT PRIMARY KEY NOT NULL,
`passport` VARCHAR(255) NOT NULL UNIQUE COMMENT '唯一登录秘钥_创建用户后得到',
`token` VARCHAR(255) NOT NULL UNIQUE COMMENT '登录秘钥',
`name` VARCHAR(45) NOT NULL COMMENT '客户端名称',
`type` INT NOT NULL COMMENT '客户端类型',
`location_type` INT NOT NULL,
`localtion` VARCHAR(45) NOT NULL,
`usb_product` VARCHAR(255) NOT NULL COMMENT '连接的设备',
`money` INT NOT NULL COMMENT '余额',
`tansformed_monry` INT NOT NULL COMMENT '已提现的金额',
`is_online` INT NOT NULL COMMENT '占位（是否在线）',
`ban` INT NOT NULL DEFAULT false COMMENT '是否禁止登录',
`expiration_time` DATETIME NOT NULL COMMENT '到期时间')
COMMENT = '在未登录状态下，使用token连接并调用注册方法时，
自动创建表返回打印机服务器账户信息';
/**
@table: PrintServer
@description: 在未登录状态下，使用token连接并调用注册方法时，
自动创建表返回打印机服务器账户信息
*/

CREATE TABLE `MyDevice` (
`user_id` INT NOT NULL COMMENT '用户id',
`prints` INT NOT NULL COMMENT '有管理权限的打印机');

CREATE TABLE `BanedTokens` (
`id` INT PRIMARY KEY NOT NULL,
`token` VARCHAR(255) NOT NULL COMMENT '被阻止登录的token',
`baned_time` DATETIME NOT NULL,
`allowed` INT NOT NULL COMMENT '是否已允许',
`dc` INT NOT NULL COMMENT '备注')
COMMENT = '已阻止的连接';
/**
@table: BanedTokens
@description: 已阻止的连接
*/

CREATE TABLE `ABTokens` (
`id` INT PRIMARY KEY NOT NULL,
`tokens` VARCHAR(255) NOT NULL,
`allow_type` INT NOT NULL,
`dc` TEXT NOT NULL COMMENT '备注')
COMMENT = '管理token连接';
/**
@table: ABTokens
@description: 管理token连接
*/

CREATE TABLE `TransferRecord` (
`id` INT PRIMARY KEY NOT NULL,
`transfer_money` DOUBLE NOT NULL,
`status` INT(3) NOT NULL,
`operate_user` INT NOT NULL,
`PrinterId` INT NOT NULL);
/**
@table: TransferRecord
@columnsDescription:  id() transfer_money() status() operate_user() PrinterId()
*/

ALTER TABLE `MyDevice` ADD CONSTRAINT `MyDevice_user_id_user_id` FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE SET NULL ON UPDATE NO ACTION;
ALTER TABLE `MyDevice` ADD CONSTRAINT `MyDevice_prints_PrintServer_id` FOREIGN KEY (`prints`) REFERENCES `PrintServer`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE `TransferRecord` ADD CONSTRAINT `TransferRecord_operate_user_user_id` FOREIGN KEY (`operate_user`) REFERENCES `user`(`id`) ON DELETE SET NULL ON UPDATE NO ACTION;
ALTER TABLE `TransferRecord` ADD CONSTRAINT `TransferRecord_PrinterId_PrintServer_id` FOREIGN KEY (`PrinterId`) REFERENCES `PrintServer`(`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;