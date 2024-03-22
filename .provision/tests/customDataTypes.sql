CREATE TABLE `customDataTypes`
(
    `id`                    BIGINT(20) unsigned    NOT NULL AUTO_INCREMENT,
    `typeUnixSeconds`       BIGINT(20) unsigned    NOT NULL,
    `typeUnixSecondsSigned` BIGINT(20)             NOT NULL,
    `typeUnixMillis`        BIGINT(20) unsigned    NOT NULL,
    `typeUnixMillisSigned`  BIGINT(20)             NOT NULL,
    `typeElapsedNanos`      BIGINT(20) unsigned    NOT NULL,
    `enabled`               ENUM ('true', 'false') NOT NULL,
    `json`                  TEXT                   NOT NULL,
    PRIMARY KEY (`id`)
);