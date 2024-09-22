INSERT INTO employees.professional_info(employee_id, department_id, job_title_id, employment_status_id)
SELECT employee_id, department_id, job_title_id, employment_status_id FROM employees.employees;

