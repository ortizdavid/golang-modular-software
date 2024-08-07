-- Schema: configuration 

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

-- default configs
-- Inserting a record into the 'configurations.email_configuration' table
INSERT INTO configurations.email_configuration (smtp_server, smtp_port, sender_email, sender_password, unique_id, created_at, updated_at)
VALUES
('smtp.example.com', '587', 'noreply@example.com', 'password123', 'email-config-001', NOW(), NOW());

-- Inserting a record into the 'configurations.basic_configuration' table
INSERT INTO configurations.basic_configuration (app_name, app_acronym, max_records_per_page, max_admin_users, max_super_admin_users, unique_id, created_at, updated_at)
VALUES
('MyApp', 'MA', 50, 10, 2, 'basic-config-001', NOW(), NOW());

-- Inserting a record into the 'configurations.company_configuration' table
INSERT INTO configurations.company_configuration (company_name, company_acronym, company_main_color, company_logo, company_email, company_phone, unique_id, created_at, updated_at)
VALUES
('Example Corp', 'EC', '#123456', 'logo.png', 'contact@example.com', '+1234567890', 'company-config-001', NOW(), NOW());
