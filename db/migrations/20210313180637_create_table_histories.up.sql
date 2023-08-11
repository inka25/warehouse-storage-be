CREATE TABLE IF NOT EXISTS `histories` (
                             `product_id` int(10) unsigned DEFAULT NULL,
                             `type` varchar(255) NOT NULL,
                             `previous_info` json NOT NULL,
                             `new_info` json NOT NULL,
                             `description` varchar(255) DEFAULT NULL,
                             `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             `updated_by` varchar(255) NOT NULL,
                             KEY `product_id` (`product_id`),
                             CONSTRAINT `histories_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`) ON UPDATE CASCADE
)