-- Schema: authentication
DROP TABLE IF EXISTS authentication.roles;
CREATE TABLE authentication.roles (
    role_id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    role_name VARCHAR(100) NOT NULL,
    description VARCHAR(200),
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

DROP TABLE IF EXISTS authentication.users;
CREATE TABLE authentication.users (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    user_image VARCHAR(100), 
    is_active BOOLEAN DEFAULT TRUE,
    is_logged BOOLEAN,
    token VARCHAR(150) UNIQUE,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

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

DROP TABLE IF EXISTS authentication.permissions;
CREATE TABLE authentication.permissions (
    permission_id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    permission_name VARCHAR(100) NOT NULL,
    description VARCHAR(200),
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

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


DROP TYPE IF EXISTS TYPE_ACTIVITY_STATUS;
CREATE TYPE TYPE_ACTIVITY_STATUS AS ENUM('Online', 'Offline');

DROP TABLE IF EXISTS authentication.login_activity;
CREATE TABLE authentication.login_activity (
    login_id SERIAL PRIMARY KEY,
    user_id INT UNIQUE NOT NULL,
    status TYPE_ACTIVITY_STATUS DEFAULT 'Offline',
    host VARCHAR(150),
    browser VARCHAR(150),
    ip_address VARCHAR(50),
    device VARCHAR(150),
    location VARCHAR(150),
    last_login TIMESTAMP DEFAULT NOW(),
    last_logout TIMESTAMP DEFAULT NOW(),
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES  authentication.users(user_id)
);

-- Initial inserts
-- roles
INSERT INTO authentication.roles (code, role_name) VALUES ('super-admin', 'Super Administrator');
INSERT INTO authentication.roles (code, role_name) VALUES ('admin', 'Administrator');
INSERT INTO authentication.roles (code, role_name) VALUES ('employee', 'Employee');
-- Default user: used to manage application
INSERT INTO authentication.users (user_name, email,  password) VALUES ('admin@user.com', 'admin@user.com', '$2a$10$9VE1S3YfjRPA5Hu7ZAV.ROy9M8aQsEAy0t2AgrCnzoDpEqhbunspq');