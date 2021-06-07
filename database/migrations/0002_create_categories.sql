-- +migrate Up
CREATE TABLE IF NOT EXISTS categories(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    service_id INT UNSIGNED NOT NULL,
    name VARCHAR(20) NOT NULL,
    name_ja VARCHAR(20) NOT NULL,
    FOREIGN KEY fk_service_id(service_id) REFERENCES services(id)
);

-- +migrate Down
DROP TABLE IF EXISTS categories;
