-- views for 'employees' schema 

-- View: view_employee_data
CREATE OR REPLACE VIEW employees.view_employee_data AS
SELECT emp.employee_id, emp.unique_id,
    emp.first_name, emp.last_name,
    emp.identification_number,
    emp.gender, 
    TO_CHAR(emp.date_of_birth, 'YYYY-MM-DD') AS date_of_birth,
    TO_CHAR(emp.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(emp.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    it.type_id AS identification_type_id, 
    it.type_name AS identification_type_name,
    co.country_id, co.country_name,
    ms.status_id AS marital_status_id,
    ms.status_name AS marital_status_name,
    dpt.department_id, dpt.department_name,
    es.status_id AS employment_status_id,
    es.status_name AS employment_status_name,
    jt.job_title_id, jt.title_name AS job_title_name
FROM employees.employees emp 
LEFT JOIN reference.identification_types it ON(it.type_id = emp.identification_type_id)
LEFT JOIN reference.countries co ON(co.country_id = emp.country_id)
LEFT JOIN reference.marital_statuses ms ON(ms.status_id = emp.marital_status_id)
LEFT JOIN reference.employment_statuses es ON(es.status_id = emp.employment_status_id)
LEFT JOIN company.departments dpt ON(dpt.department_id = emp.department_id)
LEFT JOIN employees.job_titles jt ON(jt.job_title_id = emp.job_title_id)
ORDER BY emp.created_at;



-- view: view_employee_data
CREATE OR REPLACE VIEW employees.view_document_data AS
SELECT doc.document_id, doc.unique_id,
    doc.document_name, doc.document_number,
    TO_CHAR(doc.expiration_date, 'YYYY-MM-DD') AS expiration_date,
    doc.file_name,
    doc.status,
    TO_CHAR(doc.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(doc.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    emp.employee_id, 
    emp.unique_id AS employee_unique_id,
    emp.first_name, emp.last_name,
    dt.type_id AS document_type_id,
    dt.type_name AS document_type_name
FROM employees.documents doc 
LEFT JOIN employees.employees emp ON(emp.employee_id = doc.employee_id)
LEFT JOIN employees.document_types dt ON (dt.type_id = doc.document_type_id);



-- view: view_statistics_data
CREATE OR REPLACE VIEW employees.view_statistics_data AS
SELECT 
    (SELECT COUNT(*) FROM employees.job_titles) AS job_titles,
    (SELECT COUNT(*) FROM employees.employees) AS employees,
    (SELECT COUNT(*) FROM employees.document_types) AS document_types;



