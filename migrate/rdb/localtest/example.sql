-- example: table
DROP TABLE IF EXISTS `example`;

CREATE TABLE `example`
(
    `id`         bigint                                  NOT NULL COMMENT 'id',
    `name`       varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'name',
    `created_at` timestamp(6)                            NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `updated_at` timestamp(6)                            NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='測試用';


insert ignore into `example` (id, name) values (1,'');
insert ignore into `example` (id, name) values (2,'');


