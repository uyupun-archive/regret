-- +migrate Up
CREATE TABLE IF NOT EXISTS inquiry_validations(
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    service_id INT UNSIGNED NOT NULL,
    is_required_subject BOOLEAN NOT NULL,
    is_required_email BOOLEAN NOT NULL,
    is_required_category BOOLEAN NOT NULL,
    is_required_text BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY fk_service_id(service_id) REFERENCES services(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS inquiry_validations;
