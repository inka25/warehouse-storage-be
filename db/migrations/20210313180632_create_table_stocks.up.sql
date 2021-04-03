CREATE TABLE IF NOT EXISTS `stocks` (
                          `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                          `product_id` int(10) unsigned NOT NULL,
                          `warehouse_id` int(10) unsigned NOT NULL,
                          `stocks` bigint(20) DEFAULT '0',
                          `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `updated_by` varchar(255) NOT NULL,
                          PRIMARY KEY (`id`),
                          KEY `products` (`product_id`,`warehouse_id`,`stocks`),
                          KEY `warehouse` (`warehouse_id`,`product_id`,`stocks`),
                          KEY `stock` (`stocks`)
)