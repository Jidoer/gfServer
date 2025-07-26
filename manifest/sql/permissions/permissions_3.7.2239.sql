DROP TABLE IF EXISTS `role_permissions`;
DROP TABLE IF EXISTS `roles`;
DROP TABLE IF EXISTS `permissions`;

--角色 common admin ... 后期加入createby等等
CREATE TABLE `roles` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `rolesID` INT NOT NULL UNIQUE, ---> 1:common 2:admin 此id用于前端判断
    `name` VARCHAR(50) NOT NULL UNIQUE,
    `description` VARCHAR(255),
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

--权限 btn.link btn.edit ...
CREATE TABLE `permissions` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(50) NOT NULL UNIQUE,
    `description` VARCHAR(255)
);

--角色-权限关系表 admin ----> [btn.edit, btn.link]
CREATE TABLE `role_permissions` (
    `role_id` INT NOT NULL,
    `permission_id` INT NOT NULL,
    PRIMARY KEY (`role_id`, `permission_id`),
    CONSTRAINT `role_permissions_role_id_roles_id` FOREIGN KEY (`role_id`) 
    REFERENCES `roles`(`id`) ON DELETE CASCADE,
    CONSTRAINT `role_permissions_permission_id_permissions_id` FOREIGN KEY (`permission_id`) 
    REFERENCES `permissions`(`id`) ON DELETE CASCADE
);
