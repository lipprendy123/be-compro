CREATE TABLE IF NOT EXISTS about_company_keynotes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    about_company_id INT,
    keypoint TEXT,
    path_image TEXT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (about_company_id) REFERENCES about_companies(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE INDEX idx_about_company_keynotes_about_company_id ON about_company_keynotes(about_company_id);
