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
  `title` varchar(255) NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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

CREATE TABLE `relations` (
  `follow_id` bigint PRIMARY KEY AUTO_INCREMENT,
  `followed_id` bigint NOT NULL,
  `follower_id` bigint NOT NULL,
  `deleted` tinyint(4) NOT NULL DEFAULT 0
);

ALTER TABLE `relations` ADD FOREIGN KEY (`followed_id`) REFERENCES `users` (`user_id`);
ALTER TABLE `relations` ADD FOREIGN KEY (`follower_id`) REFERENCES `users` (`user_id`);

-- comment model
CREATE TABLE `comments`  (
    `comment_id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint NOT NULL,
    `video_id` bigint NOT NULL,
    `content` varchar(255) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    PRIMARY KEY (`comment_id`),
    CONSTRAINT `comment_user` FOREIGN KEY (`user_id`) REFERENCES `simple_douyin`.`users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `comment_video` FOREIGN KEY (`video_id`) REFERENCES `simple_douyin`.`videos` (`video_id`) ON DELETE CASCADE ON UPDATE CASCADE
);
