-- phpMyAdmin SQL Dump
-- version 5.2.2
-- https://www.phpmyadmin.net/
--
-- 主机： 47.93.100.223:3306
-- 生成日期： 2025-07-26 04:20:09
-- 服务器版本： 10.11.10-MariaDB-log
-- PHP 版本： 8.2.29

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `hellos`
--

-- --------------------------------------------------------

--
-- 表的结构 `ABTokens`
--

CREATE TABLE `ABTokens` (
  `id` int(10) UNSIGNED NOT NULL,
  `token` varchar(255) NOT NULL,
  `allow_type` tinyint(4) NOT NULL COMMENT '允许类型',
  `remark` text DEFAULT NULL COMMENT '备注'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci COMMENT='管理Token连接';

-- --------------------------------------------------------

--
-- 表的结构 `BannedTokens`
--

CREATE TABLE `BannedTokens` (
  `id` int(10) UNSIGNED NOT NULL,
  `token` varchar(255) NOT NULL COMMENT '被阻止登录的token',
  `banned_time` datetime NOT NULL DEFAULT current_timestamp() COMMENT '禁用时间',
  `allowed` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否已允许',
  `remark` text DEFAULT NULL COMMENT '备注'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci COMMENT='被禁止的Token';

-- --------------------------------------------------------

--
-- 表的结构 `permissions`
--

CREATE TABLE `permissions` (
  `id` int(11) NOT NULL,
  `name` varchar(50) NOT NULL,
  `description` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci;

--
-- 转存表中的数据 `permissions`
--

INSERT INTO `permissions` (`id`, `name`, `description`) VALUES
(32, 'btn.add', '基本权限'),
(33, 'btn.del', '基本权限'),
(34, 'btn.edit', '基本权限'),
(35, 'btn.link', '基本权限'),
(36, 'PARTFORM_ADMIN', '平台管理员');

-- --------------------------------------------------------

--
-- 表的结构 `PrintServer`
--

CREATE TABLE `PrintServer` (
  `id` int(10) UNSIGNED NOT NULL,
  `passport` varchar(255) NOT NULL COMMENT '唯一登录秘钥',
  `token` varchar(255) NOT NULL COMMENT '登录秘钥 可以设定为序列号 可以相同',
  `name` varchar(45) NOT NULL COMMENT '客户端名称',
  `type` tinyint(4) NOT NULL COMMENT '客户端类型',
  `location_type` tinyint(4) NOT NULL COMMENT '位置类型',
  `location` varchar(100) NOT NULL COMMENT '地址信息',
  `usb_product` varchar(255) NOT NULL COMMENT '连接的设备',
  `balance` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '余额',
  `withdrawn_money` decimal(10,2) NOT NULL DEFAULT 0.00 COMMENT '已提现的金额',
  `is_online` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否在线',
  `ban` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否禁止登录',
  `expiration_time` datetime NOT NULL COMMENT '到期时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci COMMENT='打印服务器表';

-- --------------------------------------------------------

--
-- 表的结构 `roles`
--

CREATE TABLE `roles` (
  `id` int(11) NOT NULL,
  `rolesID` int(11) NOT NULL COMMENT '给用户使用唯一ID',
  `name` varchar(50) NOT NULL COMMENT '名称',
  `role_sign` varchar(50) NOT NULL COMMENT '标识',
  `rank` int(11) DEFAULT 50 COMMENT '权重',
  `status` tinyint(4) DEFAULT NULL COMMENT '状态码 0正常 1禁用',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci;

--
-- 转存表中的数据 `roles`
--

INSERT INTO `roles` (`id`, `rolesID`, `name`, `role_sign`, `rank`, `status`, `description`, `created_at`, `updated_at`) VALUES
(4, 10, '超级管理员', 'admin', 50, 1, '平台管理员', '2025-03-08 07:00:14', '2025-03-08 07:07:53'),
(5, 0, '普通用户', 'common', 50, 1, '测试', '2025-03-08 14:11:56', '2025-03-16 16:36:00');

-- --------------------------------------------------------

--
-- 表的结构 `role_permissions`
--

CREATE TABLE `role_permissions` (
  `role_id` int(11) NOT NULL,
  `permission_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci;

--
-- 转存表中的数据 `role_permissions`
--

INSERT INTO `role_permissions` (`role_id`, `permission_id`) VALUES
(4, 32),
(4, 33),
(4, 34),
(4, 35),
(4, 36),
(5, 32),
(5, 35);

-- --------------------------------------------------------

--
-- 表的结构 `SystemMenus`
--

CREATE TABLE `SystemMenus` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '主键 ID',
  `parent_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级路由 ID，0 表示顶级路由',
  `path` varchar(255) NOT NULL COMMENT '路由路径',
  `name` varchar(255) NOT NULL COMMENT '路由名称',
  `component` varchar(255) NOT NULL COMMENT '组件路径',
  `redirect` varchar(255) DEFAULT NULL COMMENT '重定向路径',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `is_link` varchar(255) DEFAULT NULL COMMENT '外链地址',
  `is_hide` tinyint(1) DEFAULT 0 COMMENT '是否隐藏',
  `is_keep_alive` tinyint(1) DEFAULT 0 COMMENT '是否缓存',
  `is_affix` tinyint(1) DEFAULT 0 COMMENT '是否固定标签',
  `is_iframe` tinyint(1) DEFAULT 0 COMMENT '是否嵌套 iframe',
  `roles` text DEFAULT NULL COMMENT '允许访问的角色，JSON 存储',
  `icon` varchar(255) NOT NULL DEFAULT 'https://....' COMMENT '图标',
  `created_at` timestamp NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='路由表';

--
-- 转存表中的数据 `SystemMenus`
--

INSERT INTO `SystemMenus` (`id`, `parent_id`, `path`, `name`, `component`, `redirect`, `title`, `is_link`, `is_hide`, `is_keep_alive`, `is_affix`, `is_iframe`, `roles`, `icon`, `created_at`, `updated_at`) VALUES
(0, 0, '/home', 'message.router.home', '/home/index', NULL, 'message.router.home', NULL, 0, 0, 0, 0, NULL, 'https://....', '2025-03-16 09:27:04', '2025-03-17 00:20:04'),
(362, 0, '/system', 'system', '/layout/routerView/parent', '/system/menu', 'message.router.system', '', 0, 1, 0, 0, '[\"admin\"]', 'iconfont icon-xitongshezhi', '2025-03-16 01:27:30', '2025-03-16 01:27:30'),
(363, 0, '/limits', 'limits', '/layout/routerView/parent', '/limits/frontEnd', 'message.router.limits', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-quanxian', '2025-03-16 01:27:30', '2025-03-16 01:27:30'),
(365, 0, '/fun', 'funIndex', '/layout/routerView/parent', '/fun/tagsView', 'message.router.funIndex', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-crew_feature', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(366, 0, '/pages', 'pagesIndex', '/layout/routerView/parent', '/pages/filtering', 'message.router.pagesIndex', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-fuzhiyemian', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(367, 0, '/make', 'makeIndex', '/layout/routerView/parent', '/make/selector', 'message.router.makeIndex', '', 0, 1, 0, 0, '[\"admin\"]', 'iconfont icon-siweidaotu', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(368, 0, '/params', 'paramsIndex', '/layout/routerView/parent', '/params/common', 'message.router.paramsIndex', '', 0, 1, 0, 0, '[\"admin\"]', 'iconfont icon-zhongduancanshu', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(369, 0, '/visualizing', 'visualizingIndex', '/layout/routerView/parent', '/visualizing/visualizingLinkDemo1', 'message.router.visualizingIndex', '', 0, 1, 0, 0, '[\"admin\"]', 'ele-ChatLineRound', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(370, 0, '/chart', 'chartIndex', '/chart/index', '', 'message.router.chartIndex', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-ico_shuju', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(371, 0, '/personal', 'personal', '/personal/index', '', 'message.router.personal', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-gerenzhongxin', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(372, 0, '/tools', 'tools', '/tools/index', '', 'message.router.tools', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-gongju', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(373, 0, '/link', 'layoutLinkView', '/layout/routerView/link', '', 'message.router.layoutLinkView', 'https://element-plus.gitee.io/#/zh-CN/component/installation', 0, 0, 0, 0, '[\"admin\"]', 'iconfont icon-caozuo-wailian', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(376, 362, '/system/menu', 'systemMenu', '/system/menu/index', '', 'message.router.systemMenu', '', 0, 1, 0, 0, '[\"admin\"]', 'iconfont icon-caidan', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(377, 362, '/system/role', 'systemRole', '/system/role/index', '', 'message.router.systemRole', '', 0, 1, 0, 0, '[\"admin\"]', 'ele-ColdDrink', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(378, 362, '/system/user', 'systemUser', '/system/user/index', '', 'message.router.systemUser', '', 0, 1, 0, 0, '[\"admin\"]', 'iconfont icon-icon-', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(379, 362, '/system/auth', 'systemDept', '/system/auth/index', '', 'message.router.limits', '', 0, 1, 0, 0, '[\"admin\"]', 'iconfont icon-quanxian', '2025-03-16 01:27:31', '2025-03-17 00:34:58'),
(380, 362, '/system/dic', 'systemDic', '/system/dic/index', '', 'message.router.systemDic', '', 0, 1, 0, 0, '[\"admin\"]', 'ele-SetUp', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(381, 363, '/limits/frontEnd', 'limitsFrontEnd', '/layout/routerView/parent', '/limits/frontEnd/page', 'message.router.limitsFrontEnd', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', '', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(382, 363, '/limits/backEnd', 'limitsBackEnd', '/layout/routerView/parent', '', 'message.router.limitsBackEnd', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', '', '2025-03-16 01:27:31', '2025-03-16 01:27:31'),
(385, 365, '/fun/tagsView', 'funTagsView', '/fun/tagsView/index', '', 'message.router.funTagsView', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'ele-Pointer', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(386, 365, '/fun/countup', 'funCountup', '/fun/countup/index', '', 'message.router.funCountup', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'ele-Odometer', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(387, 365, '/fun/wangEditor', 'funWangEditor', '/fun/wangEditor/index', '', 'message.router.funWangEditor', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-fuwenbenkuang', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(388, 365, '/fun/cropper', 'funCropper', '/fun/cropper/index', '', 'message.router.funCropper', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-caijian', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(389, 365, '/fun/qrcode', 'funQrcode', '/fun/qrcode/index', '', 'message.router.funQrcode', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-ico', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(390, 365, '/fun/echartsMap', 'funEchartsMap', '/fun/echartsMap/index', '', 'message.router.funEchartsMap', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-ditu', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(391, 365, '/fun/printJs', 'funPrintJs', '/fun/printJs/index', '', 'message.router.funPrintJs', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'ele-Printer', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(392, 365, '/fun/clipboard', 'funClipboard', '/fun/clipboard/index', '', 'message.router.funClipboard', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'ele-DocumentCopy', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(393, 365, '/fun/gridLayout', 'funGridLayout', '/fun/gridLayout/index', '', 'message.router.funGridLayout', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-tuodong', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(394, 365, '/fun/splitpanes', 'funSplitpanes', '/fun/splitpanes/index', '', 'message.router.funSplitpanes', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon--chaifenlie', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(395, 366, '/pages/filtering', 'pagesFiltering', '/pages/filtering/index', '', 'message.router.pagesFiltering', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'ele-Sell', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(396, 366, '/pages/filtering/details1', 'pagesFilteringDetails1', '/pages/filtering/details1', '', 'message.router.pagesFilteringDetails1', '', 1, 0, 0, 0, '[\"admin\",\"common\"]', 'ele-Sunny', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(397, 366, '/pages/iocnfont', 'pagesIocnfont', '/pages/iocnfont/index', '', 'message.router.pagesIocnfont', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'ele-Present', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(398, 366, '/pages/element', 'pagesElement', '/pages/element/index', '', 'message.router.pagesElement', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'ele-Eleme', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(399, 366, '/pages/awesome', 'pagesAwesome', '/pages/awesome/index', '', 'message.router.pagesAwesome', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'ele-SetUp', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(400, 366, '/pages/formAdapt', 'pagesFormAdapt', '/pages/formAdapt/index', '', 'message.router.pagesFormAdapt', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-biaodan', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(401, 366, '/pages/tableRules', 'pagesTableRules', '/pages/tableRules/index', '', 'message.router.pagesTableRules', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-jiliandongxuanzeqi', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(402, 366, '/pages/formI18n', 'pagesFormI18n', '/pages/formI18n/index', '', 'message.router.pagesFormI18n', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-diqiu', '2025-03-16 01:27:32', '2025-03-16 01:27:32'),
(403, 366, '/pages/formRules', 'pagesFormRules', '/pages/formRules/index', '', 'message.router.pagesFormRules', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-shuxing', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(404, 366, '/pages/listAdapt', 'pagesListAdapt', '/pages/listAdapt/index', '', 'message.router.pagesListAdapt', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-chazhaobiaodanliebiao', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(405, 366, '/pages/waterfall', 'pagesWaterfall', '/pages/waterfall/index', '', 'message.router.pagesWaterfall', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-zidingyibuju', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(406, 366, '/pages/steps', 'pagesSteps', '/pages/steps/index', '', 'message.router.pagesSteps', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-step', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(407, 366, '/pages/preview', 'pagesPreview', '/pages/preview/index', '', 'message.router.pagesPreview', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-15tupianyulan', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(408, 366, '/pages/waves', 'pagesWaves', '/pages/waves/index', '', 'message.router.pagesWaves', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-bolangneng', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(409, 366, '/pages/tree', 'pagesTree', '/pages/tree/index', '', 'message.router.pagesTree', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-shuxingtu', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(410, 366, '/pages/drag', 'pagesDrag', '/pages/drag/index', '', 'message.router.pagesDrag', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'ele-Pointer', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(411, 366, '/pages/lazyImg', 'pagesLazyImg', '/pages/lazyImg/index', '', 'message.router.pagesLazyImg', '', 0, 1, 0, 0, '[\"admin\"]', 'ele-PictureFilled', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(412, 366, '/pages/dynamicForm', 'pagesDynamicForm', '/pages/dynamicForm/index', '', 'message.router.pagesDynamicForm', '', 0, 1, 0, 0, '[\"admin\"]', 'iconfont icon-wenducanshu-05', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(413, 366, '/pages/workflow', 'pagesWorkflow', '/pages/workflow/index', '', 'message.router.pagesWorkflow', '', 0, 1, 0, 0, '[\"admin\"]', 'ele-Connection', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(414, 367, '/make/selector', 'makeSelector', '/make/selector/index', '', 'message.router.makeSelector', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-xuanzeqi', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(415, 367, '/make/noticeBar', 'makeNoticeBar', '/make/noticeBar/index', '', 'message.router.makeNoticeBar', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'ele-Bell', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(416, 367, '/make/svgDemo', 'makeSvgDemo', '/make/svgDemo/index', '', 'message.router.makeSvgDemo', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'fa fa-thumbs-o-up', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(417, 367, '/make/tableDemo', 'makeTableDemo', '/make/tableDemo/index', '', 'message.router.makeTableDemo', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', 'iconfont icon-shuju', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(418, 368, '/params/common', 'paramsCommon', '/params/common/index', '', 'message.router.paramsCommon', '', 0, 1, 0, 0, '[\"admin\"]', 'iconfont icon-putong', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(419, 368, '/params/common/details', 'paramsCommonDetails', '/params/common/details', '', 'message.router.paramsCommonDetails', '', 1, 1, 0, 0, '[\"admin\"]', 'ele-Comment', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(420, 368, '/params/dynamic', 'paramsDynamic', '/params/dynamic/index', '', 'message.router.paramsDynamic', '', 0, 1, 0, 0, '[\"admin\"]', 'iconfont icon-dongtai', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(421, 368, '/params/dynamic/details/:t/:id/:tagsViewName', 'paramsDynamicDetails', '/params/dynamic/details', '', 'message.router.paramsDynamicDetails', '', 1, 1, 0, 0, '[\"admin\"]', 'ele-Lightning', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(422, 369, '/visualizing/visualizingLinkDemo1', 'visualizingLinkDemo1', '/layout/routerView/link', '', 'message.router.visualizingLinkDemo1', '/visualizingDemo1', 0, 0, 0, 0, '[\"admin\"]', 'iconfont icon-caozuo-wailian', '2025-03-16 01:27:33', '2025-03-16 01:27:33'),
(423, 369, '/visualizing/visualizingLinkDemo2', 'visualizingLinkDemo2', '/layout/routerView/link', '', 'message.router.visualizingLinkDemo2', '/visualizingDemo2', 0, 0, 0, 0, '[\"admin\"]', 'iconfont icon-caozuo-wailian', '2025-03-16 01:27:34', '2025-03-16 01:27:34'),
(424, 381, '/limits/frontEnd/page', 'limitsFrontEndPage', '/limits/frontEnd/page/index', '', 'message.router.limitsFrontEndPage', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', '', '2025-03-16 01:27:34', '2025-03-16 01:27:34'),
(425, 381, '/limits/frontEnd/btn', 'limitsFrontEndBtn', '/limits/frontEnd/btn/index', '', 'message.router.limitsFrontEndBtn', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', '', '2025-03-16 01:27:34', '2025-03-16 01:27:34'),
(426, 382, '/limits/backEnd/page', 'limitsBackEndEndPage', '/limits/backEnd/page/index', '', 'message.router.limitsBackEndEndPage', '', 0, 1, 0, 0, '[\"admin\",\"common\"]', '', '2025-03-16 01:27:34', '2025-03-16 01:27:34');

-- --------------------------------------------------------

--
-- 表的结构 `TransferRecord`
--

CREATE TABLE `TransferRecord` (
  `id` int(10) UNSIGNED NOT NULL,
  `transfer_money` decimal(10,2) NOT NULL COMMENT '转账金额',
  `status` tinyint(4) NOT NULL COMMENT '状态',
  `operate_user` int(10) UNSIGNED DEFAULT NULL COMMENT '操作用户ID',
  `printer_id` int(10) UNSIGNED DEFAULT NULL COMMENT '打印机ID',
  `create_at` datetime NOT NULL DEFAULT current_timestamp() COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci COMMENT='资金转账记录';

-- --------------------------------------------------------

--
-- 表的结构 `User`
--

CREATE TABLE `User` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '用户ID',
  `passport` varchar(45) NOT NULL COMMENT '账号uid',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `nickname` varchar(45) NOT NULL COMMENT '用户昵称',
  `role` tinyint(4) NOT NULL DEFAULT 0 COMMENT '角色(0-普通用户 10管理员)',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像url',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(45) NOT NULL DEFAULT '' COMMENT '手机号',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '账户状态 (1: Normal, 2: Blocked)',
  `create_at` datetime NOT NULL DEFAULT current_timestamp(),
  `update_at` datetime NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci COMMENT='用户信息表';

--
-- 转存表中的数据 `User`
--

INSERT INTO `User` (`id`, `passport`, `password`, `nickname`, `role`, `avatar`, `email`, `phone`, `status`, `create_at`, `update_at`) VALUES
(1, 'Admin123', '123456', 'Admin123', 10, '', '', '', 1, '2025-03-04 07:00:31', '2025-03-04 20:26:58'),
(2, 'string', 'string', 'string', 0, '', '', '', 1, '2025-03-04 07:56:25', '2025-03-04 07:56:25'),
(3, 'strsssing', 'string', 'strsing', 0, '', '', '', 1, '2025-03-04 07:57:38', '2025-03-04 07:57:38');

-- --------------------------------------------------------

--
-- 表的结构 `UserDevice`
--

CREATE TABLE `UserDevice` (
  `id` int(10) UNSIGNED NOT NULL COMMENT '关联表主键',
  `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户ID',
  `device_id` int(10) UNSIGNED NOT NULL COMMENT '设备ID',
  `bind_time` datetime DEFAULT current_timestamp() COMMENT '绑定时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci COMMENT='用户与设备关联表';

--
-- 转储表的索引
--

--
-- 表的索引 `ABTokens`
--
ALTER TABLE `ABTokens`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `token` (`token`);

--
-- 表的索引 `BannedTokens`
--
ALTER TABLE `BannedTokens`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `token` (`token`);

--
-- 表的索引 `permissions`
--
ALTER TABLE `permissions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- 表的索引 `PrintServer`
--
ALTER TABLE `PrintServer`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `passport` (`passport`),
  ADD UNIQUE KEY `idx_passport` (`passport`);

--
-- 表的索引 `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `rolesID` (`rolesID`),
  ADD UNIQUE KEY `role_sign` (`role_sign`);

--
-- 表的索引 `role_permissions`
--
ALTER TABLE `role_permissions`
  ADD PRIMARY KEY (`role_id`,`permission_id`),
  ADD KEY `role_permissions_permission_id_permissions_id` (`permission_id`);

--
-- 表的索引 `SystemMenus`
--
ALTER TABLE `SystemMenus`
  ADD PRIMARY KEY (`id`),
  ADD KEY `parent_id` (`parent_id`);

--
-- 表的索引 `TransferRecord`
--
ALTER TABLE `TransferRecord`
  ADD PRIMARY KEY (`id`),
  ADD KEY `operate_user` (`operate_user`),
  ADD KEY `printer_id` (`printer_id`);

--
-- 表的索引 `User`
--
ALTER TABLE `User`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `passport` (`passport`),
  ADD UNIQUE KEY `idx_passport` (`passport`);

--
-- 表的索引 `UserDevice`
--
ALTER TABLE `UserDevice`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`device_id`),
  ADD KEY `device_id` (`device_id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `ABTokens`
--
ALTER TABLE `ABTokens`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `BannedTokens`
--
ALTER TABLE `BannedTokens`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `permissions`
--
ALTER TABLE `permissions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=37;

--
-- 使用表AUTO_INCREMENT `PrintServer`
--
ALTER TABLE `PrintServer`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `roles`
--
ALTER TABLE `roles`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=40;

--
-- 使用表AUTO_INCREMENT `SystemMenus`
--
ALTER TABLE `SystemMenus`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键 ID', AUTO_INCREMENT=432;

--
-- 使用表AUTO_INCREMENT `TransferRecord`
--
ALTER TABLE `TransferRecord`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `User`
--
ALTER TABLE `User`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `UserDevice`
--
ALTER TABLE `UserDevice`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '关联表主键';

--
-- 限制导出的表
--

--
-- 限制表 `role_permissions`
--
ALTER TABLE `role_permissions`
  ADD CONSTRAINT `role_permissions_permission_id_permissions_id` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `role_permissions_role_id_roles_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE;

--
-- 限制表 `SystemMenus`
--
ALTER TABLE `SystemMenus`
  ADD CONSTRAINT `SystemMenus_ibfk_1` FOREIGN KEY (`parent_id`) REFERENCES `SystemMenus` (`id`) ON DELETE CASCADE;

--
-- 限制表 `TransferRecord`
--
ALTER TABLE `TransferRecord`
  ADD CONSTRAINT `TransferRecord_ibfk_1` FOREIGN KEY (`operate_user`) REFERENCES `User` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  ADD CONSTRAINT `TransferRecord_ibfk_2` FOREIGN KEY (`printer_id`) REFERENCES `PrintServer` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;

--
-- 限制表 `UserDevice`
--
ALTER TABLE `UserDevice`
  ADD CONSTRAINT `UserDevice_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `User` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `UserDevice_ibfk_2` FOREIGN KEY (`device_id`) REFERENCES `PrintServer` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
