
DROP TYPE IF EXISTS TYPE_USER_STATUS;
CREATE TYPE TYPE_USER_STATUS AS ENUM('Yes', 'No');


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
    role_id INT NOT NULL,
    user_name VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(150) NOT NULL,
    user_image VARCHAR(100), 
    active TYPE_USER_STATUS NOT NULL DEFAULT 'Yes',
    token VARCHAR(150) UNIQUE NOT NULL,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_user_role FOREIGN KEY(role_id) REFERENCES  authentication.roles(role_id)
);


DROP TABLE IF EXISTS authentication.permissions;
CREATE TABLE authentication.permissions (
    permission_id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE,
    permission_name VARCHAR(100) NOT NULL,
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

-- Initial inserts
-- roles
INSERT INTO authentication.roles (code, role_name) VALUES ('super-admin', 'Super Administrator');
INSERT INTO authentication.roles (code, role_name) VALUES ('administrator', 'Administrator');
INSERT INTO authentication.roles (code, role_name) VALUES ('employee', 'Employee');

-- permissions
INSERT INTO authentication.permissions (code, permission_name) VALUES ('change-configuration', 'Change Configuration');
INSERT INTO authentication.permissions (code, permission_name) VALUES ('list-configuration', 'List Configurations');
INSERT INTO authentication.permissions (code, permission_name) VALUES ('lock-application', 'Lock Application');
INSERT INTO authentication.permissions (code, permission_name) VALUES ('create-role', 'Create Roles');
INSERT INTO authentication.permissions (code, permission_name) VALUES ('update-role', 'Update Roles');
INSERT INTO authentication.permissions (code, permission_name) VALUES ('assign-role', 'Assign Roles');
INSERT INTO authentication.permissions (code, permission_name) VALUES ('list-roles', 'List Roles');
INSERT INTO authentication.permissions (code, permission_name) VALUES ('create-user', 'Create User');
INSERT INTO authentication.permissions (code, permission_name) VALUES ('delete-user', 'Delete User');
INSERT INTO authentication.permissions (code, permission_name) VALUES ('update-user', 'Update User');
INSERT INTO authentication.permissions (code, permission_name) VALUES ('list-users', 'List Users');


