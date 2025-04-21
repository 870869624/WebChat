-- -------------------------------------------------------------
-- TablePlus 6.0.0(550)
--
-- https://tableplus.com/
--
-- Database: go_gin_chat
-- Generation Time: 2025-04-21 下午3:08:34.5020
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


CREATE TABLE `friends` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '用户ID',
  `to_user_id` int DEFAULT '0' COMMENT '私聊用户ID',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `messages` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL COMMENT '用户ID',
  `room_id` int NOT NULL COMMENT '房间ID',
  `to_user_id` int DEFAULT '0' COMMENT '私聊用户ID',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '聊天内容',
  `image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '图片URL',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `password` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '密码',
  `avatar_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '1' COMMENT '头像ID',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC;

INSERT INTO `friends` (`id`, `user_id`, `to_user_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(7, 1, 2, '2025-04-21 14:07:52', '2025-04-21 14:07:52', NULL),
(8, 1, 4, '2025-04-21 14:28:45', '2025-04-21 14:28:45', NULL),
(9, 4, 3, '2025-04-21 14:57:02', '2025-04-21 14:57:02', NULL);

INSERT INTO `messages` (`id`, `user_id`, `room_id`, `to_user_id`, `content`, `image_url`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 1, 1, 0, '11111', '', '2025-04-21 13:18:26', '2025-04-21 13:18:26', NULL),
(2, 2, 1, 0, '123', '', '2025-04-21 13:21:15', '2025-04-21 13:21:15', NULL),
(3, 1, 1, 0, '111', '', '2025-04-21 13:27:50', '2025-04-21 13:27:50', NULL),
(4, 1, 1, 0, '111', '', '2025-04-21 13:52:28', '2025-04-21 13:52:28', NULL),
(5, 3, 1, 0, '123123123', '', '2025-04-21 14:28:12', '2025-04-21 14:28:12', NULL),
(6, 4, 1, 0, 'qweqweqw', '', '2025-04-21 14:28:26', '2025-04-21 14:28:26', NULL);

INSERT INTO `users` (`id`, `username`, `password`, `avatar_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, '123', 'e10adc3949ba59abbe56e057f20f883e', '1', '2025-04-21 13:18:07', '2025-04-21 14:57:05', NULL),
(2, '456', 'e10adc3949ba59abbe56e057f20f883e', '1', '2025-04-21 13:20:59', '2025-04-21 13:20:59', NULL),
(3, '789', 'e10adc3949ba59abbe56e057f20f883e', '1', '2025-04-21 14:28:05', '2025-04-21 14:28:05', NULL),
(4, 'qwe', 'e10adc3949ba59abbe56e057f20f883e', '1', '2025-04-21 14:28:20', '2025-04-21 14:56:57', NULL);



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;