-- View: view_user_data
DROP VIEW IF EXISTS authentication.view_user_data;
CREATE VIEW authentication.view_user_data AS 
SELECT us.user_id, us.unique_id,
	us.user_name, us.password, 
	us.active, us.user_image,
    us.token, 
	TO_CHAR(us.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(us.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
	ro.role_id, ro.role_name,
	ro.code AS role_code
FROM authentication.users us
JOIN authentication.roles ro ON(ro.role_id = us.role_id)
ORDER BY created_at DESC;


-- View: view_role_permission
DROP VIEW IF EXISTS authentication.view_user_permissions_data;
CREATE VIEW authentication.view_user_permissions_data AS 
SELECT  pe.permission_id, pe.unique_id,
    pe.permission_name, pe.code,
    pe.description,
    TO_CHAR(pe.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(pe.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    ro.role_id, ro.role_name,
    ro.code AS role_code,
    us.user_id, us.user_name
FROM authentication.permissions pe
JOIN authentication.permission_roles pr ON(pr.permission_id = pe.permission_id)
JOIN authentication.roles ro ON(ro.role_id = pr.role_id)
JOIN authentication.users us ON(ro.role_id = us.role_id)
ORDER BY created_at DESC;

