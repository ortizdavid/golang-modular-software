package entities

type ModuleInfo struct {
    Id   int
    Code string
}

var (
    ModuleAuthentication = ModuleInfo{Id: 1, Code: "authentication"}
    ModuleCompany        = ModuleInfo{Id: 2, Code: "company"}
    ModuleEmployees      = ModuleInfo{Id: 3, Code: "employees"}
    ModuleReferences     = ModuleInfo{Id: 4, Code: "references"}
    ModuleReports        = ModuleInfo{Id: 5, Code: "reports"}
    ModuleConfigurations = ModuleInfo{Id: 6, Code: "configurations"}
)
