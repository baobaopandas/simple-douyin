CREATE TABLE `users` (
  `user_id` bigint PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `follow_count` bigint DEFAULT 0,
  `follower_count` bigint DEFAULT 0 
);

CREATE TABLE `videos` (
  `video_id` bigint PRIMARY KEY AUTO_INCREMENT,
  `author` bigint NOT NULL,
  `play_url` varchar(255) NOT NULL,
  `cover_url` varchar(255) NOT NULL,
  `favorite_count` bigint DEFAULT 0,
  `comment_count` bigint DEFAULT 0, 
  `title` varchar(255) NOT NULL
);

CREATE TABLE `favorite` (
  `favorite_id` bigint PRIMARY KEY AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`),
  FOREIGN KEY (`video_id`) REFERENCES `videos` (`video_id`),
  `statement`  BOOLEAN NOT NULL
);

ALTER TABLE `videos` ADD FOREIGN KEY (`author`) REFERENCES `users` (`user_id`);
