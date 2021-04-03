CREATE TABLE IF NOT EXISTS `products` (
                            `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                            `code` varchar(255) NOT NULL,
                            `name` varchar(255) NOT NULL,
                            `brand_id` int(10) unsigned NOT NULL,
                            `country_id` int(10) unsigned NOT NULL,
                            `description` varchar(255) DEFAULT NULL,
                            `product_type_id` int(10) unsigned NOT NULL,
                            `deleted` smallint(6) DEFAULT '0',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `code` (`code`),
                            KEY `brand_id` (`brand_id`),
                            KEY `product_type_id` (`product_type_id`),
                            KEY `country_id` (`country_id`),
                            CONSTRAINT `products_ibfk_1` FOREIGN KEY (`brand_id`) REFERENCES `brands` (`id`) ON UPDATE CASCADE,
                            CONSTRAINT `products_ibfk_2` FOREIGN KEY (`product_type_id`) REFERENCES `product_types` (`id`) ON UPDATE CASCADE,
                            CONSTRAINT `products_ibfk_3` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`) ON UPDATE CASCADE
)