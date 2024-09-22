-- tables for 'authentication' schema 

DROP TYPE IF EXISTS authentication.TYPE_ROLE_STATUS;
CREATE TYPE authentication.TYPE_ROLE_STATUS AS ENUM('Enabled', 'Disabled');

-- Table: roles
DROP TABLE IF EXISTS authentication.roles;
CREATE TABLE authentication.roles (
    role_id SERIAL PRIMARY KEY,
    code VARCHAR(30) UNIQUE NOT NULL,
    role_name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    status authentication.TYPE_ROLE_STATUS DEFAULT 'Enabled',
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
-- Inserts
INSERT INTO authentication.roles (role_id, code, role_name, description, status, unique_id) VALUES 
(1, 'role_super_admin', 'Super Administrator', 'Has full access to all system features and settings.', 'Enabled', '0c4e2b1f-59ba-49b6-ba4f-81622f33732d'),
(2, 'role_admin', 'Administrator', 'Manages users, permissions, and overall system configuration.', 'Enabled', '0c8e2b1f-87ba-49b6-ba4f-81634f33732d'),
(3, 'role_manager', 'Manager', 'Oversees specific departments or projects with restricted admin capabilities.', 'Enabled', '0c8e2b1f-39ba-49b6-ba4f-81622f33732d'),
(4, 'role_employee', 'Employee', 'Accesses daily tasks and data relevant to their role.', 'Enabled', '0c8e2b1f-99ba-49b6-ba4f-81622f33732d'),
(5, 'role_customer', 'Customer', 'Interacts with the system for purchasing and account management.', 'Disabled', '0c8e2b1f-11ba-49b6-ba4f-81622f33732d'),
(6, 'role_supplier', 'Supplier', 'Manages supply-related transactions and information.', 'Disabled', '0c4e2b1f-59ba-49b6-ba4f-81622f33733e'),
(7, 'role_support', 'Support', 'Provides assistance and resolves issues for other users.', 'Disabled', '0c8e2b1f-87ba-49b6-ba4f-81634f33733e'),
(8, 'role_developer', 'Developer', 'Works on system development and maintenance tasks.', 'Enabled', '0c8e2b1f-39ba-49b6-ba4f-81622f33733e'),
(9, 'role_guest', 'Guest', 'Has limited access to view non-sensitive parts of the system.', 'Disabled', '0c8e2b1f-11ba-49b6-ba4f-81622f33733e');



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
-- inserts
INSERT INTO authentication.users (user_name, email, password, user_image, is_active, token, unique_id, created_at, updated_at) VALUES
('admin01', 'admin01@example.com', '$2a$10$WK73KU34gno.h1TqJFLrmux5uVIrNwS5TfgKxLcKxeSO15DP.McwO', NULL, TRUE, NULL, '0c8e2b1f-89ba-49b6-ba4f-81622f33732d', NOW(), NOW()),
('admin02', 'admin02@example.com', '$2a$10$Rb44LaGqdM9R4Lx3zg59Z.bZGAlP05OGU5cR9Vni7W35EksJOuW/a', NULL, TRUE, NULL, '30823080-e83a-462d-ba41-88daff6e016d', NOW(), NOW()),
('employee01', 'employee01@example.com', '$2a$10$AlQU9C64eQgiXGTcn2/gLuszJWfw31VkPkP4TI6OpgKjmzST6h1/a', NULL, TRUE, NULL, '1fbe2e02-8f87-4312-9059-1d14f3cef623', NOW(), NOW());



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


DROP TABLE IF EXISTS authentication.user_associations;
CREATE TABLE authentication.user_associations (
    association_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES  authentication.users(user_id)
);


-- Table: permissions
DROP TABLE IF EXISTS authentication.permissions;
CREATE TABLE authentication.permissions (
    permission_id SERIAL PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    permission_name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
-- I

-- Inserts
INSERT INTO authentication.permissions (code, permission_name, description) VALUES
('authentication.users.index', 'List Users', 'Handles the index route for listing users.'),
('authentication.users.details', 'View User Details', 'Handles the route for displaying user details.'),
('authentication.users.create', 'Create User', 'Handles the creation of a new user.'),
('authentication.users.edit', 'Edit User', 'Handles the editing of an existing user.'),
('authentication.users.assignRole', 'Assign Role', 'Handles assigning a role to a user.'),
('authentication.users.removeRole', 'Remove Role', 'Handles removing a role from a user.'),
('authentication.users.search', 'Search Users', 'Handles searching for users.'),
('authentication.users.deactivate', 'Deactivate User', 'Handles the deactivation of a user.'),
('authentication.users.activate', 'Activate User', 'Handles the activation of a user.'),
('authentication.users.getAllActiveUsers', 'Get All Active Users', 'Handles the retrieval of all active users.'),
('authentication.users.getAllInactiveUsers', 'Get All Inactive Users', 'Handles the retrieval of all inactive users.'),
('authentication.users.getAllOnlineUsers', 'Get All Online Users', 'Handles the retrieval of all online users.'),
('authentication.users.getAllOfflineUsers', 'Get All Offline Users', 'Handles the retrieval of all offline users.'),
('authentication.roles.index', 'List Roles', 'Access the index view for roles'),
('authentication.roles.details', 'View Role Details', 'View details of a specific role'),
('authentication.roles.create', 'Create Role', 'Create a new role'),
('authentication.roles.edit', 'Edit Role', 'Edit an existing role'),
('authentication.roles.delete', 'Delete Role', 'Delete an existing role'),
('authentication.roles.search', 'Search Roles', 'Search for roles'),
('authentication.roles.assignPermission', 'Assign Permission', 'Assign a permission to a role'),
('authentication.roles.removePermission', 'Remove Permission', 'Remove a permission from a role'),
('authentication.permissions.index', 'List Permissions', 'Access the index view for permissions'),
('authentication.permissions.details', 'View Permission Details', 'View details of a specific permission'),
('authentication.permissions.create', 'Create Permission', 'Create a new permission'),
('authentication.permissions.edit', 'Edit Permission', 'Edit an existing permission'),
('authentication.permissions.delete', 'Delete Permission', 'Delete an existing permission'),
('authentication.permissions.search', 'Search Permissions', 'Search for permissions'),
('authentication.login-activities.index', 'List Login Activities', 'Access the index view for login activities'),
('authentication.login-activities.details', 'View Login Activity Details', 'View details of a specific login activity'),
('authentication.login-activities.search', 'Search Login Activities', 'Search for login activities'),
('account.uploadUserImage', 'Upload User Image', 'Upload a new image for the user'),
('account.changePassword', 'Change Password', 'Change the user password'),
('account.userData', 'User Data', 'View user data'),
('configurations.basic-configurations.index', 'View Basic Configurations', 'Access the index view for basic configurations'),
('configurations.basic-configurations.edit', 'Edit Basic Configuration', 'Update basic configuration settings'),
('configurations.company-configurations.index', 'View Company Configurations', 'Access the index view for company configurations'),
('configurations.company-configurations.edit', 'Edit Company Configurations', 'Update company configuration settings'),
('configurations.email-configurations.index', 'View Email Configurations', 'Access the index view for email configurations'),
('configurations.email-configurations.edit', 'Edit Email Configurations', 'Access the form to edit email configurations'),
('configurations.email-configurations.update', 'Update Email Configurations', 'Update email configuration settings'),
('configurations.module-flags.index', 'View Module Flags', 'Access the index view for module flags'),
('configurations.module-flags.manage', 'Manage Module Flags', 'Access the form to manage module flags'),
('configurations.module-flags.update', 'Update Module Flags', 'Update module flag settings'),
('company.branches.index', 'View Branches', 'Access the index view for branches'),
('company.branches.details', 'View Branch Details', 'Access the details view for a branch'),
('company.branches.create', 'Create Branch', 'Create a new branch'),
('company.branches.edit', 'Edit Branch', 'Edit an existing branch'),
('company.branches.search', 'Search Branches', 'Search for branches'),
('company.departments.index', 'View Departments', 'Access the index view for departments'),
('company.departments.details', 'View Department Details', 'Access the details view for a department'),
('company.departments.create', 'Create Department', 'Create a new department'),
('company.departments.edit', 'Edit Department', 'Edit an existing department'),
('company.departments.search', 'Search Departments', 'Search for departments'),
('company.offices.index', 'View Offices', 'Access the index view for offices'),
('company.offices.details', 'View Office Details', 'Access the details view for an office'),
('company.offices.create', 'Create Office', 'Create a new office'),
('company.offices.edit', 'Edit Office', 'Edit an existing office'),
('company.offices.search', 'Search Offices', 'Search for offices'),
('company.policies.index', 'View Policies', 'Access the index view for policies'),
('company.policies.details', 'View Policy Details', 'Access the details view for a policy'),
('company.policies.create', 'Create Policy', 'Create a new policy'),
('company.policies.edit', 'Edit Policy', 'Edit an existing policy'),
('company.policies.search', 'Search Policies', 'Search for policies'),
('company.policies.remove', 'Remove Policy', 'Remove an existing policy'),
('company.projects.index', 'View Projects', 'Access the index view for projects'),
('company.projects.details', 'View Project Details', 'Access the details view for a project'),
('company.projects.create', 'Create Project', 'Create a new project'),
('company.projects.edit', 'Edit Project', 'Edit an existing project'),
('company.projects.search', 'Search Projects', 'Search for projects'),
('company.projects.remove', 'Remove Project', 'Remove an existing project'),
('company.rooms.index', 'View Rooms', 'Access the index view for rooms'),
('company.rooms.details', 'View Room Details', 'Access the details view for a room'),
('company.rooms.create', 'Create Room', 'Create a new room'),
('company.rooms.edit', 'Edit Room', 'Edit an existing room'),
('company.rooms.search', 'Search Rooms', 'Search for rooms'),
('company.company-info.index', 'View Company Info', 'Access the index view for company information'),
('company.company-info.details', 'View Company Details', 'Access the details view for a company'),
('company.company-info.create', 'Create Company Info', 'Create new company information'),
('company.company-info.edit', 'Edit Company Info', 'Edit existing company information'),
('references.countries.index', 'View Countries', 'Access the index view for countries'),
('references.countries.details', 'View Country Details', 'Access the details view for a country'),
('references.countries.create', 'Create Country', 'Create new country information'),
('references.countries.edit', 'Edit Country', 'Edit existing country information'),
('references.countries.search', 'Search Countries', 'Search for countries'),
('references.countries.search-results', 'View Country Search Results', 'View the results of country searches'),
('references.countries.delete', 'Delete Country', 'Remove a country'),
('references.currencies.index', 'View Currencies', 'Access the index view for currencies'),
('references.currencies.details', 'View Currency Details', 'Access the details view for a currency'),
('references.currencies.create', 'Create Currency', 'Create new currency information'),
('references.currencies.edit', 'Edit Currency', 'Edit existing currency information'),
('references.currencies.search', 'Search Currencies', 'Search for currencies'),
('references.currencies.search-results', 'View Currency Search Results', 'View the results of currency searches'),
('references.currencies.delete', 'Delete Currency', 'Remove a currency'),
('references.approval-statuses.index', 'View Approval Statuses', 'Access the index view for approval statuses'),
('references.approval-statuses.details', 'View Approval Status Details', 'Access the details view for an approval status'),
('references.approval-statuses.create', 'Create Approval Status', 'Create new approval status information'),
('references.approval-statuses.edit', 'Edit Approval Status', 'Edit existing approval status information'),
('references.approval-statuses.search', 'Search Approval Statuses', 'Search for approval statuses'),
('references.approval-statuses.search-results', 'View Approval Status Search Results', 'View the results of approval status searches'),
('references.approval-statuses.delete', 'Delete Approval Status', 'Remove an approval status'),
('references.contact-types.index', 'View Contact Types', 'Access the index view for contact types'),
('references.contact-types.details', 'View Contact Type Details', 'Access the details view for a contact type'),
('references.contact-types.create', 'Create Contact Type', 'Create new contact type information'),
('references.contact-types.edit', 'Edit Contact Type', 'Edit existing contact type information'),
('references.contact-types.search', 'Search Contact Types', 'Search for contact types'),
('references.contact-types.search-results', 'View Contact Type Search Results', 'View the results of contact type searches'),
('references.contact-types.delete', 'Delete Contact Type', 'Remove a contact type'),
('references.document-statuses.index', 'View Document Statuses', 'Access the index view for document statuses'),
('references.document-statuses.details', 'View Document Status Details', 'Access the details view for a document status'),
('references.document-statuses.create', 'Create Document Status', 'Create new document status information'),
('references.document-statuses.edit', 'Edit Document Status', 'Edit existing document status information'),
('references.document-statuses.search', 'Search Document Statuses', 'Search for document statuses'),
('references.document-statuses.search-results', 'View Document Status Search Results', 'View the results of document status searches'),
('references.document-statuses.delete', 'Delete Document Status', 'Remove a document status'),
('references.evaluation-statuses.index', 'View Evaluation Statuses', 'Access the index view for evaluation statuses'),
('references.evaluation-statuses.details', 'View Evaluation Status Details', 'Access the details view for an evaluation status'),
('references.evaluation-statuses.create', 'Create Evaluation Status', 'Create new evaluation status information'),
('references.evaluation-statuses.edit', 'Edit Evaluation Status', 'Edit existing evaluation status information'),
('references.evaluation-statuses.search', 'Search Evaluation Statuses', 'Search for evaluation statuses'),
('references.evaluation-statuses.search-results', 'View Evaluation Status Search Results', 'View the results of evaluation status searches'),
('references.evaluation-statuses.delete', 'Delete Evaluation Status', 'Remove an evaluation status'),
('references.identification-types.index', 'View Identification Types', 'Access the index view for identification types'),
('references.identification-types.details', 'View Identification Type Details', 'Access the details view for an identification type'),
('references.identification-types.create', 'Create Identification Type', 'Create new identification type information'),
('references.identification-types.edit', 'Edit Identification Type', 'Edit existing identification type information'),
('references.identification-types.search', 'Search Identification Types', 'Search for identification types'),
('references.identification-types.search-results', 'View Identification Type Search Results', 'View the results of identification type searches'),
('references.identification-types.delete', 'Delete Identification Type', 'Remove an identification type'),
('references.marital-statuses.index', 'View Marital Statuses', 'Access the index view for marital statuses'),
('references.marital-statuses.details', 'View Marital Status Details', 'Access the details view for a marital status'),
('references.marital-statuses.create', 'Create Marital Status', 'Create new marital status information'),
('references.marital-statuses.edit', 'Edit Marital Status', 'Edit existing marital status information'),
('references.marital-statuses.search', 'Search Marital Statuses', 'Search for marital statuses'),
('references.marital-statuses.search-results', 'View Marital Status Search Results', 'View the results of marital status searches'),
('references.marital-statuses.delete', 'Delete Marital Status', 'Remove a marital status'),
('references.task-statuses.index', 'View Task Statuses', 'Access the index view for task statuses'),
('references.task-statuses.details', 'View Task Status Details', 'Access the details view for a task status'),
('references.task-statuses.create', 'Create Task Status', 'Create new task status information'),
('references.task-statuses.edit', 'Edit Task Status', 'Edit existing task status information'),
('references.task-statuses.search', 'Search Task Statuses', 'Search for task statuses'),
('references.task-statuses.search-results', 'View Task Status Search Results', 'View the results of task status searches'),
('references.task-statuses.delete', 'Delete Task Status', 'Remove a task status'),
('references.user-statuses.index', 'View User Statuses', 'Access the index view for user statuses'),
('references.user-statuses.details', 'View User Status Details', 'Access the details view for a user status'),
('references.user-statuses.create', 'Create User Status', 'Create new user status information'),
('references.user-statuses.edit', 'Edit User Status', 'Edit existing user status information'),
('references.user-statuses.search', 'Search User Statuses', 'Search for user statuses'),
('references.user-statuses.search-results', 'View User Status Search Results', 'View the results of user status searches'),
('references.user-statuses.delete', 'Delete User Status', 'Remove a user status'),
('references.workflow-statuses.index', 'View Workflow Statuses', 'Access the index view for workflow statuses'),
('references.workflow-statuses.details', 'View Workflow Status Details', 'Access the details view for a workflow status'),
('references.workflow-statuses.create', 'Create Workflow Status', 'Create new workflow status information'),
('references.workflow-statuses.edit', 'Edit Workflow Status', 'Edit existing workflow status information'),
('references.workflow-statuses.search', 'Search Workflow Statuses', 'Search for workflow statuses'),
('references.workflow-statuses.search-results', 'View Workflow Status Search Results', 'View the results of workflow status searches'),
('references.workflow-statuses.delete', 'Delete Workflow Status', 'Remove a workflow status');


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
    unique_id VARCHAR(50),
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
    x_user_id VARCHAR(50) UNIQUE,
    x_api_key VARCHAR(150) UNIQUE,
    is_active BOOLEAN DEFAULT TRUE,
    created_by INT,
    expires_at TIMESTAMP NOT NULL,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_user_key FOREIGN KEY(user_id) REFERENCES  authentication.users(user_id)
);

