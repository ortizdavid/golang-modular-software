package entities

type RoleInfo struct {
    Id   int
    Code string
}

var (
    RoleSuperAdmin = RoleInfo{Id: 1, Code: "role_super_admin"}
    RoleAdmin      = RoleInfo{Id: 2, Code: "role_admin"}
    RoleManager    = RoleInfo{Id: 3, Code: "role_manager"}
    RoleEmployee   = RoleInfo{Id: 4, Code: "role_employee"}
    RoleCustomer   = RoleInfo{Id: 5, Code: "role_customer"}
    RoleSupplier   = RoleInfo{Id: 6, Code: "role_supplier"}
    RoleSupport    = RoleInfo{Id: 7, Code: "role_support"}
    RoleDeveloper  = RoleInfo{Id: 8, Code: "role_developer"}
    RoleGuest      = RoleInfo{Id: 9, Code: "role_guest"}
)
