package entities

type ModuleInfo struct {
    Id   int
    Code string
}

var (
    ModuleAuthentication = ModuleInfo{Code: "authentication"}
    ModuleConfigurations = ModuleInfo{Code: "configurations"}
    ModuleReferences     = ModuleInfo{Code: "references"}
    ModuleCompany        = ModuleInfo{Code: "company"}
    ModuleEmployees      = ModuleInfo{Code: "employees"}
    ModuleReports        = ModuleInfo{Code: "reports"}
)


