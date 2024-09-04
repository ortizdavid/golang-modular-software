package entities

type CoreEntityInfo struct {
	Id   int
	Code string
}

var (
	// Module: Authentication
	CoreEntityUsers            = CoreEntityInfo{Code: "authentication.users"}
	CoreEntityActiveUsers      = CoreEntityInfo{Code: "authentication.active_users"}
	CoreEntityInactiveUsers    = CoreEntityInfo{Code: "authentication.inactive_users"}
	CoreEntityOnlineUsers      = CoreEntityInfo{Code: "authentication.online_users"}
	CoreEntityOfflineUsers     = CoreEntityInfo{Code: "authentication.offline_users"}
	CoreEntityRoles            = CoreEntityInfo{Code: "authentication.roles"}
	CoreEntityPermissions      = CoreEntityInfo{Code: "authentication.permissions"}
	CoreEntityLoginActivity    = CoreEntityInfo{Code: "authentication.login_activity"}

	// Module: Configurations
	CoreEntityBasicConfigurations  = CoreEntityInfo{Code: "configurations.basic_configurations"}
	CoreEntityCompanyConfigurations = CoreEntityInfo{Code: "configurations.company_configurations"}
	CoreEntityEmailConfigurations   = CoreEntityInfo{Code: "configurations.email_configurations"}
	CoreEntityModules               = CoreEntityInfo{Code: "configurations.modules"}
	CoreEntityCoreEntities          = CoreEntityInfo{Code: "configurations.core_entities"}
	CoreEntityModuleFlags           = CoreEntityInfo{Code: "configurations.module_flags"}
	CoreEntityCoreEntityFlags       = CoreEntityInfo{Code: "configurations.core_entity_flags"}

	// Module: References
	CoreEntityCountries           = CoreEntityInfo{Code: "references.countries"}
	CoreEntityCurrencies          = CoreEntityInfo{Code: "references.currencies"}
	CoreEntityIdentificationTypes = CoreEntityInfo{Code: "references.identification_types"}
	CoreEntityContactTypes        = CoreEntityInfo{Code: "references.contact_types"}
	CoreEntityMaritalStatuses     = CoreEntityInfo{Code: "references.marital_statuses"}
	CoreEntityTaskStatuses        = CoreEntityInfo{Code: "references.task_statuses"}
	CoreEntityApprovalStatuses    = CoreEntityInfo{Code: "references.approval_statuses"}
	CoreEntityDocumentStatuses    = CoreEntityInfo{Code: "references.document_statuses"}
	CoreEntityWorkflowStatuses    = CoreEntityInfo{Code: "references.workflow_statuses"}
	CoreEntityEvaluationStatuses  = CoreEntityInfo{Code: "references.evaluation_statuses"}
	CoreEntityUserStatuses        = CoreEntityInfo{Code: "references.user_statuses"}
	CoreEntityEmploymentStatuses  = CoreEntityInfo{Code: "references.employment_statuses"}

	// Module: Company
	CoreEntityCompanyInfo = CoreEntityInfo{Code: "company.company_info"}
	CoreEntityBranches    = CoreEntityInfo{Code: "company.branches"}
	CoreEntityOffices     = CoreEntityInfo{Code: "company.offices"}
	CoreEntityDepartments = CoreEntityInfo{Code: "company.departments"}
	CoreEntityRooms       = CoreEntityInfo{Code: "company.rooms"}
	CoreEntityProjects    = CoreEntityInfo{Code: "company.projects"}
	CoreEntityPolicies    = CoreEntityInfo{Code: "company.policies"}

	// Module: Employees
	CoreEntityEmployees = CoreEntityInfo{Code: "employees.employees"}
	CoreEntityJobTitles = CoreEntityInfo{Code: "employees.job_titles"}

	// Module: Reports
	CoreEntityUserReports          = CoreEntityInfo{Code: "reports.user_reports"}
	CoreEntityConfigurationReports = CoreEntityInfo{Code: "reports.configuration_reports"}
	CoreEntityCompanyReports       = CoreEntityInfo{Code: "reports.company_reports"}
	CoreEntityEmployeeReports      = CoreEntityInfo{Code: "reports.employee_reports"}
	CoreEntityReferenceReports     = CoreEntityInfo{Code: "reports.reference_reports"}
)
