-- tables for 'employees' schema 

-- TYPES
DROP TYPE IF EXISTS TYPE_GENDER;
CREATE TYPE TYPE_GENDER AS ENUM('Male', 'Female');

-- Table: job_titles
DROP TABLE IF EXISTS employees.job_titles;
CREATE TABLE employees.job_titles(
    job_title_id SERIAL PRIMARY KEY,
    title_name VARCHAR(150) NOT NULL,
    description TEXT,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
-- Example inserts
INSERT INTO employees.job_titles (title_name, description) VALUES
('Human Resources Manager', 'Oversees HR functions including recruitment, employee relations, and compliance.'),
('Recruitment Specialist', 'Handles the recruitment process, from sourcing candidates to conducting interviews.'),
('HR Business Partner', 'Works closely with business units to align HR strategies with business goals.'),
('Payroll Specialist', 'Manages payroll processing, tax calculations, and compliance with wage laws.'),
('HR Coordinator', 'Supports HR operations by managing employee records, onboarding, and administrative tasks.'),
('Benefits Administrator', 'Manages employee benefits programs, including health insurance and retirement plans.'),
('Training and Development Manager', 'Develops and implements training programs to enhance employee skills and performance.'),
('Employee Relations Specialist', 'Addresses employee concerns and mediates conflicts to ensure a positive workplace environment.'),
('Compensation Analyst', 'Analyzes and develops compensation structures and salary benchmarking.'),
('HRIS Analyst', 'Maintains and optimizes the Human Resources Information System (HRIS) for data management and reporting.');


-- Table: employees
DROP TABLE IF EXISTS employees.employees;
CREATE TABLE employees.employees(
    employee_id SERIAL PRIMARY KEY,
    user_id INT,
    identification_type_id INT NOT NULL,
    country_id INT,
    marital_status_id INT,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    identification_number VARCHAR(30) UNIQUE NOT NULL,
    gender TYPE_GENDER,
    date_of_birth DATE,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_identification FOREIGN KEY(identification_type_id) REFERENCES reference.identification_types(type_id),
    CONSTRAINT fk_country FOREIGN KEY(country_id) REFERENCES reference.countries(country_id),
    CONSTRAINT fk_marital_status FOREIGN KEY(marital_status_id) REFERENCES reference.marital_statuses(status_id)
);
-- Indexes:
DROP INDEX IF EXISTS idx_employees_user;
CREATE INDEX idx_employees_user ON employees.employees(user_id);
DROP INDEX IF EXISTS idx_employees_first_name;
CREATE INDEX  IF EXISTS idx_employees_first_name ON employees.employees(first_name);
DROP INDEX IF EXISTS idx_employees_last_name;
CREATE INDEX idx_employees_last_name ON employees.employees(last_name);
DROP INDEX IF EXISTS idx_employees_country;
CREATE INDEX idx_employees_country ON employees.employees(country_id);


-- Table: professional_info
DROP TABLE IF EXISTS employees.professional_info;
CREATE TABLE employees.professional_info(
    professional_id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL,
    department_id INT,
    job_title_id INT,
    employment_status_id INT NOT NULL,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_professional_employee FOREIGN KEY (employee_id) REFERENCES employees.employees(employee_id),
    CONSTRAINT fk_department FOREIGN KEY(department_id) REFERENCES company.departments(department_id),
    CONSTRAINT fk_job_title FOREIGN KEY(job_title_id) REFERENCES employees.job_titles(job_title_id),
    CONSTRAINT fk_employment_status FOREIGN KEY(employment_status_id) REFERENCES reference.employment_statuses(status_id)
);


-- Table: document_types
DROP TABLE IF EXISTS employees.document_types;
CREATE TABLE employees.document_types(
    type_id SERIAL PRIMARY KEY,
    type_name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    unique_id VARCHAR(50) UNIQUE DEFAULT uuid_generate_v4()::text,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
-- Insert
INSERT INTO employees.document_types (type_name, description) VALUES
('Identification', 'Documents used for personal identification'),
('Employment', 'Documents related to employment'),
('Financial', 'Documents related to financial information'),
('Compliance', 'Documents related to compliance and legal requirements'),
('Training/Certifications', 'Documents related to training and certifications'),
('Health', 'Documents related to employee health and safety');



DROP TYPE IF EXISTS TYPE_DOCUMENT_STATUS;
CREATE TYPE TYPE_DOCUMENT_STATUS AS ENUM('Expired', 'Active');

-- Table: documents
DROP TABLE IF EXISTS employees.documents;
CREATE TABLE employees.documents(
    document_id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL,
    document_type_id INT NOT NULL,
    document_name VARCHAR(200),
    document_number VARCHAR(40),
    expiration_date DATE,
    file_name VARCHAR(150),
    status TYPE_DOCUMENT_STATUS DEFAULT 'Active',
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_document_employee FOREIGN KEY (employee_id) REFERENCES employees.employees(employee_id),
    CONSTRAINT fk_document_type FOREIGN KEY (document_type_id) REFERENCES employees.document_types(type_id)
);
-- Indexes
DROP INDEX IF EXISTS idx_employee_id;
CREATE INDEX idx_employee_id ON employees.documents(employee_id);
DROP INDEX IF EXISTS idx_document_type_id;
CREATE INDEX idx_document_type_id ON employees.documents(document_type_id);


-- Table: employee_phones
DROP TABLE IF EXISTS employees.employee_phones;
CREATE TABLE employees.employee_phones(
    phone_id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL,
    contact_type_id INT NOT NULL,
    phone_number VARCHAR(30) UNIQUE NOT NULL,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_contact_employee FOREIGN KEY (employee_id) REFERENCES employees.employees(employee_id),
    CONSTRAINT fk_contact_type FOREIGN KEY (contact_type_id) REFERENCES reference.contact_types(type_id)
);

-- Table: employee_emails
DROP TABLE IF EXISTS employees.employee_emails;
CREATE TABLE employees.employee_emails(
    email_id SERIAL PRIMARY KEY,
    employee_id INT NOT NULL,
    contact_type_id INT NOT NULL,
    email_address VARCHAR(150) UNIQUE NOT NULL,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_contact_employee FOREIGN KEY (employee_id) REFERENCES employees.employees(employee_id),
    CONSTRAINT fk_contact_type FOREIGN KEY (contact_type_id) REFERENCES reference.contact_types(type_id)
);

