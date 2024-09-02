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
(1, 'Authentication', 'authentication', 'Handles user authentication and authorization.'),
(2, 'Configurations', 'configurations', 'Manages system configurations.'),
(3, 'References', 'references', 'Stores and manages reference data.'),
(4, 'Company', 'company', 'Manages company-related information and settings.'),
(5, 'Employees', 'employees', 'Manages employee records and details.'),
(6, 'Reports', 'reports', 'Generates and manages various reports.');



-- FEATURE FLAGS
DROP TYPE IF EXISTS configurations.TYPE_FLAG_STATUS;
CREATE TYPE configurations.TYPE_FLAG_STATUS AS ENUM ('Enabled', 'Disabled');

-- Table: module_flag
DROP TABLE IF EXISTS configurations.module_flag;
CREATE TABLE configurations.module_flag(
    flag_id SERIAL PRIMARY KEY,
    module_id INT NOT NULL,
    status coonfigurations.TYPE_FLAG_STATUS DEFAULT 'Disabled',
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_flag_module FOREIGN KEY(module_id) REFERENCES configurations.modules(module_id)
);
-- Insert module flags
INSERT INTO configurations.module_flag (flag_id, module_id, status) VALUES 
(1, 1, 'Enabled'),  -- Authentication
(2, 2, 'Enabled'),  -- Configurations
(3, 3, 'Disabled'),  -- References
(4, 4, 'Disabled'),  -- Company
(5, 5, 'Disabled'),  -- Employees
(6, 6, 'Disabled'); -- Reports


-- Table: core_entities
DROP TABLE IF EXISTS configurations.core_entities;
CREATE TABLE configurations.core_entities (
    entity_id SERIAL PRIMARY KEY,
    module_id INT NOT NULL,
    entity_name VARCHAR(100) UNIQUE NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL,
    description TEXT,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_module_entity FOREIGN KEY(module_id) REFERENCES configurations.modules(module_id)
);
-- Insert entities into core_entities with module_code.entity_code format
-- Module: Authentication
INSERT INTO configurations.core_entities (entity_id, module_id, entity_name, code, description) VALUES
-- Module: Authentication
(1, 1, 'Users', 'authentication.users', 'Enabled or disabled'),
(2, 1, 'Active Users', 'authentication.active_users', 'Users currently active'),
(3, 1, 'Inactive Users', 'authentication.inactive_users', 'Users currently inactive'),
(4, 1, 'Online Users', 'authentication.online_users', 'Users currently online'),
(5, 1, 'Offline Users', 'authentication.offline_users', 'Users currently offline'),
(6, 1, 'Roles', 'authentication.roles', 'User roles in the system'),
(7, 1, 'Permissions', 'authentication.permissions', 'User permissions in the system'),
(8, 1, 'Login Activity', 'authentication.login_activity', 'User login activity'),
-- Module: Configurations
(9, 2, 'Basic Configurations', 'configurations.basic_configurations', 'Basic system configurations'),
(10, 2, 'Company Configurations', 'configurations.company_configurations', 'Company-related configurations'),
(11, 2, 'Email Configurations', 'configurations.email_configurations', 'Email system configurations'),
(12, 2, 'Modules', 'configurations.modules', 'System modules'),
(13, 2, 'Core Entities', 'configurations.core_entities', 'Core entities used in the system'),
(14, 2, 'Module Flags', 'configurations.module_flags', 'Flags for various modules'),
(15, 2, 'Core Entity Flags', 'configurations.core_entity_flags', 'Flags for core entities'),
-- Module: References
(16, 3, 'Countries', 'references.countries', 'List of countries'),
(17, 3, 'Currencies', 'references.currencies', 'List of currencies'),
(18, 3, 'Identification Types', 'references.identification_types', 'Types of identification'),
(19, 3, 'Contact Types', 'references.contact_types', 'Types of contact information'),
(20, 3, 'Marital Statuses', 'references.marital_statuses', 'Marital status options'),
(21, 3, 'Task Statuses', 'references.task_statuses', 'Statuses of tasks'),
(22, 3, 'Approval Statuses', 'references.approval_statuses', 'Statuses of document approvals'),
(23, 3, 'Document Statuses', 'references.document_statuses', 'Statuses of documents'),
(24, 3, 'Workflow Statuses', 'references.workflow_statuses', 'Statuses in workflows'),
(25, 3, 'Evaluation Statuses', 'references.evaluation_statuses', 'Statuses of evaluations'),
(26, 3, 'User Statuses', 'references.user_statuses', 'Statuses of users'),
(27, 3, 'Employment Statuses', 'references.employment_statuses', 'Statuses of employment'),
-- Module: Company
(28, 4, 'Company Info', 'company.company_info', 'Information about the company'),
(29, 4, 'Branches', 'company.branches', 'Company branches'),
(30, 4, 'Offices', 'company.offices', 'Company offices'),
(31, 4, 'Departments', 'company.departments', 'Company departments'),
(32, 4, 'Rooms', 'company.rooms', 'Company rooms'),
(33, 4, 'Projects', 'company.projects', 'Company projects'),
(34, 4, 'Policies', 'company.policies', 'Company policies'),
-- Module: Employees
(35, 5, 'Employees', 'employees.employees', 'Employee records'),
(36, 5, 'Job Titles', 'employees.job_titles', 'Job titles within the company'),
-- Module: Reports
(37, 6, 'User Reports', 'reports.user_reports', 'Reports related to users'),
(38, 6, 'Configuration Reports', 'reports.configuration_reports', 'Reports related to system configurations'),
(39, 6, 'Company Reports', 'reports.company_reports', 'Reports related to company information'),
(40, 6, 'Employee Reports', 'reports.employee_reports', 'Reports related to employees'),
(41, 6, 'Reference Reports', 'reports.reference_reports', 'Reports related to reference data');




