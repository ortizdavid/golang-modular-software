-- views for 'reference' schema 

-- View: view_company_data
DROP VIEW IF EXISTS company.view_company_data;
CREATE VIEW company.view_company_data AS
SELECT co.company_id, co.unique_id,
    co.company_name, co.company_acronym,
    co.company_type, co.industry,
    TO_CHAR(co.founded_date, 'YYYY-MM-DD') AS founded_date,
    co.address,
    co.phone, co.email,
    co.website_url,
    TO_CHAR(co.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(co.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
FROM company.companies co
ORDER BY created_at DESC;


-- View: view_branch_data
DROP VIEW IF EXISTS company.view_branch_data;
CREATE VIEW company.view_branch_data AS
SELECT br.branch_id, br.unique_id,
    br.branch_name, br.code,
    br.address,br.phone, 
    br.email,
    TO_CHAR(br.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(br.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    co.company_id, co.company_name
FROM company.branches br 
JOIN company.companies co ON(co.company_id = br.company_id)
ORDER BY created_at DESC;

