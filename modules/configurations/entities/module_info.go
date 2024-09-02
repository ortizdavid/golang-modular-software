package entities

type ModuleInfo struct {
    Id   int
    Code string
}

var (
    ModuleAuthentication = ModuleInfo{Id: 1, Code: "authentication"}
    ModuleConfigurations = ModuleInfo{Id: 2, Code: "configurations"}
    ModuleReferences     = ModuleInfo{Id: 3, Code: "references"}
    ModuleCompany        = ModuleInfo{Id: 4, Code: "company"}
    ModuleEmployees      = ModuleInfo{Id: 5, Code: "employees"}
    ModuleReports        = ModuleInfo{Id: 6, Code: "reports"}
)


