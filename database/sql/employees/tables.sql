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
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Table: employees
DROP TABLE IF EXISTS employees.employees;
CREATE TABLE employees.employees(
    employee_id SERIAL PRIMARY KEY,
    identification_type_id INT NOT NULL,
    nationality INT,
    marital_status_id INT,
    department_id INT,
    job_title_id INT,
    employment_status_id INT NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    gender TYPE_GENDER,
    date_of_birth DATE,
    identification_number VARCHAR(30) UNIQUE NOT NULL,
    unique_id VARCHAR(50) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_identification FOREIGN KEY(identification_type_id) REFERENCES reference.identification_types(type_id),
    CONSTRAINT fk_nationality FOREIGN KEY(nationality) REFERENCES reference.countries(country_id),
    CONSTRAINT fk_marital_status FOREIGN KEY(marital_status_id) REFERENCES reference.marital_statuses(status_id),
    CONSTRAINT fk_department FOREIGN KEY(department_id) REFERENCES company.departments(department_id),
    CONSTRAINT fk_job_title FOREIGN KEY(job_title_id) REFERENCES employees.job_titles(job_title_id),
    CONSTRAINT fk_employment_status FOREIGN KEY(employment_status_id) REFERENCES reference.employment_statuses(status_id)
);
-- Indexes
CREATE INDEX idx_nataionality ON employees.employees(nationality);

