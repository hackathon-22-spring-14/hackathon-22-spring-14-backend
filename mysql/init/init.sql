DROP DATABASE IF EXISTS hackathon;
CREATE DATABASE hackathon;
USE hackathon;

CREATE TABLE IF NOT EXISTS `stamps` (
  `id` char(36) NOT NULL,
  `name` varchar(32) NOT NULL,
  `image_url` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- TODO: テストデータを後で消す
INSERT INTO `stamps` (`id`, `name`, `image_url`) VALUES
('f33255b9-294a-4fb2-a00b-34a40ddfba8e', 'stamp1', 'https://example.com/stamp1.png'),
('d493daea-a0a3-447e-85c8-52331e3f018c', 'stamp2', 'https://example.com/stamp2.png');

CREATE TABLE IF NOT EXISTS `users` (
  `id` varchar(32) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- TODO: テストデータを後で消す
INSERT INTO `users` (`id`, `password`) VALUES
('user1', '$2a$10$.pn8EY6zCytzgV3JW5dXYeZ2xnsUI2cmuCsFbbYlsuGGotKx4qOhO'), -- password="password"
('user2', '$2a$10$777RHmBriDm7ilr64wim8OaQCAMiqzVS.Dwn.UnqlOSJDwcWxUi1m'); -- password="password2",bcrypt
