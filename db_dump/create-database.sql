CREATE DATABASE IF NOT EXISTS `drankbeurs`;
USE `drankbeurs`;

CREATE TABLE `bar` (
  `id` varchar(45) NOT NULL,
  `name` varchar(45) NOT NULL,
  `password` varchar(45) NOT NULL,
  `token` varchar(45) NOT NULL,
  `super_admin` tinyint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  UNIQUE KEY `token_UNIQUE` (`token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `category` (
  `id` varchar(45) NOT NULL,
  `name` varchar(45) NOT NULL,
  `bar_id` varchar(45) NOT NULL,
  `color` varchar(45) NOT NULL DEFAULT 'FFFFFF',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `implemented_by_idx` (`bar_id`),
  CONSTRAINT `implemented_by` FOREIGN KEY (`bar_id`) REFERENCES `bar` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `product` (
  `id` varchar(45) NOT NULL,
  `name` varchar(45) NOT NULL,
  `bar_id` varchar(45) NOT NULL,
  `start_price` decimal(10,2) NOT NULL,
  `current_price` decimal(10,2) NOT NULL,
  `rise_multiplier` decimal(10,5) NOT NULL,
  `tag` varchar(16) NOT NULL,
  `category_id` varchar(45) NOT NULL,
  `drop_multiplier` decimal(10,5) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`),
  KEY `room_idx` (`bar_id`),
  KEY `belongs_to_idx` (`category_id`),
  CONSTRAINT `belongs_to` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`),
  CONSTRAINT `sold_by` FOREIGN KEY (`bar_id`) REFERENCES `bar` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `device` (
  `id` varchar(45) NOT NULL,
  `bar_id` varchar(45) NOT NULL,
  `name` varchar(45) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `room_idx` (`bar_id`),
  CONSTRAINT `belongs_to_room` FOREIGN KEY (`bar_id`) REFERENCES `bar` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `order` (
  `id` varchar(45) NOT NULL,
  `device_id` varchar(45) NOT NULL,
  `product_id` varchar(45) NOT NULL,
  `timestamp` varchar(45) NOT NULL,
  `amount` int NOT NULL DEFAULT '0',
  `price_per_product` decimal(10,2) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `product_idx` (`product_id`),
  KEY `ordered_by_idx` (`device_id`),
  CONSTRAINT `contains` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`) ON UPDATE RESTRICT,
  CONSTRAINT `ordered_by` FOREIGN KEY (`device_id`) REFERENCES `device` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
