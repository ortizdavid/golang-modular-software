-- Schema: authentication

-- View: view_user_data
DROP VIEW IF EXISTS authentication.view_user_data;
CREATE VIEW authentication.view_user_data AS 
SELECT us.user_id, us.unique_id,
	us.user_name, us.email,
    us.password, us.is_active, 
    us.user_image, us.token, 
	TO_CHAR(us.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(us.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
FROM authentication.users us
ORDER BY created_at DESC;

-- View: view_user_role_data
DROP VIEW IF EXISTS authentication.view_user_role_data;
CREATE VIEW authentication.view_user_role_data AS 
SELECT  ur.user_role_id, ur.unique_id,
    TO_CHAR(ur.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(ur.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    ro.role_id, ro.role_name,
    ro.code AS role_code,
    us.user_id, us.user_name
FROM authentication.user_roles ur
JOIN authentication.roles ro ON(ro.role_id = ur.role_id)
JOIN authentication.users us ON(us.user_id = ur.user_id)
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
JOIN authentication.users us ON(us.user_id = pr.role_id)
ORDER BY created_at DESC;
