-- view: view_country_data
CREATE OR REPLACE VIEW reference.view_country_data AS 
SELECT 
    co.country_id, 
    co.unique_id,
    co.country_name,
    co.iso_code,
    LOWER(co.iso_code) AS iso_code_lower,
    co.dialing_code,
    TO_CHAR(co.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(co.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
FROM reference.countries co
ORDER BY co.created_at DESC; 



-- view: view_statistics_data
CREATE OR REPLACE VIEW reference.view_statistics_data AS
SELECT 
    (SELECT COUNT(*) FROM reference.countries) AS countries,
    (SELECT COUNT(*) FROM reference.currencies) AS currencies,
    (SELECT COUNT(*) FROM reference.identification_types) AS identification_types,
    (SELECT COUNT(*) FROM reference.contact_types) AS contact_types,
    (SELECT COUNT(*) FROM reference.marital_statuses) AS marital_statuses,
    (SELECT COUNT(*) FROM reference.task_statuses) AS task_statuses,
    (SELECT COUNT(*) FROM reference.approval_statuses) AS approval_statuses,
    (SELECT COUNT(*) FROM reference.document_statuses) AS document_statuses,
    (SELECT COUNT(*) FROM reference.workflow_statuses) AS workflow_statuses,
    (SELECT COUNT(*) FROM reference.evaluation_statuses) AS evaluation_statuses,
    (SELECT COUNT(*) FROM reference.user_statuses) AS user_statuses,
    (SELECT COUNT(*) FROM reference.employment_statuses) AS employment_statuses;

