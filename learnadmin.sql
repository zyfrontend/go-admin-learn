/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31)
 Source Host           : localhost:3306
 Source Schema         : learnadmin

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31)
 File Encoding         : 65001

 Date: 26/01/2023 14:14:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_users
-- ----------------------------
DROP TABLE IF EXISTS `app_users`;
CREATE TABLE `app_users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户名',
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '密码',
  `address` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '钱包地址',
  `salt` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '加密串',
  `chain` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '安全密码',
  `amount` bigint NULL DEFAULT NULL COMMENT '用户余额',
  `pay_u` bigint NULL DEFAULT NULL COMMENT '开启机器人投资Usdt余额',
  `amount_usdt` bigint NULL DEFAULT NULL COMMENT 'lp添池Usdt余额',
  `amount_lp` bigint NULL DEFAULT NULL COMMENT '开启机器人Lp量',
  `level` bigint NULL DEFAULT NULL COMMENT '等级',
  `recommend_user_id` bigint NULL DEFAULT NULL COMMENT '推荐人id',
  `is_broker` bigint NULL DEFAULT NULL COMMENT '是否矿工',
  `is_member` bigint NULL DEFAULT NULL COMMENT '是否会员',
  `status` bigint NULL DEFAULT NULL COMMENT '状态',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_app_users_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_app_users_username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of app_users
-- ----------------------------
INSERT INTO `app_users` VALUES (1, '2023-01-22 15:06:14.922', '2023-01-22 15:06:14.922', NULL, 'frontend', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (2, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (3, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (4, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (5, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (6, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (7, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (8, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (9, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (10, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (11, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (12, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (13, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (14, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (15, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (16, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (17, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (18, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (19, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'coder', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);
INSERT INTO `app_users` VALUES (20, '2023-01-22 22:03:25.985', '2023-01-22 22:03:25.985', NULL, 'enduser', '11111111', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
