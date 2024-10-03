package repositories

import (
	"context"
	"fmt"
	"github.com/ortizdavid/golang-modular-software/database"
	"github.com/ortizdavid/golang-modular-software/modules/employees/entities"
)

type EmployeeCompleteDataRepository struct {
    db *database.Database
}

func NewEmployeeCompleteDataRepository(db *database.Database) *EmployeeCompleteDataRepository {
    return &EmployeeCompleteDataRepository{
        db: db,
    }
}

// GetByUniqueId retrieves complete employee data by a given unique identifier.
func (repo *EmployeeCompleteDataRepository) GetByUniqueId(ctx context.Context, uniqueId string) (entities.EmployeeCompleteData, error) {
    return repo.findBy(ctx, "unique_id", uniqueId)
}

// GetByIdentificationNumber retrieves complete employee data by a given identification number.
func (repo *EmployeeCompleteDataRepository) GetByIdentificationNumber(ctx context.Context, identNumber string) (entities.EmployeeCompleteData, error) {
    return repo.findBy(ctx, "identification_number", identNumber)
}

// findBy retrieves employee data dynamically based on the provided field (either 'unique_id' or 'identification_number'). It returns the complete data if found, otherwise returns an empty result.
func (repo *EmployeeCompleteDataRepository) findBy(ctx context.Context, conditionField, conditionValue string) (entities.EmployeeCompleteData, error) {
    var rawDataList []entities.EmployeeCompleteDataRaw
    var completeData entities.EmployeeCompleteData

    query := fmt.Sprintf("SELECT * FROM employees.view_employee_complete_data WHERE %s=?", conditionField)
    result := repo.db.WithContext(ctx).Raw(query, conditionValue).Scan(&rawDataList)
    if result.Error != nil {
        return completeData, result.Error
    }
    // If no data is found, return an empty structure.
    if len(rawDataList) == 0 {
        return completeData, nil 
    }
    return repo.buildCompleteEmployeeData(rawDataList), nil
}

// buildCompleteEmployeeData aggregates multiple pieces of employee-related data (documents, phones, emails, addresses) into a single EmployeeCompleteData structure.
func (repo *EmployeeCompleteDataRepository) buildCompleteEmployeeData(rawDataList []entities.EmployeeCompleteDataRaw) entities.EmployeeCompleteData {
    var completeData entities.EmployeeCompleteData
    firstRow := rawDataList[0]

    // Populate the main employee, personal, and professional data from the first row.
    addSingleData(&completeData, firstRow)

    // Maps to track existing entries and prevent duplicates in documents, phones, emails, and addresses.
    phoneMap := make(map[string]bool)
    emailMap := make(map[string]bool)
    documentMap := make(map[string]bool)
    addressMap := make(map[string]bool)

    // Iterate through each row and append unique documents, phones, emails, and addresses to the completeData.
    for _, row := range rawDataList {
        addDocumentIfNotExists(&completeData, row, documentMap)
        addPhoneIfNotExists(&completeData, row, phoneMap)
        addEmailIfNotExists(&completeData, row, emailMap)
        addAddressIfNotExists(&completeData, row, addressMap)
    }

    return completeData
}

// addSingleData populates the main fields of the EmployeeCompleteData structure with the basic employee, personal, and professional information.
func addSingleData(completeData *entities.EmployeeCompleteData, firstRow entities.EmployeeCompleteDataRaw) {
    completeData.EmployeeId = firstRow.EmployeeId
    completeData.UniqueId = firstRow.UniqueId
    completeData.CreatedAt = firstRow.CreatedAt
    completeData.UpdatedAt = firstRow.UpdatedAt

    // Personal information
    completeData.PersonalInfo = entities.PersonalBasic{
        FirstName:            firstRow.FirstName,
        LastName:             firstRow.LastName,
        IdentificationNumber: firstRow.IdentificationNumber,
        Gender:               firstRow.Gender,
        DateOfBirth:          firstRow.DateOfBirth,
        IdentificationTypeName: firstRow.IdentificationTypeName,
        CountryName:          firstRow.CountryName,
        MaritalStatusName:    firstRow.MaritalStatusName,
    }

    // Professional information
    completeData.ProfessionalInfo = entities.ProfessionalBasic{
        DepartmentName:       firstRow.DepartmentName,
        EmploymentStatusName: firstRow.EmploymentStatusName,
        JobTitleName:         firstRow.JobTitleName,
    }

    // User account information
    completeData.UserAccount = entities.AccountBasic{
        UserName:  firstRow.UserName,
        UserEmail: firstRow.UserEmail,
    }
}

// Helper functions for adding unique entries of documents, phones, emails, and addresses ---------------------------------------

// addDocumentIfNotExists appends a document to the EmployeeCompleteData only if it doesn't already exist in the documentMap.
func addDocumentIfNotExists(data *entities.EmployeeCompleteData, row entities.EmployeeCompleteDataRaw, docMap map[string]bool) {
    docKey := row.DocumentName + row.DocumentNumber
    if !docMap[docKey] {
        data.Documents = append(data.Documents, entities.DocumentBasic{
            DocumentName:     row.DocumentName,
            DocumentNumber:   row.DocumentNumber,
            ExpirationDate:   row.ExpirationDate,
            FileName:         row.FileName,
            Status:           row.DocumentStatus,
            DocumentTypeName: row.DocumentTypeName,
        })
        docMap[docKey] = true
    }
}

// addPhoneIfNotExists appends a phone number to the EmployeeCompleteData only if it doesn't already exist in the phoneMap.
func addPhoneIfNotExists(data *entities.EmployeeCompleteData, row entities.EmployeeCompleteDataRaw, phoneMap map[string]bool) {
    if !phoneMap[row.PhoneNumber] {
        data.Phones = append(data.Phones, entities.PhoneBasic{
            PhoneNumber:     row.PhoneNumber,
            ContactTypeName: row.PhoneContactTypeName,
        })
        phoneMap[row.PhoneNumber] = true
    }
}

// addEmailIfNotExists appends an email to the EmployeeCompleteData only if it doesn't already exist in the emailMap.
func addEmailIfNotExists(data *entities.EmployeeCompleteData, row entities.EmployeeCompleteDataRaw, emailMap map[string]bool) {
    if !emailMap[row.EmailAddress] {
        data.Emails = append(data.Emails, entities.EmailBasic{
            EmailAddress:    row.EmailAddress,
            ContactTypeName: row.EmailContactTypeName,
        })
        emailMap[row.EmailAddress] = true
    }
}

// addAddressIfNotExists appends an address to the EmployeeCompleteData only if it doesn't already exist in the addressMap.
func addAddressIfNotExists(data *entities.EmployeeCompleteData, row entities.EmployeeCompleteDataRaw, addressMap map[string]bool) {
    addressKey := row.State + row.City + row.Street
    if !addressMap[addressKey] {
        data.Addresses = append(data.Addresses, entities.AddressBasic{
            State:            row.State,
            City:             row.City,
            Neighborhood:     row.Neighborhood,
            Street:           row.Street,
            HouseNumber:      row.HouseNumber,
            PostalCode:       row.PostalCode,
            CountryCode:      row.CountryCode,
            AdditionalDetails: row.AdditionalDetails,
            IsCurrent:        row.IsCurrent,
        })
        addressMap[addressKey] = true
    }
}
