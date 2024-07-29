-- Schema: authentication

-- View: view_user_data
DROP VIEW IF EXISTS authentication.view_user_data;
CREATE VIEW authentication.view_user_data AS 
SELECT 
    us.user_id,
    us.unique_id,
    us.user_name,
    us.email,
    us.is_active,
    us.is_logged,
    us.user_image,
    us.token,
    TO_CHAR(us.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(us.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    COALESCE(la.status, 'Offline') AS status, 
    COALESCE(la.host, 'Unknown') AS host,
    COALESCE(la.browser, 'Unknown') AS browser,
    COALESCE(la.ip_address, 'Unknown') AS ip_address,
    COALESCE(la.device, 'Unknown') AS device, 
    COALESCE(la.location, 'Unknown') AS location, 
    COALESCE(TO_CHAR(la.last_login, 'YYYY-MM-DD HH24:MI:SS'), 'N/A') AS last_login,
    COALESCE(TO_CHAR(la.last_logout, 'YYYY-MM-DD HH24:MI:SS'), 'N/A') AS last_logout
FROM authentication.users us
LEFT JOIN authentication.login_activity la 
ON us.user_id = la.user_id
ORDER BY us.created_at DESC;


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

