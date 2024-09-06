-- views for 'employees' schema 

-- View: view_employee_data
DROP VIEW IF EXISTS employees.view_employee_data;
CREATE VIEW employees.view_employee_data AS
SELECT emp.employee_id, emp.unique_id,
    emp.first_name, emp.last_name,
    emp.identification_number,
    emp.gender, emp.date_of_birth,
    TO_CHAR(emp.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(emp.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    it.type_id AS identification_type_id, 
    it.type_name AS identification_type_name,
    co.country_id, co.country_name,
    ms.status_id AS marital_status_id,
    ms.status_name AS marital_status_name,
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



