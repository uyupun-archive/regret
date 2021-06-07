-- +migrate Up
CREATE TABLE IF NOT EXISTS categories(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    name_ja VARCHAR(20) NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS categories;