-- Table: core_entity_flag
DROP TABLE IF EXISTS configurations.core_entity_flag;
CREATE TABLE configurations.core_entity_flag(
    flag_id SERIAL PRIMARY KEY,
    entity_id INT NOT NULL,
    module_id INT NOT NULL,
    status TYPE_FLAG_STATUS DEFAULT 'Disabled',
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_flag_entity FOREIGN KEY(entity_id) REFERENCES configurations.core_entities(entity_id),
    CONSTRAINT fk_module_flag_entity FOREIGN KEY(module_id) REFERENCES configurations.modules(module_id)
);
-- Module: Authentication
INSERT INTO configurations.core_entity_flag (flag_id, entity_id, module_id, status) VALUES
(1, 1, 1, 'Disabled'), -- Users
(2, 2, 1, 'Disabled'), -- ActiveUsers
(3, 3, 1, 'Disabled'), -- InactiveUsers
(4, 4, 1, 'Disabled'), -- OnlineUsers
(5, 5, 1, 'Disabled'), -- OfflineUsers
(6, 6, 1, 'Disabled'), -- Roles
(7, 7, 1, 'Disabled'), -- Permissions
(8, 8, 1, 'Disabled'); -- LoginActivity
-- Module: Configurations
INSERT INTO configurations.core_entity_flag (flag_id, entity_id, module_id, status) VALUES
(9, 9, 2, 'Enabled'), -- BasicConfigurations
(10, 10, 2, 'Enabled'), -- CompanyConfigurations
(11, 11, 2, 'Enabled'), -- EmailConfigurations
(12, 12, 2, 'Enabled'), -- Modules
(13, 13, 2, 'Enabled'), -- CoreEntities
(14, 14, 2, 'Enabled'), -- ModuleFlags
(15, 15, 2, 'Enabled'); -- CoreEntityFlags
-- Module: References
INSERT INTO configurations.core_entity_flag (flag_id, entity_id, module_id, status) VALUES
(16, 16, 3, 'Disabled'), -- Countries
(17, 17, 3, 'Disabled'), -- Currencies
(18, 18, 3, 'Disabled'), -- IdentificationTypes
(19, 19, 3, 'Disabled'), -- ContactTypes
(20, 20, 3, 'Disabled'), -- MaritalStatuses
(21, 21, 3, 'Disabled'), -- TaskStatuses
(22, 22, 3, 'Disabled'), -- ApprovalStatuses
(23, 23, 3, 'Disabled'), -- DocumentStatuses
(24, 24, 3, 'Disabled'), -- WorkflowStatuses
(25, 25, 3, 'Disabled'), -- EvaluationStatuses
(26, 26, 3, 'Disabled'), -- UserStatuses
(27, 27, 3, 'Disabled'); -- EmploymentStatuses
-- Module: Company
INSERT INTO configurations.core_entity_flag (flag_id, entity_id, module_id, status) VALUES
(28, 28, 4, 'Disabled'), -- CompanyInfo
(29, 29, 4, 'Disabled'), -- Branches
(30, 30, 4, 'Disabled'), -- Offices
(31, 31, 4, 'Disabled'), -- Departments
(32, 32, 4, 'Disabled'), -- Rooms
(33, 33, 4, 'Disabled'), -- Projects
(34, 34, 4, 'Disabled'); -- Policies
-- Module: Employees
INSERT INTO configurations.core_entity_flag (flag_id, entity_id, module_id, status) VALUES
(35, 35, 5, 'Disabled'), -- Employees
(36, 36, 5, 'Disabled'); -- JobTitles
-- Module: Reports
INSERT INTO configurations.core_entity_flag (flag_id, entity_id, module_id, status) VALUES
(37, 37, 6, 'Disabled'), -- UserReports
(38, 38, 6, 'Disabled'), -- ConfigurationReports
(39, 39, 6, 'Disabled'), -- CompanyReports
(40, 40, 6, 'Disabled'), -- EmployeeReports
(41, 41, 6, 'Disabled'); -- ReferenceReports

