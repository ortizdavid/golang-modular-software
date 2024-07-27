-- SCHEMAS 

-- for all configurations
CREATE SCHEMA IF NOT EXISTS configurations;

-- user, roles, permission, authentication management
CREATE SCHEMA IF NOT EXISTS authentication;

-- company management
CREATE SCHEMA IF NOT EXISTS company;

-- reference management
CREATE SCHEMA IF NOT EXISTS reference;

-- seed
CREATE TABLE IF NOT EXISTS seeds (
    id SERIAL PRIMARY KEY,
    script_name VARCHAR(255) UNIQUE,
    executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
