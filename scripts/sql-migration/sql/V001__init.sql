CREATE TABLE `badget`
(
    `id`            BIGINT PRIMARY KEY AUTO_INCREMENT,
    `name`          VARCHAR(64) NOT NULL,
    `image`         VARCHAR(64) NOT NULL,

    `category`      VARCHAR(64) NOT NULL,
    `description`   VARCHAR(64) NOT NULL,
    `rarity`        VARCHAR(64) NOT NULL,
    `trigger_event` VARCHAR(64) NOT NULL,

    `condition`     int         NOT NULL,

    `status`        VARCHAR(64) NOT NULL,

    `created_at`    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `user_badget_asset`
(
    `id`            BIGINT PRIMARY KEY AUTO_INCREMENT,
    `user_id`       BIGINT    NOT NULL,
    `badget_id`     BIGINT    NOT NULL,
    `current_state` INT       NOT NULL,
    `created_at`    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);