-- views for 'reference' schema 

-- View: view_company_data
DROP VIEW IF EXISTS view_company_data;
CREATE VIEW view_company_data AS
SELECT co.company_id, co.unique_id,
    co.company_name, co.company_acronym,
    co.company_type, co.industry,
    co.founded_date, co.address,
    co.phone, co.email,
    co.website_url,
    TO_CHAR(co.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(co.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
FROM company.companies co
ORDER BY created_at DESC;

