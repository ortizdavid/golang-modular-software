-- views for 'authentication' schema 

-- View: view_user_data
DROP VIEW IF EXISTS authentication.view_user_data;
CREATE VIEW authentication.view_user_data AS 
SELECT 
    us.user_id,
    us.unique_id,
    us.user_name,
    us.email,
    us.password, 
    CASE WHEN us.is_active THEN 'Yes' ELSE 'No' END AS is_active,
    us.user_image,
    us.initial_role,
    us.token,
    TO_CHAR(us.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(us.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    us.status
FROM authentication.users us
ORDER BY us.created_at ASC;


-- View: view_role_data
DROP VIEW IF EXISTS authentication.view_role_data;
CREATE VIEW authentication.view_role_data AS 
SELECT  ro.role_id, ro.unique_id,
    ro.role_name, ro.code,
    ro.description, ro.status,
    ro.created_at,
    ro.updated_at
FROM authentication.roles ro
ORDER BY created_at ASC;


-- View: view_user_role_data
DROP VIEW IF EXISTS authentication.view_user_role_data;
CREATE VIEW authentication.view_user_role_data AS 
SELECT  ur.user_role_id, ur.unique_id,
    TO_CHAR(ur.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(ur.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    ro.role_id, ro.role_name,
    ro.code AS role_code,
    ro.status AS role_status,
    us.user_id, us.unique_id AS user_unique_id,
    us.user_name
FROM authentication.user_roles ur
JOIN authentication.roles ro ON(ro.role_id = ur.role_id)
JOIN authentication.users us ON(us.user_id = ur.user_id)
ORDER BY created_at DESC;


-- View: view_user_association_data
DROP VIEW IF EXISTS authentication.view_user_association_data;
CREATE VIEW authentication.view_user_association_data AS 
SELECT  ua.association_id, entity_id,
    ua.unique_id, ua.entity_name,
    TO_CHAR(ua.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(ua.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    us.user_id, us.unique_id AS user_unique_id,
    us.user_name
FROM authentication.user_associations ua
JOIN authentication.users us ON(us.user_id = ua.user_id)
ORDER BY created_at DESC;


-- View: view_permission_data
DROP VIEW IF EXISTS authentication.view_permission_data;
CREATE VIEW authentication.view_permission_data AS 
SELECT  pe.permission_id, pe.unique_id,
    pe.permission_name, pe.code,
    pe.description,
    TO_CHAR(pe.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(pe.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at
FROM authentication.permissions pe
ORDER BY created_at ASC;


-- View: view_permission_role_data
DROP VIEW IF EXISTS authentication.view_permission_role_data;
CREATE VIEW authentication.view_permission_role_data AS 
SELECT pr.permission_role_id, pr.unique_id,  
    TO_CHAR(pr.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(pr.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    pe.permission_id, pe.permission_name, 
    pe.code AS permission_code,
    ro.role_id, ro.unique_id AS role_unique_id,
    ro.role_name
FROM authentication.permissions pe
JOIN authentication.permission_roles pr ON(pr.permission_id = pe.permission_id)
JOIN authentication.roles ro ON(ro.role_id = pr.role_id)
ORDER BY created_at DESC;


-- View: view_login_activity_data
DROP VIEW IF EXISTS authentication.view_login_activity_data;
CREATE VIEW authentication.view_login_activity_data AS
SELECT la.login_id, la.unique_id,
    la.status, 
    COALESCE(la.host, 'Unknown') AS host,
    COALESCE(la.browser, 'Unknown') AS browser,
    COALESCE(la.ip_address, 'Unknown') AS ip_address,
    COALESCE(la.device, 'Unknown') AS device, 
    COALESCE(la.location, 'Unknown') AS location, 
    COALESCE(TO_CHAR(la.last_login, 'YYYY-MM-DD HH24:MI:SS'), 'N/A') AS last_login,
    COALESCE(TO_CHAR(la.last_logout, 'YYYY-MM-DD HH24:MI:SS'), 'N/A') AS last_logout,
    la.total_login,
    la.total_logout,
    TO_CHAR(la.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(la.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    us.user_id, us.user_name,
    us.email
FROM authentication.login_activity la
JOIN authentication.users us ON(la.user_id = us.user_id)
ORDER BY created_at DESC;



--- View: view_user_api_key_data
DROP VIEW IF EXISTS authentication.view_user_api_key_data;
CREATE VIEW authentication.view_user_api_key_data AS 
SELECT uak.api_key_id, uak.unique_id,
    uak.x_user_id, uak.x_api_key, 
    CASE WHEN uak.is_active THEN 'Yes' ELSE 'No' END AS is_active,
    uak.created_by,
    TO_CHAR(uak.expires_at, 'YYYY-MM-DD HH24:MI:SS') AS expires_at,
    TO_CHAR(uak.created_at, 'YYYY-MM-DD HH24:MI:SS') AS created_at,
    TO_CHAR(uak.updated_at, 'YYYY-MM-DD HH24:MI:SS') AS updated_at,
    us.user_id, us.user_name,
    us.email
FROM authentication.user_api_key uak
JOIN authentication.users us ON(uak.user_id = us.user_id)
ORDER BY created_at DESC;


-- view: view_statistics_data
DROP VIEW IF EXISTS authentication.view_statistics_data;
CREATE VIEW authentication.view_statistics_data AS
SELECT 
    COUNT(user_id) AS users,
    SUM(CASE WHEN is_active THEN 1 ELSE 0 END) AS active_users,
    SUM(CASE WHEN NOT is_active THEN 1 ELSE 0 END) AS inactive_users,
    -- Calculate online and offline users without referencing the main query's tables
    (SELECT COUNT(DISTINCT user_id) FROM authentication.users WHERE status = 'Online') AS online_users, 
    (SELECT COUNT(DISTINCT user_id) FROM authentication.users WHERE status = 'Offline') AS offline_users, 
    (SELECT COUNT(role_id) FROM authentication.roles) AS roles,
    (SELECT COUNT(permission_id) FROM authentication.permissions) AS permissions, 
    (SELECT COUNT(*) FROM authentication.login_activity) AS login_activity 
FROM authentication.users;


