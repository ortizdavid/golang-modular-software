-- tables for 'configurations' schema 

-- Table: email_configuration
DROP TABLE IF EXISTS configurations.email_configuration;
CREATE TABLE configurations.email_configuration (
    configuration_id SERIAL PRIMARY KEY,
    smtp_server VARCHAR(50),
    smtp_port VARCHAR(4),
    sender_email VARCHAR(100),
    sender_password VARCHAR(150),
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
-- Inserting a record into the 'configurations.email_configuration' table
INSERT INTO configurations.email_configuration (smtp_server, smtp_port, sender_email, sender_password, unique_id, created_at, updated_at)
VALUES
('smtp.example.com', '587', 'noreply@example.com', 'password123', 'email-config-001', NOW(), NOW());


-- Table: basic_configuration
DROP TABLE IF EXISTS configurations.basic_configuration;
CREATE TABLE configurations.basic_configuration (
    configuration_id SERIAL PRIMARY KEY,
    app_name VARCHAR(100) NOT NULL,
    app_acronym VARCHAR(50),
    max_records_per_page INT,
    max_admin_users INT,
    max_super_admin_users INT,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
-- Inserting a record into the 'configurations.basic_configuration' table
INSERT INTO configurations.basic_configuration (app_name, app_acronym, max_records_per_page, max_admin_users, max_super_admin_users, unique_id, created_at, updated_at)
VALUES
('MyApp', 'MA', 50, 10, 2, 'basic-config-001', NOW(), NOW());


-- Table: company_configuration
DROP TABLE IF EXISTS configurations.company_configuration;
CREATE TABLE configurations.company_configuration (
    configuration_id SERIAL PRIMARY KEY,
    company_name VARCHAR(100) NOT NULL,
    company_acronym VARCHAR(50),
    company_main_color VARCHAR(10),
    company_logo VARCHAR(100),
    company_email VARCHAR(100),
    company_phone VARCHAR(20),
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
-- Inserting a record into the 'configurations.company_configuration' table
INSERT INTO configurations.company_configuration (company_name, company_acronym, company_main_color, company_logo, company_email, company_phone, unique_id, created_at, updated_at)
VALUES
('Example Corp', 'EC', '#123456', 'logo.png', 'contact@example.com', '+1234567890', 'company-config-001', NOW(), NOW());


-- Table: modules
DROP TABLE IF EXISTS configurations.modules;
CREATE TABLE configurations.modules (
    module_id SERIAL PRIMARY KEY,
    module_name VARCHAR(100) UNIQUE NOT NULL,
    code VARCHAR(30) UNIQUE NOT NULL,
    description TEXT,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO configurations.modules (module_id, module_name, code, description) VALUES 
(1, 'Authentication', 'configurations', 'Handles user authentication and authorization.'),
(2, 'Company', 'company', 'Manages company-related information and settings.'),
(3, 'Employees', 'employees', 'Manages employee records and details.'),
(4, 'References', 'references', 'Stores and manages reference data.'),
(5, 'Reports', 'reports', 'Generates and manages various reports.'),
(6, 'Configurations', 'configurations', 'Manages system configurations.');


-- Table: features
DROP TABLE IF EXISTS configurations.features;
CREATE TABLE configurations.features (
    feature_id SERIAL PRIMARY KEY,
    module_id INT NOT NULL,
    feature_name VARCHAR(100) UNIQUE NOT NULL,
    code VARCHAR(30) UNIQUE NOT NULL,
    description TEXT,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_module_feature FOREIGN KEY(module_id) REFERENCES configurations.modules(module_id)
);


-- FEATURE FLAGS
DROP TYPE IF EXISTS TYPE_FLAG_STATUS;
CREATE TYPE TYPE_FLAG_STATUS AS ENUM ('Enabled', 'Disabled');

-- Table: module_flag
DROP TABLE IF EXISTS configurations.module_flag;
CREATE TABLE configurations.module_flag(
    flag_id SERIAL PRIMARY KEY,
    module_id INT NOT NULL,
    status TYPE_FLAG_STATUS DEFAULT 'Disabled',
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_flag_module FOREIGN KEY(module_id) REFERENCES configurations.modules(module_id)
);
-- Insert module flags
INSERT INTO configurations.module_flag (module_id, status) VALUES 
(1, 'Enabled'),  -- Authentication
(2, 'Disabled'),  -- Company
(3, 'Disabled'),  -- Employees
(4, 'Disabled'),  -- References
(5, 'Disabled'),  -- Reports
(6, 'Enabled');   -- Configuration


-- Table: feature_flag
DROP TABLE IF EXISTS configurations.feature_flag;
CREATE TABLE configurations.feature_flag(
    flag_id SERIAL PRIMARY KEY,
    feature_id INT NOT NULL,
    status TYPE_FLAG_STATUS DEFAULT 'Disabled',
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_flag_feature FOREIGN KEY(feature_id) REFERENCES configurations.features(feature_id)
);

