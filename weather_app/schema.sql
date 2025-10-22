CREATE DATABASE IF NOT EXISTS weatherdb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE weatherdb;

CREATE TABLE IF NOT EXISTS weather_queries (
    id INT AUTO_INCREMENT PRIMARY KEY,
    city VARCHAR(100) NOT NULL,
    temperature FLOAT NOT NULL,
    description VARCHAR(255) NOT NULL,
    queried_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    ip_address VARCHAR(45)
);

CREATE USER IF NOT EXISTS 'weatheruser'@'%' IDENTIFIED BY 'password123';

GRANT ALL PRIVILEGES ON weatherdb.* TO 'weatheruser'@'%';

FLUSH PRIVILEGES;
