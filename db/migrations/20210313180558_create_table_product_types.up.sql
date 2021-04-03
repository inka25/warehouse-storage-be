CREATE TABLE IF NOT EXISTS  `product_types` (
                                 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                 `name` varchar(255) NOT NULL,
                                 `deleted` smallint(6) DEFAULT '0',
                                 PRIMARY KEY (`id`)
);