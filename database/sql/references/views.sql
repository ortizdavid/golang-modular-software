-- view: view_country_data
DROP VIEW IF EXISTS reference.view_country_data;
CREATE VIEW reference.view_country_data AS 
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

