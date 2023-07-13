CREATE DATABASE `tulisaja` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */;

-- tulisaja.likes definition

CREATE TABLE `likes` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` bigint(20) unsigned DEFAULT NULL,
  `counter` bigint(20) unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT current_timestamp(3),
  `updated_at` datetime(3) DEFAULT current_timestamp(3),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- tulisaja.users definition

CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `full_name` longtext DEFAULT NULL,
  `username` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT current_timestamp(3),
  `updated_at` datetime(3) DEFAULT current_timestamp(3),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- tulisaja.articles definition

CREATE TABLE `articles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `content` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT current_timestamp(3),
  `updated_at` datetime(3) DEFAULT current_timestamp(3),
  PRIMARY KEY (`id`),
  KEY `fk_users_article` (`user_id`),
  CONSTRAINT `fk_users_article` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- tulisaja.comments definition

CREATE TABLE `comments` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `article_id` bigint(20) unsigned DEFAULT NULL,
  `comment` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT current_timestamp(3),
  `updated_at` datetime(3) DEFAULT current_timestamp(3),
  PRIMARY KEY (`id`),
  KEY `fk_comments_user` (`user_id`),
  KEY `fk_articles_comment` (`article_id`),
  CONSTRAINT `fk_articles_comment` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`),
  CONSTRAINT `fk_comments_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- tulisaja.followers definition

CREATE TABLE `followers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned DEFAULT NULL,
  `follow_user_id` bigint(20) unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT current_timestamp(3),
  `updated_at` datetime(3) DEFAULT current_timestamp(3),
  PRIMARY KEY (`id`),
  KEY `fk_followers_follow_user` (`follow_user_id`),
  KEY `fk_users_followers` (`user_id`),
  CONSTRAINT `fk_followers_follow_user` FOREIGN KEY (`follow_user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_users_followers` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;