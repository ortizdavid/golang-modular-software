-- triggers for 'authentication' schema 


-- Lock delete assigned roles-------------------------------------------------------------------------------------
-- Drop the trigger if it already exists
DROP TRIGGER IF EXISTS lock_delete_assigned_roles ON authentication.roles;
DROP FUNCTION IF EXISTS authentication.check_role_assigned();
-- Create the trigger function
CREATE OR REPLACE FUNCTION authentication.check_role_assigned()
RETURNS TRIGGER AS $$
BEGIN
    -- Check if the role is assigned to any user
    IF EXISTS (
        SELECT 1
        FROM authentication.user_roles
        WHERE role_id = OLD.role_id
    ) THEN
        -- Raise an exception to prevent deletion
        RAISE EXCEPTION 'Cannot delete role % because it is currently assigned to one or more users.', OLD.id;
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- Create the trigger that uses the function
CREATE TRIGGER lock_delete_assigned_roles
BEFORE DELETE ON authentication.roles
FOR EACH ROW
EXECUTE FUNCTION authentication.check_role_assigned();



-- Lock delete assigned permissions-------------------------------------------------------------------------------------------
-- Drop the trigger if it already exists
DROP TRIGGER IF EXISTS lock_delete_assigned_permissions ON authentication.permissions;
DROP FUNCTION IF EXISTS authentication.check_permission_assigned();
-- Create the trigger function
CREATE OR REPLACE FUNCTION authentication.check_permission_assigned()
RETURNS TRIGGER AS $$
BEGIN
    -- Check if the permission is assigned to any role
    IF EXISTS (
        SELECT 1
        FROM authentication.permission_roles
        WHERE permission_id = OLD.permission_id
    ) THEN
        -- Raise an exception to prevent deletion
        RAISE EXCEPTION 'Cannot delete permission % because it is currently assigned to one or more roles.', OLD.id;
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;
-- Create the trigger that uses the function
CREATE TRIGGER lock_delete_assigned_permissions
BEFORE DELETE ON authentication.permissions
FOR EACH ROW
EXECUTE FUNCTION authentication.check_permission_assigned();

