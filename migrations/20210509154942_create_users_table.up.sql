CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(40) PRIMARY KEY,
    username VARCHAR(244),
    email VARCHAR(255),
    password_hash VARCHAR(255),
    attributes JSON,
    verified TINYINT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    project_id VARCHAR(40),
    client_id VARCHAR(255),
    FOREIGN KEY (project_id) REFERENCES projects (id) ON DELETE CASCADE,
    FOREIGN KEY (client_id) REFERENCES clients (id) ON DELETE CASCADE
);