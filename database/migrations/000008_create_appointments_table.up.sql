CREATE TABLE IF NOT EXISTS appointments (
    id INT AUTO_INCREMENT PRIMARY KEY,
    service_id INT,
    name VARCHAR(150),
    phone_number VARCHAR(15),
    email VARCHAR(150),
    brief TEXT,
    meet_at TIMESTAMP NOT NULL,
    budget DECIMAL(10,1) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (service_id) REFERENCES service_sections(id) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE INDEX idx_appointments_service_id ON appointments(service_id);
