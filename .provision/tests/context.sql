CREATE TABLE `context`
(
    `id`        bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `uuid`      char(36)            NOT NULL,
    `uuidHash`  int(10) unsigned    NOT NULL,
    `createdAt` bigint(20) unsigned NOT NULL,
    PRIMARY KEY (`id`),
    KEY `uuidIdx` (`uuidHash`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;