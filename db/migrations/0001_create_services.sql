-- +migrate Up
CREATE TABLE IF NOT EXISTS services(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    name_ja VARCHAR(20) NOT NULL,
    description VARCHAR(100),
    access_token VARCHAR(20) NOT NULL UNIQUE
);

-- +migrate Down
DROP TABLE IF EXISTS services;
