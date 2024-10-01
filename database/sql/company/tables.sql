-- Table: companies
DROP TABLE IF EXISTS company.companies;
CREATE TABLE company.companies (
    company_id SERIAL PRIMARY KEY,
    company_name VARCHAR(100) NOT NULL,
    company_acronym VARCHAR(20),
    company_type VARCHAR(50),
    industry VARCHAR(50),
    founded_date DATE,
    address TEXT,
    phone VARCHAR(20),
    email VARCHAR(100),
    website_url VARCHAR(100),
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
CREATE INDEX idx_company_name ON company.companies(company_name);


-- Table: branches
DROP TABLE IF EXISTS company.branches;
CREATE TABLE company.branches (
    branch_id SERIAL PRIMARY KEY,
    company_id INT NOT NULL,
    branch_name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE,
    address TEXT,
    phone VARCHAR(20),
    email VARCHAR(100),
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_company_branch FOREIGN KEY (company_id) REFERENCES company.companies(company_id)
);
CREATE INDEX idx_branch_name ON company.branches(branch_name);


-- Table: offices
DROP TABLE IF EXISTS company.offices;
CREATE TABLE company.offices (
    office_id SERIAL PRIMARY KEY,
    company_id INT NOT NULL,
    office_name VARCHAR(100) NOT NULL,
    code VARCHAR(20) UNIQUE,
    address TEXT,  
    phone VARCHAR(20),  
    email VARCHAR(100),  
    unique_id VARCHAR(50) UNIQUE,  
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES company.companies(company_id)
);
CREATE INDEX idx_office_name ON company.offices(office_name);


-- Table: departments
DROP TABLE IF EXISTS company.departments;
CREATE TABLE company.departments (
    department_id SERIAL PRIMARY KEY,  -- Fixed typo here
    company_id INT NOT NULL,
    department_name VARCHAR(150) NOT NULL,
    acronym VARCHAR(20),
    description TEXT,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_company_department FOREIGN KEY (company_id) REFERENCES company.companies(company_id)
);
CREATE INDEX idx_department_name ON company.departments(department_name);


-- Table: rooms
DROP TABLE IF EXISTS company.rooms;
CREATE TABLE company.rooms (
    room_id SERIAL PRIMARY KEY,
    company_id INT NOT NULL,
    branch_id INT NOT NULL,
    room_name VARCHAR(50),
    number VARCHAR(10), 
    capacity INT NOT NULL,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_branch_room FOREIGN KEY (branch_id) REFERENCES company.branches(branch_id),
    CONSTRAINT fk_company_room FOREIGN KEY (company_id) REFERENCES company.companies(company_id)
);
CREATE INDEX idx_room_name ON company.rooms(room_name);


-- Table: projects
DROP TABLE IF EXISTS company.projects;
CREATE TABLE company.projects (
    project_id SERIAL PRIMARY KEY,
    project_name VARCHAR(100) NOT NULL,
    description TEXT,
    start_date DATE,
    end_date DATE,
    status VARCHAR(50),
    company_id INT,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES company.companies(company_id)
);

-- Table: project_attachments
DROP TABLE IF EXISTS company.project_attachments;
CREATE TABLE company.project_attachments (
    attachment_id SERIAL PRIMARY KEY,
    project_id INT NOT NULL,
    company_id INT,
    attachment_name VARCHAR(100) NOT NULL,
    file_name VARCHAR(150) NOT NULL,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES company.companies(company_id),
    FOREIGN KEY (project_id) REFERENCES company.projects(project_id)
);


-- Table: policies
DROP TABLE IF EXISTS company.policies;
CREATE TABLE company.policies (
    policy_id SERIAL PRIMARY KEY,
    company_id INT,
    policy_name VARCHAR(100) NOT NULL,
    description TEXT,
    effective_date DATE,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES company.companies(company_id)
);


-- Table: policy_attachments
DROP TABLE IF EXISTS company.policy_attachments;
CREATE TABLE company.policy_attachments (
    attachment_id SERIAL PRIMARY KEY,
    policy_id INT NOT NULL,
    company_id INT,
    attachment_name VARCHAR(100) NOT NULL,
    file_name VARCHAR(150) NOT NULL,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (company_id) REFERENCES company.companies(company_id),
    FOREIGN KEY (policy_id) REFERENCES company.policies(policy_id)
);
