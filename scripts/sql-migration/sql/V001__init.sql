CREATE TABLE `badge`
(
    `id`            BIGINT PRIMARY KEY AUTO_INCREMENT,
    `name`          VARCHAR(64) NOT NULL,
    `image`         TEXT NOT NULL,

    `category`      VARCHAR(64) NOT NULL,
    `description`   VARCHAR(64) NOT NULL,
    `rarity`        VARCHAR(64) NOT NULL,
    `trigger_event` VARCHAR(64) NOT NULL,

    `condition`     int         NOT NULL,

    `status`        VARCHAR(64) NOT NULL,

    `created_at`    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO badge VALUES
(1, '100 day runner', 'test', 'login', 'login for more than 100 days', '1', 'login', 100, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(2, '200 day runner', 'test', 'login', 'login for more than 200 days', '2', 'login', 200, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(3, '500 day runner', 'test', 'login', 'login for more than 500 days', '3', 'login', 500, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(4, '1000 day runner', 'test', 'login', 'login for more than 1000 days', '4', 'login', 1000, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(5, '2000 day runner', 'test', 'login', 'login for more than 2000 days', '5', 'login', 2000, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(6, '1 year anniversary', 'test', 'anniversary login', 'login for 1 year anniversary', '1', 'anniversary login', 1, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(7, '2 year anniversary', 'test', 'anniversary login', 'login for 2 year anniversary', '1', 'anniversary login', 2, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(8, '3 year anniversary', 'test', 'anniversary login', 'login for 3 year anniversary', '2', 'anniversary login', 3, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(9, '4 year anniversary', 'test', 'anniversary login', 'login for 4 year anniversary', '2', 'anniversary login', 4, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(10, '5 year anniversary', 'test', 'anniversary login', 'login for 5 year anniversary', '3', 'anniversary login', 5, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(11, '6 year anniversary', 'test', 'anniversary login', 'login for 6 year anniversary', '3', 'anniversary login', 6, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(12, '200 hour fighter', 'test', 'long time watch', 'watch time longer than 200 hours', '1', 'anniversary login', 200, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(13, '500 hour fighter', 'test', 'long time watch', 'watch time longer than 500 hours', '2', 'anniversary login', 500, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(14, '1000 hour fighter', 'test', 'long time watch', 'watch time longer than 1000 hours', '3', 'anniversary login', 1000, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(15, '2000 hour fighter', 'test', 'long time watch', 'watch time longer than 2000 hours', '4', 'anniversary login', 2000, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(16, '3000 hour fighter', 'test', 'long time watch', 'watch time longer than 3000 hours', '5', 'anniversary login', 3000, 'PUBLISHED', '2023-05-13', '2023-05-13'),
(17, 'birthday', 'test', 'birthday login', 'login in on birthday', '3', 'birthday login', 1, 'PUBLISHED', '2023-05-13', '2023-05-13');



CREATE TABLE `user_badge_asset`
(
    `id`            BIGINT PRIMARY KEY AUTO_INCREMENT,
    `user_id`       BIGINT    NOT NULL,
    `badge_id`     BIGINT    NOT NULL,
    `current_state` INT       NOT NULL,
    `created_at`    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO user_badge_asset VALUES
(1, 10000, 1, 1, '2023-05-13', '2023-05-13'),
(2, 10000, 2, 1, '2023-05-13', '2023-05-13'),
(3, 10000, 3, 1, '2023-05-13', '2023-05-13'),
(4, 10000, 6, 1, '2023-05-13', '2023-05-13'),
(5, 10000, 7, 1, '2023-05-13', '2023-05-13'),
(6, 10000, 12, 1, '2023-05-13', '2023-05-13'),
(7, 10000, 17, 1, '2023-05-13', '2023-05-13'),
(8, 20000, 1, 1, '2023-05-13', '2023-05-13'),
(9, 20000, 2, 1, '2023-05-13', '2023-05-13'),
(10, 20000, 6, 1, '2023-05-13', '2023-05-13'),
(11, 20000, 12, 1, '2023-05-13', '2023-05-13'),
(12, 20000, 13, 1, '2023-05-13', '2023-05-13'),
(13, 20000, 4, 1, '2023-05-13', '2023-05-13');