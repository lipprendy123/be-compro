CREATE TABLE IF NOT EXISTS portofolio_testimonials (
    id INT AUTO_INCREMENT PRIMARY KEY,
    portofolio_section_id INT,
    client_name TEXT,
    thumbnail VARCHAR(200),
    message TEXT,
    role VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (portofolio_section_id) REFERENCES portofolio_sections(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE INDEX idx_portofolio_testimonials_portofolios_section_id ON portofolio_testimonials(portofolio_section_id);
