CREATE TABLE IF NOT EXISTS portofolio_details (
    id INT AUTO_INCREMENT PRIMARY KEY,
    portofolio_section_id INT,
    category VARCHAR(150),
    client_name TEXT,
    project_date TIMESTAMP,
    project_url TEXT NULL,
    title VARCHAR(200),
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (portofolio_section_id) REFERENCES portofolio_sections(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE INDEX idx_portofolio_details_portofolios_section_id ON portofolio_details(portofolio_section_id);
