CREATE TABLE IF NOT EXISTS service_sections (
    id INT AUTO_INCREMENT PRIMARY KEY,
    path_icon TEXT,
    name VARCHAR(150),
    tagline TEXT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL
) ENGINE=InnoDB;
