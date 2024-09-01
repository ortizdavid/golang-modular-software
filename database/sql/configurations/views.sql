-- views for 'configurations' schema 


-- View: view_module_flag_data
DROP VIEW IF EXISTS configurations.view_module_flag_data;
CREATE VIEW configurations.view_module_flag_data AS 
SELECT mf.flag_id, mf.unique_id,
    mf.status,
    TO_CHAR(mf.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(mf.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    md.module_id, md.module_name, 
    md.code
FROM configurations.module_flag mf
JOIN configurations.modules md ON (md.module_id = mf.module_id)
ORDER BY module_name;

-- View: view_core_entity_data
DROP VIEW IF EXISTS configurations.view_core_entity_data;
CREATE VIEW configurations.view_core_entity_data AS 
SELECT ce.entity_id, ce.unique_id,
    ce.entity_name, ce.code,
    ce.description,
    TO_CHAR(ce.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(ce.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    md.module_id, md.module_name
FROM configurations.core_entities ce
JOIN configurations.modules md ON (md.module_id = ce.module_id)
ORDER BY entity_name;

