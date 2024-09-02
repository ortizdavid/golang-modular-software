package entities

type CoreEntityInfo struct {
	Id   int
	Code string
}

var (
	// Module: Authentication
	CoreEntityUsers            = CoreEntityInfo{Id: 1, Code: "authentication.users"}
	CoreEntityActiveUsers      = CoreEntityInfo{Id: 2, Code: "authentication.active_users"}
	CoreEntityInactiveUsers    = CoreEntityInfo{Id: 3, Code: "authentication.inactive_users"}
	CoreEntityOnlineUsers      = CoreEntityInfo{Id: 4, Code: "authentication.online_users"}
	CoreEntityOfflineUsers     = CoreEntityInfo{Id: 5, Code: "authentication.offline_users"}
	CoreEntityRoles            = CoreEntityInfo{Id: 6, Code: "authentication.roles"}
	CoreEntityPermissions      = CoreEntityInfo{Id: 7, Code: "authentication.permissions"}
	CoreEntityLoginActivity    = CoreEntityInfo{Id: 8, Code: "authentication.login_activity"}

	// Module: Configurations
	CoreEntityBasicConfigurations  = CoreEntityInfo{Id: 9, Code: "configurations.basic_configurations"}
	CoreEntityCompanyConfigurations = CoreEntityInfo{Id: 10, Code: "configurations.company_configurations"}
	CoreEntityEmailConfigurations   = CoreEntityInfo{Id: 11, Code: "configurations.email_configurations"}
	CoreEntityModules               = CoreEntityInfo{Id: 12, Code: "configurations.modules"}
	CoreEntityCoreEntities          = CoreEntityInfo{Id: 13, Code: "configurations.core_entities"}
	CoreEntityModuleFlags           = CoreEntityInfo{Id: 14, Code: "configurations.module_flags"}
	CoreEntityCoreEntityFlags       = CoreEntityInfo{Id: 15, Code: "configurations.core_entity_flags"}

	// Module: References
	CoreEntityCountries           = CoreEntityInfo{Id: 16, Code: "references.countries"}
	CoreEntityCurrencies          = CoreEntityInfo{Id: 17, Code: "references.currencies"}
	CoreEntityIdentificationTypes = CoreEntityInfo{Id: 18, Code: "references.identification_types"}
	CoreEntityContactTypes        = CoreEntityInfo{Id: 19, Code: "references.contact_types"}
	CoreEntityMaritalStatuses     = CoreEntityInfo{Id: 20, Code: "references.marital_statuses"}
	CoreEntityTaskStatuses        = CoreEntityInfo{Id: 21, Code: "references.task_statuses"}
	CoreEntityApprovalStatuses    = CoreEntityInfo{Id: 22, Code: "references.approval_statuses"}
	CoreEntityDocumentStatuses    = CoreEntityInfo{Id: 23, Code: "references.document_statuses"}
	CoreEntityWorkflowStatuses    = CoreEntityInfo{Id: 24, Code: "references.workflow_statuses"}
	CoreEntityEvaluationStatuses  = CoreEntityInfo{Id: 25, Code: "references.evaluation_statuses"}
	CoreEntityUserStatuses        = CoreEntityInfo{Id: 26, Code: "references.user_statuses"}
	CoreEntityEmploymentStatuses  = CoreEntityInfo{Id: 27, Code: "references.employment_statuses"}

	// Module: Company
	CoreEntityCompanyInfo = CoreEntityInfo{Id: 28, Code: "company.company_info"}
	CoreEntityBranches    = CoreEntityInfo{Id: 29, Code: "company.branches"}
	CoreEntityOffices     = CoreEntityInfo{Id: 30, Code: "company.offices"}
	CoreEntityDepartments = CoreEntityInfo{Id: 31, Code: "company.departments"}
	CoreEntityRooms       = CoreEntityInfo{Id: 32, Code: "company.rooms"}
	CoreEntityProjects    = CoreEntityInfo{Id: 33, Code: "company.projects"}
	CoreEntityPolicies    = CoreEntityInfo{Id: 34, Code: "company.policies"}

	// Module: Employees
	CoreEntityEmployees = CoreEntityInfo{Id: 35, Code: "employees.employees"}
	CoreEntityJobTitles = CoreEntityInfo{Id: 36, Code: "employees.job_titles"}

	// Module: Reports
	CoreEntityUserReports          = CoreEntityInfo{Id: 37, Code: "reports.user_reports"}
	CoreEntityConfigurationReports = CoreEntityInfo{Id: 38, Code: "reports.configuration_reports"}
	CoreEntityCompanyReports       = CoreEntityInfo{Id: 39, Code: "reports.company_reports"}
	CoreEntityEmployeeReports      = CoreEntityInfo{Id: 40, Code: "reports.employee_reports"}
	CoreEntityReferenceReports     = CoreEntityInfo{Id: 41, Code: "reports.reference_reports"}
)
