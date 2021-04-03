CREATE TABLE IF NOT EXISTS `countries` (
                             `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                             `country` char(2) NOT NULL,
                             `deleted` smallint(6) DEFAULT '0',
                             PRIMARY KEY (`id`),
                             UNIQUE KEY `country` (`country`)
) ;