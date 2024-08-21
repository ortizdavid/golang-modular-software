-- tables for 'authentication' schema 

-- Table: roles
DROP TABLE IF EXISTS authentication.roles;
CREATE TABLE authentication.roles (
    role_id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    role_name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Table: users
DROP TABLE IF EXISTS authentication.users;
CREATE TABLE authentication.users (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    user_image VARCHAR(100), 
    is_active BOOLEAN DEFAULT TRUE,
    token VARCHAR(150) UNIQUE,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Table: user_roles
DROP TABLE IF EXISTS authentication.user_roles;
CREATE TABLE authentication.user_roles (
    user_role_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    role_id INT NOT NULL,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES  authentication.users(user_id),
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES  authentication.roles(role_id)
);

-- Table: permissions
DROP TABLE IF EXISTS authentication.permissions;
CREATE TABLE authentication.permissions (
    permission_id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    permission_name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Table: permission_roles
DROP TABLE IF EXISTS authentication.permission_roles;
CREATE TABLE authentication.permission_roles (
    permission_role_id SERIAL PRIMARY KEY,
    role_id INT NOT NULL,
    permission_id INT NOT NULL,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES  authentication.roles(role_id),
    CONSTRAINT fk_permission FOREIGN KEY(permission_id) REFERENCES  authentication.permissions(permission_id)
);

-- TYPE_ACTIVITY_STATUS
DROP TYPE IF EXISTS TYPE_ACTIVITY_STATUS;
CREATE TYPE TYPE_ACTIVITY_STATUS AS ENUM('Online', 'Offline');

-- Table: login_activity
DROP TABLE IF EXISTS authentication.login_activity;
CREATE TABLE authentication.login_activity (
    login_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    status TYPE_ACTIVITY_STATUS DEFAULT 'Offline',
    host VARCHAR(150),
    browser VARCHAR(150),
    ip_address VARCHAR(50),
    device VARCHAR(150),
    location VARCHAR(150),
    last_login TIMESTAMP DEFAULT NOW(),
    last_logout TIMESTAMP DEFAULT NOW(),
    total_login INT DEFAULT 0,
    total_logout INT DEFAULT 0,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES  authentication.users(user_id)
);
-- Index
DROP INDEX IF EXISTS idx_login_user_id;
CREATE INDEX idx_login_user_id ON authentication.login_activity(user_id);

-- Table: user_api_key
DROP TABLE IF EXISTS authentication.user_api_key;
CREATE TABLE authentication.user_api_key (
    api_key_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    key VARCHAR(150) UNIQUE,
    is_active BOOLEAN DEFAULT TRUE,
    created_by INT,
    expires_at TIMESTAMP NOT NULL,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_user_key FOREIGN KEY(user_id) REFERENCES  authentication.users(user_id)
);
-- Index
DROP INDEX IF EXISTS idx_apikey_user_id;
CREATE INDEX idx_apikey_user_id ON authentication.user_api_key(user_id);


-- Initial inserts
-- roles
INSERT INTO authentication.roles (code, role_name, unique_id) VALUES ('SUPER_ADMIN', 'Super Administrator', '0c4e2b1f-59ba-49b6-ba4f-81622f33732d');
INSERT INTO authentication.roles (code, role_name, unique_id) VALUES ('ADMIN', 'Administrator', '0c8e2b1f-87ba-49b6-ba4f-81634f33732d');
INSERT INTO authentication.roles (code, role_name, unique_id) VALUES ('EMPLOYEE', 'Employee', '0c8e2b1f-39ba-49b6-ba4f-81622f33732d');

-- test users
-- Inserting the users into the 'authentication.users' table
INSERT INTO authentication.users (user_name, email, password, user_image, is_active, token, unique_id, created_at, updated_at) VALUES
('admin01', 'admin01@example.com', '$2a$10$WK73KU34gno.h1TqJFLrmux5uVIrNwS5TfgKxLcKxeSO15DP.McwO', NULL, TRUE, NULL, '0c8e2b1f-89ba-49b6-ba4f-81622f33732d', NOW(), NOW()),
('admin02', 'admin02@example.com', '$2a$10$Rb44LaGqdM9R4Lx3zg59Z.bZGAlP05OGU5cR9Vni7W35EksJOuW/a', NULL, TRUE, NULL, '30823080-e83a-462d-ba41-88daff6e016d', NOW(), NOW()),
('employee01', 'employee01@example.com', '$2a$10$AlQU9C64eQgiXGTcn2/gLuszJWfw31VkPkP4TI6OpgKjmzST6h1/a', NULL, TRUE, NULL, '1fbe2e02-8f87-4312-9059-1d14f3cef623', NOW(), NOW());
