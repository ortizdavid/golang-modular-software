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


-- View: view_office_data
DROP VIEW IF EXISTS company.view_office_data;
CREATE VIEW company.view_office_data AS
SELECT of.office_id, of.unique_id,
    of.office_name, of.address,
    of.phone, of.email,
    TO_CHAR(of.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(of.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    co.company_id, co.company_name
FROM company.offices of 
JOIN company.companies co ON(co.company_id = of.company_id)
ORDER BY created_at DESC;


-- View: view_department_data
DROP VIEW IF EXISTS company.view_department_data;
CREATE VIEW company.view_department_data AS
SELECT dpt.department_id, dpt.unique_id,
    dpt.department_name, dpt.acronym,
    dpt.description,
    TO_CHAR(dpt.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(dpt.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    co.company_id, co.company_name
FROM company.departments dpt 
JOIN company.companies co ON(co.company_id = dpt.company_id)
ORDER BY created_at DESC;


-- View: view_room_data
DROP VIEW IF EXISTS company.view_room_data;
CREATE VIEW company.view_room_data AS
SELECT rm.room_id, rm.unique_id,
    rm.room_name, rm.number,
    rm.capacity,
    TO_CHAR(rm.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(rm.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    co.company_id, co.company_name,
    br.branch_id, br.branch_name
FROM company.rooms rm 
JOIN company.companies co ON(co.company_id = rm.company_id)
JOIN company.branches br ON(br.branch_id = rm.branch_id)
ORDER BY created_at DESC;

