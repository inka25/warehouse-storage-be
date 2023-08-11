CREATE TABLE IF NOT EXISTS `inventories` (
                          `product_id` int(10) unsigned NOT NULL,
                          `warehouse_id` int(10) unsigned NOT NULL,
                          `stock` bigint(20) DEFAULT '0',
                          `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `updated_by` varchar(255) NOT NULL,
                          KEY `products` (`product_id`,`warehouse_id`,`stock`),
                          KEY `warehouse` (`warehouse_id`,`product_id`,`stock`),
                          KEY `stock` (`stock`)
)