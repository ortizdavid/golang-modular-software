package entities

type RoleInfo struct {
    Id      int
    Code    string
}

var (
    RoleSuperAdmin = RoleInfo { Id: 1, Code: "role_super_admin"}
    RoleAdmin      = RoleInfo { Id: 2, Code: "role_role_admin"}
    RoleEmployee   = RoleInfo { Id: 3, Code: "role_employee"}
)
